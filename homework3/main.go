package main

import ( 
     	"encoding/json"
          "io/ioutil"
        "net/http"
		"fmt"
)

type Weather struct{
	 Data  WeatherInfo  `json:"data"`
	 Status int          `json:"status"`
	 Desc string         `json:"desc"`
}

type WeatherInfo struct {

    Yesterday Yester   `json:"yesterday"`
	City string        `json:"city"`
	Forecast []Foreca   `json:"forecast"`
	Ganmao string       `json:"ganmao"`
	Wendu string        `json:"wendu"`
}

type Yester struct{
	Date string  `json:"date"`
	High string   `json:"high"`
	Fx string    `json:"fx"`
	Low string    `json:"low"`
	Fl string      `json:"fl"`
	Type string    `json:"type"`
} 

type Foreca struct{
	Date string  `json:"date"`
	High string   `json:"high"`
	Fl string      `json:"fengli"`
	Low string    `json:"low"`
	Fx string    `json:"fengxiang"`
	Type string    `json:"type"`
} 


func GetWeather(city string){
 

	str:= "http://wthrcdn.etouch.cn/weather_mini?city="+city
	resp, err := http.Get(str)  //使用get方法访问
	if err != nil {
		return
	}

	defer resp.Body.Close()
	input, err1 := ioutil.ReadAll(resp.Body)    //读取流数据
	if err1 != nil {
		return
	}


	var jsonWeather Weather
	err2:=json.Unmarshal(input, &jsonWeather)   //解析json数据
	if err2 != nil {
		return
	}

	

	fmt.Println(city,"天气信息")
	fmt.Printf("status: %d \n",jsonWeather.Status)
	fmt.Printf("desc: %s \n",jsonWeather.Desc)
	
	for i := 0; i < 5; i++ {
	    fmt.Printf("--> date:%s  high:%s  风力:%s  low:%s  风向:%s  type:%s \n", jsonWeather.Data.Forecast[i].Date,jsonWeather.Data.Forecast[i].High,jsonWeather.Data.Forecast[i].Fl,jsonWeather.Data.Forecast[i].Low, jsonWeather.Data.Forecast[i].Fx,jsonWeather.Data.Forecast[i].Type)
	}
    
}

func main(){
	
	GetWeather("深圳")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	GetWeather("广州")
}
