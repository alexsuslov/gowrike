package gowrike

import (
	"context"
	"io"
	"os"
)

var CONTACTS = "/contacts"

// ContactsRaw Contacts
func ContactsRaw(ctx context.Context, id *string) (res io.ReadCloser, err error) {
	URL := os.Getenv("WRIKE_URL") + CONTACTS
	if id != nil {
		URL = os.Getenv("WRIKE_URL") + CONTACTS + "/" + *id
	}

	body, _, err := Request(ctx, "GET", URL, nil, nil)
	return body, err
}
