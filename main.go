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
	for n:=1 ; n>0; n--{
		list := worklist

		for _,link := range <-worklist{
			if !seen[link]{
				seen[link] =true
				n++
				go func (link string, basedomain string)  {
				foundLinks := crawl(link,basedomain)
				if foundLinks != nil{
					worklist<- foundLinks
				}

				}()
			}
		}
	}


}