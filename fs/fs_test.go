package fs

import (
	"fmt"
	"os"
	"testing"
)

func TestFile_get_contents(t *testing.T) {
	res, _ := FileGetContents("fs.go")
	if len(res) == 0 {
		t.FailNow()
	}
	md5_str, err := FileMd5("./fs.go")
	fmt.Println("fs.go md5:", md5_str, err)
}
func TestFile_put_contents(t *testing.T) {
	test_data := "hello"
	if FileExists("aaa") {
		os.Remove("aaa")
	}
	FilePutContents("aaa", []byte(test_data))
	res, _ := FileGetContents("aaa")
	if string(res) != test_data {
		t.FailNow()
	}
	FilePutContents("aaa", []byte("nihao"), FILE_APPEND)
	res, _ = FileGetContents("aaa")
	if string(res) != test_data+"nihao" {
		t.FailNow()
	}
	os.Remove("aaa")
}
