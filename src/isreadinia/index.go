package isreadinia

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func ReadIni(route *map[string]string)  {
	cfg,er := ini.Load("C:\\Users\\wanghang\\project\\index\\betel\\config\\config.ini");
	if er != nil {
		fmt.Println(er)
	}
	for _,v := range cfg.Section("host").Keys() {
		(*route)[v.Name()] = v.String()
	}
}

