package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/spiermar/zerocoin/blockchain"
)

func main() {
	blockchain.GenerateGenesisBlock("Hello, World!")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/mineBlock", func(c echo.Context) error {
		newBlock := blockchain.GenerateNextBlock(c.FormValue("data"))
		out, err := json.Marshal(newBlock)
		if err != nil {
			panic(err)
		}
		return c.String(http.StatusOK, string(out))
	})

	e.Start(":1323")
}
