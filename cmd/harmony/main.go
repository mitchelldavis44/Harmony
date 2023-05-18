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
	SecurityGroupId string `hcl:"security_group_id"`
	KeyPairName       string `hcl:"key_pair_name"`
	SubnetId          string `hcl:"subnet_id"`
	IamInstanceProfile string `hcl:"iam_instance_profile"`
	VpcId             string `hcl:"vpc_id"`
	Tags map[string]string `hcl:"tags"`
}

type Root struct {
	AWSInstances []AWSInstance `hcl:"resource,block"`
}

func main() {
    var filename string
    if len(os.Args) > 2 {
        filename = os.Args[2]
    } else {
        filename = "infrastructure.hcl"
    }

    var root Root
    fmt.Println("Starting to decode file...")
    err := hclsimple.DecodeFile(filename, nil, &root)  // use the filename variable here
    if err != nil {
        fmt.Printf("Error decoding file: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Successfully decoded file")
    fmt.Printf("Root: %+v\n", root)

    var infra infrastructure.Infrastructure
    infra = awsprovider.NewAWSProvider()

    operation := os.Args[1] // this should be either "create" or "delete"

    for _, instance := range root.AWSInstances {
        if operation == "create" {
            fmt.Printf("Creating instance %s with instance type %s and image ID %s\n",
                instance.Name, instance.InstanceType, instance.ImageID)
            err := infra.CreateResource(instance.Name, instance.InstanceType, instance.ImageID, instance.SecurityGroupId, instance.KeyPairName, instance.SubnetId, instance.IamInstanceProfile, instance.VpcId)
            if err != nil {
                fmt.Printf("Error creating resource: %v\n", err)
                os.Exit(1)
            }
            fmt.Printf("Successfully created resource: %s\n", instance.Name)
        } else if operation == "delete" {
            fmt.Printf("Deleting instance %s\n", instance.Name)
            err := infra.DeleteResource(instance.Name)
            if err != nil {
                fmt.Printf("Error deleting resource: %v\n", err)
                os.Exit(1)
            }
            fmt.Printf("Successfully deleted resource: %s\n", instance.Name)
        }
    }
}
