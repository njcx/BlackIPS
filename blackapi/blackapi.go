package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/json-iterator/go"
	"net/http"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var client = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6381",
	Password: "123456",
	DB:       0,
})

type JsonR struct {
	Code    int    `json:"code"`
	Ip      string `json:"ip"`
	Message string `json:"message"`
}

func JsonRF() *JsonR {
	return &JsonR{}
}

func checkIP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	param_ip, found := req.Form["ip"]
	if !(found) {
		fmt.Fprint(w, "请添加参数ip")
		return
	}
	result := JsonRF()
	ip := param_ip[0]
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "--CheckIP  is running--" + ip)
	val2, _ := client.Get(ip).Result()

	if len(val2) > 1 {
		result.Code = 1
		result.Ip = ip
		result.Message = val2
	} else {
		result.Code = 0
		result.Ip = ip
		//result.Message =
	}
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}


func delIP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	param_ip, found := req.Form["ip"]
	if !(found) {
		fmt.Fprint(w, "请添加参数ip")
		return
	}
	result := JsonRF()
	ip := param_ip[0]
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "--Add Whitelist--" + ip)
	val2, _ := client.Del(ip).Result()

	if val2 == 1 {
		result.Code = 1
		result.Ip = ip
		result.Message = "IP已加白"
	} else {
		result.Code = 0
		result.Ip = ip
		result.Message = "IP不存在"
	}
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}



func main() {
	fmt.Println("This is web server~")
	http.HandleFunc("/check", checkIP)
	http.HandleFunc("/whitelist", delIP)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Listen And Serve error: ", err.Error())
	}
}