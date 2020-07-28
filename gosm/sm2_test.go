/*
Copyright Zhigui.com. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gosm

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhigui-projects/gm-plugins/primitive"
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
	ok, _ := gs.Verify(pri.Public(), sign, msg, nil)
	if ok != true {
		t.Fatal("Verify error")
	} else {
		t.Log("Verify ok")
	}
}

func TestSM2Crypt(t *testing.T) {
	sk, err := kd.GenPrivateKey()
	assert.NoError(t, err)

	pk := sk.Public()

	msg := []byte("testing sm2 crypt")

	ciphertext, err := gs.Encrypt(pk, msg)
	assert.NoError(t, err)

	plaintext, err := gs.Decrypt(sk, ciphertext, nil)
	assert.NoError(t, err)
	assert.Equal(t, msg, plaintext)
}

func TestPKCS8(t *testing.T) {
	sk, err := kd.GenPrivateKey()
	assert.NoError(t, err)

	der, err := kd.MarshalPKCS8PrivateKey(sk, []byte("password"))
	assert.NoError(t, err)

	pri, err := kd.ParsePKCS8PrivateKey(der, []byte("password"))
	assert.NoError(t, err)

	if reflect.TypeOf(pri) != reflect.TypeOf(&primitive.Sm2PrivateKey{}) {
		t.Errorf(" decoded PKCS#8 returned unexpected key type: %T", pri)
	}
}
func TestPKIX(t *testing.T) {
	sk, err := kd.GenPrivateKey()
	assert.NoError(t, err)

	pk := sk.Public()

	der, err := kd.MarshalPKIXPublicKey(pk)
	assert.NoError(t, err)

	pub, err := kd.ParsePKIXPublicKey(der)
	assert.NoError(t, err)

	if reflect.TypeOf(pub) != reflect.TypeOf(&primitive.Sm2PublicKey{}) {
		t.Errorf(" decoded PKIX returned unexpected key type: %T", pub)
	}
}
