package main

import ( 
	"fmt"
	"encoding/json"
	"flag"
	"os"
	"github.com/ghanatava/adventureTime"
)

func main(){
	filename := flag.String("file","gopher.json","the-JSON-FILE with adventure stroy")
	flag.Parse()
	fmt.Printf("Using the story in %s",filename)

	f,err := os.Open(*filename)
	if err!=nil{
		panic(err)
	}

	d := json.NewDecoder(f)
	var story adventureTime.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n",story)

}