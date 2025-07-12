package main

import (
	"github.com/gin-gonic/gin"
	"github.com/islamuzaqpai/notes-app/internal/db"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
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

	router.POST("/users", registrationHandler)
	router.GET("/users", getUsersHandler)
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

func registrationHandler(c *gin.Context) {
	var user db.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong JSON"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	err = db.InsertUser(conn, user.Username, user.Email, string(hashedPassword))
	if err != nil {
		c.JSON(500, gin.H{"error": "Error inserting user"})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

func getUsersHandler(c *gin.Context) {
	users, err := db.GetUsers(conn)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error when getting users"})
		return
	}
	c.JSON(200, users)
}

type LoginRequest struct {
	login    string
	password string
}

func loginHandler(c *gin.Context) {
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong JSON"})
		return
	}

	dbUser, err := db.GetUserByEmail(conn, user.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email error or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

}
