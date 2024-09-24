package currentflights

import (
	"context"
	"fmt"
	"log"

	firestore "cloud.google.com/go/firestore"
	"cloud.google.com/go/logging"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/googleapis/google-cloudevents-go/cloud/firestoredata"
	"google.golang.org/protobuf/proto"
)

// set the GOOGLE_CLOUD_PROJECT environment variable when deploying.
// var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
var projectID = "brucearctor-demo-ingestion"

// client is a Firestore client, reused between function invocations.
var client *firestore.Client

func init() {
	functions.CloudEvent("SetMostRecentFlights", setMostRecentFlights)
}

// HelloFirestore is triggered by a change to a Firestore document.
func setMostRecentFlights(ctx context.Context, event event.Event) error {

	// ctx := context.Background()

	logName := "my-currentflights-log"
	projectID := "brucearctor-demo-ingestion"

	logClient, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer logClient.Close()

	logger := logClient.Logger(logName).StandardLogger(logging.Info)
	logger.Println("hello world")

	var data firestoredata.DocumentEventData
	if err := proto.Unmarshal(event.Data(), &data); err != nil {
		return fmt.Errorf("proto.Unmarshal: %w", err)
	}

	logger.Printf("Function triggered by change to: %v\n", event.Source())
	logger.Printf("Old value: %+v\n", data.GetOldValue())
	logger.Printf("New value: %+v\n", data.GetValue())
	// logger.Printf(data.ProtoMessage())
	fmt.Printf("Function triggered by change to: %v\n", event.Source())
	fmt.Printf("Old value: %+v\n", data.GetOldValue())
	fmt.Printf("New value: %+v\n", data.GetValue())
	return nil
}
