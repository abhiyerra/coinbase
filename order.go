package coinbase

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	OrderCompleted = "completed"
	OrderCanceled  = "canceled"
)

type Order struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`

	TotalBtc struct {
		Cents       int    `json:"cents"`
		CurrencyIso string `json:"currency_iso"`
	} `json:"total_btc"`

	TotalNative struct {
		Cents       int    `json:"cents"`
		CurrencyIso string `json:"currency_iso"`
	} `json:"total_native"`

	Custom         string `json:"custom"`
	ReceiveAddress string `json:"receive_address"`

	Button struct {
		Type       string `json:"type"`
		Name       string `json:"name"`
		Decription string `json:"description"`
		Id         string `json:"id"`
	} `json:"button"`

	Transaction struct {
		Id           string `json:"id"`
		Hash         string `json:"hash"`
		Configations int    `json:"confirmations"`
	} `json:"transaction"`

	Customer struct {
		Email           string   `json:"email"`
		ShippingAddress []string `json:"shipping_address"`
	} `json:"customer"`
}

type Callback struct {
	Order Order `json:"order"`
}

func ParseCallback(callback string) Order {
	c := Callback{}
	err := json.Unmarshal([]byte(callback), &c)
	fmt.Printf("%v", c.Order)
	if err != nil {
		log.Println(err)
		return Order{}
	}
	return c.Order
}
