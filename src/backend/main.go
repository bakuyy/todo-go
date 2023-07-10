package main // imports pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http" //imports
	"os"
	// "github.com/rs/cors"
	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
)

type Task struct { // type is used to define a custom type, Task is name, struct is keyword used to def a strucuted data type that can hold multiple fields
	ID        string `json:"id"` // created three fields within a Task - ID, Task, Completed
	Task      string `json:"tasks"`
	Completed string `json:"completed"`
}

// allows you to represent a JSON database where the field can hold multiple tasks
type Database struct { // another custom type, named Database,
	Tasks []Task `json:"tasks"` // Tasks is  the name of field in the Database struct (like ID, task above), contains Tasks, "json: tasks" is part of struct tag, which provides medta data abt the field
} // in the case above, it specifies that the field should be labled as tasks with enc or decoding the struct to JSON. 


func handleTasks(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		getTasks(c)
	case http.MethodPost:
		addTasks(c)
	case http.MethodDelete:
		deleteTask(c)
	default:
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}
}




func getTasks(c *gin.Context) {
	err := loadTasksFromJSON(&db) // load tasks from json database; 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, db.Tasks)
}

func addTasks(c *gin.Context){
	var newTask Task 
	err := c.BindJSON(&newTask) 
	
	// json.NewDecorder(r.Body) creates new json.Decoder
	// Decode(&newTask) decodes json data from r.Body (reuqest body)
	// error handling: decode method returns an error  (occurs if json data doenst match structure of newTask var)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return // displays error if needed
	}
	fmt.Println( newTask)

	errs := loadTasksFromJSON(&db) // laods tasks frmo json database into db variable, then returns loaded tasks

	if errs != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	db.Tasks = append(db.Tasks, newTask) //adds new task to existing list

	err = saveTasksToJSON(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task added successfully"})


}

func deleteTask(c *gin.Context) {
	taskID := c.Query("id")
	// R.url represents the URL of the incoming request
	//Query() retrives query parameters from URL
	// get("id") retrieves value of query parameter, and this value is assigned to "taskID"
	if taskID =="" { // checks to see if the ID is empty
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task ID is required"})		
		return
	}

	err := loadTasksFromJSON(&db)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})


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
	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.Use(CORSMiddleware())

	router.POST("/tasks", addTasks)
	router.DELETE("/tasks", deleteTask)
	
	// Run the server
	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
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
