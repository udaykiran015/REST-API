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
   events,err:=models.GetAllEvents()
   if err !=nil{
      context.JSON(http.StatusInternalServerError,gin.H{"message":"could not fetch event .Try again later"})
      return
   }
   context.JSON(http.StatusOK,events)
}
func createEvent(context *gin.Context){
   var event models.Event
   err:=context.ShouldBindBodyWithJSON(&event)
   if err!=nil{
      context.JSON(http.StatusBadRequest,gin.H{
         "message":"could not parse request",
      })
      return
      
   }
   
   event.ID=1
   event.UserID=1
   err=event.Save()
   if err!=nil{
      context.JSON(http.StatusInternalServerError,gin.H{"message":"could not save event"})
      return
   }
   context.JSON(http.StatusCreated,gin.H{"message":"event created ","event":event})
}