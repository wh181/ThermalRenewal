package iswebsocketa
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"time"
)

var md5_str = "";
var cwdPath string
func Wstart(rootpath string,port string)  {
	start(rootpath)
	http.Handle("/conn",websocket.Handler(upper));
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start(rootpath string)  {
	cwdPath = rootpath
	need_dir_time(cwdPath);
}
func need_dir_time(cwd string){
	fd,err := os.OpenFile(cwd,os.O_RDONLY,os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
	many_dir,error := fd.Readdir(-1)
	if error != nil {
		fmt.Println(error)
	}
	for _,y := range many_dir{
		if !y.IsDir() {
			md5_str = md5_myself(y.ModTime().String() + md5_str)
		}else{
			if y.Name() != "runtime" && y.Name() != "thinkphp" && y.Name() != "vendor"{
				need_dir_time(cwd+"/"+y.Name())
			}
		}
	}
}

func md5_myself(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}


func upper(ws *websocket.Conn) {
	var err error
	fmt.Println("has connect")
	if err = websocket.Message.Send(ws, "over"); err != nil {
		fmt.Println(err)
	}
	current_md5_str := md5_str
	for {
		md5_str = ""
		need_dir_time(cwdPath)
		if current_md5_str != md5_str {
			current_md5_str = md5_str
			if err = websocket.Message.Send(ws, "reload"); err != nil {
				ws.Close()
				return
			}
		}
		time.Sleep(time.Second)
	}
}




