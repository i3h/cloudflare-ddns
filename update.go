package main

import "fmt"

func Update() {
	recs := GetRecords()
	ip := GetIP()
	var count int
	for _, r := range recs {
		if r.Type == TYPE {
			count++
		}
	}
	fmt.Println("Count: ", count)
	if count == 0 {
		CreateRecord(ip)
	} else if count > 1 {
		ClearRecords()
		CreateRecord(ip)
	} else if count == 1 {
		for _, r := range recs {
			if r.Type == TYPE {
				if r.Content != ip {
					ClearRecords()
					CreateRecord(ip)
				}
			}
		}
	}
}
