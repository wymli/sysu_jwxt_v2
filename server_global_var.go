package main

import "log"
import "net/http"

type Client struct {
	*http.Client
	loginForm map[string][]string
	isLogin   bool
}

var client *Client
var urlLists urlList
var _campus_id map[string]string
var _selectedType map[string]string
var _selectedCate map[string]string

const (
	WEBVPN = 1
	NORMAL = 2
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	log.SetPrefix("[@STrelitziA@]")
	urlLists.init(NORMAL)
	client = getDefaultClient()

	_campus_id = map[string]string{
		"东校园":  "5063559",
		"北校园":  "5062202",
		"南校园":  "5062201",
		"深圳校区": "333291143",
		"珠海校区": "5062203",
	}
	_selectedType = map[string]string{
		"本专业":  "1",
		"校级公选": "4",
		"跨专业":  "2",
	}
	_selectedCate = map[string]string{
		"专必":     "11",
		"专选":     "21",
		"院内公选":   "30",
		"公必(体育)": "10",
		"公必(大英)": "10",
		"公必(其他)": "10",
	}
}

// 分类: 东校园, 南校园 等等
func getCampusId(campus string) string {
	return _campus_id[campus]
}

// 分类: 本专业,校级公选,跨专业 三类
func getSelectedType(courseType string) string {
	return _selectedType[courseType]
}

// 当selectedType== 1 时有效 分类: 本专业的课,包括专必专选,体育英语
func getSelectedCate(courseCate string) string {
	return _selectedCate[courseCate]
}
