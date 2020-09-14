package main

func Update() {
	recs := GetRecords()
	ip := GetIP()
	var count int
	for _, r := range recs {
		if r.Type == TYPE {
			count++
		}
	}
	if count == 0 {
		CreateRecord(ip)
		log.Info("Update record | ", "Type: ", TYPE, " | IP: ", ip)
	} else if count > 1 {
		ClearRecords()
		CreateRecord(ip)
		log.Info("Update record | ", "Type: ", TYPE, " | IP: ", ip)
	} else if count == 1 {
		for _, r := range recs {
			if r.Type == TYPE {
				if r.Content != ip {
					ClearRecords()
					CreateRecord(ip)
					log.Info("Update record | ", "Type: ", TYPE, " | IP: ", ip)
				} else {
					log.Info("Already up-to-date | ", "Type: ", TYPE, " | IP: ", ip)
				}
			}
		}
	}
}
