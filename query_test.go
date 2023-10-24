package qvsjson_test

import (
	"fmt"
	"net/url"
	"testing"
)

func BenchmarkQuery(b *testing.B) {
	q := make(url.Values)

	for i := 0; i < b.N; i++ {
		q.Set("count", fmt.Sprint(100))
		q.Set("offset", fmt.Sprint(100))
		q.Encode()
	}
}
