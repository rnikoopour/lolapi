package lolAPI

type LoLAPI struct {
	region      string
	riotToken   string
	SummonerAPI *SummonerAPI
}

func NewLolAPI(region, riotToken string) *LoLAPI {
	return &LoLAPI{
		region:      region,
		riotToken:   riotToken,
		SummonerAPI: NewSummonerAPI(region, riotToken)}
}
