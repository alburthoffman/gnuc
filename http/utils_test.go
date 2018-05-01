package http

import (
	"testing"
)

func TestGet(t *testing.T) {
	bingUrl := "https://www.bing.com"
	httpResponse := Get(bingUrl)

	if httpResponse.httpStatus != 200 {
		t.Errorf("http status should be 200, while it is %d", httpResponse.httpStatus)
	}
}

