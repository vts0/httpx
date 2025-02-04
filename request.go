package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RequestOptions struct {
	Client  *http.Client
	Method  string
	URL     string
	Headers map[string]string
	Query   interface{}
	Body    interface{}
}

func Get[T any](ctx context.Context, url string, options *RequestOptions) (*T, error) {
	return performRequest[T](ctx, http.MethodGet, url, nil, options)
}

func Post[T any](ctx context.Context, url string, body interface{}, options *RequestOptions) (*T, error) {
	return performRequest[T](ctx, http.MethodPost, url, body, options)
}

func Put[T any](ctx context.Context, url string, body interface{}, options *RequestOptions) (*T, error) {
	return performRequest[T](ctx, http.MethodPut, url, body, options)
}

func Delete[T any](ctx context.Context, url string, options *RequestOptions) (*T, error) {
	return performRequest[T](ctx, http.MethodDelete, url, nil, options)
}

func Patch[T any](ctx context.Context, url string, body interface{}, options *RequestOptions) (*T, error) {
	return performRequest[T](ctx, http.MethodPatch, url, body, options)
}

func performRequest[T any](ctx context.Context, method, url string, body interface{}, options *RequestOptions) (*T, error) {
	if options == nil {
		options = &RequestOptions{}
	}
	if options.Client == nil {
		options.Client = http.DefaultClient
	}
	if options.Headers == nil {
		options.Headers = make(map[string]string)
	}

	req, err := buildRequest(ctx, method, url, body, options)
	if err != nil {
		return nil, err
	}

	resp, err := options.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request %s failed with status: %s", req.URL.String(), resp.Status)
	}

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

func buildRequest(ctx context.Context, method, url string, body interface{}, options *RequestOptions) (*http.Request, error) {
	bodyReader, err := serializeToReader(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	if options.Query != nil {
		queryString, err := serializeToQueryString(options.Query)
		if err != nil {
			return nil, err
		}
		req.URL.RawQuery = queryString
	}

	return req, nil
}

func serializeToReader(data interface{}) (io.Reader, error) {
	if data == nil {
		return nil, nil
	}
	bodyBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize data: %w", err)
	}
	return bytes.NewReader(bodyBytes), nil
}

func serializeToQueryString(query interface{}) (string, error) {
	if query == nil {
		return "", nil
	}
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return "", fmt.Errorf("failed to serialize query: %w", err)
	}
	var queryMap map[string]interface{}
	if err := json.Unmarshal(queryBytes, &queryMap); err != nil {
		return "", fmt.Errorf("failed to unmarshal query: %w", err)
	}
	values := url.Values{}
	for key, value := range queryMap {
		if value != nil {
			values.Add(key, fmt.Sprint(value))
		}
	}
	return values.Encode(), nil
}
