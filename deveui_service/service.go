package deveui_service

import (
	"fmt"
	"machine_max_deveui_generator/deveui_service/domain_models"
	"machine_max_deveui_generator/lorawan_provider_client"
	"machine_max_deveui_generator/lorawan_provider_client/client_models"
	"sync"
)

//go:generate moq -out service_moq_test.go . iClient

type iClient interface {
	RegisterDevEUI(eui *domain_models.DevEUI) (*domain_models.DevEUI, error)
}

type Service struct {
	Client iClient
}

func NewDefaultService() *Service {
	return &Service{
		Client: lorawan_provider_client.NewDefaultClient(),
	}
}

func (s *Service) RegisterBatch(amount int) (*domain_models.DevEUIBatch, error) {
	// First we need to generate a batch each with a unique shortcode
	devEUIs := domain_models.GenerateBatchWithUniqueShortCodes(amount)

	wg := &sync.WaitGroup{}
	maxRoutines := 10 // this number signifies the amount of go routines that are allowed to be running concurrently
	routineStop := make(chan struct{}, maxRoutines) // the routineStop keeps track of the amount of go routines currently running

	// Start registering the DevEUIs asynchronously
	for i, devEUI := range devEUIs.Batch {
		wg.Add(1)
		routineStop <- struct{}{} // signals a new go routine has started. This blocks if the maxRoutines number has been reached

		// register a DevEUI asynchronously
		go func(i int, devEUI *domain_models.DevEUI, devEUIs *domain_models.DevEUIBatch, wg *sync.WaitGroup) {
			defer wg.Done()
			s.registerDevEUI(i, devEUI, devEUIs, 0)
			<- routineStop // signals the go routine has ended. If the max was reached a new routine can start now
		}(i, devEUI, devEUIs, wg)
	}

	wg.Wait()

	return devEUIs, nil
}

func (s *Service) registerDevEUI(skipIndex int, inputEUI *domain_models.DevEUI, batch *domain_models.DevEUIBatch, numberOfTries int) {
	var err error
	numberOfTries++

	if numberOfTries > 10 {
		panic(fmt.Sprintf("#%d of the batch has been retried 10 times. Something is wrong", skipIndex))
	}

	// Register the DevEUI at the LoraWan provider
	batch.Batch[skipIndex], err = s.Client.RegisterDevEUI(inputEUI)

	if err == client_models.ErrDevEUIAlreadyExists { // If we get a EUI already exists error, create a new DevEUI with a unique short code and try again
		// Lock the batch so we are sure this new record will be unique
		batch.Lock.Lock()
		// Generate a new DevEUI with a unique short code
		inputEUI = domain_models.GenerateUniqueShortCode(skipIndex, inputEUI, batch)
		batch.Lock.Unlock()

		// Try registering again
		s.registerDevEUI(skipIndex, inputEUI, batch, numberOfTries)
	} else if err != nil { // If we get a different error try again
		s.registerDevEUI(skipIndex, inputEUI, batch, numberOfTries)
	}
}
