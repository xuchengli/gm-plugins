/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gm_plugins_test

import (
	"encoding/hex"
	"fmt"
	"log"

	gmplugins "github.com/zhigui-projects/gm-plugins"
)

func Example_sm2() {
	suite := gmplugins.GetSmCryptoSuite()
	msg := []byte("hello")
	pri, err := suite.GenPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	sign, err := suite.Sign(pri, msg, nil)
	if err != nil {
		log.Fatal(err)
	}
	ok, err := suite.Verify(suite.PublicKey(pri), sign, msg, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("verify result is: ", ok)
	// Output:
	// verify result is: true
}

func Example_sm3() {
	suite := gmplugins.GetSmCryptoSuite()
	h := suite.NewSm3()
	h.Write([]byte("hello"))
	val := h.Sum(nil)
	fmt.Println(hex.EncodeToString(val))
	// Output:
	// becbbfaae6548b8bf0cfcad5a27183cd1be6093b1cceccc303d9c61d0a645268
}

func Example_sm4() {
	suite := gmplugins.GetSmCryptoSuite()
	sk := []byte("1234567890abcdef")
	block, err := suite.NewSm4Cipher(sk)
	if err != nil {
		log.Fatal(err)
	}
	d0 := make([]byte, 16)
	data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
	block.Encrypt(d0, data)
	d1 := make([]byte, 16)
	block.Decrypt(d1, d0)
	fmt.Println(d1)
	// Output:
	// [1 35 69 103 137 171 205 239 254 220 186 152 118 84 50 16]
}
