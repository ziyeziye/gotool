package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	open, err := os.Open("/Users/fanyanan/Downloads/gotool")
	if err != nil {
		t.Fatal(err)
	}
	files := []*os.File{open}
	flag, err := gotool.ZipUtils.Compress(files, "/Users/fanyanan/Downloads/test.zip")
	if err != nil {
		t.Fatal(err)
	}
	if flag {
		fmt.Println("压缩成功")
	}
}

func TestDeCompress(t *testing.T) {
	compress, err := gotool.ZipUtils.DeCompress("/Users/fanyanan/Downloads/test.zip", "/Users/fanyanan/Downloads")
	if err != nil {
		t.Fatal(err)
	}
	if compress {
		fmt.Println("解压成功")
	}
}
