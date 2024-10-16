package main

import (
	"net"
	"strings"

	"github.com/oschwald/geoip2-golang"
	"github.com/oschwald/maxminddb-golang"
)

func getCounries() (metadata maxminddb.Metadata, countryMap map[string][]*net.IPNet, err error) {
	database, err := maxminddb.Open("./data/Country.mmdb")
	if err != nil {
		return
	}
	defer database.Close()

	metadata = database.Metadata
	networks := database.Networks(maxminddb.SkipAliasedNetworks)
	countryMap = make(map[string][]*net.IPNet)
	var country geoip2.Enterprise
	var tttgtg geoip2.ASN
	_ = tttgtg
	var ipNet *net.IPNet
	for networks.Next() {
		ipNet, err = networks.Network(&country)
		if err != nil {
			return
		}
		code := strings.ToLower(country.RegisteredCountry.IsoCode)
		countryMap[code] = append(countryMap[code], ipNet)
	}

	err = networks.Err()
	return
}

/* func prep() {
	var includedCodes []string
	files, err := os.ReadDir("./data/ips")
	for _, file := range files {
		read, err := os.Open("./data/ips/" + file.Name())
		if err != nil {
			return nil, err
		}
		defer read.Close()

		code := strings.TrimSuffix(file.Name(), ".txt")
		includedCodes = append(includedCodes, code)

		scanner := bufio.NewScanner(read)
		for scanner.Scan() {
			line := scanner.Text()
			_, ipNet, err := net.ParseCIDR(line)
			if err != nil {
				return nil, err
			}
			dataMap[code] = append(dataMap[code], ipNet)
		}
	}
	return includedCodes, err

} */

func main() {
	getCounries()

}
