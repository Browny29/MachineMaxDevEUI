package deveui_service

import (
	"github.com/stretchr/testify/assert"
	"machine_max_deveui_generator/deveui_service/domain_models"
	"testing"
)

func TestService_RegisterBatch_Mocked(t *testing.T) {
	const amount = 5000

	s := NewDefaultService()

	mockedClient := &iClientMock{
		RegisterDevEUIFunc: func(eui *domain_models.DevEUI) (*domain_models.DevEUI, error) {
			return eui, nil
		},
	}

	s.Client = mockedClient

	result, err := s.RegisterBatch(amount)
	assert.NoError(t, err)
	assert.Equal(t, amount, len(result))
}

func TestService_RegisterBatch(t *testing.T) {
	const amount = 3

	s := NewDefaultService()

	result, err := s.RegisterBatch(amount)
	assert.NoError(t, err)
	assert.Equal(t, amount, len(result))
}
