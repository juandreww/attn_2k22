package controller

func listCurrencyRate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is listcurrencyrate api: ", r.Method)
	var list []configconvertrate

	rows, err := con.Query("SELECT cf.name currencyfrom, ct.name currencyto, round(p.rate,2) rate FROM currencyrate p LEFT JOIN currency cf ON cf.id = p.currencyfrom LEFT JOIN currency ct ON ct.id = p.currencyto ORDER BY p.id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		cur := configconvertrate{}
		err := rows.Scan(&cur.CurrencyFrom, &cur.CurrencyTo, &cur.Rate)

		IsError := HandleErrorOfSelect(w, err)
		if (IsError == true) {
			fmt.Println("row is not exist")
			return
		}

		list = append(list, cur)
	}

	tpl.ExecuteTemplate(w, "listcurrencyrate.gohtml", list)
}

func addCurrencyRate(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPost) {
		fmt.Println("This is add conversion rate api: ", r.Method)
		data := configconvertrate{
			r.FormValue("currencyfrom"),
			r.FormValue("currencyto"),
			r.FormValue("rate"),
		}

		var str string
		var intval int
		check1 := currency{}

		sqlStatement := `SELECT count(id) id FROM currency WHERE (id = ? OR id = ?);`
		row := con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo)
		err := row.Scan(&check1.ID)
		IsError := HandleErrorOfSelect(w, err)
		if (IsError == true) {
			tmp := currency{
				"error",
				"All CurrencyID is not found in database",
			}
			tpl.ExecuteTemplate(w, "addcurrencyrate.gohtml", tmp)
			return
		} else {
			intval, err = strconv.Atoi(check1.ID)
			if intval < 2 {
				tmp := currency{
					"error",
					"One of the CurrencyID is not found in database",
				}
				tpl.ExecuteTemplate(w, "addcurrencyrate.gohtml", tmp)
				return
			}
		}
		
		sqlStatement = `SELECT count(id) id FROM currencyrate 
						WHERE ((currencyfrom = ? AND currencyto = ?) OR (currencyfrom = ? AND currencyto = ?))`
		row = con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo, data.CurrencyTo, data.CurrencyFrom)
		err = row.Scan(&str)
		IsError = HandleErrorOfSelect(w, err)
		if (IsError == true) {
			if str != "0" {
				tmp := currency{
					"error",
					"CurrencyRate is not exist in the database",
				}
				tpl.ExecuteTemplate(w, "addcurrencyrate.gohtml", tmp)
				return
			}
		} else {
			intval, err = strconv.Atoi(str)
			if intval >  0 {
				tmp := currency{
					"error",
					"CurrencyRate already exist in the database",
				}
				tpl.ExecuteTemplate(w, "addcurrencyrate.gohtml", tmp)
				return
			}
		}

		sqlStatement = `SELECT nullif(max(id),0) id FROM currencyrate`
		row = con.QueryRow(sqlStatement)
		err = row.Scan(&str)
		IsError = HandleErrorOfSelect(w, err)
		if (IsError == true) {
			str = "0"
		}

		intval, err = strconv.Atoi(str)
		intval++
		sqlStatement = `
			INSERT INTO currencyrate (id, currencyfrom, currencyto, rate)
			VALUES (?, ?, ?, ?)`
		_, err = con.Exec(sqlStatement, intval, data.CurrencyFrom, data.CurrencyTo, data.Rate)
		if err != nil {
			panic(err)
		}
		
		tmp := currency{
			"succeed",
			"Currency Rate added",
		}
		tpl.ExecuteTemplate(w, "addcurrencyrate.gohtml", tmp)
	} else {
		fmt.Println("This is add conversion rate api: ", r.Method)

		tpl.ExecuteTemplate(w, "addcurrencyrate.gohtml", nil)
	}
}

func convertCurrency(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPost) {
		fmt.Println("This is add convertcurrency api: ", r.Method)
		data := configconvertrate{
			r.FormValue("currencyfrom"),
			r.FormValue("currencyto"),
			r.FormValue("amount"),
		}
		fmt.Println(data)

		var str string
		var intval int
		var floatval, amount float64
		check1 := currency{}

		sqlStatement := `SELECT count(id) id FROM currency WHERE (id = ? OR id = ?);`
		row := con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo)
		err := row.Scan(&check1.ID)
		IsError := HandleErrorOfSelect(w, err)
		if (IsError == true) {
			tmp := currency{
				"error",
				"All CurrencyID is not found in database",
			}
			tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
			return
		} else {
			intval, err = strconv.Atoi(check1.ID)
			if intval < 2 {
				tmp := currency{
					"error",
					"One of the CurrencyID is not found in database",
				}
				tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
				return
			}
		}
		
		sqlStatement = `SELECT count(id) id FROM currencyrate 
						WHERE ((currencyfrom = ? AND currencyto = ?) OR (currencyfrom = ? AND currencyto = ?))`
		row = con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo, data.CurrencyTo, data.CurrencyFrom)
		err = row.Scan(&str)
		IsError = HandleErrorOfSelect(w, err)
		if (IsError == true) {
			if str != "0" {
				tmp := currency{
					"error",
					"CurrencyRate is not exist in the database (1)",
				}
				tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
				return
			}
		}

		var val1, val2 string
		sqlStatement = `SELECT currencyfrom, currencyto,rate FROM currencyrate 
					WHERE ((currencyfrom = ? AND currencyto = ?) OR (currencyfrom = ? AND currencyto = ?))
					LIMIT 1`
		row = con.QueryRow(sqlStatement, data.CurrencyFrom, data.CurrencyTo, data.CurrencyTo, data.CurrencyFrom)
		err = row.Scan(&val1, &val2, &str)
		IsError = HandleErrorOfSelect(w, err)
		if (IsError == true) {
			tmp := currency{
				"error",
				"CurrencyRate is not exist in the database (2)",
			}
			tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
			return 
		}

		floatval, err = strconv.ParseFloat(str, 64)
		amount, err = strconv.ParseFloat(data.Rate, 64)
		if data.CurrencyFrom == val1 {
			amount = amount * floatval
		} else {
			amount = amount / floatval
		}
		
		str = fmt.Sprintf("%.2f", amount)
		
		tmp := currency{
			"succeed",
			"Congrats! You converted rate to " + str,
		}
		tpl.ExecuteTemplate(w, "convertcurrency.gohtml", tmp)
	} else {
		fmt.Println("This is add convertcurrency api: ", r.Method)

		tpl.ExecuteTemplate(w, "convertcurrency.gohtml", nil)
	}
}