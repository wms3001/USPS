package USPS

import "encoding/xml"

type USPSErrorResponse struct {
	XMLName     xml.Name `xml:"Error"`
	Number      string   `xml:"Number"`
	Source      string   `xml:"Source"`
	Description string   `xml:"Description"`
	HelpFile    string   `xml:"HelpFile"`
	HelpContext string   `xml:"HelpContext"`
}
