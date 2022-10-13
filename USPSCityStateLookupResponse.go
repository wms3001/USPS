package USPS

import "encoding/xml"

type USPSCityStateZipCode struct {
	Zip5  string `xml:"Zip5"`
	City  string `xml:"City"`
	State string `xml:"State"`
}

type USPSCityStateLookupResponse struct {
	XMLName xml.Name             `xml:"CityStateLookupResponse"`
	ZipCode USPSCityStateZipCode `xml:"ZipCode"`
}
