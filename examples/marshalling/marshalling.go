package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type creditCard struct {
	Number         string    `json:"number"`
	securityCode   string    `json:"-"`
	OwnerName      string    `json:"owner_name"`
	ExpirationDate time.Time `json:"expiration_date"`
}

func main() {
	cards := []creditCard{
		{
			Number:         "5627865283",
			securityCode:   "1255",
			OwnerName:      "Brutto",
			ExpirationDate: time.Now().AddDate(10, 5, 0),
		},
		{
			Number:         "4657154851",
			securityCode:   "7924",
			OwnerName:      "Alessia",
			ExpirationDate: time.Now().AddDate(10, 9, 0),
		},
		{
			Number:         "7845821241",
			securityCode:   "9713",
			OwnerName:      "Kebrage",
			ExpirationDate: time.Now().AddDate(7, 11, 0),
		},
	}

	bytesCards, err := json.MarshalIndent(cards, "", "\t")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("bytesCards: \n%s\n", bytesCards)

	jsonCreditCards := `
[
    {
        "number": "5627865283",
        "owner_name": "Brutto",
        "expiration_date": "2036-09-26T17:53:44.808803754+02:00"
    },
    {
        "number": "4657154851",
        "owner_name": "Alessia",
        "expiration_date": "2037-01-26T17:53:44.808963287+01:00"
    },
    {
        "number": "7845821241",
        "owner_name": "Kebrage",
        "expiration_date": "2034-03-26T17:53:44.808963531+02:00"
    }
]
	`

	var creditCardsFromJson []creditCard

	if err := json.Unmarshal([]byte(jsonCreditCards), &creditCardsFromJson); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Credit cards")

	for _, v := range creditCardsFromJson {
		fmt.Println(v)
	}
}
