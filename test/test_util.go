package book

import (
	"context"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/scott-dn/go-boilerplate/api"
	"github.com/scott-dn/go-boilerplate/internal/app"
)

var application *app.App //nolint:gochecknoglobals

func init() { //nolint:gochecknoinits
	_, filename, _, _ := runtime.Caller(0) //nolint:dogsled
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	// init env
	os.Setenv("GO_ENV", "production")

	// init app
	application = app.Init()
	go api.StartHTTPServer(application)

	// request health
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"http://localhost:8080/health",
		nil,
	)
	if err != nil {
		panic(err)
	}

	// poll until server is up
	for {
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode == http.StatusOK {
			break
		}
		resp.Body.Close()
	}
}
