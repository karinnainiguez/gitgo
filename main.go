package main

import (
	flag "github.com/ogier/pflag"
)

var (
	user string
)

func main() {
	// Parse Flags
	flag.Parse()
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
