/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gosm

import (
		"testing"

		"github.com/stretchr/testify/assert"
)

func TestSm4(t *testing.T) {
		s4 := &GoSm4{}
		sk := []byte("1234567890abcdef")
		block, err := s4.NewSm4Cipher(sk)
		assert.NoError(t, err)
		d0 := make([]byte, 16)
		data := []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xfe, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x10}
		block.Encrypt(d0, data)
		d1 := make([]byte, 16)
		block.Decrypt(d1, d0)
		assert.Equal(t, data, d1)
}
