/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package primitive

import (
	"crypto"
	"crypto/elliptic"
)

type Sm2Crypto interface {
	// 签名验证，验证成功返回true，error为nil; 验证失败返回false，并返回错误信息
	Verify(k *Sm2PublicKey, sig, msg []byte, opts crypto.SignerOpts) (bool, error)
	// 数据签名, 成功返回数字签名，error为nil; 失败返回nil，并返回错误信息
	Sign(k *Sm2PrivateKey, msg []byte, opts crypto.SignerOpts) ([]byte, error)
	// 数据加密, 成功返回密文，error为nil; 失败返回nil，并返回错误信息
	Encrypt(k *Sm2PublicKey, plaintext []byte) ([]byte, error)
	// 数据解密，成功返回明文，error为nil; 失败返回nil，并返回错误信息
	Decrypt(k *Sm2PrivateKey, ciphertext []byte, opts crypto.DecrypterOpts) ([]byte, error)
	// 返回sm2算法使用的椭圆曲线
	Sm2P256Curve() elliptic.Curve
}
