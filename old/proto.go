package main

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"os"
	_ "test/proto/gen"
)

//ProtoJSON

func main() {
	bstr, _ := os.ReadFile("proto/teststring.json")
	bint, _ := os.ReadFile("proto/testint.json")
	var m gen.Test
	err := protojson.Unmarshal(bstr, &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	newbstr, err := protojson.Marshal(&m)
	if err != nil {
		log.Fatal(err)
	}
	_ = os.WriteFile("proto/teststring_out.json", newbstr, 0644)
	fmt.Println(bint)
	//err = protojson.Unmarshal(bint, &m)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//newbint, err := protojson.Marshal(&m)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_ = os.WriteFile("proto/testint_out.json", newbint, 0644)
	//fmt.Println(m)
}
