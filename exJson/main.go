package main

import (
	"encoding/json"
	"fmt"
)

type SessionInfo struct {
	SSNID string
	RMID  int32
	UID   int64
}

type JoinReq struct {
	SessionInfo
	Name  string
	Buyin int64
}

func main() {
	DataToText()
	fmt.Println()
	TextToData()
}

func DataToText() {
	data := JoinReq{
		SessionInfo: SessionInfo{
			SSNID: "ABCDE",
		},
		Name:  "John",
		Buyin: 1000,
	}

	info, _ := json.Marshal(data)

	fmt.Println(string(info))
}

func TextToData() {
	text := "{\"SSNID\":\"abcde\",\"RMID\":3,\"UID\":100,\"Name\":\"Kim\",\"Buyin\":2000}"
	data := JoinReq{}

	fmt.Println(text)
	json.Unmarshal([]byte(text), &data)
	fmt.Printf("%v\n", data)

	fmt.Println(data.SSNID, data.RMID, data.UID, data.Name, data.Buyin)
}
