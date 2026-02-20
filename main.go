package main

import (
	"net/http"
	"rest-api/db"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)
func main(){
   db.InitDB()
   server := gin.Default()
   server.GET("/events",getEvent)
   server.POST("/events",createEvent)
   server.Run(":8080")//localhost:8080

}
func getEvent(context *gin.Context){
   events:=models.GetAllEvents()
   context.JSON(http.StatusOK,events)
}
func createEvent(context *gin.Context){
   var event models.Event
   err:=context.ShouldBindBodyWithJSON(&event)
   if err!=nil{
      context.JSON(http.StatusBadRequest,gin.H{
         "message":"could not parse request",
      })
      
   }
   
   event.ID=1
   event.UserID=1
   event.Save()
   context.JSON(http.StatusCreated,gin.H{"message":"event created ","event":event})
}