package gowrike

import (
	"context"
	"fmt"
	"io"
	"os"
)

var TIMELOGS = "/timelogs"
var CONTACTS_TIMELOGS = "/contacts/%s/timelogs"
var FOLDERS_TIMELOGS = "/folders/%s/timelogs"
var TASKS_TIMELOGS = "/folders/%s/timelogs"
var TIMELOG_CATEGORIES_TIMELOGS = "/timelog_categories/%s/timelogs"

func TimelogsRaw(ctx context.Context, query *string) (res io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + TIMELOGS
	if query != nil && *query != "" {
		URL += "/" + *query
	}
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

func FoldersTimelogsRaw(ctx context.Context, folderId *string) (res io.ReadCloser, err error) {
	if *folderId == "" {
		return nil, fmt.Errorf("no folderId")
	}

	URL := os.Getenv("WRIKE_URL") + fmt.Sprintf(FOLDERS_TIMELOGS, *folderId)

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}

func TasksTimelogsRaw(ctx context.Context, taskId *string) (res io.ReadCloser, err error) {
	if *taskId == "" {
		return nil, fmt.Errorf("no taskId")
	}

	URL := os.Getenv("WRIKE_URL") + fmt.Sprintf(TASKS_TIMELOGS, *taskId)

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}

func CategoriesTimelogsRaw(ctx context.Context, categoryId *string) (res io.ReadCloser, err error) {
	if *categoryId == "" {
		return nil, fmt.Errorf("no categoryId")
	}

	URL := os.Getenv("WRIKE_URL") + fmt.Sprintf(TASKS_TIMELOGS, *categoryId)

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
