package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
)

func TestStr(t *testing.T) {
	fullFilename := "test.5465465"
	suffix, err := gotool.StrUtils.RemoveSuffix(fullFilename)
	if err != nil {

	}
	fmt.Println(suffix)
}
