package gowrike

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

/**

https://developers.wrike.com/api/v4/comments/

*/
var COMMENTS = "/comments"
var TASK_COMMENTS = "/tasks/%s/comments"

// CreateCommentRaw Create Comment Raw
func CreateCommentRaw(ctx context.Context, id string, req io.ReadCloser) (res io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + fmt.Sprintf(COMMENTS, id)

	body, _, err := Request(ctx, "POST", URL, req, nil)
	return body, err
}

//CreateComment Create Comment
func CreateComment(ctx context.Context, id string, text string) (res CreateCommentResponse, err error) {
	values := url.Values{}
	values.Add("plainText", "true")
	values.Add("text", text)
	data := values.Encode()
	r := io.NopCloser(strings.NewReader(data))
	body, err := CreateCommentRaw(ctx, id, r)
	if err != nil {
		return
	}
	defer body.Close()
	res = CreateCommentResponse{}
	return res, json.NewDecoder(body).Decode(&res)
}

//CommentResponse CommentResponse
type CommentResponse struct {
	ID            string   `json:"id"`
	AuthorID      string   `json:"authorId"`
	Text          string   `json:"text"`
	CreatedDate   string   `json:"createdDate"`
	UpdatedDate   string   `json:"updatedDate"`
	TaskID        string   `json:"taskId"`
	AttachmentIds []string `json:"attachmentIds"`
}

//CreateCommentResponse CreateCommentResponse
type CreateCommentResponse struct {
	Kind string            `json:"kind"`
	Data []CommentResponse `json:"data"`
}

func TaskCommentsRaw(ctx context.Context, taskId *string) (body io.ReadCloser, err error) {
	if taskId == nil && *taskId != "" {
		return nil, fmt.Errorf("no taskId")
	}
	URL := os.Getenv("WRIKE_URL") + fmt.Sprintf(TASK_COMMENTS, *taskId)

	body, _, err = Request(ctx, "GET", URL, nil, nil)
	return
}

func CommentsRaw(ctx context.Context) (body io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + COMMENTS
	body, _, err = Request(ctx, "GET", URL, nil, nil)
	return
}

func CommentsByID(ctx context.Context, id *string) (res CommentResponse, err error) {
	body, err := TaskCommentsRaw(ctx, id)
	if err != nil {
		return
	}
	res = CommentResponse{}
	return res, json.NewDecoder(body).Decode(&res)
}
