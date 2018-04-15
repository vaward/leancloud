package lean

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//https://wrwgnhp8.api.lncld.net/1.1/classes/Post
//https://wrwgnhp8.api.lncld.net/1.1/files/hello.txt
const (
	ApiHost    = ".api.lncld.net"
	ApiVersion = "1.1"
)

/*
 * Generate RESTful URL
 */
func ApiRestURL(apiKey, class string) string {
	domain := strings.ToLower(string([]byte(apiKey)[:8]))
	if class == "users" {
		return "https://" + domain + ApiHost + "/" + ApiVersion + "/" + class
	} else {
		return "https://" + domain + ApiHost + "/" + ApiVersion + "/classes/" + class
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

	if restConfig.AppID != "" {
		req.Header.Add("X-LC-Id", restConfig.AppID)
		req.Header.Add("X-LC-Key", restConfig.RestKey)
	}
	//Append Token if given
	if restreq.Token != "" {
		req.Header.Add("X-LC-Session", restreq.Token)
	}

	if restreq.Type != "" {
		req.Header.Add("Content-Type", restreq.Type)
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
