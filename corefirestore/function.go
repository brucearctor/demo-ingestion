package corefirestore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// TODO: get this to import from buf.build rather than local

	fp "github.com/brucearctor/demo-ingestion/corefirestore/_go/proto"
	"google.golang.org/protobuf/proto"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	// firebase "firebase.google.com/go"
	firestore "cloud.google.com/go/firestore"
	"cloud.google.com/go/logging"
)

func init() {

	functions.HTTP("ReceivePushAndInsert", receivePushAndInsert)
}

func receivePushAndInsert(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	logName := "my-currentflights-log"
	projectID := "brucearctor-demo-ingestion"

	logClient, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer logClient.Close()

	logger := logClient.Logger(logName).StandardLogger(logging.Info)
	logger.Println("hello world")

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Error reading message body: %v", err)
		return
	}
	logger.Println("WOW")

	// // Parse the message into your Protobuf type
	var msg fp.PostFlightStatusRequest
	if err := proto.Unmarshal(body, &msg); err != nil {
		logger.Printf("Error parsing Protobuf message: %v", err)
		fmt.Fprintf(w, "Error parsing Protobuf message: %v", err)
		return
	}
	logger.Println("FLIGHT ID:")
	logger.Println(msg.FlightId)

	// Convert Protobuf to JSON
	fmt.Fprintf(w, "stringed: %v", &msg)

	jsonData, err := json.Marshal(&msg)
	if err != nil {
		// logger.Printf("Error marshalling to JSON: %v", err)
		fmt.Fprintf(w, "Error marshalling to JSON: %v", err)
		return
	}
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		// logger.Printf("Error unmarshalling from JSON: %v", err)
		fmt.Fprintf(w, "Error unmarshalling from JSON: %v", err)
		return
	}
	logger.Printf("Successfully processed message: %v", string(jsonData))
	logger.Printf("jsonData: %v", jsonData)
	fmt.Printf("stringed: %v", string(jsonData))
	fmt.Printf("jsondata: %v", string(jsonData))

	// TODO:  Can firestore client be global var, and in init()
	// TODO: database name as VAR, provided by terraform
	firestoreClient, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	colRef := firestoreClient.Collection("flights")
	_, err = colRef.NewDoc().Set(ctx, data)

	if err != nil {
		log.Fatalf("Error creating document: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}
