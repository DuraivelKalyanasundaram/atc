// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/creds"
	"github.com/concourse/atc/db"
)

type FakeResourceConfigFactory struct {
	FindOrCreateResourceConfigStub        func(logger lager.Logger, resourceType string, source atc.Source, resourceTypes creds.VersionedResourceTypes) (db.ResourceConfig, error)
	findOrCreateResourceConfigMutex       sync.RWMutex
	findOrCreateResourceConfigArgsForCall []struct {
		logger        lager.Logger
		resourceType  string
		source        atc.Source
		resourceTypes creds.VersionedResourceTypes
	}
	findOrCreateResourceConfigReturns struct {
		result1 db.ResourceConfig
		result2 error
	}
	findOrCreateResourceConfigReturnsOnCall map[int]struct {
		result1 db.ResourceConfig
		result2 error
	}
	FindResourceConfigByIDStub        func(int) (db.ResourceConfig, bool, error)
	findResourceConfigByIDMutex       sync.RWMutex
	findResourceConfigByIDArgsForCall []struct {
		arg1 int
	}
	findResourceConfigByIDReturns struct {
		result1 db.ResourceConfig
		result2 bool
		result3 error
	}
	findResourceConfigByIDReturnsOnCall map[int]struct {
		result1 db.ResourceConfig
		result2 bool
		result3 error
	}
	CleanUnreferencedConfigsStub        func() error
	cleanUnreferencedConfigsMutex       sync.RWMutex
	cleanUnreferencedConfigsArgsForCall []struct{}
	cleanUnreferencedConfigsReturns     struct {
		result1 error
	}
	cleanUnreferencedConfigsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfig(logger lager.Logger, resourceType string, source atc.Source, resourceTypes creds.VersionedResourceTypes) (db.ResourceConfig, error) {
	fake.findOrCreateResourceConfigMutex.Lock()
	ret, specificReturn := fake.findOrCreateResourceConfigReturnsOnCall[len(fake.findOrCreateResourceConfigArgsForCall)]
	fake.findOrCreateResourceConfigArgsForCall = append(fake.findOrCreateResourceConfigArgsForCall, struct {
		logger        lager.Logger
		resourceType  string
		source        atc.Source
		resourceTypes creds.VersionedResourceTypes
	}{logger, resourceType, source, resourceTypes})
	fake.recordInvocation("FindOrCreateResourceConfig", []interface{}{logger, resourceType, source, resourceTypes})
	fake.findOrCreateResourceConfigMutex.Unlock()
	if fake.FindOrCreateResourceConfigStub != nil {
		return fake.FindOrCreateResourceConfigStub(logger, resourceType, source, resourceTypes)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findOrCreateResourceConfigReturns.result1, fake.findOrCreateResourceConfigReturns.result2
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigCallCount() int {
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	return len(fake.findOrCreateResourceConfigArgsForCall)
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigArgsForCall(i int) (lager.Logger, string, atc.Source, creds.VersionedResourceTypes) {
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	return fake.findOrCreateResourceConfigArgsForCall[i].logger, fake.findOrCreateResourceConfigArgsForCall[i].resourceType, fake.findOrCreateResourceConfigArgsForCall[i].source, fake.findOrCreateResourceConfigArgsForCall[i].resourceTypes
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigReturns(result1 db.ResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigStub = nil
	fake.findOrCreateResourceConfigReturns = struct {
		result1 db.ResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) FindOrCreateResourceConfigReturnsOnCall(i int, result1 db.ResourceConfig, result2 error) {
	fake.FindOrCreateResourceConfigStub = nil
	if fake.findOrCreateResourceConfigReturnsOnCall == nil {
		fake.findOrCreateResourceConfigReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfig
			result2 error
		})
	}
	fake.findOrCreateResourceConfigReturnsOnCall[i] = struct {
		result1 db.ResourceConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfigFactory) FindResourceConfigByID(arg1 int) (db.ResourceConfig, bool, error) {
	fake.findResourceConfigByIDMutex.Lock()
	ret, specificReturn := fake.findResourceConfigByIDReturnsOnCall[len(fake.findResourceConfigByIDArgsForCall)]
	fake.findResourceConfigByIDArgsForCall = append(fake.findResourceConfigByIDArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("FindResourceConfigByID", []interface{}{arg1})
	fake.findResourceConfigByIDMutex.Unlock()
	if fake.FindResourceConfigByIDStub != nil {
		return fake.FindResourceConfigByIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findResourceConfigByIDReturns.result1, fake.findResourceConfigByIDReturns.result2, fake.findResourceConfigByIDReturns.result3
}

func (fake *FakeResourceConfigFactory) FindResourceConfigByIDCallCount() int {
	fake.findResourceConfigByIDMutex.RLock()
	defer fake.findResourceConfigByIDMutex.RUnlock()
	return len(fake.findResourceConfigByIDArgsForCall)
}

func (fake *FakeResourceConfigFactory) FindResourceConfigByIDArgsForCall(i int) int {
	fake.findResourceConfigByIDMutex.RLock()
	defer fake.findResourceConfigByIDMutex.RUnlock()
	return fake.findResourceConfigByIDArgsForCall[i].arg1
}

func (fake *FakeResourceConfigFactory) FindResourceConfigByIDReturns(result1 db.ResourceConfig, result2 bool, result3 error) {
	fake.FindResourceConfigByIDStub = nil
	fake.findResourceConfigByIDReturns = struct {
		result1 db.ResourceConfig
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfigFactory) FindResourceConfigByIDReturnsOnCall(i int, result1 db.ResourceConfig, result2 bool, result3 error) {
	fake.FindResourceConfigByIDStub = nil
	if fake.findResourceConfigByIDReturnsOnCall == nil {
		fake.findResourceConfigByIDReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfig
			result2 bool
			result3 error
		})
	}
	fake.findResourceConfigByIDReturnsOnCall[i] = struct {
		result1 db.ResourceConfig
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfigFactory) CleanUnreferencedConfigs() error {
	fake.cleanUnreferencedConfigsMutex.Lock()
	ret, specificReturn := fake.cleanUnreferencedConfigsReturnsOnCall[len(fake.cleanUnreferencedConfigsArgsForCall)]
	fake.cleanUnreferencedConfigsArgsForCall = append(fake.cleanUnreferencedConfigsArgsForCall, struct{}{})
	fake.recordInvocation("CleanUnreferencedConfigs", []interface{}{})
	fake.cleanUnreferencedConfigsMutex.Unlock()
	if fake.CleanUnreferencedConfigsStub != nil {
		return fake.CleanUnreferencedConfigsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cleanUnreferencedConfigsReturns.result1
}

func (fake *FakeResourceConfigFactory) CleanUnreferencedConfigsCallCount() int {
	fake.cleanUnreferencedConfigsMutex.RLock()
	defer fake.cleanUnreferencedConfigsMutex.RUnlock()
	return len(fake.cleanUnreferencedConfigsArgsForCall)
}

func (fake *FakeResourceConfigFactory) CleanUnreferencedConfigsReturns(result1 error) {
	fake.CleanUnreferencedConfigsStub = nil
	fake.cleanUnreferencedConfigsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) CleanUnreferencedConfigsReturnsOnCall(i int, result1 error) {
	fake.CleanUnreferencedConfigsStub = nil
	if fake.cleanUnreferencedConfigsReturnsOnCall == nil {
		fake.cleanUnreferencedConfigsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cleanUnreferencedConfigsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfigFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateResourceConfigMutex.RLock()
	defer fake.findOrCreateResourceConfigMutex.RUnlock()
	fake.findResourceConfigByIDMutex.RLock()
	defer fake.findResourceConfigByIDMutex.RUnlock()
	fake.cleanUnreferencedConfigsMutex.RLock()
	defer fake.cleanUnreferencedConfigsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceConfigFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceConfigFactory = new(FakeResourceConfigFactory)
