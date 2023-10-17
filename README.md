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
	meta, err = td.Extract("YOUR_META")
	if err != nil {
		log.Println(err)
	}
	log.Println(meta.FazzpassId)
```
What Inside ?
```go
type Meta struct {
	FazpassId       string            
	IsActive        bool              
	Scoring         float64           
	RiskLevel       string            
	TimeStamp       string            
	Platform        string            
	IsRooted        bool              
	IsEmulator      bool              
	IsGpsSpoof      bool              
	IsAppTempering  bool              
	IsVpn           bool              
	IsCloneApp      bool              
	IsScreenSharing bool              
	IsDebug         bool              
	Application     string            
	Device          Device            
	SimSerial       []string          
	SimOperator     []string          
	Geolocation     map[string]string 
	ClientIp        string            
	LinkedDevices   []Device          
}

type Device struct {
	Name      string 
	OSVersion string 
	Series    string 
	CPU       string 
	ID        string 
}
```
