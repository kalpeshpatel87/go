package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const charset = "0123456789"

func generateRandomString(prefix string, length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return prefix + string(result)
}

type Product struct {
	Name  string
	Price float64
}

type CartItem struct {
	Product Product
	Qty     int
}

type Cart struct {
	Items map[int]CartItem
}

func NewCart() *Cart {
	return &Cart{Items: make(map[int]CartItem)}
}

func (c *Cart) AddToCart(index int, product Product, qty int) {
	if item, exists := c.Items[index]; exists {
		item.Qty += qty
		c.Items[index] = item
	} else {
		c.Items[index] = CartItem{Product: product, Qty: qty}
	}
}

func (c *Cart) PrintInvoice(customerName string, mobileNumber int, billNo, date, day, timeStr string) {
	nextSerial := serialGen()

	fmt.Println("\n-----------------------------------------------------------------")
	fmt.Printf("\n%23s%s%23s\n", "", "SKY Stationery", "")
	fmt.Println("\n-------------------------| Invoice |-----------------------------")
	fmt.Printf("Date: %s (%s) %s\n", date, day, timeStr)
	fmt.Printf("Bill Number: %s\n", billNo)
	fmt.Printf("Customer Name: %s\n", customerName)
	fmt.Printf("Mobile Number: %d\n", mobileNumber)

	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf("%-10s %-15s %-15s %-15s\n", "S.No", "Product", "Quantity", "Total")

	finalTotal := 0.0
	for _, item := range c.Items {
		total := float64(item.Qty) * item.Product.Price
		fmt.Printf("%-10d %-15s %-15d %-18.2f\n", nextSerial(), item.Product.Name, item.Qty, total)
		finalTotal += total
	}

	fmt.Println("-----------------------------------------------------------------")
	fmt.Printf("Final Total: %.2f\n", finalTotal)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("\n-------------------| Thanks for Visit our Shop |-----------------")
	fmt.Println("\n--------------------------| Software by |------------------------")
	fmt.Println("Atyantik Technologies Pvt. Ltd")
	fmt.Println("501, Privilege Avenue, Vadodara, Gujarat 390023")
	fmt.Println("-----------------------------------------------------------------")
}

func serialGen() func() int {
	serial := 0
	return func() int {
		serial++
		return serial
	}
}

func handleSelection(products []Product, cart *Cart) {
	var prodNum, qty int

	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("Enter product number to buy (0 to exit):")
	fmt.Scan(&prodNum)

	if prodNum == 0 {
		return
	}

	if prodNum < 1 || prodNum > len(products) {
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("XXX Invalid product number. XXX")
		handleSelection(products, cart)
		return
	}

	fmt.Print("Enter Quantity: ")
	fmt.Scan(&qty)

	if qty <= 0 {
		fmt.Println("Invalid Quantity.")
		handleSelection(products, cart)
		return
	}

	cart.AddToCart(prodNum, products[prodNum-1], qty)

	handleSelection(products, cart)
}

func main() {
	billNo := generateRandomString("SKY", 5)

	var mobileNumber int
	reader := bufio.NewReader(os.Stdin)

	currentTime := time.Now()
	currentDate := currentTime.Format("02-Jan-2006")
	currentDay := currentTime.Format("Monday")
	currentTimeFormatted := currentTime.Format("03:04:05 PM")

	fmt.Print("Enter Customer Name: ")
	nameInput, _ := reader.ReadString('\n')
	customerName := strings.TrimSpace(nameInput)

	fmt.Printf("Enter Mobile Number: ")
	fmt.Scanln(&mobileNumber)

	products := []Product{
		{"Pencils", 5},
		{"Eraser", 5},
		{"Shapner", 10},
		{"Pens", 10},
		{"Staplers", 45},
		{"Notebooks", 30},
		{"Folders", 20},
		{"Tapes", 10},
		{"Envelopes", 5},
		{"Covers", 40},
		{"Compass Box", 50},
	}

	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("Available Items:")
	fmt.Println("-----------------------------------------------------------------")
	for i, p := range products {
		fmt.Printf("%d. %s - %.2f\n", i+1, p.Name, p.Price)
	}

	cart := NewCart()
	handleSelection(products, cart)
	cart.PrintInvoice(customerName, mobileNumber, billNo, currentDate, currentDay, currentTimeFormatted)
}
