package main

import (
	"github.com/cloudflare/cloudflare-go"
)

func GetNewRecord(rr cloudflare.DNSRecord) cloudflare.DNSRecord {
	ip := GetIP()
	rr.Content = ip
	return rr
}

func ClearRecord() {
}
