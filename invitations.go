package gowrike

import (
	"context"
	"io"
	"os"
)

var INVITATIONS = "/invitations"

func InvitationsRaw(ctx context.Context) (res io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + INVITATIONS

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
