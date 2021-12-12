package gowrike

import (
	"context"
	"io"
	"os"
)

var ACCOUNT = "/account"

func AccountRaw(ctx context.Context) (res io.ReadCloser, err error) {

	URL := os.Getenv("WRIKE_URL") + ACCOUNT
	if err != nil {
		return
	}
	body, _, err := Request(ctx, "GET", URL, nil, nil)
	if err != nil {
		return
	}

	return body, err
}
