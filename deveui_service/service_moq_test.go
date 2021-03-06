// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package deveui_service

import (
	"machine_max_deveui_generator/deveui_service/domain_models"
	"sync"
)

// Ensure, that iClientMock does implement iClient.
// If this is not the case, regenerate this file with moq.
var _ iClient = &iClientMock{}

// iClientMock is a mock implementation of iClient.
//
//     func TestSomethingThatUsesiClient(t *testing.T) {
//
//         // make and configure a mocked iClient
//         mockediClient := &iClientMock{
//             RegisterDevEUIFunc: func(eui *domain_models.DevEUI) (*domain_models.DevEUI, error) {
// 	               panic("mock out the RegisterDevEUI method")
//             },
//         }
//
//         // use mockediClient in code that requires iClient
//         // and then make assertions.
//
//     }
type iClientMock struct {
	// RegisterDevEUIFunc mocks the RegisterDevEUI method.
	RegisterDevEUIFunc func(eui *domain_models.DevEUI) (*domain_models.DevEUI, error)

	// calls tracks calls to the methods.
	calls struct {
		// RegisterDevEUI holds details about calls to the RegisterDevEUI method.
		RegisterDevEUI []struct {
			// Eui is the eui argument value.
			Eui *domain_models.DevEUI
		}
	}
	lockRegisterDevEUI sync.RWMutex
}

// RegisterDevEUI calls RegisterDevEUIFunc.
func (mock *iClientMock) RegisterDevEUI(eui *domain_models.DevEUI) (*domain_models.DevEUI, error) {
	if mock.RegisterDevEUIFunc == nil {
		panic("iClientMock.RegisterDevEUIFunc: method is nil but iClient.RegisterDevEUI was just called")
	}
	callInfo := struct {
		Eui *domain_models.DevEUI
	}{
		Eui: eui,
	}
	mock.lockRegisterDevEUI.Lock()
	mock.calls.RegisterDevEUI = append(mock.calls.RegisterDevEUI, callInfo)
	mock.lockRegisterDevEUI.Unlock()
	return mock.RegisterDevEUIFunc(eui)
}

// RegisterDevEUICalls gets all the calls that were made to RegisterDevEUI.
// Check the length with:
//     len(mockediClient.RegisterDevEUICalls())
func (mock *iClientMock) RegisterDevEUICalls() []struct {
	Eui *domain_models.DevEUI
} {
	var calls []struct {
		Eui *domain_models.DevEUI
	}
	mock.lockRegisterDevEUI.RLock()
	calls = mock.calls.RegisterDevEUI
	mock.lockRegisterDevEUI.RUnlock()
	return calls
}
