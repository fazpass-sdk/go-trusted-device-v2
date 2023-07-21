package main

import (
	gotdv2 "github.com/fazpass-sdk/go-trusted-device-v2"
	"log"
)

func main() {
	td, err := gotdv2.Initialize("./sample/private.key")
	if err != nil {
		log.Println(err)
	}
	_, err = td.Extract("YOUR_META")
	if err != nil {
		log.Println(err)
	}
}
