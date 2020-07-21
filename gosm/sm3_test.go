/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gosm

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSm3(t *testing.T) {
	s3 := &GoSm3{}
	h := s3.NewSm3()
	h.Write([]byte("hello"))
	val := h.Sum(nil)
	assert.Equal(t, hex.EncodeToString(val), "becbbfaae6548b8bf0cfcad5a27183cd1be6093b1cceccc303d9c61d0a645268")
}
