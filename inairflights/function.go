package inairflights

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/protobuf/proto"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/logging"
	fp "github.com/brucearctor/demo-ingestion/inairflights/_go/proto"
)

func init() {
	functions.HTTP("InAirFlights", inAirFlights)
}

func inAirFlights(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	logName := "inairflights-log"
	projectID := "brucearctor-demo-ingestion"

	logClient, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer logClient.Close()

	logger := logClient.Logger(logName).StandardLogger(logging.Info)

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
	var dataMap map[string]interface{}
	err = json.Unmarshal(jsonData, &dataMap)
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
	logger.Println("OMG2")

	firestoreClient, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	inAirColRef := firestoreClient.Collection("inair")
	logger.Println("HERE2")
	// documentID := msg.FlightId

	documentID := string(msg.FlightId)
	logger.Println("HERE3")
	// TODO check what getting, AND more.
	// MAYBE something like if COUNT of documents with flightid > 0...
	// Also where eventarc comes in [ to ensure event already in flights, etc ]
	existingDoc, err := inAirColRef.Doc(documentID).Get(ctx)
	logger.Println("HERE4")

	if err != nil {
		logger.Println("WHERE!?")
		logger.Printf("collection.Doc().Get: %v", err)
		//log.Fatalf("collection.Doc().Get: %v", err)
	}
	logger.Println("HERE5")
	logger.Println(existingDoc.Data()["current_timestamp"])
	logger.Println("TYPEOF -->")
	logger.Println(reflect.TypeOf(existingDoc.Data()["current_timestamp"]))
	existingEventTime, ok := existingDoc.Data()["current_timestamp"].(int64)

	if !ok {
		// This shouldn't occur, but ...
		// TODO: probably want a DLQ
		fmt.Println("event_time field is not correct")
		logger.Println("event_time field is not correct")
		return
	}

	// TODO: Do i need to cast these?
	// This IF statement determines whether doc needs to be updated
	logger.Println("checking timestamp diff ...")
	// TODO: Better logic ... how about when one of these is missing?
	if existingEventTime < msg.CurrentTimestamp {
		logger.Println("in IF statement")
		doc, err := inAirColRef.Doc(documentID).Set(ctx, dataMap)
		if err != nil {
			logger.Printf("collection.Doc().Get: %v", err)
			// log.Fatalf("collection.Doc().Get: %v", err)
			// This is not a fatal error, as it is possible that the document doesn't yet exist
			// More complicated handling could be that FIRST insert ONLY if a 'takeoff' event
			// For now, we tried, and we update.
			// Invites questions of how to handle after already landed.
		}
		logger.Println("DOC ------>")
		fmt.Println(doc)
		logger.Println(doc)
	}
	logger.Println("WHEREAMI?")

	doc, err := inAirColRef.Doc(documentID).Set(ctx, dataMap)
	if err != nil {
		logger.Printf("collection.Doc().Set: %v", err)
		log.Fatalf("collection.Doc().Set: %v", err)
	}
	fmt.Println(doc)
	logger.Println(doc)
	// Access the document data

	logger.Println("DATA2")

	w.WriteHeader(http.StatusOK)
}
