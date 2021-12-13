package gowrike

import (
	"context"
	"encoding/json"
	"github.com/alexsuslov/gowrike/model"
	"io"
	"os"
)

var TASKS = "/tasks"

type TasksResponse struct {
	Kind string       `json:"kind"`
	Data []model.Task `json:"data"`
}

// TasksRaw Tasks Raw
func TasksRaw(ctx context.Context, query *string) (res io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + TASKS
	if *query != "" {
		URL = os.Getenv("WRIKE_URL") + TASKS + "/" + *query
	}
	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}

// TaskByID TaskByID
func TaskByID(ctx context.Context, query *string) (res TasksResponse, err error) {
	res = TasksResponse{}
	body, err := TasksRaw(ctx, query)
	if err != nil {
		return
	}
	defer body.Close()
	return res, json.NewDecoder(body).Decode(res)
}
