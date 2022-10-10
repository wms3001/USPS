package USPS

import "encoding/xml"

type USPSTrackInfo struct {
	ID           string   `xml:"ID,attr"`
	TrackSummary string   `xml:"TrackSummary"`
	TrackDetail  []string `xml:"TrackDetail"`
}

type USPSTrackResponse struct {
	XMLName   xml.Name      `xml:"TrackResponse"`
	TrackInfo USPSTrackInfo `xml:"TrackInfo"`
}
