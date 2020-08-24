package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"strconv"
	// "time"

	// "os"
	"strings"
)

type row struct {
	MainClassesID        string      `json:"mainClassesID"`
	TeachingClassID      string      `json:"teachingClassId"`
	TeachingClassNum     string      `json:"teachingClassNum"`
	TeachingClassName    interface{} `json:"teachingClassName"`
	CourseNum            string      `json:"courseNum"`
	CourseName           string      `json:"courseName"`
	Credit               float64     `json:"credit"`
	ExamFormName         string      `json:"examFormName"`
	CourseUnitNum        string      `json:"courseUnitNum"`
	CourseUnitName       string      `json:"courseUnitName"`
	TeachingTeacherNum   string      `json:"teachingTeacherNum"`
	TeachingTeacherName  string      `json:"teachingTeacherName"`
	BaseReceiveNum       int         `json:"baseReceiveNum"`
	AddReceiveNum        int         `json:"addReceiveNum"`
	TeachingTimePlace    string      `json:"teachingTimePlace"`
	StudyCampusID        string      `json:"studyCampusId"`
	Week                 string      `json:"week"`
	ClassTimes           string      `json:"classTimes"`
	CourseSelectedNum    string      `json:"courseSelectedNum"`
	FilterSelectedNum    string      `json:"filterSelectedNum"`
	SelectedStatus       string      `json:"selectedStatus"`
	CollectionStatus     string      `json:"collectionStatus"`
	TeachingLanguageCode string      `json:"teachingLanguageCode"`
	PubCourseTypeCode    interface{} `json:"pubCourseTypeCode"`
	CourseCateCode       string      `json:"courseCateCode"`
	SpecialClassCode     interface{} `json:"specialClassCode"`
	SportItemID          interface{} `json:"sportItemId"`
	RecordMode           string      `json:"recordMode"`
	ClazzNum             string      `json:"clazzNum"`
	ExamFormCode         string      `json:"examFormCode"`
	CourseID             string      `json:"courseId"`
	ScheduleExamTime     interface{} `json:"scheduleExamTime"`
}

type courseSelectList struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    struct {
		Total int   `json:"total"`
		Rows  []row `json:"rows"`
	} `json:"data"`
}

type courseInfo struct {
	Code int `json:"code"`
	Data struct {
		OutlineInfo struct {
			Creator                       string  `json:"creator"`
			CreatorName                   string  `json:"creatorName"`
			CreateTime                    string  `json:"createTime"`
			Editor                        string  `json:"editor"`
			EditorName                    string  `json:"editorName"`
			EditeTime                     string  `json:"editeTime"`
			OutlineCourseInfoID           string  `json:"outlineCourseInfoId"`
			CourseID                      string  `json:"courseId"`
			CourseNum                     string  `json:"courseNum"`
			OutlineType                   string  `json:"outlineType"`
			CourseType                    string  `json:"courseType"`
			CourseTypeName                string  `json:"courseTypeName"`
			LanguageNum                   string  `json:"languageNum"`
			PlanClassSize                 string  `json:"planClassSize"`
			ReferenceBook                 string  `json:"referenceBook"`
			CourseContentInChinese        string  `json:"courseContentInChinese"`
			CourseObjectiveAndRequirement string  `json:"courseObjectiveAndRequirement"`
			TeachMethod                   string  `json:"teachMethod"`
			EvaluationMethod              string  `json:"evaluationMethod"`
			CurrentFlowNum                string  `json:"currentFlowNum"`
			AuditStatus                   string  `json:"auditStatus"`
			GiveLessSemtster              string  `json:"giveLessSemtster"`
			LecturesCreHours              float64 `json:"lecturesCreHours"`
			TutorialsCreHours             float64 `json:"tutorialsCreHours"`
			LabCreHours                   float64 `json:"labCreHours"`
			OtherLectureCreHours          float64 `json:"otherLectureCreHours"`
			CourseResource                string  `json:"courseResource"`
			LastAuditStatus               string  `json:"lastAuditStatus"`
			SubCourseType                 string  `json:"subCourseType"`
			SubCourseTypeName             string  `json:"subCourseTypeName"`
			CourseName                    string  `json:"courseName"`
			CourseEngName                 string  `json:"courseEngName"`
			Credit                        string  `json:"credit"`
			TotalHours                    string  `json:"totalHours"`
			EstablishUnitNumberName       string  `json:"establishUnitNumberName"`
		} `json:"outlineInfo"`
		ScheduleList []struct {
			OutlineTeachingScheduleID string `json:"outlineTeachingScheduleId"`
			OutlineCourseInfoID       string `json:"outlineCourseInfoId"`
			WeekNum                   int    `json:"weekNum"`
			TeachingMainContent       string `json:"teachingMainContent"`
			TeachingHours             string `json:"teachingHours"`
			Sort                      string `json:"sort"`
		} `json:"scheduleList"`
		TeacherList []struct {
			Editor               string `json:"editor"`
			EditeTime            string `json:"editeTime"`
			ID                   string `json:"id"`
			TeacherNum           string `json:"teacherNum"`
			DepartmentNum        string `json:"departmentNum"`
			Name                 string `json:"name"`
			NameSpell            string `json:"nameSpell"`
			IDCardTypeNum        string `json:"idCardTypeNum"`
			IDCardNum            string `json:"idCardNum"`
			GenderNum            string `json:"genderNum"`
			BirthDate            string `json:"birthDate"`
			NationNum            string `json:"nationNum"`
			NationalityNum       string `json:"nationalityNum"`
			PoliticsNum          string `json:"politicsNum"`
			TeacherTypeNum       string `json:"teacherTypeNum"`
			ProfessionNum        string `json:"professionNum"`
			ProfessionNumName    string `json:"professionNumName"`
			BestEducationNum     string `json:"bestEducationNum"`
			BestEducationNumName string `json:"bestEducationNumName"`
			ThisStateNum         string `json:"thisStateNum"`
			InEstablishment      string `json:"inEstablishment"`
			OnDuty               string `json:"onDuty"`
			CampusID             string `json:"campusId"`
			Email                string `json:"email"`
			MobilePhone          string `json:"mobilePhone"`
			BestDegree           string `json:"bestDegree"`
			BestDegreeName       string `json:"bestDegreeName"`
			StationType          string `json:"stationType"`
			EmployTime           string `json:"employTime"`
		} `json:"teacherList"`
	} `json:"data"`
}

