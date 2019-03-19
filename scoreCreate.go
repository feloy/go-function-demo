// Package p contains a Firestore Cloud Function.
package p

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/functions/metadata"
	fsv "github.com/feloy/go-firestore-value"
)

// FirestoreEvent is the payload of a Firestore event.
// Please refer to the docs for additional information
// regarding Firestore events.
type FirestoreEvent struct {
	OldValue fsv.FirestoreValue `json:"oldValue"`
	Value    fsv.FirestoreValue `json:"value"`
}

// ScoreCreate is triggered by a change to a Firestore document.
func ScoreCreate(ctx context.Context, e FirestoreEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function triggered by change to: %v", meta.Resource)
	log.Printf("%v", e)

	uid, err := e.Value.GetStringValue("uid")
	if err == nil {
		log.Printf("uid: %s", uid)
	} else {
		log.Printf("Error: %s", err)
	}
	return nil
}
