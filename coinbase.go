package coinbase

var (
	api_key string
)

type ButtonRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	PriceString string `json:"price_string"`
	PriceCurrencyIso string `json:"price_currency_iso"`
	Custom string `json:"custom"`
	CallbackUrl string `json:"callback_url"`
	Description string `json:"description"`
	Style string `json:"style"`
	IncludeEmail bool `json:"include_email"`

	/*
        "name": "test",
        "type": "buy_now",
        "price_string": "1.23",
        "price_currency_iso": "USD",
        "custom": "Order123",
        "callback_url": "http://www.example.com/my_custom_button_callback",
        "description": "Sample description",
        "type": "buy_now",
        "style": "custom_large",
        "include_email": true
*/

}

type ButtonResponse struct {
	
}

type Button struct {
	request ButtonRequest
	response ButtonResponse
}

func NewButton(button_request ButtonRequest) (b *Button) {
	b = &Button{}
	b.request = button_request
}
