package model

import (
	"fmt"
	"log"
	"github.com/juandreww/attn_2k22/pkg/views"
	"database/sql"
)

func readListCurrency(list []views.ConfigConvertRate) ([]views.ConfigConvertRate, error) {
	rows, err := con.Query("SELECT id, name FROM currency ORDER BY id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		cur := []views.Currency{}
		err := rows.Scan(&cur.ID, &cur.Name,)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("row is not exist")
			return
		case nil:
		default:
			panic(err)
		}

		list = append(list, cur)
	}

	return list, nil
}