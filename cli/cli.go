package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TelediskoDAO/oracle/config"
	"github.com/TelediskoDAO/oracle/odoo"
)

func main() {
	flag.Parse()
	config, err := config.NewConfigFromEnv()
	if err != nil {
		log.Fatal(("Cannot read config"))
	}

	odooClient := odoo.Dial("https://odoo.teledisko.com/jsonrpc")
	if err := odooClient.Authenticate("teledisko", config.OdooUsername, config.OdooPassword); err != nil {
		log.Fatal("Cannot login ", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/resolutions/rewards/last", func(c *gin.Context) {
		r, err := odooClient.GetRewards()
		if err != nil {
			log.Fatal("Cannot get rewards: ", err)
		}
		rr := odoo.NewRewardsResolution(r)
		c.JSON(http.StatusOK, rr)
	})
	r.Run()
}
