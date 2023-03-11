package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Server struct {
	ServerName string `xml:"serverName""`
	ServerIP   string `xml:"serverIP""`
}

type ServersList struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Servers []Server `xml:"server"`
}

func main() {
	xmldata, err := ioutil.ReadFile("./fei.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var VPNServer ServersList
	err = xml.Unmarshal(xmldata, &VPNServer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(VPNServer, "\n", string(xmldata))

}

/*
felix@MacBook-Pro 2023Project % go run main.go
{{ servers} 1 [{Shanghai_VPN 127.0.0.1} {Beijing_VPN 127.0.0.2}]}
 <?xml version="1.0" encoding="UTF-8" ?>
<servers version="1">
    <server>
        <serverName>Shanghai_VPN</serverName>
        <serverIP>127.0.0.1</serverIP>
    </server>
    <server>
        <serverName>Beijing_VPN</serverName>
        <serverIP>127.0.0.2</serverIP>
    </server>
</servers>

*/
