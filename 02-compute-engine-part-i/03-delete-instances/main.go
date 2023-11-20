package main

import (
	"context"
	"fmt"
	"log"

	"time"

	"google.golang.org/api/compute/v1"
)

func deleteInstance(projectID, instanceName, zone string) error {
	ctx := context.Background()
	fmt.Printf("Deleting instance %s in project %s\n", instanceName, projectID)

	// Create a new Compute Engine service client
	service, err := compute.NewService(ctx)
	if err != nil {
		return fmt.Errorf("failed to create service: %v", err)
	}

	// Call the Instances.Delete method to create the instance
	op, err := service.Instances.Delete(projectID, zone, instanceName).Do()
	if err != nil {
		return fmt.Errorf("failed to delete instance: %v", err)
	}

	for {
		// Check operation status
		op, err = service.ZoneOperations.Get(projectID, zone, op.Name).Do()
		if err != nil {
			log.Fatalf("Failed to get operation: %v", err)
		}
		if op.Status == "DONE" {
			break
		}
		fmt.Printf("Waiting more 10 secs for deleting instance %s\n", instanceName)
		time.Sleep(10 * time.Second)
	}

	fmt.Printf("Instance %s created successfully!\n", instanceName)
	return nil
}

func main() {
	projectID := "compute-engine-examples"
	instanceNameTemplate := "test-instance-%d"
	zone := "us-central1-a"

	//We can use go routines to create instances in parallel
	for i := 0; i < 5; i++ {
		instanceName := fmt.Sprintf(instanceNameTemplate, i)
		err := deleteInstance(projectID, instanceName, zone)
		if err != nil {
			log.Fatalf("Failed to delete instance: %v", err)
		}
	}
}
