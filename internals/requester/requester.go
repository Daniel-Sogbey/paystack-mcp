package requester

import (
	"encoding/json"
	"io"
	"log"
	"maps"
	"net/http"
	"net/url"
	"strings"
)

func Requester[T any](method, baseUrl string, header http.Header, queryParams url.Values, body io.Reader) (*T, error) {
	fullUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	log.Println("Info Requester: URL", fullUrl.String())

	if method == http.MethodGet {
		q := fullUrl.Query()

		for k, v := range queryParams {
			q.Set(k, strings.Join(v, ","))
		}

		fullUrl.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, fullUrl.String(), body)
	if err != nil {
		return nil, err
	}

	if header != nil {
		maps.Copy(req.Header, header)
	}

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t T
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		log.Println("Error Requester:", err)
		return nil, err
	}

	return &t, nil
}
