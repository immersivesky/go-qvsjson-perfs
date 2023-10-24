package qvsjson_test

import (
	"encoding/json"
	"testing"
)

type Request struct {
	Count  int `json:"count,omitempty"`
	Offset int `json:"offset,omitempty"`
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sendMessage(Request{
			Count:  100,
			Offset: 100,
		})
	}
}

func sendMessage(req Request) {
	if _, err := json.Marshal(req); err != nil {
		panic(err)
	}
}
