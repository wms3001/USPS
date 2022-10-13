package USPS

import (
	"encoding/json"
	"encoding/xml"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/http"
	"strings"
)

type USPS struct {
	UserId   string `json:"userId"`
	PassWord string `json:"passWord"`
}

var url string = "https://secure.shippingapis.com/ShippingAPI.dll"

/**
获取简易跟踪信息
*/
func (usps *USPS) Track(trno string) *goCommon.Resp {
	request := "<TrackRequest USERID=\"" +
		usps.UserId + "\">" + "<TrackID ID=\"" + trno + "\"></TrackID></TrackRequest>"
	//resp := &goCommon.Resp{}
	m := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	a := map[string]string{}
	var wHttp = http.Http{
		url + "?API=TrackV2&XML=" + request,
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	uSPSTrackResponse := &USPSTrackResponse{}
	xml.Unmarshal([]byte(resp.Data), uSPSTrackResponse)
	re, _ := json.Marshal(uSPSTrackResponse)
	resp.Data = string(re)
	return resp
}

/**
获取分字段跟踪信息
*/
func (usps *USPS) TrackFields(trno string) *goCommon.Resp {
	request := "<TrackFieldRequest USERID= \"" + usps.UserId + "\">"
	request += "<Revision>1</Revision>"
	request += "<ClientIp>127.0.0.1</ClientIp>"
	request += "<SourceId>glkj</SourceId>"
	request += "<TrackID ID=\"" + trno + "\"/>"
	request += "</TrackFieldRequest>"

	m := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	a := map[string]string{}
	var wHttp = http.Http{
		url + "?API=TrackV2&XML=" + request,
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	strings.Replace(resp.Data, "TrackSummary", "TrackDetail", -1)
	uSPSTrackFieldsResponse := &USPSTrackFieldsResponse{}
	xml.Unmarshal([]byte(resp.Data), uSPSTrackFieldsResponse)
	re, _ := json.Marshal(uSPSTrackFieldsResponse)
	resp.Data = string(re)
	return resp
}

/**
usps运费计算
*/
func (usps *USPS) DomesticRates(request USPSDomesticRateRequest) *goCommon.Resp {
	xmlStr := "<RateV4Request USERID=\"" + usps.UserId + "\">"
	xmlStr += "<Revision>2</Revision>"
	xmlStr += "<Package ID=\"" + request.PackageId + "\">"
	xmlStr += "<Service>" + request.Service + "</Service>"
	xmlStr += "<FirstClassMailType>" + request.FirstClassMailType + "</FirstClassMailType>"
	xmlStr += "<ZipOrigination>" + request.ZipOrigination + "</ZipOrigination>"
	xmlStr += "<ZipDestination>" + request.ZipDestination + "</ZipDestination>"
	xmlStr += "<Pounds>" + request.Pounds + "</Pounds>"
	xmlStr += "<Ounces>" + request.Ounces + "</Ounces>"
	xmlStr += "<Container>" + request.Container + "</Container>"
	xmlStr += "<Size>" + request.Size + "</Size>"
	xmlStr += "<Width>" + request.Width + "</Width>"
	xmlStr += "<Length>" + request.Length + "</Length>"
	xmlStr += "<Height>" + request.Height + "</Height>"
	xmlStr += "<Girth>" + request.Girth + "</Girth>"
	xmlStr += "<Machinable>" + request.Machinable + "</Machinable>"
	xmlStr += "<ReturnLocations>" + request.ReturnLocations + "</ReturnLocations>"
	xmlStr += "<ReturnServiceInfo>" + request.ReturnServiceInfo + "</ReturnServiceInfo>"
	xmlStr += "<ShipDate>" + request.ShipDate + "</ShipDate>"
	xmlStr += "<ReturnFees>" + request.ReturnFees + "</ReturnFees>"
	xmlStr += "</Package></RateV4Request>"
	m := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	a := map[string]string{}
	var wHttp = http.Http{
		url + "?API=RateV4&XML=" + xmlStr,
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	if strings.Contains(resp.Data, "Error") {
		uSPSErrorResponse := &USPSErrorResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSErrorResponse)
		re, _ := json.Marshal(uSPSErrorResponse)
		resp.Data = string(re)
	} else {
		uSPSDomesticRateResponse := &USPSDomesticRateResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSDomesticRateResponse)
		re, _ := json.Marshal(uSPSDomesticRateResponse)
		resp.Data = string(re)
	}
	return resp
}

/**
usps国际运费计算
*/
func (usps *USPS) IntlRates(request USPSIntlRatesRequest) *goCommon.Resp {
	xmlStr := "<IntlRateV2Request USERID= \"" + usps.UserId + "\">"
	xmlStr += "<Revision>2</Revision>"
	xmlStr += "<Package ID=\"" + request.PackageId + "\">"
	xmlStr += " <Pounds>" + request.Pounds + "</Pounds>"
	xmlStr += "<Ounces>" + request.Ounces + "</Ounces>"
	xmlStr += "<Machinable>" + request.Machinable + "</Machinable>"
	xmlStr += "<MailType>" + request.MailType + "</MailType>"
	xmlStr += "<GXG><POBoxFlag>" + request.POBoxFlag + "</POBoxFlag><GiftFlag>" + request.GiftFlag + "</GiftFlag></GXG>"
	xmlStr += "<ValueOfContents>" + request.ValueOfContents + "</ValueOfContents>"
	xmlStr += "<Country>" + request.Country + "</Country>"
	xmlStr += "<Container>" + request.Container + "</Container>"
	xmlStr += "<Width>" + request.Width + "</Width>"
	xmlStr += "<Length>" + request.Length + "</Length>"
	xmlStr += "<Height>" + request.Height + "</Height>"
	xmlStr += "<Girth>" + request.Girth + "</Girth>"
	xmlStr += "<OriginZip>" + request.OriginZip + "</OriginZip>"
	xmlStr += "<CommercialFlag>" + request.CommercialFlag + "</CommercialFlag>"
	xmlStr += "<AcceptanceDateTime>" + request.AcceptanceDateTime + "</AcceptanceDateTime>"
	xmlStr += "<DestinationPostalCode>" + request.DestinationPostalCode + "</DestinationPostalCode>"
	xmlStr += "</Package></IntlRateV2Request>"
	m := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	a := map[string]string{}
	var wHttp = http.Http{
		url + "?API=IntlRateV2&XML=" + xmlStr,
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	if strings.Contains(resp.Data, "Error") {
		uSPSErrorResponse := &USPSErrorResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSErrorResponse)
		re, _ := json.Marshal(uSPSErrorResponse)
		resp.Data = string(re)
	} else {
		uSPSIntlRatesResponse := &USPSIntlRatesResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSIntlRatesResponse)
		re, _ := json.Marshal(uSPSIntlRatesResponse)
		resp.Data = string(re)
	}
	return resp
}

