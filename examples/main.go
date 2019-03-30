package main

import (
	"fmt"
	"lolapi/lolAPI"
	"os"
)

type Region = string

const (
	BR   Region = "BR1"
	EUNE Region = "EUN1"
	EUW  Region = "EUW1"
	JP   Region = "JP1"
	KR   Region = "KR"
	LAN  Region = "LAN1"
	LAS  Region = "LAS2"
	NA   Region = "NA1"
	OCE  Region = "OC1"
	PBE  Region = "PBE1"
	RU   Region = "RU"
	TR   Region = "TR1"
)

func main() {
	key := os.Getenv("API_KEY")
	lolAPI := lolAPI.NewLolAPI(NA, key)
	summoner := lolAPI.SummonerAPI.ByName("deadheartsbeat")
	fmt.Println(lolAPI.SummonerAPI.ByID(summoner.ID))
}
