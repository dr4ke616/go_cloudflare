package go_cloudflare

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoCloudflare(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCloudflare Suite")
}

func makeClient(t *testing.T) *Client {
	client, err := NewClient("foobaremail", "foobartoken")

	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if client.Token != "foobartoken" {
		t.Fatalf("token not set on client: %s", client.Token)
	}

	if client.Email != "foobaremail" {
		t.Fatalf("email not set on client: %s", client.Token)
	}

	return client
}

func TestClient_NewRequest(t *testing.T) {
	RegisterTestingT(t)

	c := makeClient(t)

	params := map[string]string{
		"foo": "bar",
		"baz": "bar",
	}

	req, err := c.NewRequest(params, "POST", "baz")
	if err != nil {
		t.Fatalf("bad: %v", err)
	}

	encoded := req.URL.Query()
	if encoded.Get("foo") != "bar" {
		t.Fatalf("bad: %v", encoded)
	}

	if encoded.Get("baz") != "bar" {
		t.Fatalf("bad: %v", encoded)
	}

	if encoded.Get("baz") != "bar" {
		t.Fatalf("bad: %v", encoded)
	}
	expected := "https://www.cloudflare.com/api_json.html?a=baz&baz=bar&email=foobaremail&foo=bar&tkn=foobartoken"
	if req.URL.String() != expected {
		t.Fatalf("bad base url: %v\n\nexpected: %v", req.URL.String(), expected)
	}

	if req.Method != "POST" {
		t.Fatalf("bad method: %v", req.Method)
	}
}
