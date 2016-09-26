// This file was generated by counterfeiter
package modelsfakes

import (
	"gcp-service-broker/brokerapi/brokers/models"
	"sync"
)

type FakeAccountManager struct {
	CreateAccountInGoogleStub        func(instanceID string, bindingID string, details models.BindDetails, instance models.ServiceInstanceDetails) (models.ServiceBindingCredentials, error)
	createAccountInGoogleMutex       sync.RWMutex
	createAccountInGoogleArgsForCall []struct {
		instanceID string
		bindingID  string
		details    models.BindDetails
		instance   models.ServiceInstanceDetails
	}
	createAccountInGoogleReturns struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	DeleteAccountFromGoogleStub        func(creds models.ServiceBindingCredentials) error
	deleteAccountFromGoogleMutex       sync.RWMutex
	deleteAccountFromGoogleArgsForCall []struct {
		creds models.ServiceBindingCredentials
	}
	deleteAccountFromGoogleReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAccountManager) CreateAccountInGoogle(instanceID string, bindingID string, details models.BindDetails, instance models.ServiceInstanceDetails) (models.ServiceBindingCredentials, error) {
	fake.createAccountInGoogleMutex.Lock()
	fake.createAccountInGoogleArgsForCall = append(fake.createAccountInGoogleArgsForCall, struct {
		instanceID string
		bindingID  string
		details    models.BindDetails
		instance   models.ServiceInstanceDetails
	}{instanceID, bindingID, details, instance})
	fake.recordInvocation("CreateAccountInGoogle", []interface{}{instanceID, bindingID, details, instance})
	fake.createAccountInGoogleMutex.Unlock()
	if fake.CreateAccountInGoogleStub != nil {
		return fake.CreateAccountInGoogleStub(instanceID, bindingID, details, instance)
	} else {
		return fake.createAccountInGoogleReturns.result1, fake.createAccountInGoogleReturns.result2
	}
}

func (fake *FakeAccountManager) CreateAccountInGoogleCallCount() int {
	fake.createAccountInGoogleMutex.RLock()
	defer fake.createAccountInGoogleMutex.RUnlock()
	return len(fake.createAccountInGoogleArgsForCall)
}

func (fake *FakeAccountManager) CreateAccountInGoogleArgsForCall(i int) (string, string, models.BindDetails, models.ServiceInstanceDetails) {
	fake.createAccountInGoogleMutex.RLock()
	defer fake.createAccountInGoogleMutex.RUnlock()
	return fake.createAccountInGoogleArgsForCall[i].instanceID, fake.createAccountInGoogleArgsForCall[i].bindingID, fake.createAccountInGoogleArgsForCall[i].details, fake.createAccountInGoogleArgsForCall[i].instance
}

func (fake *FakeAccountManager) CreateAccountInGoogleReturns(result1 models.ServiceBindingCredentials, result2 error) {
	fake.CreateAccountInGoogleStub = nil
	fake.createAccountInGoogleReturns = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountManager) DeleteAccountFromGoogle(creds models.ServiceBindingCredentials) error {
	fake.deleteAccountFromGoogleMutex.Lock()
	fake.deleteAccountFromGoogleArgsForCall = append(fake.deleteAccountFromGoogleArgsForCall, struct {
		creds models.ServiceBindingCredentials
	}{creds})
	fake.recordInvocation("DeleteAccountFromGoogle", []interface{}{creds})
	fake.deleteAccountFromGoogleMutex.Unlock()
	if fake.DeleteAccountFromGoogleStub != nil {
		return fake.DeleteAccountFromGoogleStub(creds)
	} else {
		return fake.deleteAccountFromGoogleReturns.result1
	}
}

func (fake *FakeAccountManager) DeleteAccountFromGoogleCallCount() int {
	fake.deleteAccountFromGoogleMutex.RLock()
	defer fake.deleteAccountFromGoogleMutex.RUnlock()
	return len(fake.deleteAccountFromGoogleArgsForCall)
}

func (fake *FakeAccountManager) DeleteAccountFromGoogleArgsForCall(i int) models.ServiceBindingCredentials {
	fake.deleteAccountFromGoogleMutex.RLock()
	defer fake.deleteAccountFromGoogleMutex.RUnlock()
	return fake.deleteAccountFromGoogleArgsForCall[i].creds
}

func (fake *FakeAccountManager) DeleteAccountFromGoogleReturns(result1 error) {
	fake.DeleteAccountFromGoogleStub = nil
	fake.deleteAccountFromGoogleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAccountManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAccountInGoogleMutex.RLock()
	defer fake.createAccountInGoogleMutex.RUnlock()
	fake.deleteAccountFromGoogleMutex.RLock()
	defer fake.deleteAccountFromGoogleMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAccountManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ models.AccountManager = new(FakeAccountManager)
