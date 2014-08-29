package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
)

func main() {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}

	client := ec2.New(auth, aws.APNortheast)

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
