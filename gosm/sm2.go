package gosm

import (
	"crypto"
	"crypto/rand"
	"github.com/pkg/errors"
	"gitlab.ziggurat.cn/guomi/gm-common/sm"
	"gitlab.ziggurat.cn/guomi/gm-go/sm2"
)

type GoSm2 struct{}

func (gs *GoSm2) Verify(k *sm.Sm2PublicKey, sig, msg []byte, opts crypto.SignerOpts) (bool, error) {
	pub := publicKeyToSm2(k)
	if pub.Verify(msg, sig) {
		return true, nil
	}
	return false, errors.Errorf("Failed to sm2 verify signature.")
}

func (gs *GoSm2) Sign(k *sm.Sm2PrivateKey, msg []byte, opts crypto.SignerOpts) ([]byte, error) {
	pri := privateKeyToSm2(k)
	return pri.Sign(rand.Reader, msg, opts)
}

func (gs *GoSm2) Encrypt(k *sm.Sm2PublicKey, plaintext []byte) ([]byte, error) {
	pub := publicKeyToSm2(k)
	return pub.Encrypt(plaintext)
}

func (gs *GoSm2) Decrypt(k *sm.Sm2PrivateKey, ciphertext []byte, opts crypto.DecrypterOpts) ([]byte, error) {
	pri := privateKeyToSm2(k)
	return pri.Decrypt(ciphertext)
}

type KeysDerive struct{}

func (kd *KeysDerive) ParsePKCS8UnecryptedPrivateKey(der []byte) (*sm.Sm2PrivateKey, error) {
	pri, err := sm2.ParsePKCS8UnecryptedPrivateKey(der)
	if err != nil {
		return nil, err
	}
	return sm2ToPrivateKey(pri), nil
}
