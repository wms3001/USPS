package USPS

import (
	"encoding/json"
	"encoding/xml"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/http"
)

type USPS struct {
	UserId   string `json:"userId"`
	PassWord string `json:"passWord"`
}

var url string = "https://secure.shippingapis.com/ShippingAPI.dll"

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
