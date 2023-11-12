package main

//Grant the roles/serviceusage.serviceUsageViewer to the service account.
import (
	"context"
	"fmt"
	"log"
	//Run google.golang.org/api/serviceusage/v1 to retrieve the serviceusage package.
	"google.golang.org/api/serviceusage/v1"
)

func main() {
	projectID := "compute-engine-examples"

	ctx := context.Background()

	// Create a serviceusage client with the appropriate credentials
	svc, err := serviceusage.NewService(ctx)
	if err != nil {
		log.Fatalf("Failed to create serviceusage client: %v", err)
	}

	// Call the services.list method to get a list of all enabled services
	resp, err := svc.Services.List(fmt.Sprintf("projects/%s", projectID)).Filter("state:ENABLED").Do()
	if err != nil {
		log.Fatalf("Failed to list services: %v", err)
	}

	// Print the name of each enabled service
	for _, service := range resp.Services {
		fmt.Println(service.Name)
	}
}
