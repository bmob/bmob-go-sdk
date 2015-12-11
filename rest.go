package bmob

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ApiUrl     = "https://api.bmob.cn"
	FileUrl    = "https://file.bmob.cn"
	ApiVersion = "1"
)

/*
 * Generate RESTful URL
 */
func ApiRestURL(class string) string {
	if class == "users" {
		return ApiUrl + "/" + ApiVersion + "/" + class
	} else {
		return ApiUrl + "/" + ApiVersion + "/classes/" + class
	}
}

//TODO add more url generate function

/*
 * Generate RESTful Request, set App ID & Rest ID
 */
func newRestRequest(restConfig RestConfig, restreq RestRequest) (*http.Request, error) {
	req, err := http.NewRequest(restreq.Method, restreq.Path, bytes.NewReader(restreq.Body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Bmob-Application-Id", restConfig.AppID)
	req.Header.Add("X-Bmob-REST-API-Key", restConfig.RestKey)
	//Append Token if given
	if restreq.Token != "" {
		req.Header.Add("X-Bmob-Session-Token", restreq.Token)
	}

	if restreq.Type != "" {
		req.Header.Add("Content-Type", "application/json")
	} else {
		//fallback to text/plain
		req.Header.Add("Content-Type", "text/plain")
	}
	//log.Println(req)
	return req, nil
}

/*
 * Do RESTful request
 */
func DoRestReq(restConfig RestConfig, restreq RestRequest, respDst interface{}) (http.Header, error) {
	/*
	 * Generate POST Request, set App ID & Rest ID
	 */
	req, err := newRestRequest(restConfig, restreq)

	/*
	 * Do Request
	 */
	client := &http.Client{}
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	/*
	 * Read Response
	 */
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	log.Println(string(body))
	/*
	 * Generate Response Struct
	 */
	err = json.Unmarshal(body, &respDst)
	if err != nil {
		return nil, err
	}

	return resp.Header, nil
}
