package main

import (
	"log"
	"time"

	"github.com/hashicorp/vault-client-go"
)



func main() {
	vault_addr := "http://127.0.0.1:8200"
	time_out := 30
	token := "123"
	client, err := vault.New(
		vault.WithAddress(vault_addr),
		vault.WithRequestTimeout(time.Second * time.Duration(time_out)),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err:= client.SetToken(token); err != nil {
		log.Fatal(err)
	}

}