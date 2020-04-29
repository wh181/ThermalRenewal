package ishttpa

import (
	"log"
	"net/http"
	"time"
)

func Start()  {
	ishttp := http.NewServeMux();
	ishttp.Handle("/",&isHeadera{})
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
	w.Write([]byte("成功"))
}