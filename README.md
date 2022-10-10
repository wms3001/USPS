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