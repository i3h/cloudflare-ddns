package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type MyDialer net.Dialer

func (d *MyDialer) Dial(network, address string) (net.Conn, error) {
	dd := &net.Dialer{}
	if TYPE == "AAAA" {
		return dd.Dial("tcp6", address)
	} else {
		return dd.Dial("tcp4", address)
	}
}

func GetIP() string {
	d := &MyDialer{
		Timeout: 5 * time.Second,
	}
	tr := &http.Transport{
		Dial: d.Dial,
	}
	client := &http.Client{
		Transport: tr,
	}

	req, _ := http.NewRequest("GET", GET_IP_API, nil)

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
