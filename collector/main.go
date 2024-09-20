package github.com/brucearctor/demo-ingest/collector

import (
	fp "../gen/go/buf.build/gen/go/brucearctor/demo-ingestion/protocolbuffers/go"

	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/proto"
	// I like the functions Framework
	// see: https://cloud.google.com/functions/docs/functions-framework
	// "github.com/GoogleCloudPlatform/functions-framework-go/functions"
	// but, not currently using.
)

// GCP_PROJECT is a user-set environment variable.
var projectID = os.Getenv("GCP_PROJECT")
var TopicID = os.Getenv("TOPIC")

// client is a global Pub/Sub client, initialized once per instance.
var client *pubsub.Client

func init() {

	var err error
	// client is initialized with context.Background() because it should
	// persist between function invocations.
	client, err = pubsub.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("pubsub.NewClient: %v", err)
	}
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/flights", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// ProtobufMessage is from a generated pb.go with same package name
	p := &fp.PostFlightStatusRequest{}
	data, err := io.ReadAll(r.Body)
	fmt.Println(string(data))
	if err != nil {
		http.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}

	// Parse the request body to JSON
	if err := json.Unmarshal(data, &p); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Error parsing request", http.StatusBadRequest)
		return
	}

	//write proto to pb-out (pbo)
	pbo, err := proto.Marshal(p)
	m := &pubsub.Message{
		Data: pbo,
	}
	if err != nil {
		log.Printf("Marshal ERROR: %v", err)
		http.Error(w, "Error publishing message", http.StatusInternalServerError)
		return
	}

	// r.Context() used only because they are only needed for this invocation.
	// Currently assumes the topic already exists
	// TODO: extra logging or otherwise verifying topic exists?
	id, err := client.Topic(TopicID).Publish(r.Context(), m).Get(r.Context())
	if err != nil {
		log.Printf("topic(%s).Publish.Get: %v", TopicID, err)
		http.Error(w, "Error publishing message", http.StatusInternalServerError)
		return
	}

	// probably don't need to print the message id
	// in that case, change id var to _ in the call to publish, above
	fmt.Fprintf(w, "Published msg: %v", id)
}
