package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"

	"github.com/TelediskoDAO/oracle/config"
	"github.com/TelediskoDAO/oracle/odoo"
)

func RenderResolutionPayments(p *odoo.Payroll) {
	t, err := template.ParseFiles("cli/templates/resolution-payments.md")
	if err != nil {
		panic(err)
	}
	// Right now we need to copy paste, that's why we write to stdout
	err = t.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}

func main2() {
	flag.Parse()
	config, err := config.NewConfigFromEnv()
	if err != nil {
		log.Fatal(("Cannot read config"))
	}

	odooClient := odoo.Dial("https://odoo.teledisko.com/jsonrpc")
	if err := odooClient.Authenticate("teledisko", config.OdooUsername, config.OdooPassword); err != nil {
		log.Fatal("Cannot login ", err)
	}

	payroll, err := odooClient.GetPayroll()
	if err != nil {
		log.Fatal("Cannot get payroll: ", err)
	}

	payroll.PrintTable()
	RenderResolutionPayments(payroll)
}

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
	r.GET("/resolutions/rewards/last", func(c *gin.Context) {
		p, err := odooClient.GetPayroll()
		if err != nil {
			log.Fatal("Cannot get payroll: ", err)
		}
		r := odoo.NewRewardsResolution(p)
		c.JSON(http.StatusOK, r)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
