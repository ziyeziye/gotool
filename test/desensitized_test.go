package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
)

func TestDesensitized(t *testing.T) {
	//mail desensitization
	mail := "testhello@gmail.com"
	star := gotool.DesensitizedUtils.DefaultDesensitized(mail)
	fmt.Println("mail-------------------->", star)
	//phone desensitization
	phone := "13333333333"
	desensitized := gotool.DesensitizedUtils.DefaultDesensitized(phone)
	fmt.Println("phone------------------->", desensitized)
	//customize desensitization
	//customize := "sadasdasdkljkldfjlkdjflkjsdf"
	//hideStar := gotool.DesensitizedUtils.CustomizeHash(customize, 4, 14)
	//fmt.Println("customize--------------->", hideStar)
}
