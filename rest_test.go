package bmob

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

var (
	appConfig = RestConfig{
		"d237352e841d58140c0695843159d78b",
		"3f74ad08a07d61fab887724a132d33b5"}
	//TODO add more bmob infos
	apiUrl     = "https://api.bmob.cn"
	fileUrl    = "https://file.bmob.cn"
	apiVersion = "1"
	classUrl   = "classes"
	usersUrl   = "users"
)

type TestData struct {
	Score string
	//data  DataType
}

type MyRes struct {
	RestResponse
	ImageResponse
}

type TestDataRes struct {
	Score string
	MyRes
}

func TestDelete(t *testing.T) {
	log.Println("****************************************")
	var respDst = TestDataRes{}
	header, err := DoRestReq(appConfig,
		RestRequest{
			BaseReq{"DELETE",
				ApiRestURL("GameScore") + "/" + "f04403e38c",
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

func TestGet(t *testing.T) {
	log.Println("****************************************")
	var respDst = TestDataRes{}

	header, err := DoRestReq(appConfig,
		RestRequest{
			BaseReq{
				"GET",
				ApiRestURL("GameScore") + "/" + "0d12ee1c09",
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

func TestGetMultiple(t *testing.T) {
	log.Println("****************************************")
	var respDst = TestDataRes{}

	header, err := DoRestReq(appConfig,
		RestRequest{
			BaseReq{
				"GET",
				ApiRestURL("GameScore") + "/",
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

func TestPost(t *testing.T) {
	log.Println("****************************************")

	var testData = TestData{"100"}
	b, err := json.Marshal(testData)
	if err != nil {
		fmt.Println("error:", err)
	}
	var respDst = TestDataRes{}
	header, err := DoRestReq(appConfig,
		RestRequest{
			BaseReq{
				"POST",
				ApiRestURL("GameScore"),
				""},
			"application/json",
			b},
		&respDst)
	if err == nil {
		log.Println(header)
		log.Println(respDst)
	} else {
		log.Panic(err)
	}

	log.Println("****************************************")
}

func TestPut(t *testing.T) {
	log.Println("****************************************")

	var testData = TestData{"100"}
	b, err := json.Marshal(testData)
	if err != nil {
		fmt.Println("error:", err)
	}
	var respDst = TestDataRes{}
	header, err := DoRestReq(appConfig,
		RestRequest{
			BaseReq{
				"PUT",
				ApiRestURL("GameScore") + "/" + "b8c930158c",
				""},
			"application/json",
			b},
		&respDst)
	if err == nil {
		log.Println(header)
		log.Println(respDst)
	} else {
		log.Panic(err)
	}
	log.Println("****************************************")
}
