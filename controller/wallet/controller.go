package wallet

import (
	"encoding/json"
	"fmt"

	"testgoapi/service/validator"

	"github.com/gofiber/fiber/v2"
)

func UpdateBalance(c *fiber.Ctx) error {
	params := struct {
		BrandId       uint              `json:"brandid" validate:"required"`
		PlayerId      uint              `json:"playerid" validate:"required"`
		SessionId     string            `json:"sessionid" validate:"required,alphanum,max=10"`
		Amount        float64           `json:"amount" validate:"gte=0,lte=10000000000"`
		Currency      string            `json:"currency"`
		TransactionId string            `json:"transactionid"`
		Game          string            `json:"game"`
		GameRound     string            `json:"gameround"`
		At            string            `json:"at"`
		Extra         map[string]string `json:"extra"`
	}{}
	/*
		params := struct {
			BrandId uint `json:"brandid" validate:"required"`
		}{}
	*/
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := validator.Fails(&params); err != nil {
		return c.JSON(err)
	}

	if true {
		panic("dopanic")
	}

	obj, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(obj))
	//fmt.Println(c.AllParams())
	//return c.SendString(err.Error())
	return c.JSON(params)
}
