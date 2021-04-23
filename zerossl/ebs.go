package zerossl

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

var UnableToRetrieveEBS = errors.New("unable to retrieve EAB (External Account Binding) from zerossl")

type EBS struct {
	Success bool   `json:"success"`
	Kid     string `json:"eab_kid""`
	HmacKey     string `json:"eab_hmac_key"`
}

func RetrieveEBS(apiKey string) (*EBS, error) {
	resp, err := http.PostForm("https://api.zerossl.com/acme/eab-credentials?access_key="+apiKey, url.Values{})
	if err != nil {
		return nil, err
	}
	ret := &EBS{}
	if err := json.NewDecoder(resp.Body).Decode(ret); err != nil {
		return nil, err
	}
	if !ret.Success {
		return nil, UnableToRetrieveEBS
	}
	return ret, nil
}
