package shared

import (
	"time"
)

type GatewayStatus struct {
	Eui               string    `json:"eui"`
	Time              time.Time `json:"time"`
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
	GatewayEui string    `json:"gatewayEui"`
	NodeEui    string    `json:"nodeEui"`
	Time       time.Time `json:"time"`
	RawData    string    `json:"rawData"`
	Data       string    `json:"data,omitempty"`
}

type ConsumerQueues struct {
	GatewayStatuses chan *GatewayStatus
	RxPackets       chan *RxPacket
}
