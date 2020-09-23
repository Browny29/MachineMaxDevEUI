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

func GenerateBatchWithUniqueShortCodes(amount int) []DevEUI {
	devEUIs := generateBatch(amount)

	for i, eui := range devEUIs {
		devEUIs[i] = GenerateUniqueShortCode(i, eui, devEUIs)
	}
	return devEUIs
}

func GenerateUniqueShortCode(skipIndex int, evaluationEUI DevEUI, devEUIs []DevEUI) DevEUI {
	var uniqueEUI DevEUI
	var isUnique bool = true

	for i := 0 + 1; i < len(devEUIs); i++ {
		if i != skipIndex && evaluationEUI.ShortCode == devEUIs[i].ShortCode {
			uniqueEUI = GenerateUniqueShortCode(skipIndex, generateNew(), devEUIs)
			isUnique = false
		}
	}

	if isUnique {
		uniqueEUI = evaluationEUI
	}

	return uniqueEUI
}

func generateNew() DevEUI {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("ERROR DURING GENERATING EUI:", err)
		return DevEUI{}

	}
	deveui := hex.EncodeToString(bytes)
	return DevEUI{
		ID:        deveui,
		ShortCode: deveui[11:],
	}
}

func generateBatch(amount int) []DevEUI {
	bytes := make([]byte, amount * 8)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("ERROR DURING GENERATING EUI BATCH:", err)
		return nil
	}

	euiSlice := make([]DevEUI, amount)
	completeHexString := hex.EncodeToString(bytes)
	for i:=0; i < amount * 16; i += 16 {
		j := i/16
		euiSlice[j].ID = completeHexString[i:i+16]
		euiSlice[j].ShortCode = completeHexString[i+11:i+16]
	}

	return euiSlice
}