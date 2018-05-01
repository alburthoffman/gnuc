package http

import (
	"testing"
	"os"
	"net/http"
	"log"
)

func ApiServerForTesting(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/plain")
	if r.Method == http.MethodGet {
		w.Write([]byte("api server for testing."))
	} else if r.Method == http.MethodPost {
		w.Write([]byte("post success"))
	}
}

func startTestingServer() {
	http.HandleFunc("/api", ApiServerForTesting)
	err := http.ListenAndServe(":19876", nil)
	if err != nil {
		log.Fatal("error to start testing api server: ", err)
	}
}

func setupTestEnv()  {
	go startTestingServer()
}

func teardownTestEnv()  {
}

func TestMain(m *testing.M) {
	setupTestEnv()
	retCode := m.Run()
	teardownTestEnv()
	os.Exit(retCode)
}

func TestGet(t *testing.T) {
	bingUrl := "http://localhost:19876/api"
	httpResponse := Get(bingUrl)

	if httpResponse.httpStatus != 200 {
		t.Errorf("http status should be 200, while it is %d", httpResponse.httpStatus)
	}

	if httpResponse.content != "api server for testing." {
		t.Error("content from server is incorrect.")
	}
}

func TestGetFailure(t *testing.T) {
	invalidUrl := "https://www.invaliddomain.com"
	httpResponse := Get(invalidUrl)

	if httpResponse.err == nil {
		t.Error("we should have an error here")
	}
}

func TestPost(t *testing.T) {
	bingUrl := "http://localhost:19876/api"
	httpResponse := Post(bingUrl, "application/json", "")

	if httpResponse.httpStatus != 200 {
		t.Errorf("http status should be 200, while it is %d", httpResponse.httpStatus)
	}

	if httpResponse.content != "post success" {
		t.Error("content from server is not correct.")
	}
}

