package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/islamuzaqpai/notes-app/internal/config"
	"github.com/islamuzaqpai/notes-app/internal/db"
	"github.com/islamuzaqpai/notes-app/internal/handler"
	"github.com/islamuzaqpai/notes-app/internal/midlleware"
	"github.com/jackc/pgx/v5"
	"log"
)

var conn *pgx.Conn

func main() {
	var err error
	cfg := config.Load()
	conn = db.Connect(cfg)

	defer conn.Close(context.Background())

	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	noteHandler := handler.NoteHandler{conn}
	userhandler := handler.UserHandler{conn, cfg.JWTSecret}

	router.DELETE("/notes/:id", noteHandler.DeleteNote)
	router.PUT("/notes/:id", noteHandler.UpdateNote)

	router.POST("/register", userhandler.Registration)
	router.GET("/users", userhandler.GetUsers)
	router.POST("/login", userhandler.Login)

	authorized := router.Group("/")
	authorized.Use(midlleware.AuthMiddleware(cfg.JWTSecret))

	authorized.GET("/notes", noteHandler.GetNotes)
	authorized.POST("/notes", noteHandler.CreateNote)

	log.Println("Server is started on http://localhost:8080")
	router.Run(":8080")
}
