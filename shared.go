package shared

import (
	"time"
)

type GatewayStatus struct {
	Eui               string    `json:"eui,omitempty"`
	Time              time.Time `json:"time,omitempty"`
	Latitude          *float64  `json:"latitude,omitempty"`
	Longitude         *float64  `json:"longitude,omitempty"`
	Altitude          *float64  `json:"altitude,omitempty"`
	RxCount           *int      `json:"rxCount,omitempty"`
	RxOk              *int      `json:"rxOk,omitempty"`
	RxForwarded       *int      `json:"rxForwarded,omitempty"`
	AckRatio          *float64  `json:"ackRatio,omitempty"`
	DatagramsReceived *int      `json:"datagramsReceived,omitempty"`
	DatagramsSent     *int      `json:"datagramsSent,omitempty"`
}

type RxPacket struct {
	GatewayEui string    `json:"gatewayEui,omitempty"`
	NodeEui    string    `json:"nodeEui,omitempty"`
	Time       time.Time `json:"time,omitempty"`
	Data       string    `json:"data,omitempty"`
}
