package shared

type GatewayStatus struct {
	Time              string   `json:"time,omitempty"`
	Latitude          *float64 `json:"latitude,omitempty"`
	Longitude         *float64 `json:"longitude,omitempty"`
	Altitude          *float64 `json:"altitude,omitempty"`
	RxCount           *int     `json:"rxCount,omitempty"`
	RxOk              *int     `json:"rxOk,omitempty"`
	RxForwarded       *int     `json:"rxForwarded,omitempty"`
	AckRatio          *float64 `json:"ackRatio,omitempty"`
	DatagramsReceived *int     `json:"datagramsReceived,omitempty"`
	DatagramsSent     *int     `json:"datagramsSent,omitempty"`
}

type RxPacket struct {
	Time string `json:"time,omitempty"`
	Data string `json:"data,omitempty"`
}
