package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

type SEND_PLAYER struct {
	RMID    int32
	SSNID   string
	PACKET  string
	Result  int
	Payload any
}

type SEND_ROOM struct {
	RMID    int32
	PACKET  string
	Payload any
}

type UserInfo struct {
	UID  int
	Name string
	Nick string
}

var que chan any

func main() {

	que = make(chan any, 10)

	go Recv()
	go Send()

	time.Sleep(10 * time.Second)
}

func Send() {
	i := 1
	for {
		time.Sleep(100 * time.Millisecond)
		que <- &SEND_PLAYER{SSNID: "abcde", Payload: UserInfo{UID: i, Name: "Kim"}}
		time.Sleep(100 * time.Millisecond)
		que <- &SEND_ROOM{RMID: 1}
		i++
	}
}

func Recv() {
	for {
		data := <-que
		packet, _ := json.Marshal(data)

		switch v := data.(type) {
		case *SEND_PLAYER:
			player := SEND_PLAYER{}
			json.Unmarshal(packet, &player)
			fmt.Printf("SEND_PLAYER: %+v ==> %s\n", player, string(packet))

			user := &UserInfo{}
			mapstructure.Decode(v.Payload, user)
			fmt.Printf("UserInfo: %+v\n", user)

		case *SEND_ROOM:
			room := SEND_ROOM{}
			json.Unmarshal(packet, &room)
			// fmt.Printf("SEND_ROOM:   %+v ==> %s\n", room, string(packet))
		default:
			fmt.Println("ERROR:", v)
		}
	}
}
