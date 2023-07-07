package main // imports pkg

import ("net/http"  //imports
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os")

type Task struct { // type is used to define a custom type, Task is name, struct is keyword used to def a strucuted data type that can hold multiple fields
	ID        string `json:"id"` // created three fields within a Task - ID, Task, Completed
	Task      string `json:"task"`
	Completed string `json:"completed"`
}

// allows you to represent a JSON database where the field can hold multiple tasks
type Database struct { // another custom type, named Database,
	Tasks []Task `json:"tasks"` // Tasks is  the name of field in the Database struct (like ID, task above), contains Tasks, "json: tasks" is part of struct tag, which provides medta data abt the field
} // in the case above, it specifies that the field should be labled as tasks with enc or decoding the struct to JSON. 


func loadTasksFromJSON(){
	os.Open("database.json") //opens database and reads it
}

// when declaring a var, u can specify its name followed by the type of variable
var db Database // creating a variable with the features/characteristics as our Database struct

func main(){
	// http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("hello world!"))
	// })
	// http.ListenAndServe(":8080",nil)

	loadTasksFromJSON()
}
