package gotdv2

type Device struct {
	FazpassId       string            `json:"fazpass_id"`
	IsActive        bool              `json:"is_active"`
	Scoring         float64           `json:"scoring"`
	RiskLevel       string            `json:"risk_level"`
	TimeStamp       string            `json:"time_stamp"`
	Platform        string            `json:"platform"`
	IsRooted        bool              `json:"is_rooted"`
	IsEmulator      bool              `json:"is_emulator"`
	IsGpsSpoof      bool              `json:"is_gps_spoof"`
	IsAppTempering  bool              `json:"is_app_tempering"`
	IsVpn           bool              `json:"is_vpn"`
	IsCloneApp      bool              `json:"is_clone_app"`
	IsScreenSharing bool              `json:"is_screen_sharing"`
	IsDebug         bool              `json:"is_debug"`
	Application     string            `json:"application"`
	DeviceId        map[string]string `json:"device_id"`
	SimSerial       []string          `json:"sim_serial"`
	SimOperator     []string          `json:"sim_operator"`
	Geolocation     map[string]string `json:"geolocation"`
	ClientIp        string            `json:"client_ip"`
}
