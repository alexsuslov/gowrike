package gowrike

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var DEBUG = false

// Request Request
func Request(ctx context.Context, method string, url string, reader io.ReadCloser,
	head map[string]string) (body io.ReadCloser, header http.Header, err error) {

	if DEBUG {
		log.Println("URL", url)
	}

	InsecureSkipVerify := os.Getenv("INSECURE") == "YES"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: InsecureSkipVerify},
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return
	}

	//header
	if head == nil {

		req.Header.Add("Authorization", "bearer "+os.Getenv("WRIKE_TOKEN"))
	} else {
		for k, v := range head {
			req.Header.Set(k, v)
		}
	}
	if DEBUG {
		log.Println("Header", req.Header)
	}

	client := &http.Client{Transport: tr}
	r, err := client.Do(req)
	if err != nil {
		return
	}

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		data, _ := ioutil.ReadAll(r.Body)
		err = fmt.Errorf("%v:%v", r.StatusCode, string(data))
		return
	}

	return r.Body, r.Header, err
}