func (client *Client) getCourseList(payload string) ([]row, error) {
	// jsonBody := `{"pageNo":1,"pageSize":10,"param":{"semesterYear":"2019-2","selectedType":"4","selectedCate":"11","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","collectionStatus":"0"}}`
	// jsonBody := `{"pageNo":1,"pageSize":10,"param":{"semesterYear":"2019-2","selectedType":"1","selectedCate":"21","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","collectionStatus":"0"}}`

	//时间戳,不加也行
	// timestamp := time.Now().Unix()
	// var courseListUrl_t =  courseListUrl + "?_t=" + fmt.Sprintf("%d", timestamp)

	log.Println("courseListUrl : ", urlLists.courseListUrl)
	log.Println("Query params :", payload)
	courseListReq, _ := http.NewRequest("POST", urlLists.courseListUrl, strings.NewReader(payload))
	courseListReq.Header.Add("Content-Type", "application/json;charset=UTF-8")
	courseListReq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	courseListReq.Header.Add("Referer", urlLists.baseUrl+"jwxt/mk/courseSelection/")
	courseListResp, _ := client.Do(courseListReq)
	defer courseListResp.Body.Close()
	b, _ := ioutil.ReadAll(courseListResp.Body)
	// ioutil.WriteFile("a.json", b, 0777)
	courseList := courseSelectList{}
	err := json.Unmarshal(b, &courseList)
	if err != nil {
		return nil, err
	}

	totalCourses := courseList.Data.Rows
	//寻找有空位的课
	n_courses := courseList.Data.Total
	var times = n_courses / 10 // 10 is one page size
	for i := 0; i < times; i++ {
		payload = strings.ReplaceAll(payload, `"pageNo":`+fmt.Sprintf("%d", i+1), `"pageNo":`+fmt.Sprintf("%d", i+2))
		courseListReq2, _ := http.NewRequest("POST", urlLists.courseListUrl, strings.NewReader(payload))
		courseListReq2.Header.Add("Content-Type", "application/json;charset=UTF-8")
		courseListReq2.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
		courseListReq2.Header.Add("Referer", urlLists.baseUrl+"jwxt/mk/courseSelection/")

		resp2, _ := client.Do(courseListReq2)
		defer resp2.Body.Close()
		b, _ := ioutil.ReadAll(resp2.Body)
		courseList := courseSelectList{}
		err := json.Unmarshal(b, &courseList)
		if err != nil {
			return nil, err
		}
		totalCourses = append(totalCourses, courseList.Data.Rows...)
	}
	return totalCourses, nil
}

