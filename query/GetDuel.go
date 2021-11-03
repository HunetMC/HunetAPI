package query

import (
	// own packages
	"github.com/HunetMC/HunetAPI/database"
	
	// go get packages
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
)

type Duel struct {
	ID int32
	Winner string
	WinnerKit string
	Loser string
	LoserKit string
	Inventory string
	Arena string
	Start int64
	Finish int64
	Draw bool
}

// GetDuel "/duel/:id": Retrieve information about specific duel game
func GetDuel(id string) string {
	DB := database.GetDB()
	rows, err := DB.Query("SELECT * FROM `duels` WHERE `id` = ?;", id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	
	duel := Duel{}

	for rows.Next() {
		// Error handling when inventory column is null.
		var s sql.NullString
		err := rows.Scan(
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			&duel.Inventory,
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
		)
		if s.Valid {
			// Retrieving columns
			err = rows.Scan(
				&duel.ID,
				&duel.Winner,
				&duel.WinnerKit,
				&duel.Loser,
				&duel.LoserKit,
				&duel.Inventory,
				&duel.Arena,
				&duel.Start,
				&duel.Finish,
				&duel.Draw,
			)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// Retrieving columns
			err = rows.Scan(
				&duel.ID,
				&duel.Winner,
				&duel.WinnerKit,
				&duel.Loser,
				&duel.LoserKit,
				TrashScanner{},
				&duel.Arena,
				&duel.Start,
				&duel.Finish,
				&duel.Draw,
			)
			if err != nil {
				log.Fatal(err)
			}
		}
		
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	// Convert array "player" into JSON string
	var res string
	if len(duel.Winner) == 0 {
		res = "404"
	} else {
		j, _ := json.Marshal(duel)
		res = string(j)
	}
	
	return res
}