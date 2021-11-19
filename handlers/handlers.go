package handlers

import (
	"net/http"
	"pollserver/db"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PollCreationRequest struct {
	Question string       `json:"question" binding:"required"`
	Options  []db.Options `json:"options" binding:"required"`
}

func GetPolls(c *gin.Context) {
	polls, err := db.GetAllPolls()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, polls)
}

func CreatePoll(c *gin.Context) {
	var creationUrl PollCreationRequest
	if err := c.ShouldBindJSON(&creationUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	poll := &db.Poll{
		ID:          primitive.NewObjectID(),
		PollID:      xid.New().String(),
		Total_votes: 0,
		Question:    creationUrl.Question,
		Options:     creationUrl.Options,
	}

	db.CreatePoll(poll)

	c.JSON(http.StatusOK, gin.H{
		"message": "poll created",
	})
}

func Find(c *gin.Context) {
	id := c.Param("id")
	results := db.Find(id)

	c.JSON(http.StatusOK, results)

}
