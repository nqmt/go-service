package benchmark

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
	"testing"
)

func BenchmarkXid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		id := xid.New()
		_ = id.String()
	}
}


func BenchmarkUuid(b *testing.B) {
	for n := 0; n < b.N; n++ {
		id := uuid.New()
		_ = id.String()
	}
}