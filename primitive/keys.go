/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package primitive

import (
	"crypto"
	"crypto/elliptic"
	"io"
	"math/big"
)

// sm2公钥，与ecdsa公钥结构一致，用来统一sm2公钥结构，其他国密实现的sm2公钥可以统一转换成该结构体
type Sm2PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

// sm2私钥，与ecdsa私钥结构一致，用来统一sm2私钥结构，其他国密实现的sm2私钥可以统一转换成该结构体
type Sm2PrivateKey struct {
	Sm2PublicKey
	D *big.Int
}

// 返回sm2私钥对应的公钥
func (s *Sm2PrivateKey) Public() crypto.PublicKey {
	return &s.Sm2PublicKey
}

// 该方法尚未实现，由于可能对应多种不同类型私钥的签名方法，无法统一，签名方法使用 Sm2Crypto.Sign 函数
func (s *Sm2PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	panic("Do not invoke me! Please replace for Sm2Crypto.Sign.")
}

type KeysGenerator interface {
	// 生成sm2公私钥对
	GenPrivateKey() (*Sm2PrivateKey, error)
	// 从sm2私钥导出公钥
	PublicKey(k *Sm2PrivateKey) *Sm2PublicKey
	// 从未加密的pkcs8数据解析出sm2私钥
	ParsePKCS8UnecryptedPrivateKey(der []byte) (*Sm2PrivateKey, error)
	// 从加密的pkcs8数据解析出sm2私钥
	ParsePKCS8PrivateKey(der, pwd []byte) (*Sm2PrivateKey, error)
	// 根据给定密钥，加密sm2私钥，输出ASN.1编码数据
	MarshalSm2PrivateKey(k *Sm2PrivateKey, pwd []byte) ([]byte, error)
	// 从ASN.1编码数据解析出sm2公钥
	ParseSm2PublicKey(der []byte) (*Sm2PublicKey, error)
	// 序列化sm2私钥，输出ASN.1编码数据
	MarshalSm2PublicKey(k *Sm2PublicKey) ([]byte, error)
}
