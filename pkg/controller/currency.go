package controller

import (
	"fmt"
	"github.com/juandreww/attn_2k22/pkg/model"
	"net/http"
	"html/template"
)



func listCurrency(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is listcurrency api: ", r.Method)
	var list []currency
	list, err := model.readListCurrency(list)
	tpl.ExecuteTemplate(w, "listcurrency.gohtml", list)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is index api: ", r.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func saveCurrency(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is savecurrency api: ", r.Method)

	data := currency{
		r.FormValue("id"),
		r.FormValue("name"),
	}

	fmt.Println(data)

	cur := currency{}
	IsExist := false
	sqlStatement := `SELECT id, name FROM currency WHERE (id = ? OR name = ?);`
	row := con.QueryRow(sqlStatement, data.ID, data.Name)
	err := row.Scan(&cur.ID, &cur.Name,)
	switch err {
	case sql.ErrNoRows:
		IsExist = false
	case nil:
		IsExist = true
	default:
	  	panic(err)
	}
	
	if IsExist == false {
		sqlStatement = `
			INSERT INTO currency (id, name)
			VALUES (?, ?)`
		_, err := con.Exec(sqlStatement, data.ID, data.Name)
		if err != nil {
			panic(err)
		}
	} else {
		data = currency{
			"error",
			"error",
		}
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}