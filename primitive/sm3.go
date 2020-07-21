/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package primitive

import "hash"

type Sm3Crypro interface {
	// 创建sm3实例，返回hash函数接口
	NewSm3() hash.Hash
}
