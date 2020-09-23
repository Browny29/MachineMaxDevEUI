package lorawan_provider_client

import (
	"github.com/stretchr/testify/assert"
	"machine_max_deveui_generator/deveui_service/domain_models"
	"testing"
)

func TestClient_RegisterDevEUI(t *testing.T) {
	input := &domain_models.DevEUI{
		ID:        "0123456789012345",
		ShortCode: "12345",
	}

	c := NewDefaultClient()

	deveui, err := c.RegisterDevEUI(input)
	if err != nil {
		assert.Equal(t, ErrDevEUIAlreadyExists, err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, input.ID, deveui.ID)
		assert.Equal(t, input.ShortCode, deveui.ShortCode)
	}
}
