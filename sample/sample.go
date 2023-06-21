package main

import (
	gotdv2 "github.com/fazpass-sdk/go-trusted-device-v2"
	"log"
)

func main() {
	td, err := gotdv2.Initialize("./sample/private.key", "http://localhost:8080", "YOUR_KEY")
	if err != nil {
		log.Println(err)
	}
	deviceChan, errChan := td.CheckAsyncDevice("picId", "meta", "appId")
	select {
	case device := <-deviceChan:
		log.Println(device)
		// handle device
	case err := <-errChan:
		log.Println(err)
		// handle error
	}
}
