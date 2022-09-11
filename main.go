package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"github.com/juandreww/attn_2k22/pkg/controller"
	"github.com/juandreww/attn_2k22/pkg/model"
	// "os"
)

func main() {
	mux := controller.RegisterApi()
	db := model.ConnectDB()
	defer db.Close()

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", mux)
}
