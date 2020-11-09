package main

import "fmt"
import "time"
import "github.com/modbus"

func main() {
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 10
	handler.Timeout = 1 * time.Second

	defer handler.Close()

	client := modbus.NewClient(handler)
	
	for {
		results, err := client.ReadHoldingRegisters(1, 1)

		fmt.Println(results, err)

		time.Sleep(1 * time.Second)
	}
}