package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	// "strconv"
	// "time"

	// "strconv"

	// "net/http"

	"github.com/gin-gonic/gin"

	// "net/http"
	// "log"
	// "net/http"
	// "os"
	"io/ioutil"
)

func login(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		if client.isLogin {
			ctx.Redirect(302, "/index")
			return
		}
		err := client.casFirstGet("./view/pic/captcha.jpg")
		if err != nil {
			ctx.HTML(200, "errorPage.html", err.Error())
			log.Println("casFirstGet err")
		} else {
			ctx.HTML(200, "login.html", nil)
		}
		return
	}
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var user map[string]string
	json.Unmarshal(data, &user)
	username, _ := user["username"]
	password, _ := user["password"]
	captcha, _ := user["captcha"]
	fmt.Println("username:", username, "passwrod:", password, "captcha:", captcha)
	if username == "" || password == "" {
		//relogin
		fmt.Println("Empty Username or Password")
		ctx.JSON(200, gin.H{"state": "fail", "msg": "Empty Username or Password"})
		return
	}
	if captcha == "" {
		fmt.Println("Empty Captcha")
		ctx.JSON(200, gin.H{"state": "fail", "msg": "Empty Captcha"})
		return
	}
	// do real login
	client.setUser(username, password)
	client.setCaptcha(captcha)
	err := client.casLogin()
	if err != nil {
		ctx.JSON(200, gin.H{"state": "fail", "msg": err.Error()})
		return
	}

	// GetMyTeachersInfo
	// caslogin.Get
	if client.checkLoginStatus() {
		// redirect
		client.isLogin = true
		log.Println("Login OK")
		ctx.JSON(200, gin.H{"state": "success", "msg": "ok"})
	} else {
		log.Println("Login Fail")
		ctx.JSON(200, gin.H{"state": "fail", "msg": "Wrong Username or Password"})
	}

	return
}

func courseList(ctx *gin.Context) {
	if client.ifNotLoginAndReturn(ctx) {
		return
	}

	template := `{"pageNo":1,"pageSize":10,"param":{"semesterYear":"2020-1","selectedType":"%s","selectedCate":"21","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","hiddenEmptyStatus":"0","vacancySortStatus":"0","collectionStatus":"0","studyCampusId":"5063559"}}`
	var payload string
	if ctx.Query("type") == "public" {
		payload = fmt.Sprintf(template, getSelectedType("校级公选"))
	} else if ctx.Query("type") == "major" {
		payload = fmt.Sprintf(template, getSelectedType("本专业"))
	} else {
		fmt.Println("Unknown query: type=", ctx.Query("type"))
		return
	}

	rows, err := client.getCourseList(payload)
	fmt.Println("len(rows)_", ctx.Query("type"), " is:", len(rows))
	// fmt.Println("[0]:", rows[0])
	if err != nil {
		ctx.JSON(200, gin.H{"state": "fail", "msg": "getCourseList Error :" + err.Error()})
		log.Println("getCourseList error", err)
		return
	}
	ctx.JSON(200, gin.H{"state": "success", "msg": "ok", "data": rows})
}

func index(ctx *gin.Context) {
	// client.ifNotLoginAndReturn(ctx)
	// ctx.Redirect(http.StatusMovedPermanently, "/")
	// ctx.JSON(200, "Not Login")
	// ctx.HTML(200, "index.html", nil)

	// 这里只能自己写数据,如果使用ctx.HTML,那么默认是使用gin的模板,如果定义了{{}},会与vue的重复导致错误
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	src, _ := os.Open("./view/templates/index.html")
	io.Copy(ctx.Writer, src)
	return
}

func userInfo(ctx *gin.Context) {
	// if client.ifNotLoginAndReturn(ctx) {
	// 	return
	// }
	if client.isLogin == false {
		ctx.JSON(200, gin.H{"state": "fail", "msg": "Not Login"})
		fmt.Println("studentInfo :", "Not Login")
		return
	}
	info := client.getStudentInfo()
	ctx.JSON(200, gin.H{"state": "success", "msg": "ok", "data": info})
	fmt.Println("studentInfo :", info)
	return
}

