package models

type APIResponse struct {
	BRLUSD struct {
		Bid string `json:"bid"`
	} `json:"BRLUSD"`
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
	BRLEUR struct {
		Bid string `json:"bid"`
	} `json:"BRLEUR"`
	EURBRL struct {
		Bid string `json:"bid"`
	} `json:"EURBRL"`
	BTCUSD struct {
		Bid string `json:"bid"`
	}
	BTCBRL struct {
		Bid string `json:"bid"`
	} `json:"BTCBRL"`
}
