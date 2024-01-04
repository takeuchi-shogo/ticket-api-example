package auth

import (
	_ "embed"
)

//go:embed cert/secret.pem
var RawPrivateKey []byte

//go:embed cert/public.pem
var RawPublicKey []byte
