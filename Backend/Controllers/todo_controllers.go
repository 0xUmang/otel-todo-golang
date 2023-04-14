package Controllers

import (
	"Backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("todo-server")
var Todos = make(map[string]*Models.Todo)

func GetAllTodos(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(),"GetAllTodos",trace.WithAttributes(attribute.String("Method","Inside GetAllTodos")))
	defer span.End()
	//span := trace.SpanFromContext(c.Request.Context())

	var result []*Models.Todo
	for _,todo := range Todos {
		result = append(result,todo)
	}
	c.JSON(http.StatusOK, result)

	span.SetAttributes(attribute.Int("todos.count",len(result)))

}

func CreateATodo(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(),"CreateATodo",trace.WithAttributes(attribute.String("Method","Inside CreateATodo")))
	defer span.End()

	var input Models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	todo := Models.NewTodo(input.Title,input.Description,input.Status)
	Todos[todo.ID]=todo
	c.JSON(http.StatusOK,todo)

	span.SetAttributes(attribute.String("todo.id",todo.ID),attribute.String("todo.title",todo.Title),attribute.String("todo.Description",todo.Description),attribute.String("todo.Status",todo.Status))
}

func GetATodobyID(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(),"GetATodobyID",trace.WithAttributes(attribute.String("Method","Inside GetATodobyID")))
	defer span.End()

	id := c.Param("id")
	todo,exists := Todos[id]
	if !exists {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK,todo)

	span.SetAttributes(attribute.String("todo.id",todo.ID),attribute.String("todo.title",todo.Title),attribute.String("todo.Description",todo.Description),attribute.String("todo.Status",todo.Status))


}

func UpdateATodobyID(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(),"UpdateATodobyID",trace.WithAttributes(attribute.String("Method","Inside UpdateATodobyID")))
	defer span.End()

	id := c.Param("id")
	todo, exists :=Todos[id]
	if !exists {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	var input Models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	todo.Title=input.Title
	todo.Description=input.Description
	todo.Status=input.Status

	Todos[id]=todo
	c.JSON(http.StatusOK,todo)

	span.SetAttributes(attribute.String("todo.id",todo.ID),attribute.String("todo.title",todo.Title),attribute.String("todo.Description",todo.Description),attribute.String("todo.Status",todo.Status))

}

func DeleteATodobyID(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(),"DeleteATodobyID",trace.WithAttributes(attribute.String("Method","Inside DeleteATodobyID")))
	defer span.End()

	id := c.Param("id")
	_, exists :=Todos[id]
	if !exists {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	delete(Todos,id)
	c.Status(http.StatusOK)

	span.SetAttributes(attribute.String("todo.id",id))

}
