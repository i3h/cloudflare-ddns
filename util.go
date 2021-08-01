package main

import (
	"encoding/json"
	"fmt"
)

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}
