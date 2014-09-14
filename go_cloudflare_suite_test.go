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

func makeClient(t ...*testing.T) *Client {

	client, err := NewClient("foobaremail", "foobartoken")

	if t != nil {
		Ω(err).ShouldNot(HaveOccurred())
		Ω(client.Token).Should(Equal("foobartoken"))
		Ω(client.Email).Should(Equal("foobaremail"))
	}

	return client
}

func TestClient_NewRequest(t *testing.T) {
	RegisterTestingT(t)

	params := map[string]string{
		"foo": "bar",
		"baz": "bar",
	}

	client := makeClient(t)

	req, err := client.NewRequest(params, "POST", "baz")
	Ω(err).ShouldNot(HaveOccurred())

	expected := "https://www.cloudflare.com/api_json.html?a=baz&baz=bar&email=foobaremail&foo=bar&tkn=foobartoken"
	encoded := req.URL.Query()
	Ω(encoded.Get("foo")).Should(Equal("bar"))
	Ω(encoded.Get("baz")).Should(Equal("bar"))
	Ω(req.URL.String()).Should(Equal(expected))
	Ω(req.Method).Should(Equal("POST"))
}
