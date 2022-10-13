package USPS

import "encoding/xml"

type USPSTrackFieldsDetail struct {
	EventTime       string `xml:"EventTime"`
	EventDate       string `xml:"EventDate"`
	Event           string `xml:"Event"`
	EventCity       string `xml:"EventCity"`
	EventState      string `xml:"EventState"`
	EventZIPCode    string `xml:"EventZIPCode"`
	EventCountry    string `xml:"EventCountry"`
	FirmName        string `xml:"FirmName"`
	Name            string `xml:"Name"`
	AuthorizedAgent string `xml:"AuthorizedAgent"`
	EventCode       string `xml:"EventCode"`
	GMT             string `xml:"GMT"`
	GMTOffset       string `xml:"GMTOffset"`
}

type USPSTrackFieldsInfo struct {
	ID               string                  `xml:"ID,attr"`
	Class            string                  `xml:"Class"`
	ClassOfMailCode  string                  `xml:"ClassOfMailCode"`
	DestinationCity  string                  `xml:"DestinationCity"`
	DestinationState string                  `xml:"DestinationState"`
	DestinationZip   string                  `xml:"DestinationZip"`
	EmailEnabled     bool                    `xml:"EmailEnabled"`
	KahalaIndicator  bool                    `xml:"KahalaIndicator"`
	MailTypeCode     string                  `xml:"MailTypeCode"`
	MPDATE           string                  `xml:"MPDATE"`
	MPSUFFIX         string                  `xml:"MPSUFFIX"`
	OnTime           string                  `xml:"OnTime"`
	OriginCity       string                  `xml:"OriginCity"`
	OriginState      string                  `xml:"OriginState"`
	OriginZip        string                  `xml:"OriginZip"`
	PodEnabled       bool                    `xml:"PodEnabled"`
	TPodEnabled      bool                    `xml:"TPodEnabled"`
	RestoreEnabled   bool                    `xml:"RestoreEnabled"`
	RramEnabled      bool                    `xml:"RramEnabled"`
	RreEnabled       bool                    `xml:"RreEnabled"`
	Service          string                  `xml:"Service"`
	ServiceTypeCode  string                  `xml:"ServiceTypeCode"`
	Status           string                  `xml:"Status"`
	StatusCategory   string                  `xml:"StatusCategory"`
	StatusSummary    string                  `xml:"StatusSummary"`
	TABLECODE        string                  `xml:"TABLECODE"`
	TrackDetail      []USPSTrackFieldsDetail `xml:"TrackDetail"`
}

type USPSTrackFieldsResponse struct {
	XMLName   xml.Name            `xml:"TrackResponse"`
	TrackInfo USPSTrackFieldsInfo `xml:"TrackInfo"`
}
