// +build integration

package test

import (
	"log"
	"os"
	"testing"

	"github.com/vigneshv59/typesense-go/typesense"
)

var typesenseClient *typesense.Client

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	var err error
	typesenseClient, err = setupDB()
	if err != nil {
		log.Printf("failed to setup typesense db: %v\n", err)
		return 1
	}
	defer stopDB()
	return m.Run()
}
