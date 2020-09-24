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

func (s *Service) RegisterBatch(amount int) (*domain_models.DevEUIBatch, error) {
	devEUIs := domain_models.GenerateBatchWithUniqueShortCodes(amount)

	wg := &sync.WaitGroup{}
	for i, devEUI := range devEUIs.Batch {
		wg.Add(1)
		go s.registerDevEUIRoutines(i, devEUI, devEUIs, wg)
	}

	wg.Wait()

	return devEUIs, nil
}

func (s *Service) registerDevEUIRoutines(skipIndex int, inputEUI *domain_models.DevEUI, batch *domain_models.DevEUIBatch, wg *sync.WaitGroup) (*domain_models.DevEUI, error) {
	defer wg.Done()
	return s.registerDevEUI(skipIndex, inputEUI, batch)
}

func (s *Service) registerDevEUI(skipIndex int, inputEUI *domain_models.DevEUI, batch *domain_models.DevEUIBatch) (*domain_models.DevEUI, error) {
	outputEUI, err := s.Client.RegisterDevEUI(inputEUI)
	if err == client_models.ErrDevEUIAlreadyExists {
		inputEUI = domain_models.GenerateUniqueShortCode(skipIndex, inputEUI, batch)
		return s.registerDevEUI(skipIndex, inputEUI, batch)
	}
	if err != nil {
		return nil, err
	}

	return outputEUI, nil
}
