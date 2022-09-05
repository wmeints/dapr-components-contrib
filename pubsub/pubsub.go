/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pubsub

import (
	"context"
	"fmt"

	"github.com/dapr/components-contrib/health"
)

// PubSub is the interface for message buses.
type PubSub interface {
	Batcher
	Init(metadata Metadata) error
	Features() []Feature
	Publish(req *PublishRequest) error
	Subscribe(ctx context.Context, req SubscribeRequest, handler Handler) error
	Close() error
}

type Batcher interface {
	BatchPublish(req *BatchPublishRequest) BatchPublishResponse
	BatchSubscribe(ctx context.Context, req SubscribeRequest, handler BatchHandler) error
}

// Handler is the handler used to invoke the app handler.
type Handler func(ctx context.Context, msg *NewMessage) error

// BatchHandler is the handler used to invoke the app handler during batch subscribe.
type BatchHandler func(ctx context.Context, msg *NewBatchMessage) error

func Ping(pubsub PubSub) error {
	// checks if this pubsub has the ping option then executes
	if pubsubWithPing, ok := pubsub.(health.Pinger); ok {
		return pubsubWithPing.Ping()
	} else {
		return fmt.Errorf("ping is not implemented by this pubsub")
	}
}

// DefaultBatcher is default implemnetation for Batcher
type DefaultBatcher struct {
	p PubSub
}

// NewDefaultBatcher to create new DefaultBatcher for a PubSub
func NewDefaultBatcher(pubsub PubSub) DefaultBatcher {
	defaultBatcher := DefaultBatcher{}
	defaultBatcher.p = pubsub

	return defaultBatcher
}

// BatchPublish Default Implementation
func (p *DefaultBatcher) BatchPublish(req *BatchPublishRequest) BatchPublishResponse {
	return BatchPublishResponse{}
}

// BatchSubscribe Default Implementation
func (p *DefaultBatcher) BatchSubscribe(tx context.Context, req SubscribeRequest, handler BatchHandler) error {
	return nil
}
