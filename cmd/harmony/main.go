// main.go
package main

import (
	"fmt"
	"os"

	"github.com/mitchelldavis44/Harmony/pkg/infrastructure"
	"github.com/mitchelldavis44/aws-provider/awsprovider"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./harmony <action> <resource_name>")
		os.Exit(1)
	}

	action := os.Args[1]
	resourceName := os.Args[2]

	var infra infrastructure.Infrastructure
	// Here you could decide which provider to use based on command line flags,
	// configuration files, etc. For simplicity, we're always using the AWSProvider.
	infra = awsprovider.NewAWSProvider()

	switch action {
	case "create":
		err := infra.CreateResource(resourceName)
		if err != nil {
			fmt.Printf("Error creating resource: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created resource: %s\n", resourceName)
	case "delete":
		err := infra.DeleteResource(resourceName)
		if err != nil {
			fmt.Printf("Error deleting resource: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Deleted resource: %s\n", resourceName)
	default:
		fmt.Println("Unknown action:", action)
		os.Exit(1)
	}
}
