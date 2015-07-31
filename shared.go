package shared

type GatewayStatus struct {
	Time              string   `json:"time"`
	Latitude          *float64 `json:"latitude"`
	Longitude         *float64 `json:"longitude"`
	Altitude          *float64 `json:"altitude"`
	RxCount           *int     `json:"rxCount"`
	RxOk              *int     `json:"rxOk"`
	RxForwarded       *int     `json:"rxForwarded"`
	AckRatio          *float64 `json:"ackRatio"`
	DatagramsReceived *int     `json:"datagramsReceived"`
	DatagramsSent     *int     `json:"datagramsSent"`
}

type RxPacket struct {
	Time string `json:"time"`
	Data string `json:"data"`
}
