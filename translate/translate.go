// translate
package translate

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	defaultApiAddr = "https://ntrans.xfyun.cn/v2/ots"
)

var (
	Client *TranslateClient
)

func InitWithApiKey(appid string, key string, secret string, addr string) error {
	if appid == "" || key == "" || secret == "" {
		return errors.New("miss appid or key or secret")
	}

	if addr == "" {
		addr = defaultApiAddr
	}

	Client = &TranslateClient{
		appid:   appid,
		key:     key,
		secret:  secret,
		apiAddr: addr,
	}

	return nil
}

func (c *TranslateClient) Translate(fromLan string, toLan string, content string, host, uri string) (*TranslateResponse, error) {
	var params = make(map[string]interface{}, 0)
	params["common"] = map[string]string{"app_id": c.appid}
	params["business"] = map[string]string{"from": fromLan, "to": toLan}
	params["data"] = map[string]string{"text": content}

	body, err := json.Marshal(&params)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", c.apiAddr, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	c.assemblyRequestHeader(request, host, uri, body)

	client := &http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result = &TranslateResponse{}

	err = json.Unmarshal(body, result)

	return result, err

}
