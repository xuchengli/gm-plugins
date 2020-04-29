package gosm

import (
	"crypto"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/pkg/errors"
	"gitlab.ziggurat.cn/guomi/gm-go/sm2"
	"gitlab.ziggurat.cn/guomi/gm-plugins/primitive"
)

type GoSm2 struct{}

func (gs *GoSm2) Verify(k *primitive.Sm2PublicKey, sig, msg []byte, opts crypto.SignerOpts) (bool, error) {
	pub := publicKeyToSm2(k)
	if pub.Verify(msg, sig) {
		return true, nil
	}
	return false, errors.Errorf("Failed to sm2 verify signature.")
}

func (gs *GoSm2) Sign(k *primitive.Sm2PrivateKey, msg []byte, opts crypto.SignerOpts) ([]byte, error) {
	pri := privateKeyToSm2(k)
	return pri.Sign(rand.Reader, msg, opts)
}

func (gs *GoSm2) Encrypt(k *primitive.Sm2PublicKey, plaintext []byte) ([]byte, error) {
	pub := publicKeyToSm2(k)
	return pub.Encrypt(plaintext)
}

func (gs *GoSm2) Decrypt(k *primitive.Sm2PrivateKey, ciphertext []byte, opts crypto.DecrypterOpts) ([]byte, error) {
	pri := privateKeyToSm2(k)
	return pri.Decrypt(ciphertext)
}

func (gs *GoSm2) Sm2P256Curve() elliptic.Curve {
	return sm2.P256Sm2()
}

type KeysDerive struct{}

func (kd *KeysDerive) GenPrivateKey() (*primitive.Sm2PrivateKey, error) {
	pri, err := sm2.GenerateKey()
	if err != nil {
		return nil, err
	}
	return sm2ToPrivateKey(pri), nil
}

func (kd *KeysDerive) PublicKey(k *primitive.Sm2PrivateKey) *primitive.Sm2PublicKey {
	return &k.Sm2PublicKey
}

func (kd *KeysDerive) ParsePKCS8UnecryptedPrivateKey(der []byte) (*primitive.Sm2PrivateKey, error) {
	pri, err := sm2.ParsePKCS8UnecryptedPrivateKey(der)
	if err != nil {
		return nil, err
	}
	return sm2ToPrivateKey(pri), nil
}
