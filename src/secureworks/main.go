package main

import (
	"fmt"
	"secureworks/constants"
	"secureworks/maxmind"
)

//go:generate /root/generate.sh
func main() {
	fmt.Printf("IP_DB: '%s'\n", constants.IP_DB)
	fmt.Printf("SQL_DB: '%s'\n", constants.SQL_DB)
	maxmind.NewLocationFinder(constants.IP_DB)
}
