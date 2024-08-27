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
	receivedOrdersChannel := receiveOrders()
	validOrderChannel, invalidOrderChannel := validateOrders(receivedOrdersChannel)
	reserveProductInventoryChannel := reserveProductInventory(validOrderChannel)
	fillOrdersChannel := fillOrders(reserveProductInventoryChannel)

	wg.Add(2)
	go func(invalidOrderChannel <-chan models.InvalidOrder) {
		for invalidOrder := range invalidOrderChannel {
			fmt.Printf("An invalid order has been received and rejected. %v Validation Error: %v\n", invalidOrder.Order, invalidOrder.Err)
		}
		wg.Done()
	}(invalidOrderChannel)

	go func(fillOrdersChannel <-chan models.Order) {
		for order := range fillOrdersChannel {
			fmt.Printf("A valid order has been received and processed: %v\n", order)
		}
		wg.Done()
	}(fillOrdersChannel)

	wg.Wait()
}

func fillOrders(inboundChan <-chan models.Order) <-chan models.Order {
	outboundChan := make(chan models.Order)

	var wg sync.WaitGroup
	const workers = 1
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			for order := range inboundChan {
				order.Status = 4
				fmt.Printf("The inventory for an order has been filled: %v\n", order)
				outboundChan <- order
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(outboundChan)
	}()

	return outboundChan
}

func reserveProductInventory(inboundChan <-chan models.Order) <-chan models.Order {
	outboundChan := make(chan models.Order)

	var wg sync.WaitGroup
	const workers = 1
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			for order := range inboundChan {
				order.Status = 3
				fmt.Printf("The inventory for an order has been reserved: %v\n", order)
				outboundChan <- order
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(outboundChan)
	}()

	return outboundChan
}

func validateOrders(inboundChan <-chan models.Order) (<-chan models.Order, <-chan models.InvalidOrder) {
	outboundChan := make(chan models.Order)
	errorChan := make(chan models.InvalidOrder)
	go func() {
		for order := range inboundChan {
			if order.Quantity <= 0 {
				errorChan <- models.InvalidOrder{Order: order, Err: errors.New("product quantity must be greater than zero (0)")}
			} else {
				outboundChan <- order
			}
		}
		close(outboundChan)
		close(errorChan)
	}()
	return outboundChan, errorChan
}

func receiveOrders() chan models.Order {
	outbound := make(chan models.Order)
	go func() {
		for _, anOrder := range someOrders {
			var newOrder models.Order
			err := json.Unmarshal([]byte(anOrder), &newOrder)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("An order has been received: %v\n", newOrder)
			outbound <- newOrder
		}
		close(outbound)
	}()
	return outbound
}

var someOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
	`{"productCode": 5555, "quantity": -18, "status": 1}`,
	`{"productCode": 4444, "quantity": 1, "status": 1}`,
}
