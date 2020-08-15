package main

import (
	"encoding/json"
	"fmt"
	"os"

	// "fmt"
	"io/ioutil"
	// "log"
	"net/http"
	// "net/url"
	"strings"
	// "net/url"
)

// url : https://jwxt-443.webvpn.sysu.edu.cn/jwxt/evaluation-manage/evaluationMission/queryStuAllEvalMission?_t=1580878844
// payload : {"pageNo":1,"pageSize":10,"total":true,"param":{"acadYear":"2019-1"}}

// const (
// 	myTeacherInfoUrl = "https://jwxt-443.webvpn.sysu.edu.cn/jwxt/evaluation-manage/evaluationMission/queryStuAllEvalMission"
// )

// 这个网址应该是查询没有评教的(没测试)
// https://jwxt-443.webvpn.sysu.edu.cn/jwxt/evaluation-manage/evaluationMission/queryStuEvalMission?_t=1581084713

type TeachersInfoStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    struct {
		Total int `json:"total"`
		Rows  []struct {
			// ID                string      `json:"id"`
			// StuNum             interface{} `json:"stuNum"`
			CourseName string `json:"courseName"`
			// EvallndexType     string      `json:"evallndexType"`
			// EvaluationWay     string      `json:"evaluationWay"`
			Teacher    string `json:"teacher"`
			CourseType string `json:"courseType"`
			// CourseUnit        string      `json:"courseUnit"`
			// ClassNumber string `json:"classNumber"`
			// AcadYear          string      `json:"acadYear"`
			// StartTime         string      `json:"startTime"`
			TeacherNumber string `json:"teacherNumber"`
			TeacherUnit   string `json:"teacherUnit"`
			// CourseCode        string      `json:"courseCode"`
			// Score             string      `json:"score"`
			// EvallndexTypeCode string      `json:"evallndexTypeCode"`
			// EvaluationWayCode string      `json:"evaluationWayCode"`
		} `json:"rows"`
	} `json:"data"`
}

// 通过评教接口
func (client *Client) getMyTeachersInfo() (string, error) {
	req, _ := http.NewRequest("POST", urlLists.teachersInfo, strings.NewReader(urlLists.getTeachersInfojsonBody))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Add("Referer", urlLists.baseUrl+"jwxt/mk/evaluation/")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(bytes))
	info := TeachersInfoStruct{}
	err := json.Unmarshal(bytes, &info)
	if err != nil {
		return "", err
	}
	if info.Code != 200 {
		return info.Message.(string), err
	}
	var ret string
	os.Mkdir("Teachers_pic", 0777)
	for _, it := range info.Data.Rows {
		ret += it.CourseType + " " + it.TeacherUnit + " " + it.CourseName + " " + it.Teacher + it.TeacherNumber + "\n"
		client.getTeacherImg(it.TeacherNumber, "Teachers_pic/"+it.Teacher)
	}
	return ret[:len(ret)-1], nil
}

func (client *Client) getTeacherImg(teacherId, filepath string) ([]byte, error) {
	//获得老师照片     e.g. 150149
	imgurl := urlLists.teacherImgUrl + teacherId
	r, err := http.NewRequest("GET", imgurl, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	r.Header.Add("Referer", urlLists.baseUrl+"jwxt/mk/evaluation/")
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if filepath == "" {
		return bytes, nil
	} else {
		ioutil.WriteFile(filepath+teacherId+".jpg", bytes, 0777)
		return nil, nil
	}

}

func (client *Client) getTeacherInfo(courseNum, teacherName string) map[string]interface{} {
	url := "https://jwxt.sysu.edu.cn/jwxt/training-programe/courseoutline/getalloutlineinfo?courseNum=" + courseNum
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Add("Host", "jwxt.sysu.edu.cn")
	req.Header.Add("Referer", "https://jwxt.sysu.edu.cn/jwxt/mk/courseSelection/")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if data["code"].(float64) == 52000000 {
		// 无课程大纲信息
		return nil
	}
	// courseOutline := data["data"].(map[string]interface{})["outlineInfo"].(map[string]interface{})
	teacherInfo := data["data"].(map[string]interface{})["teacherList"].([]interface{})[0].(map[string]interface{})
	if teacherName != "" && teacherName != teacherInfo["name"].(string) {
		fmt.Println("getTeacherId : Name Error")
		return nil
	}
	return teacherInfo
}
