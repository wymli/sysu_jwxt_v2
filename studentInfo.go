package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
)

type studentInfo struct {
	Name   string `json:"name"`
	School string `json:"school"`
	Major  string `json:"major"`
	Id     string `json:"id"`
}

func (this *Client) getStudentInfo() studentInfo {
	url := urlLists.baseUrl + "jwxt/student-status/countrystu/studentRollView"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Add("Origin", "https://sysu.edu.cn")
	req.Header.Add("Host", "jwxt.sysu.edu.cn")
	req.Header.Add("Referer", "https://jwxt.sysu.edu.cn/jwxt/")
	resp, _ := this.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	data := make(map[string]interface{})
	json.Unmarshal(b, &data)
	realData := data["data"].(map[string]interface{})
	info := studentInfo{
		realData["basicName"].(string), realData["rollCollegeNumNAME"].(string), realData["rollGradeDirectionNAME"].(string), realData["studentNumber"].(string),
	}
	return info
}
