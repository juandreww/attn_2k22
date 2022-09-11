package model

import (
	// "fmt"
	"log"
	"database/sql"
	"github.com/juandreww/attn_2k22/pkg/views"
)

func readListCurrency(list []views.Currency) ([]views.Currency, error) {
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
			return nil, err
		case nil:
		default:
			return nil, err
		}

		list = append(list, cur)
	}

	return list, nil
}