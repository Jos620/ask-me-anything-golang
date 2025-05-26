package main

import "github.com/gin-gonic/gin"

func errorResponse(err error) map[string]any {
	return gin.H{"error": err.Error()}
}

func defaultResponse(value any) map[string]any {
	return gin.H{"response": value}
}
