package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	Url    string  `json:"url"`
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type image struct {
	Width     float64   `json:"width"`
	Height    float64   `json:"height"`
	Title     string    `json:"title"`
	Thumbnail thumbnail `json:"thumbnail"`
	Animated  bool      `json:"animated,omitempty"`
	Ids       []int64   `json:"ids"`
}

type test struct {
	MrBe []byte `json:"mrbe"`
}

func main() {
	var img image
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	e := json.Unmarshal([]byte(rcvd), &img)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(img)
}
