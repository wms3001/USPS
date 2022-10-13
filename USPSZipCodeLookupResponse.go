package USPS

import "encoding/xml"

type USPSZipCodeLookupAddress struct {
	ID           string `xml:"ID,attr"`
	FirmName     string `xml:"FirmName"`
	Address2     string `xml:"Address2"`
	Address1     string `xml:"Address1"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	Urbanization string `xml:"Urbanization"`
	Zip5         string `xml:"Zip5"`
	Zip4         string `xml:"Zip4"`
}

type USPSZipCodeLookupResponse struct {
	XMLName xml.Name                 `xml:"ZipCodeLookupResponse"`
	Address USPSZipCodeLookupAddress `xml:"Address"`
}
