package gotdv2

type TrustedDevice interface {
	CheckDevice(picId string, meta string, appId string) (*Device, error)
	CheckAsyncDevice(picId string, meta string, appId string) (<-chan *Device, <-chan error)
	EnrollDevice(picId string, meta string, appId string) (*Device, error)
	EnrollAsyncDevice(picId string, meta string, appId string) (<-chan *Device, <-chan error)
	ValidateDevice(fazpassId string, meta string, appId string) (*Device, error)
	ValidateAsyncDevice(fazpassId string, meta string, appId string) (<-chan *Device, <-chan error)
	RemoveDevice(fazpassId string, meta string, appId string) (*Device, error)
	RemoveAsyncDevice(fazpassId string, meta string, appId string) (<-chan *Device, <-chan error)
}
