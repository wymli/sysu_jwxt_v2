package main

import (
	// "fmt"
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "regexp"

	"log"
	// "math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// 一键健康申报
// 使用portal登入   重定向到 html: http://jksb.sysu.edu.cn/infoplus/form/548809/render
//    referer :  "http://jksb.sysu.edu.cn/infoplus/form/XNYQSB/start" 获取cookie及流水号
const (
	originURI = "http://jksb.sysu.edu.cn/infoplus/form/XNYQSB/start"

	interfaceURI = "http://jksb.sysu.edu.cn/infoplus/interface/start"

	boundFields = `fieldCXXXjtgjbc,fieldMQJCRxh,fieldYQJLsfjchbfy,fieldCXXXsftjhb,fieldSTQKqt,fieldSTQKglsjrq,fieldYQJLjrsfczbldqzt,fieldCXXXjtfsqtms,fieldCXXXjtfsfj,fieldJBXXjjlxrdh,fieldJBXXxm,fieldJBXXjgsjtdz,fieldYQJLsfzhbwz,fieldSTQKfrtw,fieldMQJCRxm,fieldCXXXsftjhbq,fieldSTQKqtms,fieldCXXXjtfslc,fieldJBXXlxfs,fieldJBXXxb,fieldCXXXjtfspc,fieldYQJLsfjcqtbl,fieldCXXXssh,fieldJBXXgh,fieldCNS,fieldYC,fieldSTQKfl,fieldCXXXsftjwh,fieldCXXXfxxq,fieldSTQKdqstzk,fieldSTQKhxkn,fieldSTQKqtqksm,fieldFLid,fieldJBXXjggatj,fieldYQJLjrsfczbl,fieldJBXXjjlxr,fieldCXXXfxcfsj,fieldMQJCRcjdd,fieldSQSJ,fieldSTQKfrsjrq,fieldSTQKks,fieldJBXXcsny,fieldSTQKgm,fieldJBXXnj,fieldCXXXjtzzq,fieldJBXXJG,fieldCXXXdqszd,fieldCXXXjtzzs,fieldSTQKfx,fieldSTQKfs,fieldCXXXjtfsdb,fieldCXXXcxzt,fieldCXXXjtfshc,fieldCXXXjtjtzz,fieldCXXXsftjhbs,fieldJBXXsfzh,fieldSTQKsfstbs,fieldCXXXcqwdq,fieldJBXXfdygh,fieldJBXXjgshi,fieldJBXXfdyxm,fieldCXXXjtzz,fieldJBXXjgq,fieldCXXXjtfsqt,fieldJBXXjgs,fieldSTQKfrsjsf,fieldSTQKglsjsf,fieldJBXXdw,fieldCXXXsftjhbjtdz,fieldMQJCRlxfs`

	cookieURI = "https://cas-443.webvpn.sysu.edu.cn/cas/login?service=http%3A%2F%2Fjksb.sysu.edu.cn%2Finfoplus%2Flogin%3FretUrl%3Dhttp%253A%252F%252Fjksb.sysu.edu.cn%252Finfoplus%252Fform%252FXNYQSB%252Fstart"

	finalURI1 = "http://jksb.sysu.edu.cn/infoplus/interface/listNextStepsUsers"
	finalURI2 = "http://jksb.sysu.edu.cn/infoplus/interface/doAction"

	finalFormData = `这里写你的form data`

var (
	lang     = "zh"
	actionId = "1"
	formData string
)

func init_() {
	bytes, _ := ioutil.ReadFile("jksb_formdata.txt")
	formData = string(bytes)
	log.Println("表单数据:", formData)
}

func jksb(client *http.Client) {
	init_()
	// 获取jksb的cookie
	req, _ := http.NewRequest("GET", cookieURI, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	req.Header.Add("Referer", "https://portal.sysu.edu.cn/")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(bytes)))

	//从这个页面发起请求 , 得到重定向后的页面地址

	idc, _ := dom.Find("#idc").Attr("value")
	release, _ := dom.Find("#release").Attr("value")
	csrfToken, _ := dom.Find("meta[itemscope=csrfToken]").Attr("content")
	formData := map[string]string{
		"_VAR_URL":      originURI,
		"_VAR_URL_Attr": "{}",
	}
	bs, _ := json.Marshal(formData)
	// fmt.Println("formData:" , string(bs))

	postData := url.Values{
		"idc":       {idc},
		"release":   {release},
		"csrfToken": {csrfToken},
		"formData":  {string(bs)},
	}.Encode()
	// fmt.Println("Postdata", postData)
	req2, _ := http.NewRequest("POST", interfaceURI, strings.NewReader(postData))
	req2.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"},
		"Referer":      {"http://jksb.sysu.edu.cn/infoplus/form/XNYQSB/start"},
	}

	resp2, _ := client.Do(req2)
	defer resp2.Body.Close()
	bs, _ = ioutil.ReadAll(resp2.Body)
	receiver := map[string]interface{}{}
	fmt.Println(string(bs))
	json.Unmarshal(bs, &receiver)
	// fmt.Println(receiver["entities"])
	href := receiver["entities"].([]interface{})[0].(string)
	slice := strings.Split(href, "/")
	stepId := slice[5]
	log.Println("StepId:", stepId)

	//不同去请求href , 因为参数我们都有了,直接请求最终网址
	finalPostData := url.Values{
		"stepId":      {stepId},
		"actionId":    {"1"},
		"formData":    {finalFormData},
		"timestamp":   {fmt.Sprintf("%d", time.Now().Unix())},
		"rand":        {fmt.Sprintf("%f", rand.Float64()*999)},
		"boundFields": {boundFields},
		"csrfToken":   {csrfToken},
	}.Encode()

	req3, _ := http.NewRequest("POST", finalURI1, strings.NewReader(finalPostData))
	req3.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded; charset=UTF-8"},
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"},
		"Referer":      {href},
	}
	resp3, _ := client.Do(req3)
	defer resp3.Body.Close()
	bs, _ = ioutil.ReadAll(resp3.Body)
	fmt.Println(string(bs))
	var receiverFinal1 map[string]interface{}
	json.Unmarshal(bs, &receiverFinal1)
	state1 := receiverFinal1["entities"].([]interface{})[0].(map[string]interface{})["name"]
	if state1 == "办结" {
		log.Println("Step1:", state1)
	} else {
		log.Println("1 : 失败")
		log.Println(string(bs))
	}

	//下面是action  Math.random() * 999
	finalPostData2 := url.Values{
		"stepId":      {stepId},
		"actionId":    {"1"},
		"formData":    {finalFormData},
		"timestamp":   {fmt.Sprintf("%d", time.Now().Unix())},
		"rand":        {fmt.Sprintf("%f", rand.Float64()*999)},
		"boundFields": {boundFields},
		"csrfToken":   {csrfToken},
		"nextUsers":   {"{}"},
		"remark":      {},
	}.Encode()

	req4, _ := http.NewRequest("POST", finalURI2, strings.NewReader(finalPostData2))
	req4.Header = http.Header{
		"Content-Type": {"application/x-www-form-urlencoded; charset=UTF-8"},
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"},
		"Referer":      {href},
	}
	resp4, _ := client.Do(req3)
	defer resp4.Body.Close()
	bs, _ = ioutil.ReadAll(resp4.Body)
	var receiverFinal2 map[string]interface{}
	json.Unmarshal(bs, &receiverFinal2)
	state2 := receiverFinal2["entities"].([]interface{})[0].(map[string]interface{})["name"]
	if state2 == "办结" {
		log.Println("step2:", state2)
	} else {
		log.Println("2 : 失败")
		log.Println(string(bs))
	}

	return

}
