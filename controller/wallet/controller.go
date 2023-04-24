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
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := validator.Fails(&params); err != nil {
		return c.JSON(err)
	}

	/*
		err := validator.New().Struct(&params)

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			if len(validationErrors) > 0 {
				dump.Printf(validationErrors[0])
			}
		}
	*/

	if true {
		panic("dopanic")
	}

	//	arr := map[string]int{"foo": 1, "name": 2}
	/*
		if err := validator.New().Struct(&params); err != nil {
			//return c.JSON(err.Error())
			fmt.Printf("%T", err.Error())
			//fmt.Println(err.Error())
			return c.SendString(err.Error())
		}
	*/
	obj, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(obj))
	//fmt.Println(c.AllParams())
	return c.JSON(params)
}
