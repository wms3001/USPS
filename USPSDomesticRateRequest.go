package USPS

type USPSDomesticRateRequest struct {
	PackageId          string `json:"packageId"`
	Service            string `json:"service"`
	FirstClassMailType string `json:"firstClassMailType"`
	ZipOrigination     string `json:"zipOrigination"`
	ZipDestination     string `json:"zipDestination"`
	Pounds             string `json:"pounds"`
	Ounces             string `json:"ounces"`
	Size               string `json:"size"`
	Container          string `json:"container"`
	Width              string `json:"width"`
	Length             string `json:"length"`
	Height             string `json:"height"`
	Girth              string `json:"girth"`
	Machinable         string `json:"machinable"`
	ReturnLocations    string `json:"returnLocations"`
	ReturnServiceInfo  string `json:"returnServiceInfo"`
	ShipDate           string `json:"shipDate"`
	ReturnFees         string `json:"returnFees"`
}
