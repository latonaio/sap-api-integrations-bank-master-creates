package sap_api_output_formatter

type BankMaster struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	Bank          string `json:"bank_code"`
	Deleted       string `json:"deleted"`
}

type Bank struct {
	BankCountry         string `json:"BankCountry"`
	BankInternalID      string `json:"BankInternalID"`
	BankName            string `json:"BankName"`
	Region              string `json:"Region"`
	ShortStreetName     string `json:"ShortStreetName"`
	ShortCityName       string `json:"ShortCityName"`
	SWIFTCode           string `json:"SWIFTCode"`
	BankNetworkGrouping string `json:"BankNetworkGrouping"`
	IsMarkedForDeletion bool   `json:"IsMarkedForDeletion"`
	Bank                string `json:"Bank"`
	BankBranch          string `json:"BankBranch"`
	BankCategory        string `json:"BankCategory"`
}
