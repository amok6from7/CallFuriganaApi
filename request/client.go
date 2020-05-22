package request

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	EndpointURL *url.URL
	HTTPClient  *http.Client
}

func NewClient(endpointURL string) (*Client, error) {
	parsedURL, err := url.ParseRequestURI(endpointURL)
	failOnError(err)
	client := &Client{
		EndpointURL: parsedURL,
		HTTPClient:  &http.Client{},
	}
	return client, nil
}

func (client *Client) newRequest(body io.Reader) (*http.Request, error) {
	endPointURL := *client.EndpointURL
	req, err := http.NewRequest("POST", endPointURL.String(), body)
	failOnError(err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

func decodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(out)
}

func (client *Client) CallApi(req FuriganaApiRequest) (*FuriganaApiResponse, error) {
	params := url.Values{}
	params.Add("app_id", req.AppId)
	params.Add("sentence", req.Sentence)
	params.Add("output_type", req.OutputType)
	httpRequest, err := client.newRequest(strings.NewReader(params.Encode()))
	failOnError(err)
	httpResponse, err := client.HTTPClient.Do(httpRequest)
	failOnError(err)
	var apiResponse FuriganaApiResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
