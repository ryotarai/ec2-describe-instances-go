package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

func main() {
	var regionStr string
	flag.StringVar(&regionStr, "r", "us-east-1", "Region")
	flag.Parse()

	region := aws.Regions[regionStr]
	if region.Name == "" {
		log.Fatal(errors.New("Region " + regionStr + " is unknown."))
	}

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}

	client := ec2.New(auth, region)

	resp, err := client.Instances(nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	instances := []ec2.Instance{}
	for _, reservation := range resp.Reservations {
		instances = append(instances, reservation.Instances...)
	}

	b, err := json.Marshal(instances)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}

