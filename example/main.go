package main

import (
	"encoding/json"
	"log"
	"github.com/ieee0824/go-deepmerge"
	"fmt"
)

func main() {
	aJSON := `
	{
		"hoge": "huga",
		"array": [
			0,
			1
		]
	}`

	bJSON := `
	{
		"array": [
			1,
			2,
			3,
			4,
			{
				"foo": "bar",
				"john": "doe"
			},
			0.5,
			1.1,
			"fizz"
		]
	}
	`

	var a, b interface{}

	if err := json.Unmarshal([]byte(aJSON), &a); err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal([]byte(bJSON), &b); err != nil {
		log.Fatalln(err)
	}

	c, err := deepmerge.Merge(a, b)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(result))
}
