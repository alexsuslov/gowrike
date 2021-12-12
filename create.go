package gowrike

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexsuslov/gowrike/model"
	"io"
	"io/ioutil"
)

var CREATE = "https://www.wrike.com/api/v4/folders/%s/tasks"

func CreateRaw(ctx context.Context, folderID string,
	req io.ReadCloser) (resp io.ReadCloser, err error) {
	URL := fmt.Sprintf(CREATE, folderID)

	body, _, err := Request(ctx, "POST", URL, req, nil)
	if err != nil {
		return
	}

	return body, err
}

type CreateResponse struct {
	Kind string       `json:"kind"`
	Data []model.Task `json:"data"`
}

func Create(ctx context.Context, folderID string,
	req model.CreateTicket) (res CreateResponse, err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}

	r := ioutil.NopCloser(bytes.NewReader(data))

	body, err := CreateRaw(ctx, folderID, r)
	if err != nil {
		return
	}

	defer body.Close()

	res = CreateResponse{}
	return res, json.NewDecoder(body).Decode(&res)
}
