package main

import (
	//	"encoding/json"
	//	"fmt"
	"log"

	"github.com/kassiansun/bmob-go-sdk"
)

var (
	appConfig = bmob.RestConfig{"",
		""}
)

type TestData struct {
	Score string
	//data  DataType
}

type MyRes struct {
	bmob.RestResponse
	bmob.ImageResponse
}

type TestDataRes struct {
	TestData
	MyRes
}

func main() {
	a := bmob.RestResponse{}
	log.Println(a)
	log.Println("****************************************")
	var respDst = TestDataRes{}

	header, err := bmob.DoRestReq(appConfig,
		bmob.RestRequest{
			bmob.BaseReq{
				"GET",
				bmob.ApiRestURL("GameScore") + "/",
				""},
			"application/json",
			nil},
		&respDst)
	if err == nil {
		log.Println(header)
		log.Println(respDst)
	} else {
		log.Panic(err)
	}

	log.Println("****************************************")
}
