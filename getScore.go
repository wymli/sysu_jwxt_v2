package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	// "log"
	"net/http"
	"regexp"
)

type GPA struct {
	Meta struct {
		Success    bool   `json:"success"`
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
	} `json:"meta"`
	Data struct {
		SchoolYearGPA   float64 `json:"schoolYearGPA"`
		SchoolYearGrank int     `json:"schoolYearGrank"`
		AvgGPA          float64 `json:"avgGPA"`
		AvgGPARank      int     `json:"avgGPARank"`
		StuTotal        int     `json:"stuTotal"`
		Gpa             float64 `json:"gpa"`
	} `json:"data"`
}

func getGPA(client *http.Client) (string, error) {
	req, _ := http.NewRequest("GET", urlLists.getGPAurl, nil)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	var tmp GPA
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return "", errors.New("Get GPA Error")
	}
	reg := regexp.MustCompile(`\?.*\-.*`)
	yearAndSemester := reg.FindString(urlLists.getGPAurl)
	info := fmt.Sprintf("YearAndSemester:\t%s \nThis Semester's GPA:\t%f\tRank:%d\nAverage GPA:\t%f, Rank:%d",
		yearAndSemester, tmp.Data.SchoolYearGPA, tmp.Data.SchoolYearGrank, tmp.Data.AvgGPA, tmp.Data.AvgGPARank)
	return info, nil
}

type scoreListStruct struct {
	Code int `json:"code"`
	Data []struct {
		ScoSchoolYear   string `json:"scoSchoolYear"`
		ScoSemester     string `json:"scoSemester"`
		ScoCourseNumber string `json:"scoCourseNumber"`
		ScoCourseName   string `json:"scoCourseName"`
		// ScoCourseCategory     string  `json:"scoCourseCategory"`
		// ScoCourseCategoryName string  `json:"scoCourseCategoryName"`
		ScoCredit float64 `json:"scoCredit"`
		// ScoStudentNumber      string  `json:"scoStudentNumber"`
		// TeachClassNumber      string  `json:"teachClassNumber"`
		// OriginalScore         string  `json:"originalScore"`
		ScoFinalScore  string  `json:"scoFinalScore"`
		ScoPoint       float64 `json:"scoPoint"`
		ScoTeacherName string  `json:"scoTeacherName"`
		// AccessFlag            string  `json:"accessFlag"`
		// RecordStyle           string  `json:"recordStyle"`
		// ExamCharacter         string  `json:"examCharacter"`
		TeachNumber      string `json:"teachNumber"`
		TeachClassRank   string `json:"teachClassRank"`
		GradeMajorNumber string `json:"gradeMajorNumber"`
		GradeMajorRank   string `json:"gradeMajorRank"`
		// ScoreList             []struct {
		// 	FXCJ string `json:"FXCJ"`
		// 	FXMC string `json:"FXMC"`
		// 	MRQZ int    `json:"MRQZ"`
		// } `json:"scoreList"`
	} `json:"data"`
}

func getScoreList(client *http.Client) string {
	//get方法
	// ?_t=1580453973005&addScoreFlag=true&scoSchoolYear=2019-2020&scoSemester=1&trainTypeCode=01
	req, _ := http.NewRequest("GET", urlLists.getScoreListUrl, nil)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Referer", urlLists.baseUrl+"jwxt/ang/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	var scoreList scoreListStruct
	json.Unmarshal(b, &scoreList)
	bb, _ := json.MarshalIndent(scoreList, "", "  ")
	return string(bb)
}
