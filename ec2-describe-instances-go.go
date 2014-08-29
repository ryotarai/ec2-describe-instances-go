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

	region, err := strToRegion(regionStr)
	if err == nil {
		log.Fatal(err)
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

func strToRegion(str string) (*aws.Region, error) {
	switch str {
	case "us-east-1":
		return &aws.USEast, nil
	case "us-west-1":
		return &aws.USWest, nil
	case "us-west-2":
		return &aws.USWest2, nil
	case "eu-west-1":
		return &aws.EUWest, nil
	case "ap-southeast-1":
		return &aws.APSoutheast, nil
	case "ap-southeast-2":
		return &aws.APSoutheast2, nil
	case "ap-northeast-1":
		return &aws.APNortheast, nil
	case "sa-east-1":
		return &aws.SAEast, nil
	default:
		return nil, errors.New("Region " + str + " is unknown.")
	}
}
