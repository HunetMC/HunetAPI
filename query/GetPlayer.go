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

type Player struct {
	UUID string
	Name string
	Country string
	CurrentKit string
	KitList string
	Last int32
	First int32
	Kill int32
	Death int32
	Permission string
	Bio string
	Online string
}

type NullablePlayer struct {
	UUID string
	Name string
	Country string
	CurrentKit string
	KitList string
	Last int32
	First int32
	Kill int32
	Death int32
	Permission string
	Bio sql.NullString
	Online string
}

// GetPlayer "/player/:uuid": Retrieve information about specific player
func GetPlayer(uuid string) string {
	DB := database.GetDB()
	rows, err := DB.Query("SELECT * FROM `players` WHERE `uuid` = ?;", uuid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	
	p := Player{}
	
	for rows.Next() {
		// Constructor for containing player's data
		np := NullablePlayer{}
		// Retrieving columns
		if err := rows.Scan(
			&np.UUID,
			&np.Name,
			&np.Country,
			&np.CurrentKit,
			&np.KitList,
			&np.Last,
			&np.First,
			&np.Kill,
			&np.Death,
			&np.Permission,
			&np.Bio,
			&np.Online,
		); err != nil {
			log.Fatal(err)
		}
		
		// Convert nullable constructor to real value constructor
		p = Player {
			UUID: np.UUID,
			Name: np.Name,
			Country: np.Country,
			CurrentKit: np.CurrentKit,
			KitList: np.CurrentKit,
			Last: np.Last,
			First: np.First,
			Kill: np.Kill,
			Death: np.Death,
			Permission: np.Permission,
			Bio: np.Bio.String,
			Online: np.Online,
		}
	}
	
	// Convert constructor "p" into JSON string
	var res string
	if len(p.UUID) == 0 {
		res = "404"
	} else {
		j, _ := json.Marshal(p)
		res = string(j)
	}
	
	return res
}