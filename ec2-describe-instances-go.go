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

	region := strToRegion(regionStr)
	if region == nil {
		log.Fatal(errors.New("Region " + regionStr + " is unknown."))
	}

	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}

	client := ec2.New(auth, *region)

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

func strToRegion(str string) *aws.Region {
	switch str {
	case "us-east-1":
		return &aws.USEast
	case "us-west-1":
		return &aws.USWest
	case "us-west-2":
		return &aws.USWest2
	case "eu-west-1":
		return &aws.EUWest
	case "ap-southeast-1":
		return &aws.APSoutheast
	case "ap-southeast-2":
		return &aws.APSoutheast2
	case "ap-northeast-1":
		return &aws.APNortheast
	case "sa-east-1":
		return &aws.SAEast
	}

	return nil
}
