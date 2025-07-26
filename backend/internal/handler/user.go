package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/islamuzaqpai/notes-app/internal/auth"
	"github.com/islamuzaqpai/notes-app/internal/db"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserHandler struct {
	Conn      *pgx.Conn
	JWTSecret string
}

func (Conn UserHandler) Registration(c *gin.Context) {
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

	err = db.InsertUser(Conn.Conn, user.Username, user.Email, string(hashedPassword))
	log.Println("InsertUser error:", err) // ← вот это добавь
	if err != nil {
		c.JSON(500, gin.H{"error": "Error inserting user"})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}

func (Conn UserHandler) GetUsers(c *gin.Context) {
	users, err := db.GetUsers(Conn.Conn)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error when getting users"})
		return
	}
	c.JSON(200, users)
}

func (Conn UserHandler) Login(c *gin.Context) {
	var user db.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong JSON"})
		return
	}

	dbUser, err := db.GetUserByEmail(Conn.Conn, user.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email error or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := auth.GenerateJWT(dbUser.ID, Conn.JWTSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
