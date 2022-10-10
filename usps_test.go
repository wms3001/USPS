package USPS

import (
	"log"
	"testing"
)

func TestUSPS_Track(t *testing.T) {
	usps := &USPS{}
	usps.UserId = "680CHUAN3141"
	resp := usps.Track("93055109246000000114188641")
	log.Println(resp)
}
