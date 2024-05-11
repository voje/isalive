package main

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/voje/isalive/internal/isalive"
	"github.com/voje/isalive/templates"

	log "github.com/sirupsen/logrus"
)

const envIsaliveSites = "ISALIVE_SITES"

func main() {
    log.SetLevel(log.DebugLevel)

    sitesStr, isSet := os.LookupEnv(envIsaliveSites)
    if !isSet {
	log.Panicf("Set env %s=http://site1.com,https://site2.com", envIsaliveSites)
    }
    sites := strings.Split(sitesStr, ",")


    ia := isalive.NewIsAlive()
    for _, s := range(sites) {
	ia.AddSite(s)
    }

    r := gin.Default()
    r.Static("/assets", "./assets")
    r.GET("/", func(c *gin.Context) {
	ia.CheckHTTPs()
	templates.Index(ia.Sites).Render(context.Background(), c.Writer)
    })

    go func() {
	for {
	    ia.CheckHTTPs()
	    ia.PublishState()
	    time.Sleep(time.Second * 5)
	}
    }()

    r.Run()
}
