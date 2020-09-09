package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
)

var (
	IP_API     string
	CF_EMAIL   string
	CF_API_KEY string
	CF_ZONE    string
	CF_ZONE_ID string
	LOG_FILE   string
	INTERVAL   int
	NAME       string
)

func init() {
	// Get env
	IP_API = os.Getenv("IP_API")
	if IP_API == "" {
		IP_API = "https://ifconfig.is"
	}
	CF_EMAIL = os.Getenv("CF_EMAIL")
	CF_API_KEY = os.Getenv("CF_API_KEY")
	CF_ZONE = os.Getenv("CF_ZONE")
	LOG_FILE = os.Getenv("LOG_FILE")
	INTERVAL, _ = strconv.Atoi(os.Getenv("INTERVAL"))
	NAME = os.Getenv("NAME")
}

func main() {
	// Print env
	fmt.Println("Environment variables:")
	fmt.Println("CF_EMAIL:    ", CF_EMAIL)
	fmt.Println("CF_API_KEY:  ", CF_API_KEY)
	fmt.Println("CF_ZONE:  ", CF_ZONE)
	fmt.Println("LOG_FILE: ", LOG_FILE)
	fmt.Println("INTERVAL: ", INTERVAL)
	fmt.Println("NAME:   ", NAME)
	fmt.Println()

	// Construct a new API object
	api, err := cloudflare.New(CF_API_KEY, CF_EMAIL)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch user details on the account
	u, err := api.UserDetails()
	if err != nil {
		log.Fatal(err)
	}
	// Print user details
	PPrint(u)

	// Fetch the zone ID
	CF_ZONE_ID, err := api.ZoneIDByName(CF_ZONE)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(CF_ZONE_ID)

	rr := cloudflare.DNSRecord{Name: NAME, Type: "A"}
	recs, err := api.DNSRecords(CF_ZONE_ID, rr)
	if err != nil {
		log.Fatal(err)
	}
	PPrint(recs)
	fmt.Println()

	ip := GetIP()
	fmt.Println(ip)
	new_rr := cloudflare.DNSRecord{Name: NAME, Type: "A", Content: ip}
	PPrint(new_rr)
	err = api.UpdateDNSRecord(CF_ZONE_ID, recs[0].ID, new_rr)
	if err != nil {
		fmt.Println(err)
	}
}

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
