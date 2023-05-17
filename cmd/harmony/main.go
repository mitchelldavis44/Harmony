// main.go
package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/mitchelldavis44/Harmony/pkg/infrastructure"
	"github.com/mitchelldavis44/aws-provider/awsprovider"
)

type AWSInstance struct {
	Name         string `hcl:"name,label"`
	InstanceType string `hcl:"instance_type"`
	ImageID      string `hcl:"image_id"`
}

type Root struct {
	AWSInstances []AWSInstance `hcl:"resource,block"`
}

func main() {
	var root Root
	hclsimple.DecodeFile("infrastructure.harmony", nil, &root)

	var infra infrastructure.Infrastructure
	infra = awsprovider.NewAWSProvider()

	for _, instance := range root.AWSInstances {
		fmt.Printf("Creating instance %s with instance type %s and image ID %s\n",
			instance.Name, instance.InstanceType, instance.ImageID)

		err := infra.CreateResource(instance.Name, instance.InstanceType, instance.ImageID)
		if err != nil {
			fmt.Printf("Error creating resource: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created resource: %s\n", instance.Name)
	}
}
