package test

import (
	"encoding/hex"
	"fmt"
	"github.com/ziyeziye/gotool"
	"testing"
)

func TestGenerate(t *testing.T) {
	//进行加密
	generate := gotool.BcryptUtils.Generate("admin123")
	fmt.Println(generate)
	//进行加密对比
	hash := gotool.BcryptUtils.CompareHash(generate, "123456789")
	fmt.Println(hash)
}

func TestMd5(t *testing.T) {
	md5 := gotool.BcryptUtils.MD5("123456789")
	fmt.Println(md5)
}

func TestRsa(t *testing.T) {
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey := gotool.BcryptUtils.GenRsaKey()
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "我是被加密的数据记住我哦-------------------------------"
	fmt.Println("对消息进行签名操作...")
	signData := gotool.BcryptUtils.RsaSignWithSha256([]byte(data), prvKey)
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if gotool.BcryptUtils.RsaVerySignWithSha256([]byte(data), signData, pubKey) {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext := gotool.BcryptUtils.RsaEncrypt([]byte(data), pubKey)
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := gotool.BcryptUtils.RsaDecrypt(ciphertext, prvKey)
	fmt.Println("私钥解密后的数据：", string(sourceData))
}
