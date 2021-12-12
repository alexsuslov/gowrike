package gowrike

import (
	"context"
	"encoding/json"
	"github.com/alexsuslov/gowrike/model"
	"io"
	"net/url"
	"os"
)

var TASKS = "/tasks"

type TasksResponse struct {
	Kind string       `json:"kind"`
	Data []model.Task `json:"data"`
}

func TaskByIDRaw(ctx context.Context, id string) (res io.ReadCloser, err error) {

	u, err := url.Parse(os.Getenv("WRIKE_URL") + TASKS + "/" + id)
	if err != nil {
		return
	}
	body, _, err := Request(ctx, "GET", u.String(), nil, nil)
	if err != nil {
		return
	}

	return body, err
}

func TaskByID(ctx context.Context, id string) (res TasksResponse, err error) {
	res = TasksResponse{}
	body, err := TaskByIDRaw(ctx, id)
	if err != nil {
		return
	}
	defer body.Close()
	return res, json.NewDecoder(body).Decode(res)
}
