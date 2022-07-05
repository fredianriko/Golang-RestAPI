package main

//import depedencies
import (
	// "encoding/json"
	"encoding/json"
	"errors"

	"net/http"
	"github.com/gin-gonic/gin"
)

//simulate model for database, or schema in database
type todo struct {
	ID 			string 	`json:"id"`
	Item 		string 	`json:"item"`
	Completed 	bool 	`json:"completed"`
}


//data dummy, simulate data in database
var todos = []todo{
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record video", Completed: false},
};


//controller GET
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

//controller POST
func addTodo(context *gin.Context){
	var newTodo todo

	//read request body
	// title := json.Unmarshal(newTodo)


	//equal with JSON.parse(context.Request.Body)
	if err := context.BindJSON(&newTodo); err != nil{
		return 
	}

	//after parse the request body, append it to todos array
	todos = append(todos, newTodo)
	

	//return the new todo with json format
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error){
	for i, t := range todos{
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func updateTodo(context *gin.Context){
	
	//decode the request body json and store it to decoder variabel
 	decoder := json.NewDecoder(context.Request.Body)

	//create new variabel that has the type of todo
	var t todo

	//if error value is nil, then decode the request body json and store it to t variabel
	err := decoder.Decode(&t)

	//if error value exist, stop all process and return error
	if err != nil {
        panic(err)
    }

	//get one todo from todos array where id = t.ID
	todo, err := getTodoById(t.ID)

	//if error value exist, stop all process and return error
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	//assign new value to todo with value from request body
	todo.Item = t.Item
	todo.Completed = t.Completed

	//return the updated todo with json format and status response
	context.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(context *gin.Context){
	id := context.Param("id")

	//get one todo from todos array where id = id
	_, err := getTodoById(id)


	//if no data found with id, return error
	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	//iterate ofer todos slice, delete todo from todos array where todo.ID = id
	for i, t := range todos{
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	
	//return the deleted todo with json format and status response
	context.IndentedJSON(http.StatusOK, todos)
	
}


//main app
func main(){

	//make new gin server instance
	router := gin.Default()

	//getAllTodos
	router.GET("/todos", getTodos)

	//getOneTodo
	router.GET("/todos/:id", getTodo)

	//postNewTodo
	router.POST("/todos", addTodo)

	//updateTodo	
	router.PATCH("/todos", updateTodo)

	//deleteTodo
	router.DELETE("/todos/:id", deleteTodo)


	//start server
	router.Run("localhost:8080")

}