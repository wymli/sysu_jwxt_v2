package main

import (
	"io/ioutil"
	"net/http"
)

func (c *Client) getStudentImg(id string) []byte {
	url := "https://jwxt.sysu.edu.cn/jwxt/student-status/stu-photo/photo?photoType=1&stuNumber=" + id
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	req.Header.Add("Referer", "https://jwxt.sysu.edu.cn/jwxt/mk/studentWeb/")
	resp, _ := c.Do(req)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	return b
}
