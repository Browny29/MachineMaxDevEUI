package domain_models

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type DevEUI struct {
	ID        string
	ShortCode string
}

func GenerateUniqueShortCode(skipIndex int, evaluationEUI *DevEUI, devEUIs *DevEUIBatch) *DevEUI {
	var uniqueEUI *DevEUI
	var isUnique bool = true

	for i := 0 + 1; i < len(devEUIs.Batch); i++ {
		if i != skipIndex && evaluationEUI.ShortCode == devEUIs.Batch[i].ShortCode {
			uniqueEUI = GenerateUniqueShortCode(skipIndex, generateNew(), devEUIs)
			isUnique = false
		}
	}

	if isUnique {
		uniqueEUI = evaluationEUI
	}

	return uniqueEUI
}

func generateNew() *DevEUI {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("ERROR DURING GENERATING EUI:", err)
		return &DevEUI{}

	}
	deveui := hex.EncodeToString(bytes)
	return &DevEUI{
		ID:        deveui,
		ShortCode: deveui[11:],
	}
}
