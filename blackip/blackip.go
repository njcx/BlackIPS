package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
)

var url1 = "http://127.0.0.1:9091/check?ip="                   // 威胁情报接口,修改为真实IP
var url2 = "http://127.0.0.1:9093/iplocation/getip/ip.php?ip=" // IP地理位置接口，修改为真实IP

type result struct {
	Code    int    `json:"code"`
	Ip      string `json:"ip"`
	Message string `json:"message"`
}

type result_1 struct {
	Country  string `json:"country"`
	Ip       string `json:"ip"`
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
	Isp      string `json:"isp"`
	Area     string `json:"area"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.StaticFS("/assets", http.Dir("./assets"))
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"country": "Please enter the IP.",
			"area":    "Please enter the IP.",
			"lable":   "Please enter the IP.",
		})
	})

	r.GET("gets", func(c *gin.Context) {

		ip := c.Query("ip")
		resp, err := http.Get(url1 + ip)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var res result
		_ = json.Unmarshal(body, &res)

		resp1, err := http.Get(url2 + ip)
		if err != nil {
			return
		}
		defer resp1.Body.Close()
		body1, _ := ioutil.ReadAll(resp1.Body)
		var res1 result_1
		_ = json.Unmarshal(body1, &res1)

		if res.Code == 1 {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"country": template.HTML("<font color=\"#FF0000\">" + ip + "--" + res1.Country +
					"--" + res1.Province + "--" + res1.City + "--" + res1.County + "</font>"),
				"area":  template.HTML("<font color=\"#FF0000\">" + ip + "--" + res1.Isp + "--" + res1.Area + "</font>"),
				"lable": template.HTML("<font color=\"#FF0000\">" + ip + "--该IP为恶意--" + res.Message + "</font>"),
			})
		}

		if res.Code == 0 {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"country": ip + "--" + res1.Country + "--" + res1.Province + "--" + res1.City + "--" + res1.County,
				"area":    ip + "--" + res1.Isp + "--" + res1.Area,
				"lable":   ip + "--IP 状态未知",
			})
		}
	})
	r.Run(":80")
}
