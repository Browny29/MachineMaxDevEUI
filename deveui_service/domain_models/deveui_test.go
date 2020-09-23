package domain_models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateBatch(t *testing.T) {
	const amount = 1000

	euiSlice := generateBatch(amount)

	assert.Len(t, euiSlice, amount)
	assert.Len(t, euiSlice[0].ID, 16)
	assert.Len(t, euiSlice[0].ShortCode, 5)
	assert.Equal(t, euiSlice[0].ID[11:16], euiSlice[0].ShortCode)
}

func TestGenerateNew(t *testing.T) {
	eui := generateNew()

	assert.Len(t, eui.ID, 16)
	assert.Len(t, eui.ShortCode, 5)
	assert.Equal(t, eui.ID[11:16], eui.ShortCode)
}

func TestGenerateBatchWithUniqueShortCodes(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		devEUIs := GenerateBatchWithUniqueShortCodes(3)

		assert.NotEqual(t, devEUIs[0].ShortCode, devEUIs[1].ShortCode)
		assert.NotEqual(t, devEUIs[1].ShortCode, devEUIs[2].ShortCode)
		assert.NotEqual(t, devEUIs[0].ShortCode, devEUIs[2].ShortCode)
	}
}

func TestGenerateUniqueShortCode(t *testing.T) {
	notUniqueDevEUIs := []DevEUI{
		{
			ID:        "1234567890123456",
			ShortCode: "23456",
		},
		{
			ID:        "1234567890123456",
			ShortCode: "23456",
		},
		{
			ID:        "1234567890123456",
			ShortCode: "23456",
		},
	}

	notUniqueDevEUIs[0] = GenerateUniqueShortCode(0, notUniqueDevEUIs[0], notUniqueDevEUIs)
	notUniqueDevEUIs[1] = GenerateUniqueShortCode(1, notUniqueDevEUIs[1], notUniqueDevEUIs)
	notUniqueDevEUIs[2] = GenerateUniqueShortCode(2, notUniqueDevEUIs[2], notUniqueDevEUIs)

	assert.NotEqual(t, notUniqueDevEUIs[0].ShortCode, notUniqueDevEUIs[1].ShortCode)
	assert.NotEqual(t, notUniqueDevEUIs[1].ShortCode, notUniqueDevEUIs[2].ShortCode)
	assert.NotEqual(t, notUniqueDevEUIs[0].ShortCode, notUniqueDevEUIs[2].ShortCode)
}
