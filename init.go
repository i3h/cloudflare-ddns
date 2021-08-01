package main

import (
	"os"
	"strconv"

	"github.com/cloudflare/cloudflare-go"
)

var (
	CF_API_EMAIL string
	CF_API_KEY   string
	CF_ZONE      string
	CF_ZONE_ID   string
	LOG_FILE     string
	INTERVAL     int
	NAME         string
	TYPE         string
	API          *cloudflare.API
	GET_IP_API   string
)

func init_env() {
	CF_API_EMAIL = os.Getenv("CF_API_EMAIL")
	CF_API_KEY = os.Getenv("CF_API_KEY")
	CF_ZONE = os.Getenv("CF_ZONE")
	LOG_FILE = os.Getenv("LOG_FILE")
	INTERVAL, _ = strconv.Atoi(os.Getenv("INTERVAL"))
	NAME = os.Getenv("NAME")

	TYPE = os.Getenv("TYPE")
	if TYPE == "" {
		TYPE = "A"
	}

	GET_IP_API = os.Getenv("GET_IP_API")
	if GET_IP_API == "" {
		GET_IP_API = "https://ifconfig.is"
	}
}

func init_api() {
	var err error
	API, err = cloudflare.New(CF_API_KEY, CF_API_EMAIL)
	if err != nil {
		log.Error(err)
	}
}

func init_cf_zone_id() {
	var err error
	CF_ZONE_ID, err = API.ZoneIDByName(CF_ZONE)
	if err != nil {
		log.Fatal(err)
	}
}
