package wallet

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UpdateBalance(c *fiber.Ctx) error {
	params := struct {
		BrandId       uint              `json:"brandid" validate:"required"`
		PlayerId      uint              `json:"playerid"`
		SessionId     string            `json:"sessionid" validate:"required,alphanum,max:10"`
		Amount        float64           `json:"amount" validate:"gte=0,lte=10000000000"`
		Currency      string            `json:"currency"`
		TransactionId string            `json:"transactionid"`
		Game          string            `json:"game"`
		GameRound     string            `json:"gameround"`
		At            string            `json:"at"`
		Extra         map[string]string `json:"extra"`
	}{}
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := validator.New().Struct(&params); err != nil {
		//return c.JSON(err.Error())
		return c.SendString(err.Error())
	}

	obj, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(obj))
	//fmt.Println(c.AllParams())
	return c.JSON(params)
}
