package main

import (
	"encoding/json"
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

	region, ok := aws.Regions[regionStr]
	if !ok {
		log.Fatal(fmt.Errorf("Region %s is unknown.", regionStr))
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

