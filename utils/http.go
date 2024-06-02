package utils

import (
  "os"
  "fmt"
  "time"
  "net/http"
  "net/url"
//   "log"
)

func GetTestVideoUrl(w http.ResponseWriter) {
	TimeLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		TimeLocation = time.FixedZone("CST", 8*60*60)
	}
	str_time := time.Now().In(TimeLocation).Format("2006-01-02 15:04:05")
  fmt.Fprintln(w, "#EXTM3U")
  fmt.Fprintln(w, "#EXTINF:-1 tvg-name=\""+str_time+"\" tvg-logo=\"http://epg.51zmt.top:8000/tb1/CCTV/CCTV1.png\" group-title=\"列表更新时间\","+str_time)
  fmt.Fprintln(w, "http://tipsmessage.989089.xyz:8880/hls/video/tips/2.mp4")
}

func GetLivePrefix(r *http.Request) string {
	// 尝试从环境变量读取url
	envUrl := os.Getenv("LIVE_URL")
	// log.Println("env url:", envUrl)
	if envUrl == "" {
		// 默认url
		envUrl = "https://www.goodiptv.club"
	  }
    firstUrl := DefaultQuery(r, "url", envUrl)
    realUrl, _ := url.QueryUnescape(firstUrl)
    return realUrl
}

func DefaultQuery(r *http.Request, name string, defaultValue string) string {
  param := r.URL.Query().Get(name)
  if param == "" {
    return defaultValue
  }
  return param
}
  
func Duanyan(adurl string, realurl any) string {
  var liveurl string
  if str, ok := realurl.(string); ok {
    liveurl = str
  } else {
	liveurl = adurl
  }
  // log.Println("Redirect url:", liveurl)
 return liveurl
}
