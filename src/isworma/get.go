package isworma

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func IsGet(u string) string {
	client := &http.Client{}
	req,_ := http.NewRequest("GET",u,nil)
	respont_content,_ := client.Do(req)
	content,_ := ioutil.ReadAll(respont_content.Body)
	return string(content)
}

func IsPost(u string,parase map[string]string) string {
	quert := url.Values{}
	for k,v := range parase{
		quert.Add(k,v)
	}
	client := &http.Client{}
	req,_ := http.NewRequest("POST",u,strings.NewReader(quert.Encode()))
	respont_content,_ := client.Do(req)
	content,_ := ioutil.ReadAll(respont_content.Body)
	return string(content)
}