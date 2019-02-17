package endpoint

import "net/url"

type Endpoint struct {
	Destination string `json:"destination"`
	HMAC        bool   `json:"hmac"`
}

func ParsePath(rawpath string) (*Endpoint, error) {
	len := len(rawpath)

	_, err := url.ParseRequestURI(rawpath[1 : len-4])

	if err != nil {
		return nil, err
	}

	return &Endpoint{
		Destination: rawpath[1 : len-4],
		HMAC:        rawpath[len-4:len] == "hmac",
	}, nil
}
