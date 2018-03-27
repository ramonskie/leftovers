package elbv2

import (
	"fmt"

	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
)

type LoadBalancer struct {
	client     loadBalancersClient
	name       *string
	arn        *string
	identifier string
	rtype      string
}

func NewLoadBalancer(client loadBalancersClient, name, arn *string) LoadBalancer {
	return LoadBalancer{
		client:     client,
		name:       name,
		arn:        arn,
		identifier: *name,
		rtype:      "Elastic Load Balancer V2",
	}
}

func (l LoadBalancer) Delete() error {
	_, err := l.client.DeleteLoadBalancer(&awselbv2.DeleteLoadBalancerInput{
		LoadBalancerArn: l.arn,
	})

	if err != nil {
		return fmt.Errorf("FAILED deleting %s %s: %s", l.rtype, l.identifier, err)
	}

	return nil
}

func (l LoadBalancer) Name() string {
	return l.identifier
}

func (l LoadBalancer) Type() string {
	return l.rtype
}
