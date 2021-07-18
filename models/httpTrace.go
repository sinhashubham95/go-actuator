package models

import "time"

// HTTPTraceResult is the set of useful information to trace the http request
type HTTPTraceResult struct {
	Host string `json:"host"`

	DNSStart  time.Time     `json:"-"`
	DNSDone   time.Time     `json:"-"`
	DNSLookup time.Duration `json:"dnsLookupTimeTakenInNanos"`

	TCPStart      time.Time     `json:"-"`
	TCPDone       time.Time     `json:"-"`
	TCPConnection time.Duration `json:"tcpConnectionTimeTakenInNanos"`

	Connect     time.Duration `json:"connectTimeTakenInNanos"`
	PreTransfer time.Duration `json:"preTransferTimeTakenInNanos"`

	IsTLS        bool          `json:"isTLSEnabled"`
	TLSStart     time.Time     `json:"-"`
	TLSDone      time.Time     `json:"-"`
	TLSHandshake time.Duration `json:"tlsHandshakeTimeTakenInNanos"`

	ServerStart      time.Time     `json:"-"`
	ServerDone       time.Time     `json:"-"`
	ServerProcessing time.Duration `json:"serverProcessingTimeTakenInNanos"`

	IsReused bool `json:"isConnectionReused"`

	Done bool `json:"-"`
}
