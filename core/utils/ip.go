package utils

import (
	"core/logger"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var reader *geoip2.Reader

func init() {
	db, err := geoip2.Open("config/GeoLite2-City.mmdb")
	if err != nil {
		logger.Fatal(err.Error())
		return
	}
	reader = db
}

func GetCountry(ipAddr string) (string, error) {
	if reader == nil {
		return "", nil
	}
	ip := net.ParseIP(ipAddr)
	record, err := reader.City(ip)
	if err != nil {
		return "", err
	}
	return record.Country.Names["zh-CN"], nil
}

func GetProvince(ipAddr string) (string, error) {
	if reader == nil {
		return "", nil
	}
	ip := net.ParseIP(ipAddr)
	record, err := reader.City(ip)
	if err != nil {
		return "", err
	}
	return record.Subdivisions[0].Names["zh-CN"], nil
}

func GetCity(ipAddr string) (string, error) {
	if reader == nil {
		return "", nil
	}
	ip := net.ParseIP(ipAddr)
	record, err := reader.City(ip)
	if err != nil {
		return "", err
	}
	return record.City.Names["zh-CN"], nil
}
