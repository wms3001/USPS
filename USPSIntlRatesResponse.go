package USPS

import "encoding/xml"

type USPSPackageExtraServices struct {
	ExtraService []USPSSpecialService `xml:"ExtraService"`
}

type USPSPackagePostOffice struct {
	Name                string `xml:"Name"`
	Address             string `xml:"Address"`
	City                string `xml:"City"`
	State               string `xml:"State"`
	ZipCode             string `xml:"ZipCode"`
	RetailGXGCutOffTime string `xml:"RetailGXGCutOffTime"`
	SaturDayCutOffTime  string `xml:"SaturDayCutOffTime"`
}

type USPSPackageGXGLocations struct {
	PostOffice USPSPackagePostOffice `xml:"PostOffice"`
}

type USPSIntlPackageService struct {
	ID              string                   `xml:"ID,attr"`
	Pounds          string                   `xml:"Pounds"`
	Ounces          string                   `xml:"Ounces"`
	Container       string                   `xml:"Container"`
	Zone            string                   `xml:"Zone"`
	Postage         USPSPackagePostage       `xml:"Postage"`
	Machinable      string                   `xml:"Machinable"`
	MailType        string                   `xml:"MailType"`
	Width           string                   `xml:"Width"`
	Length          string                   `xml:"Length"`
	Height          string                   `xml:"Height"`
	Girth           string                   `xml:"Girth"`
	Country         string                   `xml:"Country"`
	ExtraServices   USPSPackageExtraServices `xml:"ExtraServices"`
	ValueOfContents string                   `xml:"ValueOfContents"`
	SvcCommitments  string                   `xml:"SvcCommitments"`
	SvcDescription  string                   `xml:"SvcDescription"`
	MaxDimensions   string                   `xml:"MaxDimensions"`
	MaxWeight       string                   `xml:"MaxWeight"`
	GXGLocations    USPSPackageGXGLocations  `xml:"GXGLocations"`
}

type USPSIntlPackage struct {
	ID                     string                   `xml:"ID,attr"`
	Prohibitions           string                   `xml:"Prohibitions"`
	Restrictions           string                   `xml:"Restrictions"`
	Observations           string                   `xml:"Observations"`
	CustomsForms           string                   `xml:"CustomsForms"`
	ExpressMail            string                   `xml:"ExpressMail""`
	AreasServed            string                   `xml:"AreasServed"`
	AdditionalRestrictions string                   `xml:"AdditionalRestrictions"`
	Service                []USPSIntlPackageService `xml:"Service"`
}

type USPSIntlRatesResponse struct {
	XMLName xml.Name        `xml:"IntlRateV2Response"`
	Package USPSIntlPackage `xml:"Package"`
}
