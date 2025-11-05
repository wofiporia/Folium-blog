package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wofiporia/folium-backend-new/internal/model"
	"github.com/wofiporia/folium-backend-new/internal/repo"
	"github.com/wofiporia/folium-backend-new/internal/util"
)

func GetAllBlogs(c *gin.Context) {
	blogs, err := repo.FindAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": blogs})
}

func GetBlogByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "invalid id"})
		return
	}
	blog, err := repo.FindBlogByID(uint(id64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	if blog == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": blog})
}

// CreateBlogRequest 创建博客请求结构
type CreateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdateBlogRequest 更新博客请求结构
type UpdateBlogRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateBlog handles POST /api/blog
func CreateBlog(c *gin.Context) {
	var req CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "请求参数错误"})
		return
	}

	blog := &model.Blog{
		Title:      req.Title,
		Content:    req.Content,
		UploadDate: time.Now(),
		UpdateDate: time.Now(),
	}

	if err := repo.CreateBlog(blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "创建博客失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": blog})
}

// UpdateBlog handles PUT /api/blog/:id
func UpdateBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid blog ID"})
		return
	}

	var req UpdateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "请求参数错误"})
		return
	}

	// 查找现有博客
	blog, err := repo.FindBlogByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	if blog == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Blog not found"})
		return
	}

	// 更新博客
	blog.Title = req.Title
	blog.Content = req.Content
	blog.UpdateDate = time.Now()

	if err := repo.UpdateBlog(blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "更新博客失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": blog})
}

// DeleteBlog handles DELETE /api/blog/:id
func DeleteBlog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid blog ID"})
		return
	}

	// 检查博客是否存在
	blog, err := repo.FindBlogByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	if blog == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Blog not found"})
		return
	}

	// 删除博客
	if err := repo.DeleteBlog(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "删除博客失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "博客删除成功"})
}

// AddBlog handles POST /api/blog/add
func AddBlog(c *gin.Context) {
	// 获取Authorization头
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "未授权访问"})
		return
	}

	// 验证token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	username, err := util.ParseToken(tokenString)
	if err != nil || username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "无效的token"})
		return
	}

	var req CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "请求参数错误: " + err.Error()})
		return
	}

	// 打印请求参数
	println("收到博客创建请求: 标题=", req.Title, ", 内容长度=", len(req.Content))

	blog := &model.Blog{
		Title:      req.Title,
		Content:    req.Content,
		UploadDate: time.Now(),
		UpdateDate: time.Now(),
	}

	if err := repo.CreateBlog(blog); err != nil {
		println("创建博客失败:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "创建博客失败: " + err.Error()})
		return
	}

	println("博客创建成功, ID=", blog.ID)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "博客创建成功", "data": blog})
}

// UpdateBlogByPath handles POST /api/blog/update/:id
func UpdateBlogByPath(c *gin.Context) {
	log.Printf("UpdateBlogByPath: 开始处理更新请求")
	
	// 获取Authorization头
	authHeader := c.GetHeader("Authorization")
	log.Printf("UpdateBlogByPath: Authorization头: %s", authHeader)
	
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		log.Printf("UpdateBlogByPath: 未授权访问")
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "未授权访问"})
		return
	}

	// 验证token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	// 安全地显示token前缀
	tokenPrefix := tokenString
	if len(tokenString) > 20 {
		tokenPrefix = tokenString[:20] + "..."
	}
	log.Printf("UpdateBlogByPath: 提取的token: %s", tokenPrefix)
	
	username, err := util.ParseToken(tokenString)
	if err != nil || username == "" {
		log.Printf("UpdateBlogByPath: token验证失败: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "无效的token"})
		return
	}
	log.Printf("UpdateBlogByPath: token验证成功，用户: %s", username)

	idStr := c.Param("id")
	log.Printf("UpdateBlogByPath: 博客ID: %s", idStr)
	
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("UpdateBlogByPath: 无效的博客ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid blog ID"})
		return
	}

	var req UpdateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("UpdateBlogByPath: 请求参数绑定失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "请求参数错误"})
		return
	}
	log.Printf("UpdateBlogByPath: 请求参数 - 标题: %s, 内容长度: %d", req.Title, len(req.Content))

	// 查找现有博客
	blog, err := repo.FindBlogByID(uint(id))
	if err != nil {
		log.Printf("UpdateBlogByPath: 查找博客失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	if blog == nil {
		log.Printf("UpdateBlogByPath: 博客不存在")
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Blog not found"})
		return
	}
	log.Printf("UpdateBlogByPath: 找到博客 - ID: %d, 标题: %s", blog.ID, blog.Title)

	// 更新博客
	blog.Title = req.Title
	blog.Content = req.Content
	blog.UpdateDate = time.Now()
	log.Printf("UpdateBlogByPath: 准备更新博客")

	if err := repo.UpdateBlog(blog); err != nil {
		log.Printf("UpdateBlogByPath: 更新博客失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "更新博客失败"})
		return
	}

	log.Printf("UpdateBlogByPath: 博客更新成功")
	c.JSON(http.StatusOK, gin.H{"success": true, "data": blog})
}

// DeleteBlogByPath handles DELETE /api/blog/delete/:id
func DeleteBlogByPath(c *gin.Context) {
	// 获取Authorization头
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "未授权访问"})
		return
	}

	// 验证token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	username, err := util.ParseToken(tokenString)
	if err != nil || username == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "无效的token"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid blog ID"})
		return
	}

	// 检查博客是否存在
	blog, err := repo.FindBlogByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	if blog == nil {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Blog not found"})
		return
	}

	// 删除博客
	if err := repo.DeleteBlog(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "删除博客失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "博客删除成功"})
}

// RegisterBlogRoutes 注册博客相关路由
func RegisterBlogRoutes(rg *gin.RouterGroup) {
	blogRoutes := rg.Group("/blog")
	{
		// 公开接口（无需认证）
		blogRoutes.GET("/all", GetAllBlogs)
		blogRoutes.GET("/:id", GetBlogByID)
		
		// 与前端兼容的管理接口（需要认证，但在路由层面不强制）
		blogRoutes.POST("/add", AddBlog)
		blogRoutes.POST("/update/:id", UpdateBlogByPath)
		blogRoutes.DELETE("/delete/:id", DeleteBlogByPath)
	}
}

// AdminLogin 处理管理员登录请求
func AdminLogin(c *gin.Context) {
    // 定义登录请求结构
    type LoginRequest struct {
        Username string `json:"username" binding:"required"`
        Password string `json:"password" binding:"required"`
    }

    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "请求参数错误"})
        return
    }

    // 从数据库查询管理员
    admin, err := repo.FindAdminByUsername(req.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "服务器错误"})
        return
    }

    // 验证用户名和密码
    if admin == nil || admin.Password != req.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "用户名或密码错误"})
        return
    }

    // 登录成功，生成JWT token
    token, err := util.GenerateToken(req.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "生成token失败"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true, 
        "message": "登录成功", 
        "data": gin.H{
            "token": "Bearer " + token,
            "username": req.Username,
        },
    })
}

// RegisterAdminRoutes 注册管理员相关路由
func RegisterAdminRoutes(rg *gin.RouterGroup) {
	adminRoutes := rg.Group("/admin")
	{
		// 登录接口
		adminRoutes.POST("/login", AdminLogin)
	}
}
