package main

import (
	"fmt"
	"net/http"
	"github.com/json-iterator/go"
	"github.com/go-redis/redis"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary


var client = redis.NewClient(&redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})


type JsonR struct {
	Code    int         `json:"code"`
	Ip      string      `json:"ip"`
	Message string      `json:"message"`
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

	if len(val2)>1 {
		result.Code = 1
		result.Ip  = ip
		result.Message = val2
	} else {
		result.Code = 0
		result.Ip  = ip
		//result.Message =
	}
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}


func main() {
	fmt.Println("This is web server~")
	http.HandleFunc("/check", checkIP)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Listen And Serve error: ", err.Error())
	}
}
