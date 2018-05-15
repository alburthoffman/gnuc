package httpcommand

import (
	"crypto/tls"
	"net/http"
	"time"
	"io/ioutil"
	"strings"
)

type HttpResponse struct {
	httpStatus int
	err error
	content string
}

func Get(url string) (hr *HttpResponse) {
	client := getHttpClient()
	resp, err := client.Get(url)
	if err == nil {
		defer resp.Body.Close()
	}
	return getHttpResponse(resp, err)
}

func Post(url string, contentType string, body string) (hr *HttpResponse)  {
	client := getHttpClient()
	resp, err := client.Post(url, contentType, strings.NewReader(body))
	if err == nil {
		defer resp.Body.Close()
	}
	return getHttpResponse(resp, err)
}

func getHttpClient() *http.Client {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	tlsConfig.BuildNameToCertificate()
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Timeout: 5 * time.Second, Transport: tr}
	return client
}

func getHttpResponse(response *http.Response, err error) (hr *HttpResponse) {
	httpResponse := &HttpResponse{httpStatus: -1, err: nil, content: ""}
	if (err != nil) {
		httpResponse.err = err
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if (err != nil) {
			httpResponse.err = err
			return httpResponse
		} else {
			httpResponse.content = string(body)
			httpResponse.httpStatus = response.StatusCode
		}
	}
	return httpResponse
}