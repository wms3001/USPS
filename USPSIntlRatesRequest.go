package USPS

type USPSIntlRatesRequest struct {
	PackageId             string `json:"packageId"`
	Pounds                string `json:"pounds"`
	Ounces                string `json:"ounces"`
	Container             string `json:"container"`
	Width                 string `json:"width"`
	Length                string `json:"length"`
	Height                string `json:"height"`
	Girth                 string `json:"girth"`
	Machinable            string `json:"machinable"`
	MailType              string `json:"MailType"`
	ValueOfContents       string `json:"ValueOfContents"`
	Country               string `json:"Country"`
	OriginZip             string `json:"OriginZip"`
	CommercialFlag        string `json:"CommercialFlag"`
	AcceptanceDateTime    string `json:"AcceptanceDateTime"`
	DestinationPostalCode string `json:"DestinationPostalCode"`
	POBoxFlag             string `json:"POBoxFlag"`
	GiftFlag              string `json:"GiftFlag"`
}
