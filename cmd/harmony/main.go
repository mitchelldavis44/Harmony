// main.go
package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/mitchelldavis44/Harmony/pkg/infrastructure"
	"github.com/mitchelldavis44/aws-provider/awsprovider"
)

type AWSInstance struct {
	Type              string `hcl:"type,label"`  // Added this line
	Name              string `hcl:"name,label"`
	InstanceType      string `hcl:"instance_type"`
	ImageID           string `hcl:"image_id"`
	SecurityGroupName string `hcl:"security_group_name"`
	KeyPairName       string `hcl:"key_pair_name"`
	SubnetId          string `hcl:"subnet_id"`
	IamInstanceProfile string `hcl:"iam_instance_profile"`
}

type Root struct {
	AWSInstances []AWSInstance `hcl:"resource,block"`
}

func main() {
	var root Root
	fmt.Println("Starting to decode file...") // New line
	err := hclsimple.DecodeFile("infrastructure.hcl", nil, &root)
	if err != nil {
		fmt.Printf("Error decoding file: %v\n", err) // Print error if decoding fails
		os.Exit(1)
	}
	fmt.Println("Successfully decoded file") // New line
	fmt.Printf("Root: %+v\n", root) // New line: print the content of root

	var infra infrastructure.Infrastructure
	infra = awsprovider.NewAWSProvider()

	for _, instance := range root.AWSInstances {
		fmt.Printf("Creating instance %s with instance type %s and image ID %s\n",
			instance.Name, instance.InstanceType, instance.ImageID)

		err := infra.CreateResource(instance.Name, instance.InstanceType, instance.ImageID, instance.SecurityGroupName, instance.KeyPairName, instance.SubnetId, instance.IamInstanceProfile)
		if err != nil {
			fmt.Printf("Error creating resource: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully created resource: %s\n", instance.Name)
	}
}
