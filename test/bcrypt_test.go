package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
)

func TestGenerate(t *testing.T) {
	generate := gotool.BcryptUtils.Generate("sssssss")
	fmt.Println(generate)
}
