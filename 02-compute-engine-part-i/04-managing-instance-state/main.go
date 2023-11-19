package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/api/compute/v1"
)

func main() {
	// Check if there is at least one command-line argument
	if len(os.Args) < 5 {
		fmt.Printf("%s <projectId> <zone> <instanceName> START|STOP|RESTART\n", os.Args[0])
		return
	}
	projectId := os.Args[1]
	zone := os.Args[2]
	instanceName := os.Args[3]

	// Get the input argument from the command-line
	action := os.Args[4]

	// Check if the input argument is different from START, STOP or RESTART
	if action != "START" && action != "STOP" && action != "RESTART" {
		fmt.Println("Invalid action")
		return
	}

	err := changeInstanceState(projectId, zone, instanceName, action)
	if err != nil {
		log.Fatal(err)
	}
}

func changeInstanceState(projectID, zone, instanceName, state string) error {
	ctx := context.Background()

	// Create a new Compute Service client
	service, err := compute.NewService(ctx)
	if err != nil {
		return err
	}

	// Check the current instance state and perform the desired state change
	switch state {
	case "START":
		opr, err := service.Instances.Start(projectID, zone, instanceName).Do()
		if err != nil {
			return err
		}
		return waitForZoneOperation(service, projectID, zone, opr.Name)

	case "STOP":
		opr, err := service.Instances.Stop(projectID, zone, instanceName).Do()
		if err != nil {
			return err
		}
		return waitForZoneOperation(service, projectID, zone, opr.Name)
	case "RESTART":
		opr, err := service.Instances.Reset(projectID, zone, instanceName).Do()
		if err != nil {
			return err
		}
		return waitForZoneOperation(service, projectID, zone, opr.Name)

	}

	return nil
}

func waitForZoneOperation(service *compute.Service, projectID, zone, operationName string) error {
	for {
		// Check operation status
		op, err := service.ZoneOperations.Get(projectID, zone, operationName).Do()
		if err != nil {
			return err
		}
		if op.Status == "DONE" {
			break
		}
		fmt.Printf("Waiting more 10 secs for the operation\n")
		time.Sleep(10 * time.Second)
	}
	return nil
}
