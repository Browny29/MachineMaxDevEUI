package main

import (
	"bufio"
	"fmt"
	"machine_max_deveui_generator/deveui_service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	const amount = 100

	processDone := make(chan bool, 1)

	setupCloseHandler(processDone)
	createBatch(amount, processDone)

	print("Press enter to quit...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func setupCloseHandler(processDone chan bool) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		fmt.Println("\n- Program will quit after process has finished, please wait...")
		processDone <- true
		os.Exit(0)
	}()
}

func createBatch(amount int, processDone chan bool) {
	processDone <- true
	println("Starting DevEUI registration: registering", amount, "DevEUIs")

	s := deveui_service.NewDefaultService()

	output, err := s.RegisterBatch(amount)
	if err != nil {
		println("ERROR:", err.Error())
	}

	println("Successfully registered", amount, "DevEUIs. RESULT:")
	println()
	for _, devEUI := range output.Batch {
		println(devEUI.ID)
	}
	println()

	println("Printed all DevEUIs.")
	<- processDone
}
