package snowflake

import (
	"testing"
)

func BenchmarkGenID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenID()
	}
}

func BenchmarkGenIDParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = GenID()
		}
	})
}
