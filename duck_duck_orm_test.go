package ddorm

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestOpen(t *testing.T) {
	d := DDORM{}

	err := d.Connect(POSTGRES, os.Getenv("DB_SOURCE"))

	if err != nil {
		t.Fail()
	}
}
