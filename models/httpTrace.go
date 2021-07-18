package models

import "time"

// HTTPTraceResult is the set of useful information to trace the http request
type HTTPTraceResult struct {
	DNSStart  time.Time     `json:"-"`
	DNSDone   time.Time     `json:"-"`
	DNSLookup time.Duration `json:"dnsLookupTimeTaken"`

	TCPStart      time.Time     `json:"-"`
	TCPDone       time.Time     `json:"-"`
	TCPConnection time.Duration `json:"tcpConnectionTimeTaken"`

	Connect     time.Duration `json:"connectTimeTaken"`
	PreTransfer time.Duration `json:"preTransferTimeTaken"`

	IsTLS        bool          `json:"isTLSEnabled"`
	TLSStart     time.Time     `json:"-"`
	TLSDone      time.Time     `json:"-"`
	TLSHandshake time.Duration `json:"tlsHandshakeTimeTaken"`

	ServerStart      time.Time     `json:"-"`
	ServerDone       time.Time     `json:"-"`
	ServerProcessing time.Duration `json:"serverProcessingTimeTaken"`

	IsReused bool `json:"isConnectionReused"`

	Done bool `json:"-"`
}
