package domain_models

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
)

type DevEUIBatch struct {
	Batch []*DevEUI
	Lock  *sync.Mutex
}

func GenerateBatchWithUniqueShortCodes(amount int) *DevEUIBatch {
	devEUIs := generateBatch(amount)

	for i, eui := range devEUIs.Batch {
		devEUIs.Batch[i] = GenerateUniqueShortCode(i, eui, devEUIs)
	}
	return devEUIs
}

func generateBatch(amount int) *DevEUIBatch {
	bytes := make([]byte, amount*8)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("ERROR DURING GENERATING EUI BATCH:", err)
		return nil
	}

	euiSlice := make([]*DevEUI, amount)
	completeHexString := hex.EncodeToString(bytes)
	for i := 0; i < amount*16; i += 16 {
		j := i / 16
		euiSlice[j] = new(DevEUI)
		euiSlice[j].ID = completeHexString[i : i+16]
		euiSlice[j].ShortCode = completeHexString[i+11 : i+16]
	}

	return &DevEUIBatch{
		Batch: euiSlice,
		Lock:  new(sync.Mutex),
	}
}
