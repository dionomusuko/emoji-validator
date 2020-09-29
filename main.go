package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type emojiJson struct {
	Codes string `json:"codes"`
}

func RunetoHex(str string) []string {
	runes := []rune(str)
	hexParts := []string{}
	for _, rune := range runes {
		hexParts = append(hexParts, fmt.Sprintf("%X", rune))
	}
	return hexParts
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
	array := RunetoHex("jgasrkfakf👨‍👨‍👧‍👧")
L:
	for _, ary := range array {
		if len(ary) < 4 {
			continue
		}
		fmt.Println(ary)
		for _, emoji := range emojiJsons {
			code := emoji.Codes
			if strings.Index(ary, code) == 0 {
				fmt.Println("emoji exists")
				break L
			}
		}
	}
}
