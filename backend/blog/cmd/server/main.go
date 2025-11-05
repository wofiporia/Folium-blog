package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wofiporia/folium-backend-new/internal/config"
	dbPkg "github.com/wofiporia/folium-backend-new/internal/db"
	handlerPkg "github.com/wofiporia/folium-backend-new/internal/handler"
	"github.com/wofiporia/folium-backend-new/internal/middleware"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// init DB (safe even if DSN not provided; you can skip if only health used)
	dbPkg.Init()

	r := gin.Default()

	// Basic CORS middleware (dev-friendly)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Root route for quick browser check
	r.GET("/", func(c *gin.Context) {
		c.String(200, "backendNew OK")
	})
	// health check
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"ok": true})
	})

	// Register API routes
	api := r.Group("/api")
	handlerPkg.RegisterBlogRoutes(api)  // Register blog routes
	handlerPkg.RegisterAdminRoutes(api) // Register admin routes

	// 受保护的管理接口
	protected := api.Group("/admin")
	protected.Use(middleware.JWTMiddleware()) // 添加JWT中间件
	{
		protected.POST("/blog", handlerPkg.CreateBlog)
		protected.PUT("/blog/:id", handlerPkg.UpdateBlog)
		protected.DELETE("/blog/:id", handlerPkg.DeleteBlog)
	}

	// 添加可选JWT中间件的路由组
	optionalAuth := api.Group("/")
	optionalAuth.Use(middleware.OptionalJWTMiddleware())
	{
		// 这里可以添加一些可选认证的路由
	}

	log.Println("backendNew listening on :" + cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
