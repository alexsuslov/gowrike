package gowrike

import (
	"context"
	"io"
	"os"
	"strings"
)

var USERS = "/users"

func UsersRaw(ctx context.Context, ids ...string) (res io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + CONTACTS + "/" + strings.Join(ids, ",")

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
