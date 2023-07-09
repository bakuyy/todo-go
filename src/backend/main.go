package main // imports pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http" //imports
	"os"
)

type Task struct { // type is used to define a custom type, Task is name, struct is keyword used to def a strucuted data type that can hold multiple fields
	ID        string `json:"id"` // created three fields within a Task - ID, Task, Completed
	Task      string `json:"task"`
	Completed string `json:"completed"`
}

// allows you to represent a JSON database where the field can hold multiple tasks
type Database struct { // another custom type, named Database,
	Tasks []Task `json:"tasks"` // Tasks is  the name of field in the Database struct (like ID, task above), contains Tasks, "json: tasks" is part of struct tag, which provides medta data abt the field
} // in the case above, it specifies that the field should be labled as tasks with enc or decoding the struct to JSON. 


func handleTasks(w http.ResponseWriter, r http.Request){ // responsible for handling incoming HTTP tasks
	// w represents responsewriter for writing the response, r represents http.reqest which woul,d contian incoming req details
	switch r.Method {
	case http.MethodGet:
		getTasks(w)
	case http.MethodPost:
		addTasks(w, r)
	case http.MethodDelete:
		deleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
	}

}



func getTasks(w http.ResponseWriter) {
	err := loadTasksFromJSON(&db) // load tasks from json database; 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //if error occurs, handles error by setting resopnse status code to the 500 internal server error
		fmt.Fprintf(w, "Error loading tasks: %s", err.Error()) // lets client know abt error, uses FPRINTF
		return
	}
	w.Header().Set("Content-Type","application/json") // set response header to specify that response will be in json format, header = "content-type", value = "applicaiton/json" these r standard
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Tasks) // encodes db.Tasks into json format and writes it to the repsonse body, this ensures tasks are sent back to client as json response
}

func addTasks(w http.ResponseWriter, r http.Request) {
	var newTask Task 
	err := json.NewDecoder(r.Body).Decode(&newTask) // regcodes body json data into newTask variable, then performs decoding
	// json.NewDecorder(r.Body) creates new json.Decoder
	// Decode(&newTask) decodes json data from r.Body (reuqest body)
	// error handling: decode method returns an error  (occurs if json data doenst match structure of newTask var)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Error found the request %s", err.Error())
		return // displays error if needed
	}

	errs := loadTasksFromJSON(&db) // laods tasks frmo json database into db variable, then returns loaded tasks

	if errs != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w,"Error loading tasks %s", errs.Error())
	}
	
	db.Tasks = append(db.Tasks,newTask) //adds new task to existing list

	err = saveTasksToJSON(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error saving tasks %s", err.Error())
		return
	}
	w. WriteHeader(http.StatusCreated)
	fmt.Fprint(w,"Task added!")

}

func deleteTask(w http.ResponseWriter, r http.Request) {
	taskID := r.URL.Query().Get("id") 
	// R.url represents the URL of the incoming request
	//Query() retrives query parameters from URL
	// get("id") retrieves value of query parameter, and this value is assigned to "taskID"
	if taskID =="" { // checks to see if the ID is empty
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Task ID required")
		return
	}

	err := loadTasksFromJSON(&db)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w,"Unable to load in request")
	}

	updatedTasks :=[]Task{}

	for _, task := range db.Tasks{
		if task.ID != taskID {
			updatedTasks = append(updatedTasks, task)
		}
	}

	db.Tasks = updatedTasks

	err = saveTasksToJSON(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to save task error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"Deleted!")


}


// when declaring a var, u can specify its name followed by the type of variable
var db Database // creating a variable with the features/characteristics as our Database struct

func loadTasksFromJSON(db *Database) error{
	file, err := os.Open("database.json") //opens database and reads it // db is assigned to file, any error that occurs is assigned to err
	if err != nil {
		fmt.Println("Error opening db file.")
		return fmt.Errorf("error opening db file: %s", err.Error())// returns the function to stop it from further running fmt.Errorf creates an error value to ,, marshalling data sepcifices the error message, %s placeholder will be replaced by value of err.rError()

	}
	defer file.Close() // defers closing of file until surrounding function from loadtasks completes, ess. ensures file is closed after reading

	data, err := ioutil.ReadAll(file) // reads all the content of opened file and assigns to data variable
	if err != nil {
		fmt.Println("Error reading data from db file")
		return fmt.Errorf("error reading data from db file: %s", err.Error())
	}

	// difference btwn os and ioutil: os opens file and provides a way to interact w/ it, and ioutil reads content into the memory to be used

	err = json.Unmarshal(data, &db) // attempts to unmarshal (decade) json data stored in the data variable, into struct var called db
	// unmarshalling: u receive data and want to extract/use this data; process of converting serialized data back into og form; json.unmarshal used to deserialize json data into go data structures. takes json string as input and converts it back to go ds (struct, slice, map etc)
	// conversting data into a format that can be transmitted or stored (such as json or xml) json.marshal turns go ds into json representation,, converts it into a json string that can be sent over network
	// if there is error, error val is assigned to err var

	if err != nil {
		fmt.Println("Error unmarshalling data")
		return fmt.Errorf("error unmarshalling data: %s", err.Error())
	}
	return nil
}

func main(){
    http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
        handleTasks(w, *r)
    })
    fmt.Println("Server started on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func saveTasksToJSON( db Database) error {
	data, err := json.MarshalIndent(db,""," ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("database.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}
