package gowrike

import (
	"context"
	"io"
	"os"
)

var WORKFLOWS = "/workflows"

func WorkflowsRaw(ctx context.Context) (res io.ReadCloser, err error) {

	URL := os.Getenv("WRIKE_URL") + WORKFLOWS

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
