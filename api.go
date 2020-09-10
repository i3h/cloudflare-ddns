package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
)

func GetIP() string {
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest("GET", IP_API, nil)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	content := strings.TrimSuffix(string(body), "\n")

	return content
}

func GetNewRecord(rr cloudflare.DNSRecord) cloudflare.DNSRecord {
	ip := GetIP()
	rr.Content = ip
	return rr
}

func ClearRecord() {
}