func teacherInfo_img(ctx *gin.Context) {
	courseNum := ctx.Query("courseNum")
	teacherId := ctx.Query("id")
	if teacherId == "" {
		info := client.getTeacherInfo(courseNum, "")
		if info == nil {
			ctx.JSON(200, gin.H{"state": "fail", "msg": "无课程大纲"})
			return
		}
		teacherId = info["id"].(string)
	}
	fmt.Println("CourseNum:", courseNum, "teacherId:", teacherId)
	b, err := client.getTeacherImg(teacherId, "")
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(ctx.Writer, bytes.NewReader(b))
}

func teacherInfo_email(ctx *gin.Context) {
	courseNum := ctx.Query("courseNum")
	if courseNum == "" {
		fmt.Println("Get Email: NO query")
		ctx.JSON(200, gin.H{"state": "fail", "msg": "No Query"})
	}
	info := client.getTeacherInfo(courseNum, "")
	if info == nil {
		ctx.JSON(200, gin.H{"state": "fail", "msg": "无课程大纲"})
	}
	teacherEmail := info["email"].(string)
	fmt.Println("CourseNum:", courseNum, "teacherEmail:", teacherEmail)
	ctx.JSON(200, gin.H{"state": "success", "msg": "ok", "data": teacherEmail})
}

func teacherInfo_all(ctx *gin.Context) {
	courseNum := ctx.Query("courseNum")
	info := client.getTeacherInfo(courseNum, "")
	fmt.Println("CourseNum:", courseNum, " | teacherInfo_ALL")

	if info == nil {
		ctx.JSON(200, gin.H{"state": "fail", "msg": "Not Exist"})
	} else {
		ctx.JSON(200, gin.H{"state": "success", "msg": "ok", "data": info})
	}
}

func courseInfo_handler(ctx *gin.Context) {
	courseNum := ctx.Query("courseNum")
	if courseNum != "" {
		ctx.JSON(200, gin.H{"state": "fail", "msg": "ok", "data": client.getCourseInfo(courseNum)})
	} else {
		ctx.JSON(200, gin.H{"state": "fail", "msg": "wrong query parameters"})
	}
}

func courseChooseHandler(ctx *gin.Context) {
	clazzId := ctx.Query("clazzId")
	selectedType := ctx.Query("selectedType")
	selectedCate := ctx.Query("selectedCate")
	ok, msg := client.courseChoose(clazzId, selectedType, selectedCate)
	if ok {
		ctx.JSON(200, gin.H{"state": "success", "msg": msg})
	} else {
		ctx.JSON(200, gin.H{"state": "fail", "msg": msg})
	}
}

func courseCancelHandler(ctx *gin.Context) {
	clazzId := ctx.Query("clazzId")
	selectedType := ctx.Query("selectedType")
	courseId := ctx.Query("courseId")
	ok, msg := client.courseCancel(courseId, clazzId, selectedType)
	if ok {
		ctx.JSON(200, gin.H{"state": "success", "msg": msg})
	} else {
		ctx.JSON(200, gin.H{"state": "fail", "msg": msg})
	}
}

func selectCourseInfoHandler(ctx *gin.Context) {
	data := client.selectCourseInfo()
	if data == nil {
		return
	}
	realCode := data["data"].(map[string]interface{})["code"]
	if data["code"].(float64) == 200 && realCode.(float64) == 200 {
		// ok
		log.Println(data["data"].(map[string]interface{})["electiveCourseStageName"].(string))
		ctx.JSON(200, gin.H{"state": "success", "data": data["data"]})
	} else {
		ctx.JSON(200, gin.H{"state": "fail", "data": data["data"]})
	}
}

func studentImgHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	img := client.getStudentImg(id)
	img_base64 := base64.RawStdEncoding.EncodeToString(img)
	// fmt.Println(img_base64)
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(200, `<form action="#"><span>id:</span><input name="id"></input><br><button type="submit">搜索</button></form><p>学号:%s</p><img src="data:image/jpeg;charset=utf-8;base64,%s" ></img>`, id, img_base64)
}

