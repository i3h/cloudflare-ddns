package main

import (
	"fmt"
)

func init() {
	init_log()
	init_env()
	init_cf_api()
	init_cf_zone_id()
}

func main() {
	/*
		rr := cloudflare.DNSRecord{Name: NAME}
		recs, err := CfApi.DNSRecords(CF_ZONE_ID, rr)
		if err != nil {
			log.Error(err)
		}

		fmt.Println(len(recs))
		for _, r := range recs {
			fmt.Printf("%s: %s\n", r.Name, r.Content)
		}
	*/

	fmt.Println(GetIP())

	/*
		d := time.Duration(INTERVAL) * time.Second
		c := time.Tick(d)
		for _ = range c {
			Update()
		}
	*/
}
