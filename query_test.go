package ddorm

import (
	"testing"
)

type messages struct {
	id      int
	message string
}

func TestFind(t *testing.T) {
	m := messages{id: 1}

	// Currently fails
	D.Find(&m)

}
