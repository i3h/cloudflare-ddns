package main

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"
)

func init() {
	init_log()
	init_env()
	init_cf_api()
	init_cf_zone_id()
}

func main() {
	rr := cloudflare.DNSRecord{Name: NAME}
	recs, err := CfApi.DNSRecords(CF_ZONE_ID, rr)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(len(recs))
	for _, r := range recs {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}

	/*
		d := time.Duration(INTERVAL) * time.Second
		c := time.Tick(d)
		for _ = range c {
			Update()
		}
	*/
}
