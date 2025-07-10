package main

import (
	"github.com/gin-gonic/gin"
	"github.com/islamuzaqpai/notes-app/internal/db"
	"github.com/jackc/pgx/v5"
	"log"
	"strconv"
)

var conn *pgx.Conn

func main() {
	var err error
	conn, err = db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}

	router := gin.Default()

	router.GET("/notes", getNotesHandler)
	router.POST("/notes", createNoteHandler)
	router.DELETE("/notes/:id", deleteNoteHandler)
	router.PUT("/notes/:id", updateNoteHandler)

	log.Println("Server is started on http://localhost:8080")
	router.Run(":8080")
}

func getNotesHandler(c *gin.Context) {
	notes, err := db.GetNotes(conn)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error when  getting notes"})
		return
	}
	c.JSON(200, notes)
}

func createNoteHandler(c *gin.Context) {
	var note db.Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(400, gin.H{"error": "Wrong JSON"})
		return
	}

	err := db.InsertNote(conn, note.Title, note.Content)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error when creating a note"})
		return
	}
	c.JSON(201, gin.H{"status": "ok"})
}

func updateNoteHandler(c *gin.Context) {
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

	err = db.UpdateNote(conn, id, note.Title, note.Content)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error Updating Note"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func deleteNoteHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	err = db.DeleteNote(conn, id)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error when deleting the note"})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
