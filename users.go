package gowrike

import (
	"context"
	"fmt"
	"io"
	"os"
)

var USERS = "/users"

func UsersRaw(ctx context.Context, id *string) (res io.ReadCloser, err error) {
	if *id == "" {
		return nil, fmt.Errorf("no userId")
	}

	URL := os.Getenv("WRIKE_URL") + USERS + "/" + *id

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
