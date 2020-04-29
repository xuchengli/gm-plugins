package gosm

import (
	"gitlab.ziggurat.cn/guomi/gm-common/sm"
	"gitlab.ziggurat.cn/guomi/gm-go/sm2"
)

func publicKeyToSm2(k *sm.Sm2PublicKey) *sm2.PublicKey {
	return &sm2.PublicKey{
		Curve: k.Curve,
		X:     k.X,
		Y:     k.Y,
	}
}

func privateKeyToSm2(k *sm.Sm2PrivateKey) *sm2.PrivateKey {
	priv := new(sm2.PrivateKey)
	priv.PublicKey.Curve = k.Sm2PublicKey.Curve
	priv.PublicKey.X, priv.PublicKey.Y = k.Sm2PublicKey.X, k.Sm2PublicKey.Y
	priv.D = k.D
	return priv
}

func sm2ToPublicKey(k *sm2.PublicKey) *sm.Sm2PublicKey {
	return &sm.Sm2PublicKey{
		Curve: k.Curve,
		X:     k.X,
		Y:     k.Y,
	}
}

func sm2ToPrivateKey(k *sm2.PrivateKey) *sm.Sm2PrivateKey {
	priv := new(sm.Sm2PrivateKey)
	priv.Sm2PublicKey.Curve = k.PublicKey.Curve
	priv.Sm2PublicKey.X, priv.Sm2PublicKey.Y = k.PublicKey.X, k.PublicKey.Y
	priv.D = k.D
	return priv
}
