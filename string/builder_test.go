package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var x, y, z int = 1, 20, 309

func UseFmt() string {
	return fmt.Sprintf("%d_%d_%d_group", x, y, z)
}

func UseBuilder() string {
	var b strings.Builder
	b.WriteString(strconv.Itoa(x))
	b.WriteString("_")
	b.WriteString(strconv.Itoa(y))
	b.WriteString("_")
	b.WriteString(strconv.Itoa(z))
	b.WriteString("_group")
	return b.String()
}
func BenchmarkFmt(b *testing.B) {
	for b.Loop() {
		UseFmt()
	}
}

func BenchmarkBuilder(b *testing.B) {
	for b.Loop() {
		UseBuilder()
	}
}

// go test -bench=. -benchmem builder_test.go
