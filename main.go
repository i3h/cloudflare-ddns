package main

import "time"

func init() {
	init_env()
	init_log()
	init_api()
	init_cf_zone_id()
}

func main() {
	d := time.Duration(INTERVAL) * time.Second
	c := time.Tick(d)
	for _ = range c {
		Update()
	}
}
