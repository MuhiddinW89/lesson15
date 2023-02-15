package main

import (
	"fmt"
	"lesson15/models"
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


	// fmt.Println(store.User.GetList())
	// fmt.Println(store.User.FindUsr("be"))
	// fmt.Println(store.User.GetById("c57aa672-902f-44c8-af9d-dfa02f62541a"))
	// fmt.Println(store.User.DleteUsr("32a803b5-3f1c-495e-8ce7-4649e4cbe3b1"))

		user:= models.Usr{
		Id: "",
		FirstName: "Sherqozi",
		LastName: "Goziyev",
		Gender: "Female",
		Card_number: "999",
		Birthday: "5555 - 11 -22",
		Profession: "Diktr",
		}

		fmt.Println(store.User.CreateUsr(user))
	
}