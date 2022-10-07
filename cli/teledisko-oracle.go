package main

import (
	"flag"
	"log"

	"github.com/TelediskoDAO/oracle"
	"github.com/TelediskoDAO/oracle/blockchain"
	"github.com/TelediskoDAO/oracle/blockchain/neokingdom"
	"github.com/TelediskoDAO/oracle/config"
	"github.com/TelediskoDAO/oracle/odoo"
)

type Clients struct {
	DAO  blockchain.DAO
	Odoo *odoo.Client
}

func LoadClients() Clients {
	config, err := config.NewConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	w, err := blockchain.NewWalletFromPrivateKey(config.EthPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	w.Dial(config.EthEndpoint)
	log.Println("Oracle address:", w.Address)

	daoClient := neokingdom.NewDAO(config.ERC20Address)
	daoClient.Connect(w)
	odooClient := odoo.Dial("https://odoo.teledisko.com/jsonrpc")
	if err := odooClient.Authenticate("teledisko", config.OdooUsername, config.OdooPassword); err != nil {
		log.Fatal("Cannot login ", err)
	}

	return Clients{
		DAO:  daoClient,
		Odoo: odooClient,
	}
}

func main() {
	flag.Parse()

	clients := LoadClients()

	payroll, err := clients.Odoo.GetPayroll()
	if err != nil {
		log.Fatal("Cannot get payroll: ", err)
	}
	payroll.PrintTable()

	if err := oracle.MintTokens(clients.Odoo, clients.DAO, payroll); err != nil {
		log.Fatal(err)
	}
}
