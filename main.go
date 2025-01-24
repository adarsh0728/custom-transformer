package main

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/numaproj/numaflow-go/pkg/sourcetransformer"
)

// PanicTransformer does two things: change mssg event time to time.Now(), logs it and panics when
// mssg counter is multiple of 500
type PanicTransformer struct {
	counter uint64
}

func (a *PanicTransformer) Transform(ctx context.Context, keys []string, d sourcetransformer.Datum) sourcetransformer.Messages {
	// Increment the counter atomically
	newCount := atomic.AddUint64(&a.counter, 1)
	// Check if the counter is a multiple of 500
	if newCount%500 == 0 {
		panic("Counter reached a multiple of 500")
	}
	// Update message event time to time.Now()
	eventTime := time.Now()
	// print so that we can use it for testing.
	fmt.Printf("PanicTransformer: Assigning event time %v to message %s and count is %d\n", eventTime, string(d.Value()), newCount)
	return sourcetransformer.MessagesBuilder().Append(sourcetransformer.NewMessage(d.Value(), eventTime).WithKeys(keys))
}

func main() {
	err := sourcetransformer.NewServer(&PanicTransformer{}).Start(context.Background())
	if err != nil {
		log.Panic("Failed to start custom panic transformer server: ", err)
	} else {
		fmt.Println("Successfully started panic transformer server")
	}
}
