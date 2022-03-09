package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusank/goim/apps/gateway/service"
)

func handleDiscoverPushServer(c *gin.Context) {
	uid := c.GetHeader("uid")
	if uid == "" {
		log.Println("uid not found")
		c.JSON(http.StatusOK, gin.H{"err": "uid not found"})
		return
	}

	agentId, err := service.LoadMatchedPushServer(context.Background())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"agentId": agentId})
}
