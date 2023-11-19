package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

// Grant roles/compute.admin to the service account
func main() {
	projectID := "compute-engine-examples" // Replace with your GCP project ID

	// Create a new Compute Service client
	ctx := context.Background()
	computeService, err := compute.NewService(ctx, option.WithScopes(compute.ComputeScope))
	if err != nil {
		log.Fatalf("Failed to create Compute Service client: %v", err)
	}

	// List all VM instances in the project
	instances, err := computeService.Instances.List(projectID, "us-central1-a").Do()
	if err != nil {
		log.Fatalf("Failed to retrieve VM instances: %v", err)
	}

	if len(instances.Items) == 0 {
		fmt.Println("No VM instances found in project")
	} else {
		// Print the name of each VM instance
		fmt.Printf("INSTANCE\tNETWORK ADDRESS\tSTATUS\n")
		for _, instance := range instances.Items {
			fmt.Printf("%s\t%s\t%s\n", instance.Name, instance.NetworkInterfaces[0].NetworkIP, instance.Status)
		}
	}

}
