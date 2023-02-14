package main

import (
	"fmt"
	"lesson15/storage"
)




func main(){

	store,err:=storage.NewJson()
	if err != nil {
		fmt.Println(err)
	}

	// users := models.Usr{}
	// user:= models.Usr{
	// Id: "c57aa672-902f-44c8-af9d-dfa02f62541a",
    // FirstName: "Mamaraim",
    // LastName: "Toshmatov",
    // Gender: "Female",
    // Card_number: "000000000000000",
    // Birthday: "0000 - 00 -00",
    // Profession: "Bekorchi",
	// }
	// fmt.Println(store.User.UpdateUsr(user))


	fmt.Println(store.User.GetList())
	// fmt.Println(store.User.FindUsr("be"))

}