package main

import (
	"github.com/kataras/iris"
	jose "gopkg.in/square/go-jose.v2"
	"crypto/rsa"
	"crypto/rand"
	"log"
)

func getWellKnown(ctx iris.Context) {
	ctx.Writef("well-known")
}

func main() {
	app := iris.New()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Generating random key: %v", err)
	}

	key := jose.SigningKey{Algorithm: jose.RS256, Key: privateKey}
	signerOpts := jose.SignerOptions{}
	signerOpts.WithType("JWT")

	rsaSigner, err := jose.NewSigner(key, &signerOpts)
	if err != nil {
		log.Fatalf("failed to create signer: %+v", err)
	}

	/* Unauthenticated routes go here */
	app.Get("/.well-known/jwks", getWellKnown)
	app.Get("/connect/token", getToken(rsaSigner))

	/* Authenticate */

	/* JWT Authenticated routes here */
	app.Run(iris.Addr(":8008"))
}
