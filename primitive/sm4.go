/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package primitive

import "crypto/cipher"

type Sm4Crypro interface {
	// 根据给定密钥创建sm4分组加密实例，返回分组加密功能函数接口
	NewSm4Cipher(key []byte) (cipher.Block, error)
}
