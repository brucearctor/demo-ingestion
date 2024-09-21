package collector

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	fp "gen/go/buf.build/gen/go/brucearctor/demo-ingestion/protocolbuffers/go"

	"cloud.google.com/go/pubsub"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/protobuf/proto"
)

// GCP_PROJECT is a user-set environment variable.
var projectID = os.Getenv("GCP_PROJECT")
var topicID = os.Getenv("TOPIC")

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
	functions.HTTP("ReceiveAndPublish", receiveAndPublish)

}

func receiveAndPublish(w http.ResponseWriter, r *http.Request) {
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
	id, err := client.Topic(topicID).Publish(r.Context(), m).Get(r.Context())
	if err != nil {
		log.Printf("topic(%s).Publish.Get: %v", topicID, err)
		http.Error(w, "Error publishing message", http.StatusInternalServerError)
		return
	}

	// probably don't need to print the message id
	// in that case, change id var to _ in the call to publish, above
	fmt.Fprintf(w, "Published msg: %v", id)
}
