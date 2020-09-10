package main

import (
	"os"
	"strconv"

	"github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"
)

const (
	IP_API_DEFAULT string = "https://ifconfig.is"
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
	TYPE       string
	CfApi      *cloudflare.API
)

func init_env() {
	IP_API = os.Getenv("IP_API")
	if IP_API == "" {
		IP_API = IP_API_DEFAULT
	}
	CF_EMAIL = os.Getenv("CF_EMAIL")
	CF_API_KEY = os.Getenv("CF_API_KEY")
	CF_ZONE = os.Getenv("CF_ZONE")
	LOG_FILE = os.Getenv("LOG_FILE")
	INTERVAL, _ = strconv.Atoi(os.Getenv("INTERVAL"))
	NAME = os.Getenv("NAME")
	TYPE = os.Getenv("TYPE")
	if TYPE == "" {
		TYPE = "A"
	}
}

func init_cf_api() {
	var err error
	CfApi, err = cloudflare.New(CF_API_KEY, CF_EMAIL)
	if err != nil {
		log.Error(err)
	}
}

func init_cf_zone_id() {
	var err error
	CF_ZONE_ID, err = CfApi.ZoneIDByName(CF_ZONE)
	if err != nil {
		log.Fatal(err)
	}
}
