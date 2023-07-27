package main
import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)


type todo struct{
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}



var todos = []todo{
	{ID: "1", Item: "New blog", Completed: true},
	{ID: "2", Item: "New blog item", Completed: false},
	{ID: "3", Item: "New blog item list", Completed: true},

}
func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context){
	var newtodo todo

	if err := context.BindJSON(&newtodo); err != nil {
		return 
	}

	todos = append(todos, newtodo)
	context.IndentedJSON(http.StatusCreated, todos)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todo[i], nil
		}

	}	

	return nil, errors.New("todos not found")
}


func main()  {
	router:= gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.Run("localhost:9090")
}