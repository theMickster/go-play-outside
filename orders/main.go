package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"orders/models"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receivedOrdersChannel = make(chan models.Order)
	var validOrderChannel = make(chan models.Order)
	var invalidOrderChannel = make(chan models.InvalidOrder)
	go receiveOrders(receivedOrdersChannel)
	go validateOrders(receivedOrdersChannel, validOrderChannel, invalidOrderChannel)

	wg.Add(1)
	go func() {
		order := <-validOrderChannel
		fmt.Printf("A valid order has been received: %v\n", order)
		wg.Done()
	}()

	go func() {
		invalidOrder := <-invalidOrderChannel
		fmt.Printf("An invalid order has been received and rejected. %v Validation Error: %v\n", invalidOrder.Order, invalidOrder.Err)
		wg.Done()
	}()
	wg.Wait()

}

func validateOrders(inbound <-chan models.Order, outbound chan<- models.Order, errorChan chan<- models.InvalidOrder) {
	order := <-inbound
	if order.Quantity <= 0 {
		errorChan <- models.InvalidOrder{Order: order, Err: errors.New("product quantity must be greater than zero (0)")}
	} else {
		outbound <- order
	}
}

func receiveOrders(outbound chan models.Order) {
	for _, anOrder := range someOrders {
		var newOrder models.Order
		err := json.Unmarshal([]byte(anOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		outbound <- newOrder
	}
}

var someOrders = []string{
	`{"productCode": 1111, "quantity": -5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42, "status": 2}`,
	`{"productCode": 3333, "quantity": 19, "status": 2}`,
	`{"productCode": 4444, "quantity": 8, "status": 3}`,
}
