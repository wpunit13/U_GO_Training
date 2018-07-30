package main

import (
	"encoding/json"
	"fmt"
)

//Address is exported here*
type Address struct {
	Roomnumber string `json:"room_number"`
	Streetname string `json:"streetname"`
	Pincode    int    `json:"Pincode"`
	//This type is used for storing address
}

func main() {
	var officeaddress Address
	Homeaddr := Address{"B802", "Kesnand Road", 412207}
	fmt.Println(Homeaddr)
	bs, _ := json.Marshal(Homeaddr)
	fmt.Println(string(bs))
	bs1 := []byte(`{"room_number":"tower 5","streetname":"Magarpatta","Pincode":412207}`)
	json.Unmarshal(bs1, &officeaddress)
	fmt.Println(officeaddress.Streetname)

}
