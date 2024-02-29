package main

import "fmt"
	list := worklist
	func getRequest(targetUrl string)(error,[]string){



}

func resolveRelativeLinks(link string, baseurl string){

}


func discoverLinks(resp string,baseurl string)[]string{

}


func crawl(targetUrl string, baseUrl string) {
	fmt.Println(targetUrl);
	resp,_ := getRequest(targetUrl)
	links := discoverLinks(resp,baseUrl)
	foundUrls := []string{}
	for _,link := range links{
		resolveRelativeLinks(link,baseUrl)
	}



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