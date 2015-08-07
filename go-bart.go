package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type station struct {
	StationInfo string `xml:"date"`
}

func main() {
	stationList := getStationList()

	v := station{}
	err := xml.Unmarshal(stationList, &v)

	if err != nil {
		log.Fatal(err)
	}

	// estimate := getRealTimeEstimate()
	fmt.Printf("%v", v)
}

func getStationList() []byte {
	return get("http://api.bart.gov/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V")
}

func getRealTimeEstimate() []byte {
	return get("http://api.bart.gov/api/etd.aspx?cmd=etd&orig=RICH&key=MW9S-E7SL-26DU-VV8V")
}

func get(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return robots
}
