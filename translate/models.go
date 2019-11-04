// models
package translate

const (
	LanguageChinese = "cn"
	LanguageEnglish = "en"
)

type TranslateClient struct {
	key     string
	secret  string
	apiAddr string
	appid   string
}

type TranslateResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Sid     string        `json:"sid"` // 请求的会话ID
	Data    *ResponseData `json:"data"`
}

type ResponseData struct {
	Result *ResponseResult `json:"result"`
}

type ResponseResult struct {
	From        string       `json:"from"` // 源语言语种
	To          string       `json:"to"`   // 目标语言语种
	TransResult *TransResult `json:"trans_result"`
}

type TransResult struct {
	Dst string `json:"dst"` // 翻译结果
	Src string `json:"src"` // 翻译源
}
