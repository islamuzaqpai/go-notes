package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/islamuzaqpai/notes-app/internal/db"
	"github.com/jackc/pgx/v5"
	"strconv"
)

type NoteHandler struct {
	Conn *pgx.Conn
}

func (Conn *NoteHandler) GetNotes(c *gin.Context) {

	userIdVal, exists := c.Get("userID")
	if !exists {
		c.JSON(500, gin.H{"error": "User ID not found"})
		return
	}

	userId := userIdVal.(int)

	notes, err := db.GetNotes(Conn.Conn, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error when  getting notes"})
		return
	}
	c.JSON(200, notes)
}

func (Conn *NoteHandler) CreateNote(c *gin.Context) {
	userIdVal, exists := c.Get("userID")
	if !exists {
		c.JSON(500, gin.H{"error": "User ID not found"})
		return
	}

	userId := userIdVal.(int)

	var note db.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": "Wrong JSON"})
		return
	}
	note.UserId = userId

	err := db.InsertNote(Conn.Conn, note.Title, note.Content, note.UserId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error when creating a note"})
		return
	}
	c.JSON(201, gin.H{"status": "ok"})
}

func (Conn *NoteHandler) UpdateNote(c *gin.Context) {
	var note db.Note
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	err = c.BindJSON(&note)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	err = db.UpdateNote(Conn.Conn, id, note.Title, note.Content)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error Updating Note"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func (Conn NoteHandler) DeleteNote(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	err = db.DeleteNote(Conn.Conn, id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error when deleting the note"})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
