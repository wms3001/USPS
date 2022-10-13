# golang usps api对接
## 简介
实现对接 usps api接口
## 使用
```
go get github.com/wms3001/USPS
```
## 实例
1. 获取跟踪信息
```go
usps := &USPS{}
usps.UserId = "680CHUAN3141"
resp := usps.Track("93055109246000000114188641")
log.Println(resp)

{"XMLName":{"Space":"","Local":"TrackResponse"},"TrackInfo":{"ID":"93055109246000000114188641","TrackSummary":"Your item departed a shipping partner facility at 1:08 am on October 9, 2022 in PICO RIVERA, CA 90660. This does not indicate receipt by the USPS or the actual mailing date.","TrackDetail":["Arrived Shipping Partner Facility, USPS Awaiting Item, October 9, 2022, 12:14 am, PICO RIVERA, CA 90660","Picked Up by Shipping Partner, USPS Awaiting Item, October 8, 2022, 11:13 pm, PICO RIVERA, CA 90660","Shipping Label Created, USPS Awaiting Item, October 8, 2022, 11:13 pm, PICO RIVERA, CA 90660"]}}
```
2. 获取跟踪信息(分字段)
```go
usps := &USPS{}
usps.UserId = "680CHUAN3141"
resp := usps.TrackFields("93055109246000000114188641")
log.Println(resp)

{"XMLName":{"Space":"","Local":"TrackResponse"},"TrackInfo":{"ID":"93055109246000000114188641","Class":"Priority Mail\u003cSUP\u003e\u0026reg;\u003c/SUP\u003e","ClassOfMailCode":"PM","DestinationCity":"GOODYEAR","DestinationState":"AZ","DestinationZip":"85338","EmailEnabled":true,"KahalaIndicator":false,"MailTypeCode":"DM","MPDATE":"2022-10-09 01:12:39.000000","MPSUFFIX":"501503904","OnTime":"false","OriginCity":"PICO RIVERA","OriginState":"CA","OriginZip":"90660","PodEnabled":false,"TPodEnabled":false,"RestoreEnabled":false,"RramEnabled":false,"RreEnabled":false,"Service":"Up to $100 insurance included","ServiceTypeCode":"055","Status":"Departed Shipping Partner Facility, USPS Awaiting Item","StatusCategory":"On Its Way to USPS","StatusSummary":"Your item departed a shipping partner facility at 1:08 am on October 9, 2022 in PICO RIVERA, CA 90660. This does not indicate receipt by the USPS or the actual mailing date.","TABLECODE":"T","TrackDetail":[{"EventTime":"12:14 am","EventDate":"October 9, 2022","Event":"Arrived Shipping Partner Facility, USPS Awaiting Item","EventCity":"PICO RIVERA","EventState":"CA","EventZIPCode":"90660","EventCountry":"","FirmName":"","Name":"","AuthorizedAgent":"false","EventCode":"81","GMT":"07:14:24","GMTOffset":"-07:00"},{"EventTime":"11:13 pm","EventDate":"October 8, 2022","Event":"Picked Up by Shipping Partner, USPS Awaiting Item","EventCity":"PICO RIVERA","EventState":"CA","EventZIPCode":"90660","EventCountry":"","FirmName":"","Name":"","AuthorizedAgent":"false","EventCode":"80","GMT":"06:13:29","GMTOffset":"-07:00"},{"EventTime":"11:13 pm","EventDate":"October 8, 2022","Event":"Shipping Label Created, USPS Awaiting Item","EventCity":"PICO RIVERA","EventState":"CA","EventZIPCode":"90660","EventCountry":"","FirmName":"","Name":"","AuthorizedAgent":"false","EventCode":"GX","GMT":"06:13:27","GMTOffset":"-07:00"}]}}
```
3. 费用计算
```go
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
```
4. 费用计算(国际)
```go
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
```
5. 地址验证
```go
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
```
6. 邮编地址验证
```go
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
```
7. 邮编查询城市州
```go
usps := &USPS{}
usps.UserId = ""
resp := usps.CityStateLookup("20024")
log.Println(resp)
```