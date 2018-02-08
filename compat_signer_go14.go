// +build go1.4

package jwt

import "crypto"

type cryptoSigner interface {
	crypto.Signer
}
