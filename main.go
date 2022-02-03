package main

import (
	"encoding/json"
	"strings"
)

type user struct {
	FirstName string
	LastName  string
}

func main() {
	closure()
	run()
}

func closure() {
	for i := 0; i < 10_000; i++ {
		getString := func(name string) string {
			u := user{
				FirstName: strings.Split(name, " ")[0],
				LastName:  strings.Split(name, " ")[1],
			}

			str, err := json.Marshal(&u)
			if err != nil {
				return ""
			}
			return string(str)
		}

		getInt64 := func(name string) int64 {
			u := user{
				FirstName: strings.Split(name, " ")[0],
				LastName:  strings.Split(name, " ")[1],
			}

			str, err := json.Marshal(&u)
			if err != nil {
				return 0
			}
			return int64(len(str))
		}

		getString("First Last")
		getInt64("First Last")
	}
}

func run() {
	for i := 0; i < 10_000; i++ {
		getStringClean("First Last")
		getInt64Clean("First Last")
	}
}

func getStringClean(name string) string {
	u := user{
		FirstName: strings.Split(name, " ")[0],
		LastName:  strings.Split(name, " ")[1],
	}

	str, err := json.Marshal(&u)
	if err != nil {
		return ""
	}

	return string(str)
}

func getInt64Clean(name string) int64 {
	u := user{
		FirstName: strings.Split(name, " ")[0],
		LastName:  strings.Split(name, " ")[1],
	}

	str, err := json.Marshal(&u)
	if err != nil {
		return 0
	}

	return int64(len(str))
}
