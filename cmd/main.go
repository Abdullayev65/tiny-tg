package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"time"
	"tiny-tg/internal/config"
	"tiny-tg/internal/handler"
	"tiny-tg/internal/pkg/jwt_manager"
	"tiny-tg/internal/pkg/postgres"
	"tiny-tg/internal/repository"
	"tiny-tg/internal/service"
	"tiny-tg/internal/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO getting updates by UPDATE struct
// create group
// join, leve group
// create, update, delete message
// get message_seen
// DONE

// Next TODO
// last_active_at
// updates_from: time (action in ws)
// attachments in front section

func main() {

	db, err := postgres.ConnectDB(config.POSTGRES_URI)
	if err != nil {
		panic(err)
	}

	jwtManager := jwt_manager.New(config.JWT_SIGNING_KEY, config.JWT_EXPIRY_DURATION)
	repos := repository.New(db)
	services := service.New(repos, jwtManager)
	wsHub := ws.NewHub(services)
	handlers := handler.New(services, jwtManager, wsHub)
	wsHub.Start()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.Static("/uploads", config.UPLOADS_DIR)

	{
		r.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message":    "Server is running!!!",
				"random_num": rand.IntN(1_000),
			})
		})

		// Auth
		{
			r.POST("/register", handlers.Register)
			r.POST("/login", handlers.Login)
		}

		api := r.Group("/api", handlers.UserIdentity)

		api.GET("/ws", handlers.WS)

		// Chat
		chat := api.Group("/chat")
		{
			chat.GET("/:chat_id", handlers.GetChat)
			chat.GET("/personal/:user_id", handlers.GetPersonalChat)
			chat.GET("/search", handlers.SearchChat)
		}

	}

	log.Fatalln(r.Run(":" + config.PORT))

}
