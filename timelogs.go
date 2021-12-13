package gowrike

import (
	"context"
	"fmt"
	"io"
	"os"
)

var TIMELOGS = "/timelogs"
var CONTACTS_TIMELOGS = "/contacts/%s/timelogs"

func TimelogsRaw(ctx context.Context) (res io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + TIMELOGS
	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}

func ContactTimelogsRaw(ctx context.Context, contactId *string) (res io.ReadCloser, err error) {
	if *contactId == "" {
		return nil, fmt.Errorf("no contactId")
	}

	URL := os.Getenv("WRIKE_URL") + fmt.Sprintf(CONTACTS_TIMELOGS, *contactId)

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
