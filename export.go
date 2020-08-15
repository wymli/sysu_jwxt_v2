package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	// "log"
	"net/http"
	"strings"
)

func (client *Client) setUser(username, password string) {
	client.loginForm["username"][0] = username
	client.loginForm["password"][0] = password
}
func (client *Client) setCaptcha(captcha string) {
	client.loginForm["captcha"][0] = captcha
}

// //deprecated
// func Login(mode int) *http.Client {
// 	urlLists.init(mode)
// 	var client *http.Client
// 	var cnt int = 0
// 	for {
// 		cnt++ //避免重复太多次
// 		if cnt > 5 {
// 			return nil
// 		}
// 		var err error
// 		client, err = casLogin() // 这里不能用 :=  ,否则会创建一个局部的client
// 		if err != nil {
// 			log.Println(err.Error())
// 			continue
// 		}
// 		if info, ok := isLoginAndGetInfo(client); ok {
// 			fmt.Println(info)
// 			if mode == WEBVPN { //获取webvpn模式下 jwxt的cookie
// 				getJwxtCookieWithWebVpn(client)
// 			}
// 			break
// 		} else {
// 			log.Println("重启，尝试次数：", cnt)
// 		}
// 	}
// 	return client
// }

func GetTeacherImg(client *http.Client, teacherId string) []byte {
	//获得老师照片     e.g. 150149
	imgurl := urlLists.teacherImgUrl + teacherId
	r, err := http.NewRequest("GET", imgurl, nil)
	if err != nil {
		return nil
	}
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	r.Header.Add("Referer", urlLists.baseUrl+"jwxt/mk/evaluation/")
	resp, _ := client.Do(r)
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return bytes
}

type TeacherInfo struct {
	info string
	img  []byte
}

func GetMyTeachersInfo(client *http.Client) ([]TeacherInfo, error) {
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
		return nil, err
	}
	if info.Code != 200 {
		return nil, err
	}
	var teachersInfo []TeacherInfo
	for _, it := range info.Data.Rows {
		tmp := it.CourseType + " " + it.TeacherUnit + " " + it.CourseName + " " + it.Teacher + it.TeacherNumber + "\n"
		teachersInfo = append(teachersInfo, TeacherInfo{tmp, GetTeacherImg(client, it.TeacherNumber)})
	}
	return teachersInfo, nil
}
