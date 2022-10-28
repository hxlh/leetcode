package main

import "testing"

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findLadders("kite", "ashy", []string{"ante", "does", "jive", "achy", "kids", "kits", "cart", "ache", "ands", "ashe", "acne", "aunt", "aids", "kite", "ants", "anne", "ashy"})
	}
}
