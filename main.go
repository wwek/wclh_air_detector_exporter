package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mrflynn/go-aqi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strings"
)

var (
	listenAddr     = flag.String("listen_addr", ":9166", "监听地址，默认 :9166")
	serialPort     = flag.String("serial_port", "/dev/ttyUSB0", "串口地址")
	DeviceID       = flag.String("device_id", "wclh", "设备id")
	DeviceLocation = flag.String("device_locaiton", "home", "设备地理位置")
)

func main() {
	flag.Parse()
	fmt.Println(versionInfo())
	SetupMetrics()

	http.HandleFunc("/", httpIndexRequestHandler)
	http.HandleFunc("/healthz", httpHealthzRequestHandler)
	http.Handle("/metrics", promhttp.Handler())
	SetupDevice(*serialPort)
	dev, err := OpenDevice(*serialPort)
	if err != nil {
		log.Fatalf("Could not open '%v'", *serialPort)
	}
	defer dev.Close()

	go receive(dev)

	log.Printf("Serving metrics at http://localhost:%v/metrics", strings.Split(*listenAddr, ":")[1])
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

func receive(a *Device) {
	log.Println("开始接收：", *serialPort)

	ctx := context.Background()
	// 持续读取和解析接收的串口数据
	err := a.Process(ctx, DecodedSignal)
	if err != nil {
		log.Println(err)
	}
}

// DecodedSignal 解码串口接收到的信号📶
func DecodedSignal(line string) (stop bool) {
	stop = false

	isData := false
	if !strings.HasPrefix(line, ReceivePrefixStar) && !strings.HasPrefix(line, ReceivePrefixField) {
		isData = true
	}
	if isData {
		signal, err := ParserSignal(line)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Println(signal)
		AssignmentMetrics(*signal)
	}
	return
}

// AssignmentMetrics 对prometheus metrics进行赋值
func AssignmentMetrics(s Signal) {
	var labels = prometheus.Labels{
		"sensor_id":       *DeviceID,
		"sensor_location": *DeviceLocation,
	}

	Temperature.With(labels).Set(s.Columns[0])
	Humidity.With(labels).Set(s.Columns[1])
	ChPm1dot0.With(labels).Set(s.Columns[2])
	ChPm2dot5.With(labels).Set(s.Columns[3])
	ChPm10.With(labels).Set(s.Columns[4])
	UsPm1dot0.With(labels).Set(s.Columns[5])
	UsPm2dot5.With(labels).Set(s.Columns[6])
	UsPm10.With(labels).Set(s.Columns[7])
	Gt0dot3um.With(labels).Set(s.Columns[8])
	Gt0dot5um.With(labels).Set(s.Columns[9])
	Gt1dot0um.With(labels).Set(s.Columns[10])
	Gt2dot5um.With(labels).Set(s.Columns[11])
	Gt5dot0um.With(labels).Set(s.Columns[12])
	Gt10um.With(labels).Set(s.Columns[13])
	if s.Model == "co2" {
		Co2.With(labels).Set(s.Columns[14])
		Hcho.With(labels).Set(s.Columns[15])
		Tvoc.With(labels).Set(s.Columns[16])
	}
	if s.Model == "no_co2" {
		Hcho.With(labels).Set(s.Columns[14])
		Tvoc.With(labels).Set(s.Columns[15])
	}

	aqiResult, err := aqi.Calculate(aqi.PM25{Concentration: s.Columns[3]})
	if err != nil {
		fmt.Println(err)
		return
	}
	Aqiatpm2dot5.With(labels).Set(aqiResult.AQI)
}

func httpHealthzRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func httpIndexRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<html>
    <head><title>wclh_air_detector_exporter</title></head>
    <body>
    <h1>wclh_air_detector_exporter</h1>
	<p><a href="metrics">Metrics</a></p>
    <p><a href="healthz">Healthz</a></p>
    <p><a href="https://github.com/wwek/wclh_air_detector_exporter" target="_blank">https://github.com/wwek/wclh_air_detector_exporter</a></p>
	</body></html>`))
}
