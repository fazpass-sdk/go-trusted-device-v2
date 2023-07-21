[![codecov](https://codecov.io/gh/fazpass-sdk/go-trusted-device-v2/branch/main/graph/badge.svg?token=CHFJQBLO7T)](https://codecov.io/gh/fazpass-sdk/go-trusted-device-v2)
# go-trusted-device-v2

## Instalation
```go
go get github.com/fazpass-sdk/go-trusted-device-v2
```

## Usage
```go
td, err := gotdv2.Initialize("./sample/private.key")
	if err != nil {
		log.Println(err)
	}
	device, err = td.Extract("YOUR_META")
	if err != nil {
		log.Println(err)
	}
	log.Println(device.FazzpassId)
```