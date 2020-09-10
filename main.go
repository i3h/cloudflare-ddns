package main

func init() {
	init_log()
	init_env()
	init_cf_api()
	init_cf_zone_id()
}

func main() {
	Update()

	/*
		d := time.Duration(INTERVAL) * time.Second
		c := time.Tick(d)
		for _ = range c {
			Update()
		}
	*/
}
