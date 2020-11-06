package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

func (c Coupons) Check(code string) string {
	for _, item := range c.Coupon {
		if code == item.Code {
			return "valid"
		}
	}
	return "invalid"
}

type Result struct {
	Status string
}

var coupons Coupons

func main() {
	coupon := Coupon{
		Code: "abc",
	}

	coupons.Coupon = append(coupons.Coupon, coupon)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9092", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	valid := coupons.Check(coupon)

	result := Result{Status: valid}

	// Desafio
	service4Status := callService4("http://localhost:9093")
	if service4Status != "OK" {
		result.Status = service4Status
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error converting json")
	}

	fmt.Fprintf(w, string(jsonResult))

}

func callService4(urlMicroservice string) string {

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 2

	res, err := retryClient.Get(urlMicroservice)
	if err != nil {
		return "Servidor D fora do AR !"
	}
	defer res.Body.Close()

	return "OK"

}
