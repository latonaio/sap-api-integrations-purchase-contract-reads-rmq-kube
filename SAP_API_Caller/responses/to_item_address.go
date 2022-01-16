package responses

type ToItemAddress struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseContract           string `json:"PurchaseContract"`
			AddressID                  string `json:"AddressID"`
			PurchaseContractItem       string `json:"PurchaseContractItem"`
			CityName                   string `json:"CityName"`
			PostalCode                 string `json:"PostalCode"`
			StreetName                 string `json:"StreetName"`
			Country                    string `json:"Country"`
			CorrespondenceLanguage     string `json:"CorrespondenceLanguage"`
			Region                     string `json:"Region"`
			PhoneNumber                string `json:"PhoneNumber"`
			FaxNumber                  string `json:"FaxNumber"`
		} `json:"results"`
	} `json:"d"`
}
