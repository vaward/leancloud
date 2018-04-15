package main

import (
	"log"

	"github.com/vaward/leancloud"
)

var (
	appConfig = lean.RestConfig{"", ""}
)

type TestData struct {
	Score string
	//data  DataType
}

type MyRes struct {
	lean.RestResponse
	lean.ImageResponse
}

type TestDataRes struct {
	TestData
	MyRes
}

func main() {
	a := lean.RestResponse{}
	log.Println(a)
	log.Println("****************************************")
	header, err := lean.DoRestReq(appConfig,
		lean.RestRequest{
			lean.BaseReq{
				"GET",
				lean.ApiRestURL(appConfig.AppID, "BevaCount"),
				""},
			"application/json",
			nil},
		&a)
	if err == nil {
		log.Println(header)
		log.Println(a.Results)
	} else {
		log.Panic(err)
	}

	log.Println("****************************************")
}
