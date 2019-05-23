 
package cryptofs

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "95511870061fb3a2899aa6b2dc9838aa"

	pwd, err := os.Getwd()
	handleErr(err, t)


	f, err := os.Open("original.txt")
	handleErr(err, t)
	defer func() {
		f.Close()
	}()

	fstat, err := f.Stat()
	handleErr(err, t)

	fileInfo := &File{fstat, "txt", pwd + `/original.txt`}

	var buf []byte
	buffer := bytes.NewBuffer(buf)
	err = fileInfo.Encrypt(key, buffer)
	handleErr(err, t)

	content, err := ioutil.ReadAll(buffer)
	handleErr(err, t)


	f2, err := os.Open("enc.txt")
	handleErr(err, t)
	defer func() {
		f2.Close()
	}()

	fstat, err = f2.Stat()
	handleErr(err, t)

	fileInfo = &File{fstat, "txt", pwd + `/enc.txt`}

	buffer.Reset()
	err = fileInfo.Decrypt(key, buffer)

	content, err = ioutil.ReadAll(buffer)
	handleErr(err, t)
    f3, err := os.Open("decrypted.txt")
    err = ioutil.WriteFile("decrypted.txt", content, 0600)
    handleErr(err, t)
    f3.Close()
	if string(content) != "Hello World" {
		t.Errorf("Expect 'Hello World' but got %s", string(content))
	}

}

func TestReplaceBy(t *testing.T) {
	pwd, err := os.Getwd()
	handleErr(err, t)

	err = ioutil.WriteFile("original.txt", []byte("Hello World"), 0600)
	handleErr(err, t)

	f, err := os.Open("original.txt")
	handleErr(err, t)

	fstat, err := f.Stat()
	handleErr(err, t)
	f.Close()

	fileInfo := &File{fstat, "txt", pwd + `/original.txt`}

	err = ioutil.WriteFile("stub.txt", []byte("Hi"), 0600)
	handleErr(err, t)
	defer os.Remove("stub.txt")

	err = fileInfo.ReplaceBy("stub.txt")
	handleErr(err, t)

	content, err := ioutil.ReadFile("original.txt")
	handleErr(err, t)

	if string(content) != "Hi" {
		t.Errorf("Expect 'Hi' but got %s", string(content))
	}

}

func handleErr(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("An error ocurred: %s", err)
	}
}
