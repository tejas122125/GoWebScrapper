package main

func crawl (){
	
}
func main(){
	worklist := make(chan []string)
	basedomain := "website address"
	go func ()  {
		worklist<- []string{basedomain}
	}()
	seen :=  make (map[string]bool)
	list := worklist


}