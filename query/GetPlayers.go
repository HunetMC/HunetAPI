package query

import (
	// own packages
	"github.com/HunetMC/HunetAPI/database"

	// go get packages
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// TrashScanner: Discard unused columns in table
type TrashScanner struct{}
func (TrashScanner) Scan(interface{}) error {
	return nil
}

// GetPlayers "/players": List entire players
func GetPlayers() string {
	DB := database.GetDB()
	rows, err := DB.Query("SELECT * FROM players;")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	
	var name string
	var players []string
	
	for rows.Next() {
		// Retrieving columns
		// Discard unused columns using TrashScanner.
		err := rows.Scan(
			TrashScanner{},
			&name,
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
			TrashScanner{},
		)
		if err != nil {
			log.Fatal(err)
		}
		
		// Add players into array
		players = append(players, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	// Convert array "players" into JSON string
	j, err := json.Marshal(players)
	return string(j)
}