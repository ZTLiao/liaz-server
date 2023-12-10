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
	if len(record.Subdivisions) > 0 {
		return record.Subdivisions[0].Names["zh-CN"], nil
	}
	return "", nil
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

func GetAddress(ipAddr string) (string, error) {
	var ipRegion string
	if len(ipAddr) > 0 {
		country, err := GetCountry(ipAddr)
		if err != nil {
			return "", err
		}
		province, err := GetProvince(ipAddr)
		if err != nil {
			return "", err
		}
		city, err := GetCity(ipAddr)
		if err != nil {
			return "", err
		}
		if len(country) != 0 || len(province) != 0 || len(city) != 0 {
			ipRegion = country + COMMA + province + COMMA + city
		}
	}
	return ipRegion, nil
}

func GetLocation(ipAddr string) (country string, province string, city string, err error) {
	if len(ipAddr) > 0 {
		country, err = GetCountry(ipAddr)
		if err != nil {
			return "", "", "", err
		}
		province, err = GetProvince(ipAddr)
		if err != nil {
			return "", "", "", err
		}
		city, err = GetCity(ipAddr)
		if err != nil {
			return "", "", "", err
		}
	}
	return
}
