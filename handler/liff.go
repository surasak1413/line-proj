package handler

import (
	"encoding/json"
	"fmt"
	"line-proj/line_api"
	"line-proj/service"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductReq struct {
	UserId       string `json:"user_id"`
	ID           string `json:"id"`
	ProductImg   string `json:"product_img"`
	ProductName  string `json:"product_name"`
	ProductPrice string `json:"product_price"`
}

func BuyProductsAPI(ctx echo.Context) error {
	req := new(ProductReq)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	if err := writeJSONToFile(req); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	body := line_api.PushMessageRequest{
		To: req.UserId,
		Messages: []interface{}{
			line_api.TextMessage{
				Type: service.MessageTypeText,
				Text: fmt.Sprintf("Product ID: %s", req.ProductName),
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return ctx.JSON(200, "success")
}

func GetOrdersApi(ctx echo.Context) error {
	req := new(ProductReq)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	file, err := os.ReadFile("orders.json")
	if err != nil {
		file = []byte("[]") // If file does not exist, initialize it with an empty array
	}

	// Decode existing JSON data
	var orders []OrderData
	if err := json.Unmarshal(file, &orders); err != nil {
		return err
	}

	var respData []OrderData
	for _, item := range orders {
		if item.UserId == req.UserId {
			respData = append(respData, item)
		}
	}

	sort.Slice(respData, func(i, j int) bool {
		return respData[i].OrderTime > orders[j].OrderTime
	})

	return ctx.JSON(200, map[string]interface{}{
		"datas": respData,
	})
}

func GetProductsApi(ctx echo.Context) error {

	file, err := os.ReadFile("products.json")
	if err != nil {
		file = []byte("[]") // If file does not exist, initialize it with an empty array
	}

	// Decode existing JSON data
	var respData []ProductData
	if err := json.Unmarshal(file, &respData); err != nil {
		return err
	}

	return ctx.JSON(200, map[string]interface{}{
		"datas": respData,
	})
}

type OrderData struct {
	OrderId      string `json:"order_id"`
	OrderTime    int64  `json:"order_time"`
	UserId       string `json:"user_id"`
	ID           string `json:"id"`
	ProductImg   string `json:"product_img"`
	ProductName  string `json:"product_name"`
	ProductPrice string `json:"product_price"`
}

type ProductData struct {
	ID           string `json:"id"`
	ProductImg   string `json:"product_img"`
	ProductName  string `json:"product_name"`
	ProductPrice string `json:"product_price"`
}

func writeJSONToFile(req *ProductReq) error {
	// Open or create the orders.json file
	file, err := os.ReadFile("orders.json")
	if err != nil {
		file = []byte("[]") // If file does not exist, initialize it with an empty array
	}

	// Decode existing JSON data
	var orders []OrderData
	if err := json.Unmarshal(file, &orders); err != nil {
		return err
	}
	currentTime := time.Now().Unix()
	// Append new data
	order := OrderData{
		OrderId:      uuid.New().String(),
		OrderTime:    currentTime,
		UserId:       req.UserId,
		ID:           req.ID,
		ProductImg:   req.ProductImg,
		ProductName:  req.ProductName,
		ProductPrice: req.ProductPrice,
	}

	// Append new data
	orders = append(orders, order)

	// Encode updated data
	updatedData, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		return err
	}

	// Write to the file
	if err := os.WriteFile("orders.json", updatedData, 0644); err != nil {
		return err
	}

	fmt.Println("Data has been written to orders.json")
	return nil
}
