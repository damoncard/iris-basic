package function

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Token struct {
	User string		`json:"user"`
	Pass string		`json:"pass"`
	Admin string	`json:"admin"`
}

func Check(a, b string) (bool, string) {
	users := getUserJSON()
	var t string
	lower_a := strings.ToLower(a)

	for i := 0;  i < len(users); i++ {
		if users[i].User == lower_a && users[i].Pass == b {
			if users[i].Admin == "yes" {
				t = "admin"
			} else {
				t = "guest"
			}
			return true, t
		}
	}

	return false, ""
}

func getUserJSON() []Token {
	file, _ := ioutil.ReadFile("./user.json")
	var users []Token
	json.Unmarshal(file, &users)
	return users
}

func writeUserJSON(users []Token) {
	users_json, _ := json.MarshalIndent(users, "", "	")
	ioutil.WriteFile("user.json", users_json, os.ModePerm)
}