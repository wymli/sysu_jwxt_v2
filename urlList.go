package main

import "log"

const (
	//jwxt点击登陆,跳到->https://jwxt.sysu.edu.cn/jwxt/api/sso/cas/login?pattern=student-login,再会被重定向到下述url
	//使用GET,第一次返回JSESSIONID,解析页面获得提交的form内容,带JSID请求验证码; //这个cookie主要是为了验证码和本次请求的匹配

	//第二次请求,使用POST,后续会有几次跳转,因为教务系统是sso模式(单点登陆),验证是在cas做的,并且不同系统的服务器是不同的cookie,理论上只要获得cas的cookie就可以了
	//具体过程大致是cas返回cookie,302到jwxt的网站,我们禁止重定向,手动读取location,然后带着cookie请求,然后会返回jwxt的cookie
	loginURL1   = "https://cas.sysu.edu.cn/cas/login?service=https://jwxt.sysu.edu.cn/jwxt/api/sso/cas/login?pattern=student-login"
	captchaURL1 = "https://cas.sysu.edu.cn/cas/captcha.jsp"

	captchaURL = "https://cas-443.webvpn.sysu.edu.cn/cas/captcha.jsp"
	loginURL   = "https://cas-443.webvpn.sysu.edu.cn/cas/login?service=https://portal.sysu.edu.cn/management/shiro-cas"

	portalIndexUrl = "https://portal.sysu.edu.cn/#/index" // useless
	getGPAurl      = "https://portal.sysu.edu.cn/api/extraCard/studentInfo/GPARank?schoolYear=2019-2020&semester=1"

	//选课退课url-----post-----payload形式(json格式,不是form的url.values的键值对map[string][]string)
	courseSelectionChooseUrl = "jwxt/choose-course-front-server/classCourseInfo/course/choose"
	courseSelectionCancelUrl = "jwxt/choose-course-front-server/classCourseInfo/course/back"

	courseListUrl = "jwxt/choose-course-front-server/classCourseInfo/course/list"

	myTeacherInfoUrl = "jwxt/evaluation-manage/evaluationMission/queryStuAllEvalMission"
	teacherImgUrl    = "jwxt/evaluation-manage/evaluationMission/profile?no="

	getScoreListUrl = "jwxt/achievement-manage/score-check/list?addScoreFlag=true&scoSchoolYear=2019-2020&scoSemester=1&trainTypeCode=01"
	teachersInfo    = "jwxt/evaluation-manage/evaluationMission/queryStuAllEvalMission"
)

// jsonBody := `{"pageNo":1,"pageSize":100,"param":{"semesterYear":"2019-2","selectedType":"1","selectedCate":"21","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","collectionStatus":"0"}}`
// err = getCourseList(client, jsonBody)

// courseSelectionCancelBody := `{"courseId":"206169488","clazzId":"1201412705275330561","selectedType":"1"}`
// // courseSelectionChooseBody := `{"clazzId":"1201412705275330561","selectedType":"1","selectedCate":"21","check":true}`
// cancelCourse(client, courseSelectionCancelBody)
// cancelCourse(client, courseSelectionCancelBody)

type urlList struct {
	loginURL   string
	captchaURL string

	baseUrl string

	getGPAurl                string
	courseListUrl            string
	courseSelectionChooseUrl string
	courseSelectionCancelUrl string
	myTeacherInfoUrl         string
	teacherImgUrl            string
	getScoreListUrl          string
	teachersInfo             string

	getPublicCoursejsonBody  string
	getTeachersInfojsonBody  string
	getMajorOptionaljsonBody string

	getCourseChoosenjsonBody   string
	getCourseCancelledjsonBody string
}

func (list *urlList) init(mode int) {
	if mode == WEBVPN {
		list.baseUrl = "https://jwxt-443.webvpn.sysu.edu.cn/"
		list.loginURL = loginURL
		list.captchaURL = captchaURL
	} else if mode == NORMAL {
		list.baseUrl = "https://jwxt.sysu.edu.cn/"
		list.loginURL = loginURL1
		list.captchaURL = captchaURL1
	} else {
		log.Fatal("mode error")
	}
	baseUrl := list.baseUrl

	list.getGPAurl = getGPAurl
	list.courseListUrl = baseUrl + courseListUrl
	list.courseSelectionCancelUrl = baseUrl + courseSelectionCancelUrl
	list.courseSelectionChooseUrl = baseUrl + courseSelectionChooseUrl
	list.myTeacherInfoUrl = baseUrl + myTeacherInfoUrl
	list.teacherImgUrl = baseUrl + teacherImgUrl
	list.getScoreListUrl = baseUrl + getScoreListUrl
	list.teachersInfo = baseUrl + teachersInfo
	// 公选的参数，最后的studyCampusId指定校区
	list.getPublicCoursejsonBody = `{"pageNo":1,"pageSize":10,"param":{"semesterYear":"2020-1","selectedType":"4","selectedCate":"11","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","collectionStatus":"0","studyCampusId":"5063559"}}`
	// {"pageNo":1,"pageSize":10,"param":{"semesterYear":"2020-1","selectedType":"4","selectedCate":"11","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","hiddenEmptyStatus":"0","vacancySortStatus":"0","collectionStatus":"0"}}
	// {"pageNo":1,"pageSize":10,"param":{"semesterYear":"2020-1","selectedType":"4","selectedCate":"11","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","hiddenEmptyStatus":"0","vacancySortStatus":"0","collectionStatus":"0","studyCampusId":"5063559"}}
	list.getTeachersInfojsonBody = `{"pageNo":1,"pageSize":20,"total":true,"param":{"acadYear":"2020-1"}}`
	// 专选
	list.getMajorOptionaljsonBody = `{"pageNo":1,"pageSize":10,"param":{"semesterYear":"2020-1","selectedType":"1","selectedCate":"21","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","collectionStatus":"0"}}`
	list.getCourseChoosenjsonBody = `{"clazzId":"1208910925716574209","selectedType":"4","selectedCate":"21","check":true}` //<- 这是公选   // 专选选课的 json格式  {"clazzId":"1201412705275330561","selectedType":"1","selectedCate":"21","check":true}

	list.getCourseCancelledjsonBody = `{"courseId":"206169488","clazzId":"1201412705275330561","selectedType":"1"}` //专选
}



// 0: {campusName: "东校园", id: "5063559", campusNumber: "4"}
// 1: {campusName: "北校园", id: "5062202", campusNumber: "2"}
// 2: {campusName: "南校园", id: "5062201", campusNumber: "1"}
// 3: {campusName: "深圳校区", id: "333291143", campusNumber: "5"}
// 4: {campusName: "珠海校区", id: "5062203", campusNumber: "3"}