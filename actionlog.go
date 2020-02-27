package actionlog

import "encoding/json"

type action struct {
	Action string `json:"action"`
	Time   int    `json:"time"`
}

type actionAverage struct {
	Action string `json:"action"`
	Avg    int    `json:"avg"`
}

//Stores each time for each action
//{
// Run: [time1, time2]
// Jump: [time1, time2]
//}
var actionData = make(map[string][]int)

//AddAction Adds an action
func AddAction(actionString string) error {
	bytes := []byte(actionString)

	var e action

	err := json.Unmarshal(bytes, &e)
	if err != nil {
		return err
	}

	if actionData[e.Action] != nil {
		actionData[e.Action] = append(actionData[e.Action], e.Time)
	} else {
		actionData[e.Action] = []int{e.Time}
	}

	return nil
}

//GetStats Gets stats for all actions
func GetStats() string {
	return ""
}
