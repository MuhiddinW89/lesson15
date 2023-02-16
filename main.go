package main

import (
	"fmt"
	"lesson15/controller"
	"lesson15/storage/jsondb"
)




func main(){

	store,err:=jsondb.NewJson()
	if err != nil {
		fmt.Println(err)
	}

	c := controller.NewController(store)

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

		// user:= models.Usr{
		// Id: "",
		// FirstName: "Abduvosid",
		// LastName: "Goziyev",
		// Gender: "Female",
		// Card_number: "999",
		// Birthday: "5555 - 11 -22",
		// Profession: "Diktr",
		// }
		

		// c.CreateUsr(user)

		// toPost := models.CreatePost{
		// 	Title: "udevs",
		// 	Description: "asdasd",
		// 	UserId: "c57aa672-902f-44c8-af9d-dfa02f62541a",
		// }

		// id, err := c.CreatePost(toPost)

		// if err != nil {
		// 	fmt.Println(err)
		// }else{
		// 	fmt.Println(id)
		// }
		// fmt.Println(c.GetByIdPost("07b4ede4-4b73-4135-a8d3-f417b0cb90e4"))

		fmt.Println(c.GetUserPosts("c57aa672-902f-44c8-af9d-dfa02f62541a"))
	
}