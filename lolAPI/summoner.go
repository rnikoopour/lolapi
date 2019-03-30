package lolAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Summoner struct {
	ProfileIconId int    `json:"profileIconId"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
}

type SummonerAPI struct {
	httpClient *http.Client
	region     string
	riotToken  string
}

func NewSummonerAPI(region, riotToken string) *SummonerAPI {
	return &SummonerAPI{
		httpClient: &http.Client{},
		region:     region,
		riotToken:  riotToken}
}

func (s *SummonerAPI) SendAPIRequest(apiMethod, apiArgument string) *Summoner {
	var url string

	switch apiMethod {
	case "by-id":
		// Looking up by summoner ID is not a specific method of the API
		//  unlike the other lookup methods
		url = fmt.Sprintf("https://%v.api.riotgames.com/lol/summoner/v4/summoners/%v", s.region, apiArgument)
	default:
		url = fmt.Sprintf("https://%v.api.riotgames.com/lol/summoner/v4/summoners/%v/%v", s.region, apiMethod, apiArgument)
	}

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("X-Riot-Token", s.riotToken)
	res, err := s.httpClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	summoner := &Summoner{}
	if err := json.Unmarshal(body, summoner); err != nil {
		log.Fatal(err)
	}

	return summoner
}

func (s *SummonerAPI) ByName(summonerName string) *Summoner {
	return s.SendAPIRequest("by-name", summonerName)
}

func (s *SummonerAPI) ByAccount(accountID string) *Summoner {
	return s.SendAPIRequest("by-account", accountID)
}

func (s *SummonerAPI) ByPuuid(puuid string) *Summoner {
	return s.SendAPIRequest("by-puuid", puuid)
}

func (s *SummonerAPI) ByID(summonerID string) *Summoner {
	return s.SendAPIRequest("by-id", summonerID)
}