/**
地址验证
*/
func (usps *USPS) AddressValidation(request USPSAddressVaildRequest) *goCommon.Resp {
	xmlStr := "<AddressValidateRequest USERID=\"" + usps.UserId + "\">"
	xmlStr += "<Revision>1</Revision>"
	xmlStr += "<Address ID=\"" + request.AddressId + "\">"
	xmlStr += "<Address1>" + request.Address1 + "</Address1>"
	xmlStr += "<Address2>" + request.Address2 + "</Address2>"
	xmlStr += "<City>" + request.City + "</City>"
	xmlStr += "<State>" + request.State + "</State>"
	xmlStr += "<Zip5>" + request.Zip5 + "</Zip5>"
	xmlStr += "<Zip4>" + request.Zip4 + "</Zip4>"
	xmlStr += "</Address></AddressValidateRequest>"
	m := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	a := map[string]string{}
	var wHttp = http.Http{
		url + "?API=Verify&XML=" + xmlStr,
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	if strings.Contains(resp.Data, "Error") {
		uSPSErrorResponse := &USPSErrorResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSErrorResponse)
		re, _ := json.Marshal(uSPSErrorResponse)
		resp.Data = string(re)
	} else {
		uSPSAddressVaildResponse := &USPSAddressVaildResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSAddressVaildResponse)
		re, _ := json.Marshal(uSPSAddressVaildResponse)
		resp.Data = string(re)
	}
	return resp
}

func (usps *USPS) ZIPCodeLookup(request USPSAddressVaildRequest) *goCommon.Resp {
	xmlStr := "<ZipCodeLookupRequest USERID=\"" + usps.UserId + "\">"
	xmlStr += "<Revision>1</Revision>"
	xmlStr += "<Address ID=\"" + request.AddressId + "\">"
	xmlStr += "<Address1>" + request.Address1 + "</Address1>"
	xmlStr += "<Address2>" + request.Address2 + "</Address2>"
	xmlStr += "<City>" + request.City + "</City>"
	xmlStr += "<State>" + request.State + "</State>"
	xmlStr += "<Zip5>" + request.Zip5 + "</Zip5>"
	xmlStr += "<Zip4>" + request.Zip4 + "</Zip4>"
	xmlStr += "</Address></ZipCodeLookupRequest>"
	m := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	a := map[string]string{}
	var wHttp = http.Http{
		url + "?API=ZipCodeLookup&XML=" + xmlStr,
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	if strings.Contains(resp.Data, "Error") {
		uSPSErrorResponse := &USPSErrorResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSErrorResponse)
		re, _ := json.Marshal(uSPSErrorResponse)
		resp.Data = string(re)
	} else {
		uSPSZipCodeLookupResponse := &USPSZipCodeLookupResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSZipCodeLookupResponse)
		re, _ := json.Marshal(uSPSZipCodeLookupResponse)
		resp.Data = string(re)
	}
	return resp
}

func (usps *USPS) CityStateLookup(zip string) *goCommon.Resp {
	xmlStr := "<CityStateLookupRequest USERID=\"" + usps.UserId + "\">"
	xmlStr += "<Revision>1</Revision>"
	xmlStr += "<ZipCode ID=\"0\">"
	xmlStr += "<Zip5>" + zip + "</Zip5>"
	xmlStr += "</ZipCode></CityStateLookupRequest>"
	m := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	a := map[string]string{}
	var wHttp = http.Http{
		url + "?API=CityStateLookup&XML=" + xmlStr,
		"",
		"",
		m,
		a,
		"",
		10,
		"",
		[]byte{},
	}
	resp := wHttp.Get()
	if strings.Contains(resp.Data, "Error") {
		uSPSErrorResponse := &USPSErrorResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSErrorResponse)
		re, _ := json.Marshal(uSPSErrorResponse)
		resp.Data = string(re)
	} else {
		uSPSCityStateLookupResponse := &USPSCityStateLookupResponse{}
		xml.Unmarshal([]byte(resp.Data), uSPSCityStateLookupResponse)
		re, _ := json.Marshal(uSPSCityStateLookupResponse)
		resp.Data = string(re)
	}
	return resp
}
