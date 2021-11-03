package database

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	// go get packages
	"log"
)

// TrashScanner: Discard unused columns in table
type TrashScanner struct{}
func (TrashScanner) Scan(interface{}) error {
	return nil
}

// GetPlayers "/players": List entire players
func GetPlayers() string {
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