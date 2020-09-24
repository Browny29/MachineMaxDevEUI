package deveui_service

import (
	"github.com/stretchr/testify/assert"
	"machine_max_deveui_generator/deveui_service/domain_models"
	"machine_max_deveui_generator/lorawan_provider_client/client_models"
	"math/rand"
	"testing"
)

func TestService_RegisterBatch_Mocked(t *testing.T) {
	const amount = 1000

	s := NewDefaultService()

	mockedClient := &iClientMock{
		RegisterDevEUIFunc: func(eui *domain_models.DevEUI) (*domain_models.DevEUI, error) {
			num := rand.Intn(4)
			if num == 0 {
				return nil, client_models.ErrDevEUIAlreadyExists
			}
			return eui, nil
		},
	}

	s.Client = mockedClient

	result, err := s.RegisterBatch(amount)
	assert.NoError(t, err)
	assert.Equal(t, amount, len(result.Batch))

	for i := range result.Batch {
		for j := range result.Batch {
			if i != j {
				assert.NotEqual(t, result.Batch[i].ShortCode, result.Batch[j].ShortCode)
			}
		}
	}
}

func TestService_RegisterBatch(t *testing.T) {
	const amount = 100

	s := NewDefaultService()

	result, err := s.RegisterBatch(amount)
	assert.NoError(t, err)
	assert.Equal(t, amount, len(result.Batch))

	for i := range result.Batch {
		for j := range result.Batch {
			if i != j {
				assert.NotEqual(t, result.Batch[i].ShortCode, result.Batch[j].ShortCode)
			}
		}
	}
}
