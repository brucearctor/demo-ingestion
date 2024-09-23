package currentflights

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/logging"

	// TODO: get this to import from buf.build rather than local
	fp "github.com/brucearctor/demo-ingestion/currentflights/_go/proto"
	"google.golang.org/protobuf/proto"

	firestore "cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

var firestoreClient *firestore.Client

func init() {

	// // TODO move this to env
	var projectID = "brucearctor-demo-ingestion"

	var err error

	firestoreClient, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	functions.HTTP("ReceivePushAndInsert", receivePushAndInsert)
}

func receivePushAndInsert(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	logName := "my-log"
	projectID := "brucearctor-demo-ingestion"

	logClient, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer logClient.Close()

	logger := logClient.Logger(logName).StandardLogger(logging.Info)
	logger.Println("hello world")

	body, err := io.ReadAll(r.Body)
	// do I need to close the above?

	if err != nil {
		fmt.Fprintf(w, "Error reading message body: %v", err)
		return
	}
	logger.Println("WOW")

	// Parse the message into your Protobuf type
	var msg fp.PostFlightStatusRequest
	if err := proto.Unmarshal(body, &msg); err != nil {
		logger.Printf("Error parsing Protobuf message: %v", err)
		fmt.Fprintf(w, "Error parsing Protobuf message: %v", err)
		return
	}
	logger.Println("FLIGHT ID:")
	logger.Println(msg.FlightId)
	// Convert Protobuf to JSON

	jsonData, err := json.Marshal(&msg)
	if err != nil {
		logger.Printf("Error marshalling to JSON: %v", err)
		fmt.Fprintf(w, "Error marshalling to JSON: %v", err)
		return
	}

	logger.Printf("Successfully processed message: %v", string(jsonData))
	logger.Printf("jsonData: %v", jsonData)

	docRef := firestoreClient.Collection("flights").NewDoc()
	_, err = docRef.Set(ctx, jsonData)
	if err != nil {
		log.Fatalf("Error creating document: %v", err)
	}

	fmt.Println("Flight added successfully!")

}
