package wallet

import (
	"encoding/json"
	"fmt"

	"testgoapi/config"
	"testgoapi/service/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func UpdateBalance(c *fiber.Ctx) error {
	params := struct {
		BrandId       uint              `json:"brandid" validate:"required"`
		PlayerId      uint              `json:"playerid" validate:"required"`
		SessionId     string            `json:"sessionid" validate:"required,alphanum,max=10"`
		Amount        float64           `json:"amount" validate:"gte=0,lte=10000000000"`
		Currency      string            `json:"currency" validate:"required"`
		TransactionId string            `json:"transactionid" validate:"required"`
		Game          uint              `json:"game" validate:"required"`
		GameRound     string            `json:"gameround" validate:"required"`
		At            string            `json:"at"`
		Extra         map[string]string `json:"extra"`
	}{}

	if err := c.BodyParser(&params); err != nil {
		error := validator.BodyParserFails(err)
		return c.JSON(error)
	}

	if err := validator.Fails(&params); err != nil {
		return c.JSON(err)
	}

	if !lo.Contains(config.Brands, params.BrandId) {
		err := validator.FieldFails("brandid")
		return c.JSON(err)
	}

	// true
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
