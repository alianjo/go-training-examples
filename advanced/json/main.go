package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var P = fmt.Println

type p struct {
	Page   int
	Fruits []string
}

type p2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Encoding
	ex1 := &p{Page: 1, Fruits: []string{"Apple", "Orange", "Peach"}}
	js, _ := json.Marshal(ex1)
	P(string(js))

	// Encoding
	ex2 := p2{Page: 2, Fruits: []string{"Orange", "peach"}}
	js, _ = json.Marshal(ex2)
	P(string(js))

	// Decoding and convert to usable values
	byt := []byte(`{"stars":7,"pages":[1,2,3,4]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		P("Error unmarshaling data", err)
	}
	P(dat)
	pages := dat["pages"].([]interface{})
	P(pages)

	stars := dat["stars"].(float64)
	P(stars)

	str := `{"page":1,"fruits":["orange","apple"]}`
	res := p2{}
	json.Unmarshal([]byte(str), &res)
	P(res.Fruits[0])

	// Encode to Stdout
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(map[string]int{"apple": 5, "lettuce": 7})
}
