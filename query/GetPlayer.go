package query

import (
	// own packages
	"github.com/HunetMC/HunetAPI/database"
	
	// go get packages
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Player struct {
	UUID string
	Name string
	Last int32
	First int32
	Kill int32
	Death int32
}

// GetPlayer "/player/:uuid": Retrieve information about specific player
func GetPlayer(uuid string) string {
	DB := database.GetDB()
	rows, err := DB.Query("SELECT * FROM `players` WHERE `uuid` = ?;", uuid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	
	player := Player{}

	for rows.Next() {
		// Retrieving columns
		// Discard unused columns using TrashScanner.
		err := rows.Scan(
			&player.UUID,
			&player.Name,
			&player.Last,
			&player.First,
			&player.Kill,
			&player.Death,
		)
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
	if len(player.UUID) == 0 {
		res = "404"
	} else {
		j, _ := json.Marshal(player)
		res = string(j)
	}
	
	return res
}