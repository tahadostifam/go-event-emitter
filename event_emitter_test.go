package event_emitter

import (
	"testing"
)

type SampleEventArgs struct {
	Name string
}

func BenchmarkEventEmitter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		eventEmitter := NewEventEmitter()

		eventName := "sample-event"
		eventArgs := SampleEventArgs{Name: "EventEmitter"}

		err := eventEmitter.On(eventName, func(v interface{}) {
			args, ok := v.(SampleEventArgs)
			if !ok {
				panic("Unable to cast func args to given arg struct")
			}

			if eventArgs.Name != args.Name {
				panic("event args and incoming args are not equal!")
			}
		})
		if err != nil {
			panic(err)
		}

		err = eventEmitter.Emit("sample-event", eventArgs)
		if err != nil {
			panic(err)
		}
	}
}
