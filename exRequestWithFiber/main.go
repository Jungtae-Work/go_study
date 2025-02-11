package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
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

	cc := fiber.AcquireClient()

	agent := cc.Post("http://localhost:8080/buyin").
		JSON(RequestData{
			UID:    1001,
			RoomID: "20250211-161030-1234",
			Buyin:  100000})

	code, body, err := agent.Bytes()
	if err != nil {
		log.Fatal(err)
	}

	var res ResponseData
	if err := json.Unmarshal(body, &res); err != nil {
		log.Fatal(err)
	}

	fmt.Println(code)
	fmt.Printf("%+v\n", res)

	fiber.ReleaseClient(cc)
}
