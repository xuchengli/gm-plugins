/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gosm

import (
	"testing"
)

var kd = &KeysDerive{}
var gs = &GoSm2{}

func BenchmarkSM2(t *testing.B) {
	t.ReportAllocs()
	msg := []byte("test")

	pri, err := kd.GenPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {

		sign, err := gs.Sign(pri, msg, nil)
		if err != nil {
			t.Fatal(err)
		}
		gs.Verify(kd.PublicKey(pri), sign, msg, nil)
	}
}

func TestSm2(t *testing.T) {
	msg := []byte("test")

	pri, err := kd.GenPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	sign, err := gs.Sign(pri, msg, nil)
	if err != nil {
		t.Fatal(err)
	}
	ok, _ := gs.Verify(kd.PublicKey(pri), sign, msg, nil)
	if ok != true {
		t.Fatal("Verify error")
	} else {
		t.Log("Verify ok")
	}
}
