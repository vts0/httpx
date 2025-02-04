# HTTPX

This package provides utility functions to perform HTTP requests (GET, POST, PUT, DELETE, PATCH) with customizable options and automatic deserialization of JSON responses.

## Installation
To install the package, run:
```sh
go get github.com/vts0/httpx
```

## Example Usage
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vts0/httpx"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	ctx := context.Background()
	url := "https://jsonplaceholder.typicode.com/users/1"
	options := &httpx.RequestOptions{}

	// Send GET request
	user, err := httpx.Get[User](ctx, url, options)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}

	fmt.Printf("User ID: %d, Name: %s\n", user.ID, user.Name)
}
```

This code demonstrates how to use the httpx package to perform a GET request and deserialize the JSON response into a User struct.

## Types

### `RequestOptions`
`RequestOptions` contains configuration for HTTP requests, including client settings, headers, body, and query parameters.

Fields:
- `Client` (`*http.Client`): The HTTP client to use for the request. Defaults to the `http.DefaultClient` if `nil`.
- `Method` (`string`): The HTTP method for the request (GET, POST, PUT, DELETE, PATCH).
- `URL` (`string`): The target URL for the request.
- `Headers` (`map[string]string`): Custom HTTP headers to include in the request.
- `Query` (`interface{}`): Query parameters to be serialized and added to the URL.
- `Body` (`interface{}`): The body content to be included in the request.

## Functions

### `Get[T any](ctx context.Context, url string, options *RequestOptions) (*T, error)`
Performs a GET request to the specified `url` with the provided `options`.

**Parameters**:
- `ctx`: The context to control the request’s lifetime.
- `url`: The target URL for the request.
- `options`: An optional `RequestOptions` to customize the request behavior.

**Returns**:
- A pointer to the result (`*T`) if successful, or an error if the request fails.

### `Post[T any](ctx context.Context, url string, body interface{}, options *RequestOptions) (*T, error)`
Performs a POST request to the specified `url` with the provided `body` and `options`.

**Parameters**:
- `ctx`: The context to control the request’s lifetime.
- `url`: The target URL for the request.
- `body`: The body content for the request.
- `options`: An optional `RequestOptions` to customize the request behavior.

**Returns**:
- A pointer to the result (`*T`) if successful, or an error if the request fails.

### `Put[T any](ctx context.Context, url string, body interface{}, options *RequestOptions) (*T, error)`
Performs a PUT request to the specified `url` with the provided `body` and `options`.

**Parameters**:
- `ctx`: The context to control the request’s lifetime.
- `url`: The target URL for the request.
- `body`: The body content for the request.
- `options`: An optional `RequestOptions` to customize the request behavior.

**Returns**:
- A pointer to the result (`*T`) if successful, or an error if the request fails.

### `Delete[T any](ctx context.Context, url string, options *RequestOptions) (*T, error)`
Performs a DELETE request to the specified `url` with the provided `options`.

**Parameters**:
- `ctx`: The context to control the request’s lifetime.
- `url`: The target URL for the request.
- `options`: An optional `RequestOptions` to customize the request behavior.

**Returns**:
- A pointer to the result (`*T`) if successful, or an error if the request fails.

### `Patch[T any](ctx context.Context, url string, body interface{}, options *RequestOptions) (*T, error)`
Performs a PATCH request to the specified `url` with the provided `body` and `options`.

**Parameters**:
- `ctx`: The context to control the request’s lifetime.
- `url`: The target URL for the request.
- `body`: The body content for the request.
- `options`: An optional `RequestOptions` to customize the request behavior.

**Returns**:
- A pointer to the result (`*T`) if successful, or an error if the request fails.

### `performRequest[T any](ctx context.Context, method, url string, body interface{}, options *RequestOptions) (*T, error)`
Internal function that sends an HTTP request using the provided method (`GET`, `POST`, `PUT`, `DELETE`, `PATCH`) and options.

**Parameters**:
- `ctx`: The context to control the request’s lifetime.
- `method`: The HTTP method for the request.
- `url`: The target URL for the request.
- `body`: The body content for the request (optional).
- `options`: The request options containing client, headers, and query parameters.

**Returns**:
- A pointer to the deserialized response (`*T`), or an error if the request fails.

### `buildRequest(ctx context.Context, method, url string, body interface{}, options *RequestOptions) (*http.Request, error)`
Constructs the HTTP request based on the provided parameters, including the body and headers.

**Parameters**:
- `ctx`: The context to control the request’s lifetime.
- `method`: The HTTP method for the request.
- `url`: The target URL for the request.
- `body`: The body content for the request (optional).
- `options`: The request options containing headers and query parameters.

**Returns**:
- A pointer to the `http.Request`, or an error if the request creation fails.

### `serializeToReader(data interface{}) (io.Reader, error)`
Serializes the provided data to an `io.Reader` for the request body.

**Parameters**:
- `data`: The data to be serialized into the request body.

**Returns**:
- An `io.Reader` containing the serialized data, or an error if serialization fails.

### `serializeToQueryString(query interface{}) (string, error)`
Serializes the query parameters to a URL-encoded query string.

**Parameters**:
- `query`: The query parameters to be serialized.

**Returns**:
- A URL-encoded query string, or an error if serialization fails.

## Error Handling
- If the HTTP request fails, the error is propagated with a message indicating the issue.
- If the HTTP status code is not within the 2xx range, an error with the status code is returned.
- If the response cannot be deserialized into the expected type, an error is returned.
