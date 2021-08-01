package main

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
)

func Update() {
	ctx := context.Background()

	ip := GetIP()

	recs, err := API.DNSRecords(ctx, CF_ZONE_ID, cloudflare.DNSRecord{
		Name: NAME,
		Type: TYPE,
	})
	if err != nil {
		fmt.Println(err)
	}

	if len(recs) == 0 {
		// Not exist
		// Do not create record
		log.Info("This record does not exist. Update failed.")
	} else if len(recs) > 1 {
		// Too many records
		// Delete all and create new one
		for _, r := range recs {
			err := API.DeleteDNSRecord(ctx, CF_ZONE_ID, r.ID)
			if err != nil {
				fmt.Println(err)
			}
		}
		_, err := API.CreateDNSRecord(ctx, CF_ZONE_ID, cloudflare.DNSRecord{
			Name:    NAME,
			Type:    TYPE,
			Content: ip,
		})
		if err != nil {
			fmt.Println(err)
		} else {
			log.Info("Updated to ", ip)
		}
	} else {
		// Has one record
		// Update if necessary
		if ip == recs[0].Content {
			log.Info("Already updated!")
		} else {
			err := API.UpdateDNSRecord(ctx, CF_ZONE_ID, recs[0].ID, cloudflare.DNSRecord{
				Name:    NAME,
				Type:    TYPE,
				Content: ip,
			})
			if err != nil {
				log.Error(err)
			} else {
				log.Info("Updated to ", ip)
			}
		}
	}
}
