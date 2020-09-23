package deveui_service

import (
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

func (s *Service) RegisterBatch(amount int) ([]domain_models.DevEUI, error) {
	devEUISlice := domain_models.GenerateBatchWithUniqueShortCodes(amount)

	wg := sync.WaitGroup{}
	for i, devEUI := range devEUISlice {
		wg.Add(1)

		go func() error{
			defer wg.Done()
			outputDevEUI, err := s.registerUniqueDevEUI(i, devEUI, devEUISlice)
			if err != nil {
				return err
			}
			devEUISlice[i] = *outputDevEUI

			return nil
		} ()
	}

	wg.Wait()

	return devEUISlice, nil
}

func (s *Service) registerUniqueDevEUI(skipIndex int, inputEUI domain_models.DevEUI, devEUIs []domain_models.DevEUI) (*domain_models.DevEUI, error) {
	outputEUI, err := s.Client.RegisterDevEUI(&inputEUI)
	if err == client_models.ErrDevEUIAlreadyExists {
		inputEUI = domain_models.GenerateUniqueShortCode(skipIndex, inputEUI, devEUIs)
		return s.registerUniqueDevEUI(skipIndex, inputEUI, devEUIs)
	}
	if err != nil {
		return nil, err
	}

	return outputEUI, nil
}
