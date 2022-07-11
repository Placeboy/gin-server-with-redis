package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"os"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "I am very healthy",
	})
}

func CreateStudent(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	fmt.Printf("name = %s, age = %s\n", name, age)
	err := Create(name, age)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "Create Student successfully!",
	})
	return
}

func Create(name, age string) error {
	_, err := client.HSet(ctx, "students:", name, age).Result()
	return err
}

func DeleteStudent(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Printf("name = %s\n", name)
	err := Delete(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   0,
		"status": "Delete Student successfully!",
	})
	return
}

func Delete(name string) error {
	_, err := client.HDel(ctx, "students:", name).Result()
	return err
}

// Create Redis Client
var (
	host     = getEnv("REDIS_HOST", "localhost")
	port     = getEnv("REDIS_PORT", "6379")
	password = getEnv("REDIS_PASSWORD", "")
	ctx = context.Background()
)

var client *redis.Client

func init() {
	// Create Redis Client
	client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
}

func main() {
	r := gin.Default()
	r.POST("/createStudent", CreateStudent)
	r.POST("/deleteStudent", DeleteStudent)
	r.GET("/healthz", HealthCheck)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}