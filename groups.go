package gowrike

import (
	"context"
	"io"
	"os"
)

var GROUPS = "/groups"

func GroupsRaw(ctx context.Context, id *string) (res io.ReadCloser, err error) {

	URL := os.Getenv("WRIKE_URL") + GROUPS

	if id != nil || *id != "" {
		URL = os.Getenv("WRIKE_URL") + GROUPS + "/" + *id
	}

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
