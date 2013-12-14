package coinbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	CoinbaseApiKey         string
	CoinbaseCallbackSecret string
)

func apiKey() string {
	if CoinbaseApiKey == "" {
		CoinbaseApiKey = os.Getenv("COINBASE_API_KEY")
		if CoinbaseApiKey == "" {
			return ""
		}
	}

	return CoinbaseApiKey
}

// Documentation of the required params:
// https://coinbase.com/api/doc/1.0/buttons/create.html
type ButtonRequest struct {
	// Required
	Name             string `json:"name"`
	Type             string `json:"type,omitempty"`
	PriceString      string `json:"price_string"`
	PriceCurrencyIso string `json:"price_currency_iso"`
	// Optional
	Custom       string `json:"custom,omitempty"`
	CallbackUrl  string `json:"callback_url,omitempty"`
	Description  string `json:"description,omitempty"`
	Style        string `json:"style,omitempty"`
	Text         string `json:"text,omitempty"`
	IncludeEmail bool   `json:"include_email,omitempty"`

	VariablePrice bool `json:"variable_price,omitempty"`
	ChoosePrice   bool `json:"choose_price,omitempty"`

	Price1 string `json:"price1,omitempty"`
	Price2 string `json:"price2,omitempty"`
	Price3 string `json:"price3,omitempty"`
	Price4 string `json:"price4,omitempty"`
	Price5 string `json:"price5,omitempty"`
}

type ButtonResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors,omitempty"`
	Button  struct {
		Code        string `json:"code,omitempty"`
		Type        string `json:"type,omitempty"`
		Style       string `json:"style,omitempty"`
		Text        string `json:"text,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Custom      string `json:"custom,omitempty"`
		CallbackUrl string `json:"callback_url,omitempty"`
		Price       struct {
			Cents       uint   `json:"cents,omitempty"`
			CurrencyIso string `json:"currency_iso,omitempty"`
		}
	}
}

type Button struct {
	Request  *ButtonRequest
	Response *ButtonResponse
}

func (b *Button) ApiUrl() string {
	return fmt.Sprintf("https://coinbase.com/api/v1/buttons?api_key=%s", apiKey())
}

func GetButton(button_request *ButtonRequest) (b *Button) {
	b = &Button{
		Request: button_request,
	}

	js, _ := json.Marshal(b.Request)

	resp, err := http.Post(b.ApiUrl(), "application/json", bytes.NewBuffer(js))
	if err != nil {
		// handle error
		fmt.Printf("err")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &b.Response)

	fmt.Println(string(body))

	return b
}

func (b *Button) CoinbaseHtml(code, anchor string) string {
	if anchor == "" {
		anchor = "Pay with Bitcoin"
	}

	return fmt.Sprintf(`
<a class="coinbase-button" data-code="%s" href="https://coinbase.com/checkouts/%s">%s</a>
<script src="https://coinbase.com/assets/button.js" type="text/javascript"></script>"
`, code, code, anchor)
}
