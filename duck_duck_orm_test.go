package ddorm

import (
	"os"
	"testing"

	"log"

	_ "github.com/lib/pq"
)

var D DDORM

func TestMain(m *testing.M) {
	D = DDORM{}

	err := D.Connect(POSTGRES, os.Getenv("DB_SOURCE"))

	if err != nil {
		log.Fatalln(err.Error())
	}

	code := m.Run()
	os.Exit(code)
}
