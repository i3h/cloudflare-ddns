package main

import (
	"github.com/cloudflare/cloudflare-go"
)

func GetRecords() []cloudflare.DNSRecord {
	rr := cloudflare.DNSRecord{Name: NAME}
	recs, err := CfApi.DNSRecords(CF_ZONE_ID, rr)
	if err != nil {
		log.Error(err)
	}
	return recs
}

func GetNewRecord(rr cloudflare.DNSRecord) cloudflare.DNSRecord {
	ip := GetIP()
	rr.Content = ip
	return rr
}

func CreateRecord(ip string) {
	rr := cloudflare.DNSRecord{Name: NAME, Type: TYPE, Content: ip}
	_, err := CfApi.CreateDNSRecord(CF_ZONE_ID, rr)
	if err != nil {
		log.Error(err)
	}
}

func ClearRecords() {
	recs := GetRecords()
	for _, r := range recs {
		if r.Type == TYPE {
			err := CfApi.DeleteDNSRecord(CF_ZONE_ID, r.ID)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
