package main

import (
	"context"
	"fmt"
	"log"
	"time"

	//Run go get google.golang.org/api/serviceusage/v1 to retrieve the serviceusage package.
	"google.golang.org/api/serviceusage/v1"
)

// Enable roles/serviceusage.serviceUsageAdmin
func main() {
	projectID := "compute-engine-examples"
	apiName := "compute.googleapis.com"

	ctx := context.Background()

	// Create a serviceusage client with default credentials
	svc, err := serviceusage.NewService(ctx)
	if err != nil {
		log.Fatalf("Failed to create serviceusage client: %v", err)
	}

	// Enable the API
	op, err := svc.Services.Enable(fmt.Sprintf("projects/%s/services/%s", projectID, apiName), &serviceusage.EnableServiceRequest{}).Do()
	if err != nil {
		log.Fatalf("Failed to enable API: %v", err)
	}

	for {
		// Check operation status
		op, err = svc.Operations.Get(op.Name).Do()
		if err != nil {
			log.Fatalf("Failed to get operation: %v", err)
		}
		if op.Done {
			break
		}
		fmt.Printf("Waiting for operation to complete: %v\n", op.Name)
		time.Sleep(10 * time.Second)
	}

	fmt.Printf("API %s has been enabled for project %s\n", apiName, projectID)
}
