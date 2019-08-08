package main

import (
	"context"

	"cloud.google.com/go/pubsub"
)

var (
	client *pubsub.Client
	top    *pubsub.Topic
)

func init() {

	ctx := context.Background()
	c, e := pubsub.NewClient(ctx, prgID)
	if e != nil {
		logger.Fatalf("Error creating PubSub client: %v", e)
	}
	client = c

	t := c.Topic(topic)
	topicExists, _ := t.Exists(ctx)

	if !topicExists {
		logger.Printf("Topic %s not found, creating...", topic)
		newTop, err := c.CreateTopic(ctx, topic)
		if err != nil {
			logger.Fatalf("Unable to create topic: %s - %v", topic, err)
		}
		top = newTop
	}

	top = t

}

// push persist the content
func push(ctx context.Context, data []byte) error {
	msg := &pubsub.Message{Data: data}
	result := top.Publish(ctx, msg)
	_, err := result.Get(ctx)
	return err
}
