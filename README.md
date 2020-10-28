
--- 
 
<p align="right">
  <a href="https://github.com/wwek/wclh_air_detector_exporter/releases/latest">
    <img alt="Release" src="https://img.shields.io/github/release/wwek/wclh_air_detector_exporter.svg?style=flat-square">
  </a>
  <a href="https://github.com/wwek/wclh_air_detector_exporter/master">
    <img alt="Travis" src="https://img.shields.io/travis/wwek/wclh_air_detector_exporter/master.svg?style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/wwek/wclh_air_detector_exporter">
    <img alt="Go Report" src="https://goreportcard.com/badge/github.com/wwek/wclh_air_detector_exporter?style=flat-square" />
  </a>
  <a href="https://codecov.io/gh/wwek/wclh_air_detector_exporter">
    <img alt="Codecov branch" src="https://codecov.io/gh/wwek/wclh_air_detector_exporter/branch/master/graph/badge.svg?style=flat-square" />
  </a>
  <a href="https://godoc.org/github.com/wwek/wclh_air_detector_exporter">
    <img alt="Go Doc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square" />
  </a>
  <a href="https://github.com/wwek/wclh_air_detector_exporter/blob/master/LICENSE">
    <img alt="Software License" src="https://img.shields.io/github/license/wwek/wclh_air_detector_exporter.svg?style=flat-square" />
  </a>
</p>


wclh_air_detector_exporter读取串口数据并把数据进行结构化，然后输出metrics
 
 M5S Temperature and Humidity+lithium battery+CO2+TVOC PM2.5 CO2(S8)TEMP&HUMI Detector Haze PM2.5 sensors Laser PM2.5 detector

 M5S 家用 激光PM2.5检测仪 甲醛 CO2 空气质量 雾霾甲醛 检测仪
 
 【在售价】380.00 元(基础版)
 
 [【立即下单】点击链接立即下单：https://s.click.taobao.com/zREU4vu](https://s.click.taobao.com/zREU4vu)
 
 ![img](https://img.alicdn.com/i2/2375177132/O1CN01tI2WiW22YVDTkFZfx_!!2375177132.jpg)
 


## 下载&安装


## 自定义编译
```
make setup
make buildall
```

## 支持的微创联合M5S空气检测仪版本

* 10 M5S+温湿度+TVOC +数据导出+锂电
* 11 M5S+温湿度+TVOC +数据导出[无锂电]
* 15 M5S+温湿度+TVOC +C02+数导+锂电
* 16 M5S+温湿度+ TVOC+CO2+数导+锂电[英文版]

只支持带数据导出的版本！
```
#不带CO2的数据导出字段
*--------------------------------------- Data Output Format Definition ----------------------------------------*
TEMP HUMI CH_PM1.0 CH_PM2.5 CH_PM10 US_PM1.0 US_PM2.5 US_PM10 >0.3um >0.5um >1.0um >2.5um >5.0um >10um HCHO TVOC

#带CO2的数据导出字段
*--------------------------------------- Data Output Format Definition --------------------------------------------*
TEMP HUMI CH_PM1.0 CH_PM2.5 CH_PM10 US_PM1.0 US_PM2.5 US_PM10 >0.3um >0.5um >1.0um >2.5um >5.0um >10um CO2 HCHO TVOC
```

![img](https://img.alicdn.com/imgextra/i2/2375177132/TB24_lLqz7nBKNjSZLeXXbxCFXa_!!2375177132.jpg)

## Grafana Dashboards

 [wclh_air_detector_exporter Grafana dashboard](./grafana-dashboard.json)
