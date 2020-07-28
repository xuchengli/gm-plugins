/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package primitive

import (
	"crypto"
	"io"
)

type Encrypter interface {
	//Encrypt(rand io.Reader, plaintext []byte, opts EncrypterOpts) (ciphertext []byte, err error)
	Encrypt(plaintext []byte) (ciphertext []byte, err error)
}

type Decrypter interface {
	//Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
	Decrypt(ciphertext []byte) (plaintext []byte, err error)
}

type Verifier interface {
	Verify(msg, sig []byte) bool
}

type EncrypterOpts interface{}

type Sm2PublicKey struct {
	crypto.PublicKey
	Verifier
	Encrypter
}

type Sm2PrivateKey struct {
	crypto.PrivateKey
	crypto.Signer
	Decrypter
}

func (s *Sm2PrivateKey) Public() *Sm2PublicKey {
	return &Sm2PublicKey{
		PublicKey: s.Signer.Public(),
	}
}

func (s *Sm2PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	return s.Signer.Sign(rand, digest, opts)
}

type KeysGenerator interface {
	// 生成sm2公私钥对
	GenPrivateKey() (*Sm2PrivateKey, error)
	// 从sm2私钥导出公钥
	PublicKey(k *Sm2PrivateKey) *Sm2PublicKey
	// 从加密的pkcs8数据解析出sm2私钥
	ParsePKCS8PrivateKey(der, pwd []byte) (*Sm2PrivateKey, error)
	// 根据给定密钥，加密sm2私钥，输出ASN.1编码数据
	MarshalPKCS8PrivateKey(k *Sm2PrivateKey, pwd []byte) (der []byte, err error)
	// 从ASN.1编码数据解析出sm2公钥
	ParsePKIXPublicKey(der []byte) (*Sm2PublicKey, error)
	// 序列化sm2私钥，输出ASN.1编码数据
	MarshalPKIXPublicKey(k *Sm2PublicKey) ([]byte, error)
}
