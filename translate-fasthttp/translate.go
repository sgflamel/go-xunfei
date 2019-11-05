// translate
package translate

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
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
	params["data"] = map[string]string{"text": base64.StdEncoding.EncodeToString([]byte(content))}

	body, err := json.Marshal(&params)

	if err != nil {
		return nil, err
	}

	var request = fasthttp.AcquireRequest()

	c.assemblyRequestHeader(request, body)

	request.SetRequestURI(c.apiAddr)
	request.SetBody(body)

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}

	if err := client.Do(request, resp); err != nil {
		return nil, err
	}

	body = resp.Body()

	var result = &TranslateResponse{}

	err = json.Unmarshal(body, result)

	return result, err

}
