package USPS

import "encoding/xml"

type USPSSpecialService struct {
	ServiceID             string `xml:"ServiceID"`
	ServiceName           string `xml:"ServiceName"`
	Available             string `xml:"Available"`
	Price                 string `xml:"Price"`
	DeclaredValueRequired string `xml:"DeclaredValueRequired"`
}

type USPSSpecialServices struct {
	SpecialService []USPSSpecialService `xml:"SpecialService"`
}

type USPSPackagePostage struct {
	CLASSID         string              `xml:"CLASSID,attr"`
	MailService     string              `xml:"MailService"`
	Rate            string              `xml:"Rate"`
	CommitmentDate  string              `xml:"CommitmentDate"`
	CommitmentName  string              `xml:"CommitmentName"`
	SpecialServices USPSSpecialServices `xml:"SpecialServices"`
}

type USPSPackage struct {
	ID             string             `xml:"ID,attr"`
	ZipOrigination string             `xml:"ZipOrigination"`
	ZipDestination string             `xml:"ZipDestination"`
	Pounds         string             `xml:"Pounds"`
	Ounces         string             `xml:"Ounces"`
	Container      string             `xml:"Container"`
	Zone           string             `xml:"Zone"`
	Postage        USPSPackagePostage `xml:"Postage"`
}

type USPSDomesticRateResponse struct {
	XMLName xml.Name    `xml:"RateV4Response"`
	Package USPSPackage `xml:"Package"`
}
