/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

/*
gm-plugins包提供国密算法实现接口，包含SM2、SM3、SM4等，并且提供了基于纯Go的国密算法默认实现。

下面是一个使用gm-plugins包的示例：

	package main

	import (
		"fmt"
		gmplugins "github.com/zhigui-projects/gm-plugins"
	)

	func main() {
		suite := gmplugins.GetSmCryptoSuite()
		msg := []byte("hello")
		pri, err := suite.GenPrivateKey()
		if err != nil {
			panic(err)
		}
		sign, err := suite.Sign(pri, msg, nil)
		if err != nil {
			panic(err)
		}
		ok, err := suite.Verify(suite.PublicKey(pri), sign, msg, nil)
		fmt.Println(ok , err)
	}

Output:
	true <nil>

详情请访问： https://github.com/zhigui-projects/gm-plugins
*/
package gm_plugins
