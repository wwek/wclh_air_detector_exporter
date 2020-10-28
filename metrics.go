package main

import "github.com/prometheus/client_golang/prometheus"

var (
	Temperature  *prometheus.GaugeVec
	Humidity     *prometheus.GaugeVec
	ChPm1dot0    *prometheus.GaugeVec
	ChPm2dot5    *prometheus.GaugeVec
	ChPm10       *prometheus.GaugeVec
	UsPm1dot0    *prometheus.GaugeVec
	UsPm2dot5    *prometheus.GaugeVec
	UsPm10       *prometheus.GaugeVec
	Gt0dot3um    *prometheus.GaugeVec
	Gt0dot5um    *prometheus.GaugeVec
	Gt1dot0um    *prometheus.GaugeVec
	Gt2dot5um    *prometheus.GaugeVec
	Gt5dot0um    *prometheus.GaugeVec
	Gt10um       *prometheus.GaugeVec
	Co2          *prometheus.GaugeVec
	Hcho         *prometheus.GaugeVec
	Tvoc         *prometheus.GaugeVec
	Aqiatpm2dot5 *prometheus.GaugeVec
)

// SetupMetrics 定义指标
func SetupMetrics() {
	Temperature = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_temperature",
		Help: "Current temperature in ℃",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	Humidity = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_humidity",
		Help: "Current humidity level in %",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	ChPm1dot0 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_ch_pm1dot0",
		Help: "Current ch pm1.0 level in ug/m³ / 国标CH PM1.0 ug/m³",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	ChPm2dot5 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_ch_pm2dot5",
		Help: "Current ch pm2.5 level in ug/m³ / 国标CH PM2.5 ug/m³",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	ChPm10 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_ch_pm10",
		Help: "Current ch pm10 level in ug/m³ / 国标CH PM10 ug/m³",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	UsPm1dot0 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_us_pm1dot0",
		Help: "Current us pm1.0 level in ug/m³ / 美标US PM1.0 ug/m³",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	UsPm2dot5 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_us_pm2dot5",
		Help: "Current us pm2.5 level in ug/m³ / 美标US PM2.5 ug/m³",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	UsPm10 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_us_pm10",
		Help: "Current us pm10 level in ug/m³ / 美标US PM10 ug/m³",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	Gt0dot3um = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_gt0dot3um",
		Help: "Current >0.3um level in ug/m³ / 直径大于 >0.3um 颗粒物个数",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	Gt0dot5um = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_gt0dot5um",
		Help: "Current >0.5um level in ug/m³ / 直径大于 >0.5um 颗粒物个数",
	}, []string{
		"sensor_id",
		"sensor_location",
	})
	Gt1dot0um = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_gt1dot0um",
		Help: "Current >1.0um level in ug/m³ / 直径大于 >1.0um 颗粒物个数",
	}, []string{
		"sensor_id",
		"sensor_location",
	})

	Gt2dot5um = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_gt2dot5um",
		Help: "Current >2.5um level in ug/m³ / 直径大于 >2.5um 颗粒物个数",
	}, []string{
		"sensor_id",
		"sensor_location",
	})

	Gt5dot0um = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_gt5dot0um",
		Help: "Current >5.0um level in ug/m³ / 直径大于 >5.0um 颗粒物个数",
	}, []string{
		"sensor_id",
		"sensor_location",
	})

	Gt10um = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_gt10um",
		Help: "Current >10um level in ug/m³ / 直径大于 >10um 颗粒物个数",
	}, []string{
		"sensor_id",
		"sensor_location",
	})

	Co2 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_co2",
		Help: "Current CO2 level in ppm / 二氧化碳 CO2 < 1000 ppm",
	}, []string{
		"sensor_id",
		"sensor_location",
	})

	Hcho = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_hcho",
		Help: "Current hcho level in mg/m³ / 甲醛 HCHO < 0.08 mg/m³",
	}, []string{
		"sensor_id",
		"sensor_location",
	})

	Tvoc = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_tvoc",
		Help: "Current tvoc level in mg/m³ / 异味 TVOC < 0.8 mg/m³",
	}, []string{
		"sensor_id",
		"sensor_location",
	})

	Aqiatpm2dot5 = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meter_aqiatpm2dot5",
		Help: "Current AQI@PM2.5 / AQI@PM2.5 控制质量指数",
	}, []string{
		"sensor_id",
		"sensor_location",
	})

	prometheus.MustRegister(Temperature)
	prometheus.MustRegister(Humidity)
	prometheus.MustRegister(ChPm1dot0)
	prometheus.MustRegister(ChPm2dot5)
	prometheus.MustRegister(ChPm10)
	prometheus.MustRegister(UsPm1dot0)
	prometheus.MustRegister(UsPm2dot5)
	prometheus.MustRegister(UsPm10)
	prometheus.MustRegister(Gt0dot3um)
	prometheus.MustRegister(Gt0dot5um)
	prometheus.MustRegister(Gt1dot0um)
	prometheus.MustRegister(Gt2dot5um)
	prometheus.MustRegister(Gt5dot0um)
	prometheus.MustRegister(Gt10um)
	prometheus.MustRegister(Co2)
	prometheus.MustRegister(Hcho)
	prometheus.MustRegister(Tvoc)
	prometheus.MustRegister(Aqiatpm2dot5)
}
