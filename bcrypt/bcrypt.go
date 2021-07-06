package bcrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"github.com/druidcaesa/gotool/logs"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type BcryptUtil struct {
	logs logs.Logs
}

// Generate Password encryption 密码加密
func (b *BcryptUtil) Generate(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		b.logs.ErrorLog().Println(err)
	}
	return string(hash)
}

// CompareHash Password validation 密码验证
func (b *BcryptUtil) CompareHash(dbPassword string, loginPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(loginPassword))
	if err != nil {
		b.logs.ErrorLog().Println(err)
		return false
	}
	return true
}

// MD5 md5签名  signature
func (b *BcryptUtil) MD5(s string) (string, error) {
	hash := md5.New()
	_, err := io.WriteString(hash, s)
	if err != nil {
		b.logs.ErrorLog().Println(err)
		return "", err
	}
	return s, err
}

// SHA1 sha1加密 encryption
func (b *BcryptUtil) SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))

}

// ComputeHmacSha256 hmac_sha256 encryption
func (b *BcryptUtil) ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}

// RsaEncrypt encryption
//@param origData Encrypted byte array
//@param publicKey Public key
func (b *BcryptUtil) RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData) //RSA算法加密
}

// RsaDecrypt Decrypt
//@param ciphertext Decrypt byte array
//@param privateKey Private key
func (b *BcryptUtil) RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext) //RSA算法解密
}
