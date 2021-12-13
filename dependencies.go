package gowrike

import (
	"context"
	"fmt"
	"io"
	"os"
)

var DEPENDENCIES = "/dependencies"
var TASK_DEPENDENCIES = "/tasks/%s/dependencies"

func DependenciesRaw(ctx context.Context, query *string) (res io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + DEPENDENCIES
	if *query != "" {
		URL = os.Getenv("WRIKE_URL") + TASKS + "/" + *query
	}
	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}

func TaskDependenciesRaw(ctx context.Context, query *string) (res io.ReadCloser, err error) {

	if query == nil || *query == "" {
		return nil, fmt.Errorf("no dependencyId")
	}
	URL := os.Getenv("WRIKE_URL") + fmt.Sprintf(TASK_DEPENDENCIES, *query)

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
