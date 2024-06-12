package main

import "testing"

func newTestApplication(t *testing.T) *application {
	app := new(application)
	cfg := config{env: "testing"}
	app.config = cfg

	return app
}
