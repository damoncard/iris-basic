package function

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Task struct {
	Name string	`json:"name"`
	Date string		`json:"date"`
	Done string		`json:"done"`
}

type Data struct {
	Admin bool	
	Ctask int
	Uctask int
	Tks []Task
}

func ReadJSON (admin bool) Data {
	act := getJSON()
	completed_task := 0
	uncompleted_task := 0

	for i := 0;  i < len(act); i++ {
		if act[i].Done == "Yes" {
			completed_task++
		} else {
			uncompleted_task++
		}
	}

	d := Data {
		admin, 
		completed_task,
		uncompleted_task,
		act,
	}

	return d
}

func ReadJSONbyID (id string) Task {
	act := getJSON()
	var d Task

	for i := 0 ; i < len(act) ; i++ {
		if act[i].Name == id {
			d.Name = act[i].Name
			d.Date = act[i].Date
			d.Done = act[i].Done
		}
	}

	return d
}

func EditJSON (t Task) {
	act := getJSON()

	for i := 0 ; i < len(act) ; i++ {
		if act[i].Name == t.Name {
			act[i].Name = t.Name
			act[i].Date = t.Date
			act[i].Done = t.Done
		}
	}
	writeJSON(act)
}

func AddJSON (t Task) {
	act := getJSON()
	act = append(act, t)
	writeJSON(act)
}

func DeleteTask (t string) {
	act := getJSON()
	var new_act  []Task

	for i := 0 ; i < len(act) ; i++ {
		if act[i].Name != t {
			new_act = append(new_act, act[i])
		}
	}
	writeJSON(new_act)
}

func getJSON () []Task {
	file, _ := ioutil.ReadFile("./data.json")
	var activity []Task
	json.Unmarshal(file, &activity)
	return activity
}

func writeJSON (act []Task) {
	act_json, _ := json.MarshalIndent(act, "", "	")
	ioutil.WriteFile("data.json", act_json, os.ModePerm)
}