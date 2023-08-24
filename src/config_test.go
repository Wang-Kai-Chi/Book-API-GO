package main

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	con := MustGetConfig()

	t.Log(con)
}
