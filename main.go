package main

//import depedencies
import (
	"io/ioutil"//package to read request body
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
	body, _ := ioutil.ReadAll(context.Request.Body)
    println(string(body))


	//equal with JSON.parse(context.Request.Body)
	if err := context.BindJSON(&newTodo); err != nil{
		return 
	}

	//after parse the request body, append it to todos array
	todos = append(todos, newTodo)
	

	//return the new todo with json format
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(context *gin.Context){

	//get request param 
	id := context.Param("id")



	//loop over todos
	for _, item := range todos {
		if item.ID == id {
			context.IndentedJSON(http.StatusOK, item)
			return
		}
	}

	//return status 404 if id is not found
	context.Status(http.StatusNotFound)
}


//main app
func main(){

	//make new gin server instance
	router := gin.Default()

	//routes GET
	router.GET("/todos", getTodos)

	//route get 1 todo
	router.GET("/todos/:id", getTodoById)

	//routes POST
	router.POST("/todos", addTodo)


	//start server
	router.Run("localhost:8080")

}