// func timeTaskCreateHandler(ctx *gin.Context) {
// 	return
// 	clazzId := ctx.Query("clazzId")
// 	selectedType := ctx.Query("selectedType")
// 	selectedCate := ctx.Query("selectedCate")
// 	freq := ctx.Query("freq")
// 	dur := ctx.Query("dur")
// 	ok, err := client.createTimeTask(clazzId, selectedType, selectedCate, freq, dur)
// 	if ok {
// 		ctx.JSON(200, gin.H{"state": "success", "msg": "ok"})
// 	} else {
// 		ctx.JSON(200, gin.H{"state": "fail", "msg": err.Error()})
// 	}
// 	return
// }

// func timeTaskDeleteHandler(ctx *gin.Context) {
// 	return
// 	clazzId := ctx.Query("clazzId")
// 	ok, err := client.deleteTimeTask(clazzId)
// 	if ok {
// 		ctx.JSON(200, gin.H{"state": "success", "msg": "ok"})
// 	} else {
// 		ctx.JSON(200, gin.H{"state": "fail", "msg": err.Error()})
// 	}
// 	return
// }

// func timeTaskReportHandler(ctx *gin.Context) {
// 	clazzId := ctx.Query("clazzId")
// 	msg := client.reportTimeTask(clazzId)
// 	ctx.JSON(200, gin.H{"state": "success", "msg": msg})
// }

// func courseRefreshTest(ctx *gin.Context) {
// 	totalCnt_ := ctx.Query("totalCnt")
// 	totalCnt, err := strconv.Atoi(totalCnt_)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	log.Println("in courseRefresh Test")
// 	template := `{"pageNo":1,"pageSize":10,"param":{"semesterYear":"2020-1","selectedType":"%s","selectedCate":"21","hiddenConflictStatus":"0","hiddenSelectedStatus":"0","hiddenEmptyStatus":"0","vacancySortStatus":"0","collectionStatus":"0","studyCampusId":"5063559"}}`
// 	log.Println("Begin Refresh TEST")
// 	beg := time.Now().Unix()
// 	for i := 0; i < totalCnt; i++ {
// 		rows, err := client.getCourseList(fmt.Sprintf(template, getSelectedType("校级公选")))
// 		log.Printf("[Test %d] len(rows) = %d\n", i, len(rows))
// 		if len(rows) == 0 {
// 			log.Println("Something Error")
// 		}
// 		if err != nil {
// 			log.Println("ERROR: ", err)
// 		}
// 	}
// 	end := time.Now().Unix()
// 	log.Println("----------------------------")
// 	log.Println("Success with for-loop:", totalCnt, " in ", end-beg, "s")
// }

func noHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"hello": "world"})
}

func main() {
	if client == nil {
		log.Fatal("client == nil")
	}
	rt := gin.Default()
	// rt.LoadHTMLGlob("./view/templates/*")
	rt.LoadHTMLFiles("./view/templates/login.html", "./view/templates/errorPage.html")
	rt.Static("/view/pic", "./view/pic")
	rt.GET("/", noHandler)
	rt.GET("/hello", noHandler)
	rt.POST("/login", login)
	rt.GET("/login", login)
	// rt.GET("/index", index)
	// rt.GET("/userInfo", userInfo)
	// rt.GET("/courseList", courseList)
	// rt.GET("/courseInfo", courseInfo_handler)
	// rt.GET("/teacherInfo/img", teacherInfo_img)
	// rt.GET("/teacherInfo/email", teacherInfo_email)
	// rt.GET("/teacherInfo/all", teacherInfo_all)
	// rt.GET("/course/choose", courseChooseHandler)
	// rt.GET("/course/cancel", courseCancelHandler)
	// rt.GET("/course/selectInfo", selectCourseInfoHandler)
	rt.GET("/student/img", studentImgHandler)
	// rt.GET("/course/timeTask/create", timeTaskCreateHandler)
	// rt.GET("/course/timeTask/delete", timeTaskDeleteHandler)
	// rt.GET("/course/timeTask/report", timeTaskReportHandler)
	// rt.GET("/course/refreshTest", courseRefreshTest)
	// rt.GET("/captcha", captcha)
	// rt.POST("/login", loginHandler)
	// rt.GET("/index", indexHandler)

	rt.Run(":9090")
}
