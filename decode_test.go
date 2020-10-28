package main

import (
	"log"
	"testing"
)

func TestParserSignal(t *testing.T) {
	inputs := []string{
		"*--------------------------------------- Data Output Format Definition ----------------------------------------*",
		"TEMP HUMI CH_PM1.0 CH_PM2.5 CH_PM10 US_PM1.0 US_PM2.5 US_PM10 >0.3um >0.5um >1.0um >2.5um >5.0um >10um HCHO TVOC",
		"24.9 45.2 23 32 35 26 33 35 4419 1301 159 8 3 0 0.006 0.21",
		"25.9 36.1 20 32 35 26 33 35 4419 1301 159 8 3 0 0.007 0.23",
	}
	for _, test := range inputs {
		rs,err := ParserSignal(test)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(rs)
	}
}