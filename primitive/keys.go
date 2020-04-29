package primitive

import (
	"crypto/elliptic"
	"math/big"
)

type Sm2PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

type Sm2PrivateKey struct {
	Sm2PublicKey
	D *big.Int
}

type KeysGenerator interface {
	GenPrivateKey() (*Sm2PrivateKey, error)
	PublicKey(k *Sm2PrivateKey) *Sm2PublicKey
	ParsePKCS8UnecryptedPrivateKey(der []byte) (*Sm2PrivateKey, error)
}
