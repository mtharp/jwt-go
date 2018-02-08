// +build !go1.4

package jwt

import (
	"crypto"
	"io"
)

// crypto.Signer does not exist in go 1.3 and the Sign() method is not
// implemented on rsa or ecdsa PrivateKey types so this path is useless
type cryptoSigner interface {
	NotImplemented() cryptoSigner
	Public() crypto.PublicKey
	Sign(io.Reader, []byte, crypto.Hash) ([]byte, error)
}
