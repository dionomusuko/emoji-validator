package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type emojiJson struct {
	Codes string `json:"codes"`
}

func main() {
	req, err := http.NewRequest("GET", "https://unpkg.com/emoji.json@13.0.0/emoji.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var emojiJsons []emojiJson
	if err := json.Unmarshal(body, &emojiJsons); err != nil {
		log.Fatal(err)
	}
	fmt.Println(emojiJsons)
}
