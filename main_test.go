package main

import (
	"fmt"
	"github.com/mrflynn/go-aqi"
	"log"
	"testing"
)

func TestAqiatpm2dot5(t *testing.T) {
	aqiResult, err := aqi.Calculate(aqi.PM25{Concentration: 26.5})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(aqiResult)
}
