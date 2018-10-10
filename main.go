package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	user string
)

func main() {
	// Parse Flags
	flag.Parse()

	// if user does not supply flags:
	if flag.NFlag() == 0 {
		printUsage()
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)

	for _, u := range users {
		result := getUsers(u)
		fmt.Printf("Username: %v\n", result.Login)
		fmt.Printf("Name: %v\n", result.Name)
		fmt.Printf("Email: %v\n", result.Email)
		fmt.Printf("Bio: %v\n", result.Bio)
		fmt.Printf("")

	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
