package modifier

import (
	"net/http"
	"testing"
)

func TestHeaderModifier_ModifyRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("ORIGIN", "origin")

	headers := make(map[string]string)
	headers["TEST-HEADER"] = "test"
	headers["ORIGIN"] = "modified"

	hm := HeaderNewModifier(headers)
	hm.ModifyRequest(req)

	if req.Header.Get("TEST-HEADER") != "test" {
		t.Errorf("Header was not set to request")
	}

	if req.Header.Get("ORIGIN") != "origin" {
		t.Errorf("Origin header was changed in request")
	}
}

func TestHeaderModifier_ModifyResponse(t *testing.T) {
	h := make(http.Header)
	h.Set("ORIGIN", "origin")
	res := http.Response{Header: h}

	headers := make(map[string]string)
	headers["TEST-HEADER"] = "test"
	headers["ORIGIN"] = "modified"

	hm := HeaderNewModifier(headers)
	hm.ModifyResponse(&res)

	if res.Header.Get("TEST-HEADER") != "test" {
		t.Errorf("Header was not set to response")
	}

	if res.Header.Get("ORIGIN") != "origin" {
		t.Errorf("Origin header was changed in response")
	}
}
