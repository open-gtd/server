package sse

import "testing"

func TestCreatePushDataFuncShouldReturnCallable(t *testing.T) {
	n := NullRegisterer{}

	c := n.CreatePushDataFunc("xxx", func() {})

	c("topic", 0)
}
