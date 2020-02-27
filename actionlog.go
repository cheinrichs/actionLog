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

//Stores averages for each action
// [
// 	{"Action": Avg},
// 	{"Action": Avg},
// ]
var averages = []actionAverage{}

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

	average := calculateAverage(e.Action)

	updateAverage(e.Action, average)

	return nil
}

//calculates the average based on all the times so far
func calculateAverage(key string) int {
	var currentAction = actionData[key]
	var sumOfTimes = 0

	for i := 0; i < len(actionData[key]); i++ {
		sumOfTimes += currentAction[i]
	}

	return sumOfTimes / len(currentAction)
}

//Updates the current average for a specific action
func updateAverage(action string, avg int) {
	//Update the average for the specific action
	for i := 0; i < len(averages); i++ {
		if averages[i].Action == action {
			averages[i].Avg = avg
			return
		}
	}

	//If an actionAverage record hasn't been added yet, add it
	newActionAverage := actionAverage{
		Action: action,
		Avg:    avg,
	}
	averages = append(averages, newActionAverage)
}

//GetStats Gets stats for all actions
func GetStats() string {
	averagesJSON, err := json.Marshal(averages)
	if err != nil {
		panic(err)
	}

	return string(averagesJSON)
}
