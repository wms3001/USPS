package USPS

import (
	"log"
	"testing"
)

func TestUSPS_Track(t *testing.T) {
	usps := &USPS{}
	usps.UserId = ""
	resp := usps.Track("93055109246000000114188641")
	log.Println(resp)
}

func TestUSPS_TrackFields(t *testing.T) {
	usps := &USPS{}
	usps.UserId = ""
	resp := usps.TrackFields("93055109246000000114188641")
	log.Println(resp)
}

func TestUSPS_DomesticRates(t *testing.T) {
	usps := &USPS{}
	usps.UserId = ""
	uSPSDomesticRateRequest := USPSDomesticRateRequest{}
	uSPSDomesticRateRequest.PackageId = "0"
	uSPSDomesticRateRequest.Service = "PRIORITY"
	uSPSDomesticRateRequest.ZipOrigination = "22201"
	uSPSDomesticRateRequest.ZipDestination = "26301"
	uSPSDomesticRateRequest.Pounds = "8"
	uSPSDomesticRateRequest.Ounces = "2"
	uSPSDomesticRateRequest.Machinable = "true"
	uSPSDomesticRateRequest.ShipDate = "2022-10-12"
	resp := usps.DomesticRates(uSPSDomesticRateRequest)
	log.Println(resp)
}

func TestUSPS_IntlRates(t *testing.T) {
	usps := &USPS{}
	usps.UserId = ""
	uSPSIntlRatesRequest := USPSIntlRatesRequest{}
	uSPSIntlRatesRequest.PackageId = "2ND"
	uSPSIntlRatesRequest.Ounces = "3.12"
	uSPSIntlRatesRequest.MailType = "Package"
	uSPSIntlRatesRequest.Machinable = "true"
	uSPSIntlRatesRequest.POBoxFlag = "Y"
	uSPSIntlRatesRequest.GiftFlag = "Y"
	uSPSIntlRatesRequest.ValueOfContents = "200"
	uSPSIntlRatesRequest.Country = "Australia"
	uSPSIntlRatesRequest.Container = "VARIABLE"
	uSPSIntlRatesRequest.Width = "12"
	uSPSIntlRatesRequest.Length = "10"
	uSPSIntlRatesRequest.Height = "6"
	uSPSIntlRatesRequest.Girth = "0"
	uSPSIntlRatesRequest.OriginZip = "18701"
	uSPSIntlRatesRequest.CommercialFlag = "Y"
	uSPSIntlRatesRequest.AcceptanceDateTime = "2022-10-12T13:15:00-06:00"
	uSPSIntlRatesRequest.DestinationPostalCode = "2046"
	resp := usps.IntlRates(uSPSIntlRatesRequest)
	log.Println(resp)
}

func TestUSPS_AddressValidation(t *testing.T) {
	usps := &USPS{}
	usps.UserId = ""
	uSPSAddressVaildRequest := USPSAddressVaildRequest{}
	uSPSAddressVaildRequest.AddressId = "0"
	uSPSAddressVaildRequest.Address1 = "SUITE K"
	uSPSAddressVaildRequest.Address2 = "29851 Aventura"
	uSPSAddressVaildRequest.State = "CA"
	uSPSAddressVaildRequest.Zip5 = "92688"
	resp := usps.AddressValidation(uSPSAddressVaildRequest)
	log.Println(resp)
}

func TestUSPS_ZIPCodeLookup(t *testing.T) {
	usps := &USPS{}
	usps.UserId = ""
	uSPSAddressVaildRequest := USPSAddressVaildRequest{}
	uSPSAddressVaildRequest.AddressId = "1"
	uSPSAddressVaildRequest.Address1 = "8 Wildwood Drive"
	uSPSAddressVaildRequest.Address2 = ""
	uSPSAddressVaildRequest.City = "Old Lyme"
	uSPSAddressVaildRequest.State = "CT"
	uSPSAddressVaildRequest.Zip5 = "06371"
	resp := usps.ZIPCodeLookup(uSPSAddressVaildRequest)
	log.Println(resp)
}

func TestUSPS_CityStateLookup(t *testing.T) {
	usps := &USPS{}
	usps.UserId = ""
	resp := usps.CityStateLookup("20024")
	log.Println(resp)
}
