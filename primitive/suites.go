/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package primitive

// 国密套件所有功能接口
type Context interface {
	KeysGenerator
	Sm2Crypto
	Sm3Crypro
	Sm4Crypro
}
