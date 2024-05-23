package service

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events, try again later.",
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event fetch succesfully",
		"event":   event,
	})
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	uId := context.GetInt64("userId")
	event.UserId = uId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event, try again.",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created succesfully",
		"event":   event,
	})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse data.",
		})
		return
	}

	uId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch data.",
		})
		return
	}

	if event.UserId != uId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update this event.",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data received.",
		})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated succesfully."})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not parse data.",
		})
		return
	}
	uId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch data.",
		})
		return
	}

	if event.UserId != uId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not delete this event.",
		})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete the event.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted succesfully",
	})

}
