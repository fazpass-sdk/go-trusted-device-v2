package gotdv2

type Meta struct {
	Challenge       string            `json:"challenge"`
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
	Biometric       Biometric         `json:"biometric"`
	IsCloneApp      bool              `json:"is_clone_app"`
	IsScreenSharing bool              `json:"is_screen_sharing"`
	IsDebug         bool              `json:"is_debug"`
	Application     string            `json:"application"`
	Device          Device            `json:"device_id"`
	SimSerial       []string          `json:"sim_serial"`
	SimOperator     []string          `json:"sim_operator"`
	Geolocation     map[string]string `json:"geolocation"`
	ClientIp        string            `json:"client_ip"`
	LinkedDevices   []Device          `json:"notifiable_devices"`
	IsNotifiable    bool              `json:"is_notifiable"`
}

type Device struct {
	Name      string `json:"name"`
	OSVersion string `json:"os_version"`
	Series    string `json:"series"`
	CPU       string `json:"cpu"`
	ID        string `json:"id"`
}

type Biometric struct {
	Level      string `json:"level"`
	IsChanging bool   `json:"is_changing"`
}
