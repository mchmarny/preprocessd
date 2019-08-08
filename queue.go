package main

import (
	"context"
	"errors"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// queue pushes events to pubsub topic
type queue struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func newQueue(ctx context.Context) (q *queue, err error) {

	if ctx == nil {
		return nil, errors.New("context not set")
	}

	c, e := pubsub.NewClient(ctx, prgID)
	if e != nil {
		return nil, e
	}

	t := c.Topic(topic)
	topicExists, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}

	if !topicExists {
		logger.Printf("Topic %s not found, creating...", topic)
		t, err = c.CreateTopic(ctx, topic)
		if err != nil {
			return nil, fmt.Errorf("Unable to create topic: %s - %v", topic, err)
		}
	}

	o := &queue{
		client: c,
		topic:  t,
	}

	return o, nil
}

// push persist the content
func (q *queue) push(ctx context.Context, data []byte) error {
	msg := &pubsub.Message{Data: data}
	result := q.topic.Publish(ctx, msg)
	_, err := result.Get(ctx)
	return err
}
