package main

import (
	"fmt"
)

type Client struct {
	name string
}

func (c Client) SendMessageToPrinter(mp MessagePrinter) {
	fmt.Printf("Client %s sending message to printer...\n", c.name)
	mp.PrintMessage("Hello world!")
}

// MessagePrinter - интерфейс для печати сообщений.
type MessagePrinter interface {
	PrintMessage(string)
}

// MessagePrinterImpl - реализация печати сообщений в текстовом формате.
type MessagePrinterImpl struct{}

func (m *MessagePrinterImpl) PrintMessage(msg string) {
	fmt.Println("Printing message in Impl format:", msg)
}

// JSONPrinter - структура для печати сообщений в формате JSON.
type JSONPrinter struct{}

func (j *JSONPrinter) PrintJSONMessage(msg string) {
	fmt.Println("Printing JSON message:", msg)
}

// JSONAdapter - адаптер для использования JSONPrinter через интерфейс MessagePrinter.
type JSONAdapter struct {
	JSONPrinter *JSONPrinter
}

func (j *JSONAdapter) PrintMessage(msg string) {
	fmt.Println("Adapter converts Impl message to JSON message")
	j.JSONPrinter.PrintJSONMessage(msg)
}

func main() {
	client := &Client{name: "Vlad"}
	standardPrinter := &MessagePrinterImpl{}
	client.SendMessageToPrinter(standardPrinter)

	// Используем JSONPrinter через адаптер.
	jsonPrinter := &JSONPrinter{}
	adapter := &JSONAdapter{JSONPrinter: jsonPrinter}
	client.SendMessageToPrinter(adapter)
}
