package primitive

import (
	"crypto"
)

type Sm2Crypto interface {
	Verify(k *Sm2PublicKey, sig, msg []byte, opts crypto.SignerOpts) (bool, error)
	Sign(k *Sm2PrivateKey, msg []byte, opts crypto.SignerOpts) ([]byte, error)
	Encrypt(k *Sm2PublicKey, plaintext []byte) ([]byte, error)
	Decrypt(k *Sm2PrivateKey, ciphertext []byte, opts crypto.DecrypterOpts) ([]byte, error)
}
