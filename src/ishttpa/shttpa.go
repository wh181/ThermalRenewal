package ishttpa

import (
	"fmt"
	"io/ioutil"
	"isreadinia"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)
var http_route =make(map[string]string)
func Start()  {
	ishttp := http.NewServeMux();
	ishttp.Handle("/",&isHeadera{})
	isreadinia.ReadIni(&http_route)

	//for _, i2 := range http_route {
	//	iswebsocketa.Wstart(i2,string(8080+1))
	//}
	//配置http参数
	isServer := &http.Server{
		Addr:         ":8086",
		WriteTimeout: time.Second * 3,
		Handler:      ishttp,
	}
	log.Fatal(isServer.ListenAndServe())
}

type isHeadera struct {}

func (s *isHeadera) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	local_dir := http_route[strings.Split(r.Host,":")[0]]
	request_dir := local_dir + strings.Replace(r.URL.String(),"/","\\",-1)
	en_result := strings.Split(request_dir,"\\")
	if en_result[len(en_result)-1:][0] == "favicon.ico" {
		return
	}
	fmt.Println(request_dir)
	fd,er := os.OpenFile(request_dir,os.O_RDONLY,0600)
	if er != nil {
		fmt.Println(er)
	}
	retu,_ := ioutil.ReadAll(fd)
	w.Write([]byte(string(retu) + web_str))
}