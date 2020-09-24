package domain_models

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGenerateBatch(t *testing.T) {
	const amount = 1000

	euiSlice := generateBatch(amount)

	assert.Len(t, euiSlice.Batch, amount)
	assert.Len(t, euiSlice.Batch[0].ID, 16)
	assert.Len(t, euiSlice.Batch[0].ShortCode, 5)
	assert.Equal(t, euiSlice.Batch[0].ID[11:16], euiSlice.Batch[0].ShortCode)
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

		assert.NotEqual(t, devEUIs.Batch[0].ShortCode, devEUIs.Batch[1].ShortCode)
		assert.NotEqual(t, devEUIs.Batch[1].ShortCode, devEUIs.Batch[2].ShortCode)
		assert.NotEqual(t, devEUIs.Batch[0].ShortCode, devEUIs.Batch[2].ShortCode)
	}
}

func TestGenerateUniqueShortCode(t *testing.T) {
	notUnique := &DevEUIBatch{
		Batch: []*DevEUI{
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
			}},
		Lock: &sync.Mutex{},
	}

	notUnique.Batch[0] = GenerateUniqueShortCode(0, notUnique.Batch[0], notUnique)
	notUnique.Batch[1] = GenerateUniqueShortCode(1, notUnique.Batch[1], notUnique)
	notUnique.Batch[2] = GenerateUniqueShortCode(2, notUnique.Batch[2], notUnique)

	assert.NotEqual(t, notUnique.Batch[0].ShortCode, notUnique.Batch[1].ShortCode)
	assert.NotEqual(t, notUnique.Batch[1].ShortCode, notUnique.Batch[2].ShortCode)
	assert.NotEqual(t, notUnique.Batch[0].ShortCode, notUnique.Batch[2].ShortCode)
}