func (client *Client) courseChoose(clazzId, selectedType, selectedCate string) (bool, string) {
	// courseSelectionChooseBody := `{"clazzId":"1201412705275330561","selectedType":"1","selectedCate":"21","check":true}` //专选 ,classid是teachingclassid
	// {"clazzId":"1208910925716574209","selectedType":"4","selectedCate":"21","check":true} //公选
	// {"code":52021104,"message":"你已选择过该课程，不能再选！","data":null}
	// {"code":200,"message":null,"data":"选课成功!"}
	// {"code":52021107,"message":"不能超过公选课限选的最大门数！","data":null}
	template := `{"clazzId":"%s","selectedType":"%s","selectedCate":"%s","check":true}`
	payload := fmt.Sprintf(template, clazzId, selectedType, selectedCate)
	log.Println("Query params :", payload)
	chooseReq, _ := http.NewRequest("POST", urlLists.courseSelectionChooseUrl, strings.NewReader(payload))
	chooseReq.Header.Add("Content-Type", "application/json;charset=UTF-8")
	chooseReq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	chooseReq.Header.Add("Referer", urlLists.baseUrl+"jwxt/mk/courseSelection/")
	chooseResp, _ := client.Do(chooseReq)
	defer chooseResp.Body.Close()
	b, _ := ioutil.ReadAll(chooseResp.Body)
	log.Println(string(b))
	var data map[string]interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return false, err.Error()
	}
	if data == nil {
		return false, "Json Parse Error"
	} else {
		if data["code"].(float64) != 200 {
			return false, data["message"].(string)
		} else {
			return true, data["data"].(string)
		}
	}
}

func (client *Client) courseCancel(courseId, clazzId, selectedType string) (bool, string) {
	// courseSelectionCancelBody := `{"courseId":"206169488","clazzId":"1201412705275330561","selectedType":"1"}`
	// {"code":200,"message":null,"data":"退课成功！"}
	// 多次退课都是退课成功
	template := `{"courseId":"%s","clazzId":"%s","selectedType":"%s"}`
	payload := fmt.Sprintf(template, courseId, clazzId, selectedType)

	log.Println("Query params :", payload)
	cancelReq, _ := http.NewRequest("POST", urlLists.courseSelectionCancelUrl, strings.NewReader(payload))
	cancelReq.Header.Add("Content-Type", "application/json;charset=UTF-8")
	cancelReq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	cancelReq.Header.Add("Referer", urlLists.baseUrl+"jwxt/mk/courseSelection/")
	cancelResp, _ := client.Do(cancelReq)
	defer cancelResp.Body.Close()
	b, _ := ioutil.ReadAll(cancelResp.Body)
	log.Println(string(b))
	var data map[string]interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return false, err.Error()
	}
	if data == nil {
		return false, "Json Parse Error"
	} else {
		if data["code"].(float64) != 200 {
			return false, data["message"].(string)
		} else {
			return true, data["data"].(string)
		}
	}
}

// func (client *Client) grabCourse(payload string, timeSeperate int) {
// 	log.Println("开始抢课---|->")
// 	for {
// 		resp := client.courseChoose(payload)
// 		if resp["code"].(float64) == 200 {
// 			log.Println("抢课成功")
// 			break
// 		} else { // 52021104  52021107
// 			log.Println("失败,睡眠" + fmt.Sprintf("%d", timeSeperate) + "s")
// 			time.Sleep(time.Duration(timeSeperate) * time.Second)
// 		}
// 	}
// }

// example: {"pageNo":1,"pageSize":10,"param":{"semesterYear":"2020-1","selectedType":"4","selectedCate":"11","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","hiddenEmptyStatus":"0","vacancySortStatus":"0","collectionStatus":"0","studyCampusId":"5063559"}}
// selectedType == 1 时, selectedCate才有效
func constructPayload(pageNo, pageSize int, semesterYear, selectedType, selectedCate, campusId string) string {
	var template string = `{"pageNo":%d,"pageSize":%d,"param":{"semesterYear":"%s","selectedType":%s,"selectedCate":%s,"hiddenConflictStatus":"0","hiddenSelectedStatus":"0","hiddenEmptyStatus":"0","vacancySortStatus":"0","collectionStatus":"0","studyCampusId":%s}}`
	payload := fmt.Sprintf(template, pageNo, pageSize, semesterYear, selectedType, selectedCate, campusId)
	return payload
}

func (client *Client) getCourseInfo(courseNum string) []byte {
	url := urlLists.baseUrl + "jwxt.sysu.edu.cn/jwxt/training-programe/courseoutline/getalloutlineinfo?courseNum="
	req, _ := http.NewRequest("GET", url+courseNum, nil)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Add("Referer", urlLists.baseUrl+"jwxt/mk/courseSelection/")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}

