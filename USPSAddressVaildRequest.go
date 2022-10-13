package USPS

type USPSAddressVaildRequest struct {
	AddressId string `json:"addressId"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip5      string `json:"zip5"`
	Zip4      string `json:"zip4"`
}
