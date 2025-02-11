package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestData struct {
	UID    int64  `json:"uid"`
	RoomID string `json:"room_id"`
	Buyin  int64  `json:"buyin"`
}

type ResponseData struct {
	Result     int    `json:"result"`
	AgencyCode int    `json:"agency_code"`
	AgencyName string `json:"agency_name"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Buyin      int64  `json:"buyin"`
}

func main() {
	reqData := RequestData{
		UID:    1001,
		RoomID: "room_1",
		Buyin:  1000,
	}

	rawData, _ := json.Marshal(reqData)
	buff := bytes.NewBuffer(rawData)
	resp, err := http.Post("http://localhost:8080/buyin", "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var respData ResponseData
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
	fmt.Printf("%+v\n", respData)
}
