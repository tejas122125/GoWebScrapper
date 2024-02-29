package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)



var userAgents = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
"Mozilla/5.0 (Macintosh; Intel Mac OS X 13_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15"}



func getRandomAgent(){
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}
	func getRequest(targetUrl string)(*http.Response,error){

client:=  &http.Client{}

req,err := http.NewRequest("GET",targetUrl,nil)
if err != nil {
	return nil,err
}else{
	return res,nil
}

req.Header.Set("user-agent",getRandomAgent())
res,err := client.Do(req)
if err != nil {
	return nil,err
}else{
	return res,nil
}


}

func resolveRelativeLinks(link string, baseurl string)(bool,string){

}


func discoverLinks(resp *http.Response,baseurl string)[]string{
	if (resp !=nil){
	doc , _ := goquery.NewDocumentFromResponse(resp)
foundurls := []string{}
	if doc != nil{

	}
	return foundurls
	}else{
return []string {}
	}

}


func crawl(targetUrl string, baseUrl string) {
	fmt.Println(targetUrl);
	resp,_ := getRequest(targetUrl)
	links := discoverLinks(resp,baseUrl)
	foundUrls := []string{}
	for _,link := range links{
		ok,correctLinks := resolveRelativeLinks(link,baseUrl)
		if ok{
			if correctLinks != ""{
				foundUrls = append(foundUrls,correctLinks)
			}
		}
	}
// here we [parsse] the html content of our choice of links using go query



}
func main() {
	worklist := make(chan []string)
	basedomain := "website address"
	go func() {
		worklist <- []string{basedomain}
	}()
	seen := make(map[string]bool)
	for n := 1; n > 0; n-- {
		list := worklist

		for _, link := range <-worklist {
			if !seen[link] {
				seen[link] = true
				n++
				BadStmt

			}
		}
	}

}