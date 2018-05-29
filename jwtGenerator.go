package main

import (
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func getJwt(signer jose.Signer) () {
	builder := jwt.Signed(signer)

	return
}