func (client *Client) _courseAddCollection(classId, selectedType string) error {
	template := `{"classesID":"%s","selectedType":"%s"}`
	payload := fmt.Sprintf(template, classId, selectedType)
	url := "https://jwxt.sysu.edu.cn/jwxt/choose-course-front-server/stuCollectedCourse/create"
	req, _ := http.NewRequest("GET", url, strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	if data["code"].(float64) != 200 {
		return errors.New("_courseAddCollection:Error:resp_code!=200")
	}
	return nil
}

func (client *Client) _courseDeleteCollection(classId, selectedType string) error {
	err := client._courseAddCollection(classId, selectedType)
	if err != nil {
		errors.New(err.Error() + " in _courseDeleteCollection")
	}
	return nil
}

func (client *Client) getCollectedCourseList(classId, selectedType, selectedCate string) ([]row, error) {
	template := `{"pageNo":1,"pageSize":10,"param":{"semesterYear":"2020-1","selectedType":"%s","selectedCate":"%s","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","hiddenEmptyStatus":"0","vacancySortStatus":"0","collectionStatus":"1","studyCampusId":"5063559"}}`
	payload := fmt.Sprintf(template, selectedType, selectedCate)
	rows, err := client.getCourseList(payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rows, nil
}

// getOneCourseInLoop
func (client *Client) createTimeTask(classId, selectedType, selectedCate, freq, dur string) (bool, error) {
	freq_, _ := strconv.Atoi(freq)
	dur_, _ := strconv.Atoi(dur)
	total_cnt := dur_ / freq_
	cnt := 0
	isCollected, err := client.isCourseCollected(classId, selectedType, selectedCate)
	if err != nil {
		log.Println(err)
		return false, err
	}
	if !isCollected {
		err = client._courseAddCollection(classId, selectedType)
		if err != nil {
			log.Println(err.Error())
			return false, err
		}
	}

	timeOver := false

	// init thread state
	threadState[classId] = THREAD_RUN
	timeTaskState[classId] = TIME_TASK_FAIL

	for {
		// check if cancel ; 无需加锁,此线程只读,最多多一次循环
		if threadState[classId] == THREAD_STOP {
			break
		}

		cnt++
		if cnt >= total_cnt {
			timeOver = true
			break
		}
		isAvailable, err := client.isCourseAvailable(classId, selectedType, selectedCate)
		if err != nil {
			log.Println("createTimeTask: getOneCourse in loop:", err.Error())
			return false, err
		}
		if isAvailable {
			isOk, msg := client.courseChoose(classId, selectedType, selectedCate)
			if isOk {
				timeTaskState[classId] = TIME_TASK_SUCCESS
				return true, nil
			} else {
				log.Println("TimeTask: classId:", classId, "查询到空位,选课失败:", msg)
			}
		}
		time.Sleep(time.Second * time.Duration(freq_))
	}
	threadState[classId] = THREAD_STOP

	err = client._courseDeleteCollection(classId, selectedType)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	if isCollected, _ = client.isCourseCollected(classId, selectedType, selectedCate); isCollected {
		log.Println("courseDeleteCollection : err")
	}
	if timeOver {
		return false, errors.New("TimeLimit: 定时任务到时,选课失败")
	}
	if threadState[classId] == THREAD_STOP {
		return false, errors.New("DeleteTimeTask: cancelled initiativly")
	}
	return true, nil
}

func (client *Client) deleteTimeTask(classId string) (bool, error) {
	threadState[classId] = THREAD_STOP
	return true, nil
}

func (client *Client) reportTimeTask(classId string) map[string]interface{} {
	return map[string]interface{}{
		"isThreadRun":       threadState[classId] == THREAD_RUN,
		"isTimeTaskSuccess": timeTaskState[classId] == TIME_TASK_SUCCESS,
	}
}

func (client *Client) isCourseCollected(classId, selectedType, selectedCate string) (bool, error) {
	row, err := client.getCollectedCourseList(classId, selectedType, selectedCate)
	for _, v := range row {
		if v.TeachingClassID == classId {
			return true, err
		}
	}
	return false, err
}

func (client *Client) isCourseAvailable(classId, selectedType, selectedCate string) (bool, error) {
	row, err := client.getCollectedCourseList(classId, selectedType, selectedCate)
	for _, v := range row {
		if v.TeachingClassID == classId {
			if fmt.Sprintf("%s", v.BaseReceiveNum) == v.CourseSelectedNum {
				return false, err
			} else {
				return true, err
			}
		}
	}
	return false, err
}
