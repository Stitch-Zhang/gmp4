package gmp4

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//RemoteVideo define a web video
type RemoteVideo struct {
	url     *url.URL
	timeout time.Duration
	header  http.Header
	mvhdBox
}

// NewRemote return a RemoteVideo object
func NewRemote(rawurl string) (*RemoteVideo, error) {
	if !strings.Contains(rawurl, "mp4") {
		return nil, fmt.Errorf("%s is invalid\n make sure the type of url resource is mp4 ", rawurl)
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	return &RemoteVideo{url: u}, nil
}

// SetHeader can manually set webClient cookie UA etc
func (r *RemoteVideo) SetHeader(h http.Header) {
	r.header = h
}

// SetTimeout can manually set webClient timeout
func (r *RemoteVideo) SetTimeout(time time.Duration) {
	r.timeout = time
}

func (r *RemoteVideo) getMvhdBox() mvhdBox {
	return r.mvhdBox
}

func (r *RemoteVideo) collectData() {
	client := &http.Client{Timeout: r.timeout}
	req, err := http.NewRequest("GET", r.url.String(), nil)
	if err != nil {
		fmt.Println("remoteFile.go:Generate request failed!")
	}
	if r.header != nil {
		req.Header = r.header
	}
	var (
		i    = 1
		skip = 0
	)
	for {
		//Split video to receive quickly
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", (i-1)*128+skip, i*128+skip))
		resp, err := client.Do(req)
		//Handle resource is not mp4
		if resp.Header.Get("Content-Type") != "video/mp4" {
			panic(r.url.String() + "%s is invalid resource")
		}
		if err != nil {
			fmt.Println("remoteFile.go:send request failed!", err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		switch n, cas := check(body); cas {
		case "mvhd":
			r.metadataIndex = n
			r.metadata = body
			return
		case "mdat":
			skip = int(byteSliceToInt(body[n-5 : n-1]))
		default:
			i++
		}

	}
}
