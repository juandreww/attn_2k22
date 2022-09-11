package controller

import (
	"net/http"
	// "encoding/json"
	// "2k22go/views"
	// "fmt"
)

func RegisterApi() *http.ServeMux {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/savecurrency", saveCurrency)
	mux.HandleFunc("/listcurrency", listCurrency)
	mux.HandleFunc("/listcurrencyrate", listCurrencyRate)
	mux.HandleFunc("/addcurrencyrate", addCurrencyRate)
	mux.HandleFunc("/convertcurrency", convertCurrency)
	return mux
}