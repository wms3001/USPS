package USPS

type USPSAddressVaild struct {
	ID                   string `xml:"ID,attr"`
	Address1             string `xml:"Address1"`
	Address2             string `xml:"Address2"`
	City                 string `xml:"City"`
	CityAbbreviation     string `xml:"CityAbbreviation"`
	State                string `xml:"State"`
	Zip5                 string `xml:"Zip5"`
	Zip4                 string `xml:"Zip4"`
	DeliveryPoint        string `xml:"DeliveryPoint"`
	CarrierRoute         string `xml:"CarrierRoute"`
	Footnotes            string `xml:"Footnotes"`
	DPVConfirmation      string `xml:"DPVConfirmation"`
	DPVCMRA              string `xml:"DPVCMRA"`
	DPVFootnotes         string `xml:"DPVFootnotes"`
	Business             string `xml:"Business"`
	CentralDeliveryPoint string `xml:"CentralDeliveryPoint"`
	Vacant               string `xml:"Vacant"`
}

type USPSAddressVaildResponse struct {
	Address USPSAddressVaild `xml:"Address"`
}
