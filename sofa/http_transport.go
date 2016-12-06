package sofa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

// HTTPTransport implements Sofa pbrpc over HTTP
type HTTPTransport struct {
	endpoint   string
	httpClient *http.Client
}

// NewHTTPTransport returns a new HTTPTransport
func NewHTTPTransport(address string) (Transport, error) {
	transport := &HTTPTransport{
		endpoint:   address,
		httpClient: http.DefaultClient,
	}

	return transport, nil
}

// RequestResponse sends request and awaits response
func (t *HTTPTransport) RequestResponse(ctx context.Context, method string, req, resp proto.Message) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "failed to marshal request to JSON")
	}

	url := t.makeURL(method)

	httpRequest, err := http.NewRequest("POST", url, bytes.NewReader(reqBody))
	if err != nil {
		return errors.Wrap(err, "failed to create http request")
	}

	httpResponse, err := t.httpClient.Do(httpRequest)
	if err != nil {
		return errors.Wrap(err, "http request failed")
	}
	defer httpResponse.Body.Close()

	decoder := json.NewDecoder(httpResponse.Body)
	err = decoder.Decode(&resp)
	if err != nil {
		return errors.Wrap(err, "failed to parse response body")
	}

	return nil
}

func (t *HTTPTransport) makeURL(method string) string {
	return fmt.Sprintf("http://%s/%s", t.endpoint, method)
}
