package main

import (
	"errors"
	"fmt"
)

func main() {
	//-------------------------------------------------------------------------------------------------
	//-----------------------------------------------Maps-----------------------------------------------
	//-------------------------------------------------------------------------------------------------

	ages := make(map[string]int)
	ages["Thaer"] = 26
	fmt.Println(ages)

	ages_v2 := map[string]int{
		"Thaer": 27,
	}
	fmt.Println(ages_v2)

	names := []string{"Thaer", "Thaer2", "Thaer3"}
	phoneNumbers := []string{"+96171700996", "+96171700997", "+96171700998"}
	usersMap, _ := getUserMap(names, phoneNumbers)
	fmt.Println(usersMap["Thaer"])

	//Inserting an element into a map:
	usersMap["Lara"] = user{
		name:        "Lara",
		phoneNumber: "+96171700995",
	}
	fmt.Println(usersMap["Lara"])

	//Getting an element
	userLara := usersMap["Lara"]
	fmt.Println(userLara)

	//Deleting an element
	delete(usersMap, "Lara")
	fmt.Println(usersMap)

	_, ok := usersMap["Thaer"]
	fmt.Println(ok)

	//Shorthanded syntax for the above check:
	if _, ok := usersMap["Lara"]; !ok {
		fmt.Println("Lara not found")
	}

	//The key has to be a comparabletype, you can't use slices, maps, and functions since they point to addresses in memory
	//And you can't compare addresses, you need to compare values
	//structs are comparable
}

//-------------------------------------------------------------------------------------------------
//-----------------------------------------------Maps-----------------------------------------------
//-------------------------------------------------------------------------------------------------

func getUserMap(names []string, phoneNumbers []string) (map[string]user, error) {
	userMap := map[string]user{}
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber string
}
