package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/scottcrespo/mongobeat/beater"
)

func main() {
	err := beat.Run("mongobeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
