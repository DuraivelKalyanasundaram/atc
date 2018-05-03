// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakeVolumeFactory struct {
	GetTeamVolumesStub        func(teamID int) ([]db.CreatedVolume, error)
	getTeamVolumesMutex       sync.RWMutex
	getTeamVolumesArgsForCall []struct {
		teamID int
	}
	getTeamVolumesReturns struct {
		result1 []db.CreatedVolume
		result2 error
	}
	getTeamVolumesReturnsOnCall map[int]struct {
		result1 []db.CreatedVolume
		result2 error
	}
	CreateContainerVolumeStub        func(int, string, db.CreatingContainer, string) (db.CreatingVolume, error)
	createContainerVolumeMutex       sync.RWMutex
	createContainerVolumeArgsForCall []struct {
		arg1 int
		arg2 string
		arg3 db.CreatingContainer
		arg4 string
	}
	createContainerVolumeReturns struct {
		result1 db.CreatingVolume
		result2 error
	}
	createContainerVolumeReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 error
	}
	FindContainerVolumeStub        func(int, string, db.CreatingContainer, string) (db.CreatingVolume, db.CreatedVolume, error)
	findContainerVolumeMutex       sync.RWMutex
	findContainerVolumeArgsForCall []struct {
		arg1 int
		arg2 string
		arg3 db.CreatingContainer
		arg4 string
	}
	findContainerVolumeReturns struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}
	findContainerVolumeReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}
	FindBaseResourceTypeVolumeStub        func(int, *db.UsedWorkerBaseResourceType) (db.CreatingVolume, db.CreatedVolume, error)
	findBaseResourceTypeVolumeMutex       sync.RWMutex
	findBaseResourceTypeVolumeArgsForCall []struct {
		arg1 int
		arg2 *db.UsedWorkerBaseResourceType
	}
	findBaseResourceTypeVolumeReturns struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}
	findBaseResourceTypeVolumeReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}
	CreateBaseResourceTypeVolumeStub        func(int, *db.UsedWorkerBaseResourceType) (db.CreatingVolume, error)
	createBaseResourceTypeVolumeMutex       sync.RWMutex
	createBaseResourceTypeVolumeArgsForCall []struct {
		arg1 int
		arg2 *db.UsedWorkerBaseResourceType
	}
	createBaseResourceTypeVolumeReturns struct {
		result1 db.CreatingVolume
		result2 error
	}
	createBaseResourceTypeVolumeReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 error
	}
	FindResourceCacheVolumeStub        func(string, *db.UsedResourceCache) (db.CreatedVolume, bool, error)
	findResourceCacheVolumeMutex       sync.RWMutex
	findResourceCacheVolumeArgsForCall []struct {
		arg1 string
		arg2 *db.UsedResourceCache
	}
	findResourceCacheVolumeReturns struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}
	findResourceCacheVolumeReturnsOnCall map[int]struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}
	FindTaskCacheVolumeStub        func(teamID int, uwtc *db.UsedWorkerTaskCache) (db.CreatingVolume, db.CreatedVolume, error)
	findTaskCacheVolumeMutex       sync.RWMutex
	findTaskCacheVolumeArgsForCall []struct {
		teamID int
		uwtc   *db.UsedWorkerTaskCache
	}
	findTaskCacheVolumeReturns struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}
	findTaskCacheVolumeReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}
	CreateTaskCacheVolumeStub        func(teamID int, uwtc *db.UsedWorkerTaskCache) (db.CreatingVolume, error)
	createTaskCacheVolumeMutex       sync.RWMutex
	createTaskCacheVolumeArgsForCall []struct {
		teamID int
		uwtc   *db.UsedWorkerTaskCache
	}
	createTaskCacheVolumeReturns struct {
		result1 db.CreatingVolume
		result2 error
	}
	createTaskCacheVolumeReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 error
	}
	FindResourceCertsVolumeStub        func(workerName string, uwrc *db.UsedWorkerResourceCerts) (db.CreatingVolume, db.CreatedVolume, error)
	findResourceCertsVolumeMutex       sync.RWMutex
	findResourceCertsVolumeArgsForCall []struct {
		workerName string
		uwrc       *db.UsedWorkerResourceCerts
	}
	findResourceCertsVolumeReturns struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}
	findResourceCertsVolumeReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}
	CreateResourceCertsVolumeStub        func(workerName string, uwrc *db.UsedWorkerResourceCerts) (db.CreatingVolume, error)
	createResourceCertsVolumeMutex       sync.RWMutex
	createResourceCertsVolumeArgsForCall []struct {
		workerName string
		uwrc       *db.UsedWorkerResourceCerts
	}
	createResourceCertsVolumeReturns struct {
		result1 db.CreatingVolume
		result2 error
	}
	createResourceCertsVolumeReturnsOnCall map[int]struct {
		result1 db.CreatingVolume
		result2 error
	}
	FindVolumesForContainerStub        func(db.CreatedContainer) ([]db.CreatedVolume, error)
	findVolumesForContainerMutex       sync.RWMutex
	findVolumesForContainerArgsForCall []struct {
		arg1 db.CreatedContainer
	}
	findVolumesForContainerReturns struct {
		result1 []db.CreatedVolume
		result2 error
	}
	findVolumesForContainerReturnsOnCall map[int]struct {
		result1 []db.CreatedVolume
		result2 error
	}
	GetOrphanedVolumesStub        func() ([]db.CreatedVolume, []db.DestroyingVolume, error)
	getOrphanedVolumesMutex       sync.RWMutex
	getOrphanedVolumesArgsForCall []struct{}
	getOrphanedVolumesReturns     struct {
		result1 []db.CreatedVolume
		result2 []db.DestroyingVolume
		result3 error
	}
	getOrphanedVolumesReturnsOnCall map[int]struct {
		result1 []db.CreatedVolume
		result2 []db.DestroyingVolume
		result3 error
	}
	GetFailedVolumesStub        func() ([]db.FailedVolume, error)
	getFailedVolumesMutex       sync.RWMutex
	getFailedVolumesArgsForCall []struct{}
	getFailedVolumesReturns     struct {
		result1 []db.FailedVolume
		result2 error
	}
	getFailedVolumesReturnsOnCall map[int]struct {
		result1 []db.FailedVolume
		result2 error
	}
	GetDestroyingVolumesStub        func(workerName string) ([]string, error)
	getDestroyingVolumesMutex       sync.RWMutex
	getDestroyingVolumesArgsForCall []struct {
		workerName string
	}
	getDestroyingVolumesReturns struct {
		result1 []string
		result2 error
	}
	getDestroyingVolumesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	FindCreatedVolumeStub        func(handle string) (db.CreatedVolume, bool, error)
	findCreatedVolumeMutex       sync.RWMutex
	findCreatedVolumeArgsForCall []struct {
		handle string
	}
	findCreatedVolumeReturns struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}
	findCreatedVolumeReturnsOnCall map[int]struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeFactory) GetTeamVolumes(teamID int) ([]db.CreatedVolume, error) {
	fake.getTeamVolumesMutex.Lock()
	ret, specificReturn := fake.getTeamVolumesReturnsOnCall[len(fake.getTeamVolumesArgsForCall)]
	fake.getTeamVolumesArgsForCall = append(fake.getTeamVolumesArgsForCall, struct {
		teamID int
	}{teamID})
	fake.recordInvocation("GetTeamVolumes", []interface{}{teamID})
	fake.getTeamVolumesMutex.Unlock()
	if fake.GetTeamVolumesStub != nil {
		return fake.GetTeamVolumesStub(teamID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTeamVolumesReturns.result1, fake.getTeamVolumesReturns.result2
}

func (fake *FakeVolumeFactory) GetTeamVolumesCallCount() int {
	fake.getTeamVolumesMutex.RLock()
	defer fake.getTeamVolumesMutex.RUnlock()
	return len(fake.getTeamVolumesArgsForCall)
}

func (fake *FakeVolumeFactory) GetTeamVolumesArgsForCall(i int) int {
	fake.getTeamVolumesMutex.RLock()
	defer fake.getTeamVolumesMutex.RUnlock()
	return fake.getTeamVolumesArgsForCall[i].teamID
}

func (fake *FakeVolumeFactory) GetTeamVolumesReturns(result1 []db.CreatedVolume, result2 error) {
	fake.GetTeamVolumesStub = nil
	fake.getTeamVolumesReturns = struct {
		result1 []db.CreatedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) GetTeamVolumesReturnsOnCall(i int, result1 []db.CreatedVolume, result2 error) {
	fake.GetTeamVolumesStub = nil
	if fake.getTeamVolumesReturnsOnCall == nil {
		fake.getTeamVolumesReturnsOnCall = make(map[int]struct {
			result1 []db.CreatedVolume
			result2 error
		})
	}
	fake.getTeamVolumesReturnsOnCall[i] = struct {
		result1 []db.CreatedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) CreateContainerVolume(arg1 int, arg2 string, arg3 db.CreatingContainer, arg4 string) (db.CreatingVolume, error) {
	fake.createContainerVolumeMutex.Lock()
	ret, specificReturn := fake.createContainerVolumeReturnsOnCall[len(fake.createContainerVolumeArgsForCall)]
	fake.createContainerVolumeArgsForCall = append(fake.createContainerVolumeArgsForCall, struct {
		arg1 int
		arg2 string
		arg3 db.CreatingContainer
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateContainerVolume", []interface{}{arg1, arg2, arg3, arg4})
	fake.createContainerVolumeMutex.Unlock()
	if fake.CreateContainerVolumeStub != nil {
		return fake.CreateContainerVolumeStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createContainerVolumeReturns.result1, fake.createContainerVolumeReturns.result2
}

func (fake *FakeVolumeFactory) CreateContainerVolumeCallCount() int {
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	return len(fake.createContainerVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) CreateContainerVolumeArgsForCall(i int) (int, string, db.CreatingContainer, string) {
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	return fake.createContainerVolumeArgsForCall[i].arg1, fake.createContainerVolumeArgsForCall[i].arg2, fake.createContainerVolumeArgsForCall[i].arg3, fake.createContainerVolumeArgsForCall[i].arg4
}

func (fake *FakeVolumeFactory) CreateContainerVolumeReturns(result1 db.CreatingVolume, result2 error) {
	fake.CreateContainerVolumeStub = nil
	fake.createContainerVolumeReturns = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) CreateContainerVolumeReturnsOnCall(i int, result1 db.CreatingVolume, result2 error) {
	fake.CreateContainerVolumeStub = nil
	if fake.createContainerVolumeReturnsOnCall == nil {
		fake.createContainerVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 error
		})
	}
	fake.createContainerVolumeReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindContainerVolume(arg1 int, arg2 string, arg3 db.CreatingContainer, arg4 string) (db.CreatingVolume, db.CreatedVolume, error) {
	fake.findContainerVolumeMutex.Lock()
	ret, specificReturn := fake.findContainerVolumeReturnsOnCall[len(fake.findContainerVolumeArgsForCall)]
	fake.findContainerVolumeArgsForCall = append(fake.findContainerVolumeArgsForCall, struct {
		arg1 int
		arg2 string
		arg3 db.CreatingContainer
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("FindContainerVolume", []interface{}{arg1, arg2, arg3, arg4})
	fake.findContainerVolumeMutex.Unlock()
	if fake.FindContainerVolumeStub != nil {
		return fake.FindContainerVolumeStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findContainerVolumeReturns.result1, fake.findContainerVolumeReturns.result2, fake.findContainerVolumeReturns.result3
}

func (fake *FakeVolumeFactory) FindContainerVolumeCallCount() int {
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	return len(fake.findContainerVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindContainerVolumeArgsForCall(i int) (int, string, db.CreatingContainer, string) {
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	return fake.findContainerVolumeArgsForCall[i].arg1, fake.findContainerVolumeArgsForCall[i].arg2, fake.findContainerVolumeArgsForCall[i].arg3, fake.findContainerVolumeArgsForCall[i].arg4
}

func (fake *FakeVolumeFactory) FindContainerVolumeReturns(result1 db.CreatingVolume, result2 db.CreatedVolume, result3 error) {
	fake.FindContainerVolumeStub = nil
	fake.findContainerVolumeReturns = struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindContainerVolumeReturnsOnCall(i int, result1 db.CreatingVolume, result2 db.CreatedVolume, result3 error) {
	fake.FindContainerVolumeStub = nil
	if fake.findContainerVolumeReturnsOnCall == nil {
		fake.findContainerVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 db.CreatedVolume
			result3 error
		})
	}
	fake.findContainerVolumeReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolume(arg1 int, arg2 *db.UsedWorkerBaseResourceType) (db.CreatingVolume, db.CreatedVolume, error) {
	fake.findBaseResourceTypeVolumeMutex.Lock()
	ret, specificReturn := fake.findBaseResourceTypeVolumeReturnsOnCall[len(fake.findBaseResourceTypeVolumeArgsForCall)]
	fake.findBaseResourceTypeVolumeArgsForCall = append(fake.findBaseResourceTypeVolumeArgsForCall, struct {
		arg1 int
		arg2 *db.UsedWorkerBaseResourceType
	}{arg1, arg2})
	fake.recordInvocation("FindBaseResourceTypeVolume", []interface{}{arg1, arg2})
	fake.findBaseResourceTypeVolumeMutex.Unlock()
	if fake.FindBaseResourceTypeVolumeStub != nil {
		return fake.FindBaseResourceTypeVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findBaseResourceTypeVolumeReturns.result1, fake.findBaseResourceTypeVolumeReturns.result2, fake.findBaseResourceTypeVolumeReturns.result3
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolumeCallCount() int {
	fake.findBaseResourceTypeVolumeMutex.RLock()
	defer fake.findBaseResourceTypeVolumeMutex.RUnlock()
	return len(fake.findBaseResourceTypeVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolumeArgsForCall(i int) (int, *db.UsedWorkerBaseResourceType) {
	fake.findBaseResourceTypeVolumeMutex.RLock()
	defer fake.findBaseResourceTypeVolumeMutex.RUnlock()
	return fake.findBaseResourceTypeVolumeArgsForCall[i].arg1, fake.findBaseResourceTypeVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolumeReturns(result1 db.CreatingVolume, result2 db.CreatedVolume, result3 error) {
	fake.FindBaseResourceTypeVolumeStub = nil
	fake.findBaseResourceTypeVolumeReturns = struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolumeReturnsOnCall(i int, result1 db.CreatingVolume, result2 db.CreatedVolume, result3 error) {
	fake.FindBaseResourceTypeVolumeStub = nil
	if fake.findBaseResourceTypeVolumeReturnsOnCall == nil {
		fake.findBaseResourceTypeVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 db.CreatedVolume
			result3 error
		})
	}
	fake.findBaseResourceTypeVolumeReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolume(arg1 int, arg2 *db.UsedWorkerBaseResourceType) (db.CreatingVolume, error) {
	fake.createBaseResourceTypeVolumeMutex.Lock()
	ret, specificReturn := fake.createBaseResourceTypeVolumeReturnsOnCall[len(fake.createBaseResourceTypeVolumeArgsForCall)]
	fake.createBaseResourceTypeVolumeArgsForCall = append(fake.createBaseResourceTypeVolumeArgsForCall, struct {
		arg1 int
		arg2 *db.UsedWorkerBaseResourceType
	}{arg1, arg2})
	fake.recordInvocation("CreateBaseResourceTypeVolume", []interface{}{arg1, arg2})
	fake.createBaseResourceTypeVolumeMutex.Unlock()
	if fake.CreateBaseResourceTypeVolumeStub != nil {
		return fake.CreateBaseResourceTypeVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createBaseResourceTypeVolumeReturns.result1, fake.createBaseResourceTypeVolumeReturns.result2
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolumeCallCount() int {
	fake.createBaseResourceTypeVolumeMutex.RLock()
	defer fake.createBaseResourceTypeVolumeMutex.RUnlock()
	return len(fake.createBaseResourceTypeVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolumeArgsForCall(i int) (int, *db.UsedWorkerBaseResourceType) {
	fake.createBaseResourceTypeVolumeMutex.RLock()
	defer fake.createBaseResourceTypeVolumeMutex.RUnlock()
	return fake.createBaseResourceTypeVolumeArgsForCall[i].arg1, fake.createBaseResourceTypeVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolumeReturns(result1 db.CreatingVolume, result2 error) {
	fake.CreateBaseResourceTypeVolumeStub = nil
	fake.createBaseResourceTypeVolumeReturns = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolumeReturnsOnCall(i int, result1 db.CreatingVolume, result2 error) {
	fake.CreateBaseResourceTypeVolumeStub = nil
	if fake.createBaseResourceTypeVolumeReturnsOnCall == nil {
		fake.createBaseResourceTypeVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 error
		})
	}
	fake.createBaseResourceTypeVolumeReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindResourceCacheVolume(arg1 string, arg2 *db.UsedResourceCache) (db.CreatedVolume, bool, error) {
	fake.findResourceCacheVolumeMutex.Lock()
	ret, specificReturn := fake.findResourceCacheVolumeReturnsOnCall[len(fake.findResourceCacheVolumeArgsForCall)]
	fake.findResourceCacheVolumeArgsForCall = append(fake.findResourceCacheVolumeArgsForCall, struct {
		arg1 string
		arg2 *db.UsedResourceCache
	}{arg1, arg2})
	fake.recordInvocation("FindResourceCacheVolume", []interface{}{arg1, arg2})
	fake.findResourceCacheVolumeMutex.Unlock()
	if fake.FindResourceCacheVolumeStub != nil {
		return fake.FindResourceCacheVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findResourceCacheVolumeReturns.result1, fake.findResourceCacheVolumeReturns.result2, fake.findResourceCacheVolumeReturns.result3
}

func (fake *FakeVolumeFactory) FindResourceCacheVolumeCallCount() int {
	fake.findResourceCacheVolumeMutex.RLock()
	defer fake.findResourceCacheVolumeMutex.RUnlock()
	return len(fake.findResourceCacheVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindResourceCacheVolumeArgsForCall(i int) (string, *db.UsedResourceCache) {
	fake.findResourceCacheVolumeMutex.RLock()
	defer fake.findResourceCacheVolumeMutex.RUnlock()
	return fake.findResourceCacheVolumeArgsForCall[i].arg1, fake.findResourceCacheVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeFactory) FindResourceCacheVolumeReturns(result1 db.CreatedVolume, result2 bool, result3 error) {
	fake.FindResourceCacheVolumeStub = nil
	fake.findResourceCacheVolumeReturns = struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindResourceCacheVolumeReturnsOnCall(i int, result1 db.CreatedVolume, result2 bool, result3 error) {
	fake.FindResourceCacheVolumeStub = nil
	if fake.findResourceCacheVolumeReturnsOnCall == nil {
		fake.findResourceCacheVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatedVolume
			result2 bool
			result3 error
		})
	}
	fake.findResourceCacheVolumeReturnsOnCall[i] = struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindTaskCacheVolume(teamID int, uwtc *db.UsedWorkerTaskCache) (db.CreatingVolume, db.CreatedVolume, error) {
	fake.findTaskCacheVolumeMutex.Lock()
	ret, specificReturn := fake.findTaskCacheVolumeReturnsOnCall[len(fake.findTaskCacheVolumeArgsForCall)]
	fake.findTaskCacheVolumeArgsForCall = append(fake.findTaskCacheVolumeArgsForCall, struct {
		teamID int
		uwtc   *db.UsedWorkerTaskCache
	}{teamID, uwtc})
	fake.recordInvocation("FindTaskCacheVolume", []interface{}{teamID, uwtc})
	fake.findTaskCacheVolumeMutex.Unlock()
	if fake.FindTaskCacheVolumeStub != nil {
		return fake.FindTaskCacheVolumeStub(teamID, uwtc)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findTaskCacheVolumeReturns.result1, fake.findTaskCacheVolumeReturns.result2, fake.findTaskCacheVolumeReturns.result3
}

func (fake *FakeVolumeFactory) FindTaskCacheVolumeCallCount() int {
	fake.findTaskCacheVolumeMutex.RLock()
	defer fake.findTaskCacheVolumeMutex.RUnlock()
	return len(fake.findTaskCacheVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindTaskCacheVolumeArgsForCall(i int) (int, *db.UsedWorkerTaskCache) {
	fake.findTaskCacheVolumeMutex.RLock()
	defer fake.findTaskCacheVolumeMutex.RUnlock()
	return fake.findTaskCacheVolumeArgsForCall[i].teamID, fake.findTaskCacheVolumeArgsForCall[i].uwtc
}

func (fake *FakeVolumeFactory) FindTaskCacheVolumeReturns(result1 db.CreatingVolume, result2 db.CreatedVolume, result3 error) {
	fake.FindTaskCacheVolumeStub = nil
	fake.findTaskCacheVolumeReturns = struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindTaskCacheVolumeReturnsOnCall(i int, result1 db.CreatingVolume, result2 db.CreatedVolume, result3 error) {
	fake.FindTaskCacheVolumeStub = nil
	if fake.findTaskCacheVolumeReturnsOnCall == nil {
		fake.findTaskCacheVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 db.CreatedVolume
			result3 error
		})
	}
	fake.findTaskCacheVolumeReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) CreateTaskCacheVolume(teamID int, uwtc *db.UsedWorkerTaskCache) (db.CreatingVolume, error) {
	fake.createTaskCacheVolumeMutex.Lock()
	ret, specificReturn := fake.createTaskCacheVolumeReturnsOnCall[len(fake.createTaskCacheVolumeArgsForCall)]
	fake.createTaskCacheVolumeArgsForCall = append(fake.createTaskCacheVolumeArgsForCall, struct {
		teamID int
		uwtc   *db.UsedWorkerTaskCache
	}{teamID, uwtc})
	fake.recordInvocation("CreateTaskCacheVolume", []interface{}{teamID, uwtc})
	fake.createTaskCacheVolumeMutex.Unlock()
	if fake.CreateTaskCacheVolumeStub != nil {
		return fake.CreateTaskCacheVolumeStub(teamID, uwtc)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createTaskCacheVolumeReturns.result1, fake.createTaskCacheVolumeReturns.result2
}

func (fake *FakeVolumeFactory) CreateTaskCacheVolumeCallCount() int {
	fake.createTaskCacheVolumeMutex.RLock()
	defer fake.createTaskCacheVolumeMutex.RUnlock()
	return len(fake.createTaskCacheVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) CreateTaskCacheVolumeArgsForCall(i int) (int, *db.UsedWorkerTaskCache) {
	fake.createTaskCacheVolumeMutex.RLock()
	defer fake.createTaskCacheVolumeMutex.RUnlock()
	return fake.createTaskCacheVolumeArgsForCall[i].teamID, fake.createTaskCacheVolumeArgsForCall[i].uwtc
}

func (fake *FakeVolumeFactory) CreateTaskCacheVolumeReturns(result1 db.CreatingVolume, result2 error) {
	fake.CreateTaskCacheVolumeStub = nil
	fake.createTaskCacheVolumeReturns = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) CreateTaskCacheVolumeReturnsOnCall(i int, result1 db.CreatingVolume, result2 error) {
	fake.CreateTaskCacheVolumeStub = nil
	if fake.createTaskCacheVolumeReturnsOnCall == nil {
		fake.createTaskCacheVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 error
		})
	}
	fake.createTaskCacheVolumeReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindResourceCertsVolume(workerName string, uwrc *db.UsedWorkerResourceCerts) (db.CreatingVolume, db.CreatedVolume, error) {
	fake.findResourceCertsVolumeMutex.Lock()
	ret, specificReturn := fake.findResourceCertsVolumeReturnsOnCall[len(fake.findResourceCertsVolumeArgsForCall)]
	fake.findResourceCertsVolumeArgsForCall = append(fake.findResourceCertsVolumeArgsForCall, struct {
		workerName string
		uwrc       *db.UsedWorkerResourceCerts
	}{workerName, uwrc})
	fake.recordInvocation("FindResourceCertsVolume", []interface{}{workerName, uwrc})
	fake.findResourceCertsVolumeMutex.Unlock()
	if fake.FindResourceCertsVolumeStub != nil {
		return fake.FindResourceCertsVolumeStub(workerName, uwrc)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findResourceCertsVolumeReturns.result1, fake.findResourceCertsVolumeReturns.result2, fake.findResourceCertsVolumeReturns.result3
}

func (fake *FakeVolumeFactory) FindResourceCertsVolumeCallCount() int {
	fake.findResourceCertsVolumeMutex.RLock()
	defer fake.findResourceCertsVolumeMutex.RUnlock()
	return len(fake.findResourceCertsVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindResourceCertsVolumeArgsForCall(i int) (string, *db.UsedWorkerResourceCerts) {
	fake.findResourceCertsVolumeMutex.RLock()
	defer fake.findResourceCertsVolumeMutex.RUnlock()
	return fake.findResourceCertsVolumeArgsForCall[i].workerName, fake.findResourceCertsVolumeArgsForCall[i].uwrc
}

func (fake *FakeVolumeFactory) FindResourceCertsVolumeReturns(result1 db.CreatingVolume, result2 db.CreatedVolume, result3 error) {
	fake.FindResourceCertsVolumeStub = nil
	fake.findResourceCertsVolumeReturns = struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindResourceCertsVolumeReturnsOnCall(i int, result1 db.CreatingVolume, result2 db.CreatedVolume, result3 error) {
	fake.FindResourceCertsVolumeStub = nil
	if fake.findResourceCertsVolumeReturnsOnCall == nil {
		fake.findResourceCertsVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 db.CreatedVolume
			result3 error
		})
	}
	fake.findResourceCertsVolumeReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 db.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) CreateResourceCertsVolume(workerName string, uwrc *db.UsedWorkerResourceCerts) (db.CreatingVolume, error) {
	fake.createResourceCertsVolumeMutex.Lock()
	ret, specificReturn := fake.createResourceCertsVolumeReturnsOnCall[len(fake.createResourceCertsVolumeArgsForCall)]
	fake.createResourceCertsVolumeArgsForCall = append(fake.createResourceCertsVolumeArgsForCall, struct {
		workerName string
		uwrc       *db.UsedWorkerResourceCerts
	}{workerName, uwrc})
	fake.recordInvocation("CreateResourceCertsVolume", []interface{}{workerName, uwrc})
	fake.createResourceCertsVolumeMutex.Unlock()
	if fake.CreateResourceCertsVolumeStub != nil {
		return fake.CreateResourceCertsVolumeStub(workerName, uwrc)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createResourceCertsVolumeReturns.result1, fake.createResourceCertsVolumeReturns.result2
}

func (fake *FakeVolumeFactory) CreateResourceCertsVolumeCallCount() int {
	fake.createResourceCertsVolumeMutex.RLock()
	defer fake.createResourceCertsVolumeMutex.RUnlock()
	return len(fake.createResourceCertsVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) CreateResourceCertsVolumeArgsForCall(i int) (string, *db.UsedWorkerResourceCerts) {
	fake.createResourceCertsVolumeMutex.RLock()
	defer fake.createResourceCertsVolumeMutex.RUnlock()
	return fake.createResourceCertsVolumeArgsForCall[i].workerName, fake.createResourceCertsVolumeArgsForCall[i].uwrc
}

func (fake *FakeVolumeFactory) CreateResourceCertsVolumeReturns(result1 db.CreatingVolume, result2 error) {
	fake.CreateResourceCertsVolumeStub = nil
	fake.createResourceCertsVolumeReturns = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) CreateResourceCertsVolumeReturnsOnCall(i int, result1 db.CreatingVolume, result2 error) {
	fake.CreateResourceCertsVolumeStub = nil
	if fake.createResourceCertsVolumeReturnsOnCall == nil {
		fake.createResourceCertsVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatingVolume
			result2 error
		})
	}
	fake.createResourceCertsVolumeReturnsOnCall[i] = struct {
		result1 db.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindVolumesForContainer(arg1 db.CreatedContainer) ([]db.CreatedVolume, error) {
	fake.findVolumesForContainerMutex.Lock()
	ret, specificReturn := fake.findVolumesForContainerReturnsOnCall[len(fake.findVolumesForContainerArgsForCall)]
	fake.findVolumesForContainerArgsForCall = append(fake.findVolumesForContainerArgsForCall, struct {
		arg1 db.CreatedContainer
	}{arg1})
	fake.recordInvocation("FindVolumesForContainer", []interface{}{arg1})
	fake.findVolumesForContainerMutex.Unlock()
	if fake.FindVolumesForContainerStub != nil {
		return fake.FindVolumesForContainerStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findVolumesForContainerReturns.result1, fake.findVolumesForContainerReturns.result2
}

func (fake *FakeVolumeFactory) FindVolumesForContainerCallCount() int {
	fake.findVolumesForContainerMutex.RLock()
	defer fake.findVolumesForContainerMutex.RUnlock()
	return len(fake.findVolumesForContainerArgsForCall)
}

func (fake *FakeVolumeFactory) FindVolumesForContainerArgsForCall(i int) db.CreatedContainer {
	fake.findVolumesForContainerMutex.RLock()
	defer fake.findVolumesForContainerMutex.RUnlock()
	return fake.findVolumesForContainerArgsForCall[i].arg1
}

func (fake *FakeVolumeFactory) FindVolumesForContainerReturns(result1 []db.CreatedVolume, result2 error) {
	fake.FindVolumesForContainerStub = nil
	fake.findVolumesForContainerReturns = struct {
		result1 []db.CreatedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindVolumesForContainerReturnsOnCall(i int, result1 []db.CreatedVolume, result2 error) {
	fake.FindVolumesForContainerStub = nil
	if fake.findVolumesForContainerReturnsOnCall == nil {
		fake.findVolumesForContainerReturnsOnCall = make(map[int]struct {
			result1 []db.CreatedVolume
			result2 error
		})
	}
	fake.findVolumesForContainerReturnsOnCall[i] = struct {
		result1 []db.CreatedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) GetOrphanedVolumes() ([]db.CreatedVolume, []db.DestroyingVolume, error) {
	fake.getOrphanedVolumesMutex.Lock()
	ret, specificReturn := fake.getOrphanedVolumesReturnsOnCall[len(fake.getOrphanedVolumesArgsForCall)]
	fake.getOrphanedVolumesArgsForCall = append(fake.getOrphanedVolumesArgsForCall, struct{}{})
	fake.recordInvocation("GetOrphanedVolumes", []interface{}{})
	fake.getOrphanedVolumesMutex.Unlock()
	if fake.GetOrphanedVolumesStub != nil {
		return fake.GetOrphanedVolumesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrphanedVolumesReturns.result1, fake.getOrphanedVolumesReturns.result2, fake.getOrphanedVolumesReturns.result3
}

func (fake *FakeVolumeFactory) GetOrphanedVolumesCallCount() int {
	fake.getOrphanedVolumesMutex.RLock()
	defer fake.getOrphanedVolumesMutex.RUnlock()
	return len(fake.getOrphanedVolumesArgsForCall)
}

func (fake *FakeVolumeFactory) GetOrphanedVolumesReturns(result1 []db.CreatedVolume, result2 []db.DestroyingVolume, result3 error) {
	fake.GetOrphanedVolumesStub = nil
	fake.getOrphanedVolumesReturns = struct {
		result1 []db.CreatedVolume
		result2 []db.DestroyingVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) GetOrphanedVolumesReturnsOnCall(i int, result1 []db.CreatedVolume, result2 []db.DestroyingVolume, result3 error) {
	fake.GetOrphanedVolumesStub = nil
	if fake.getOrphanedVolumesReturnsOnCall == nil {
		fake.getOrphanedVolumesReturnsOnCall = make(map[int]struct {
			result1 []db.CreatedVolume
			result2 []db.DestroyingVolume
			result3 error
		})
	}
	fake.getOrphanedVolumesReturnsOnCall[i] = struct {
		result1 []db.CreatedVolume
		result2 []db.DestroyingVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) GetFailedVolumes() ([]db.FailedVolume, error) {
	fake.getFailedVolumesMutex.Lock()
	ret, specificReturn := fake.getFailedVolumesReturnsOnCall[len(fake.getFailedVolumesArgsForCall)]
	fake.getFailedVolumesArgsForCall = append(fake.getFailedVolumesArgsForCall, struct{}{})
	fake.recordInvocation("GetFailedVolumes", []interface{}{})
	fake.getFailedVolumesMutex.Unlock()
	if fake.GetFailedVolumesStub != nil {
		return fake.GetFailedVolumesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getFailedVolumesReturns.result1, fake.getFailedVolumesReturns.result2
}

func (fake *FakeVolumeFactory) GetFailedVolumesCallCount() int {
	fake.getFailedVolumesMutex.RLock()
	defer fake.getFailedVolumesMutex.RUnlock()
	return len(fake.getFailedVolumesArgsForCall)
}

func (fake *FakeVolumeFactory) GetFailedVolumesReturns(result1 []db.FailedVolume, result2 error) {
	fake.GetFailedVolumesStub = nil
	fake.getFailedVolumesReturns = struct {
		result1 []db.FailedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) GetFailedVolumesReturnsOnCall(i int, result1 []db.FailedVolume, result2 error) {
	fake.GetFailedVolumesStub = nil
	if fake.getFailedVolumesReturnsOnCall == nil {
		fake.getFailedVolumesReturnsOnCall = make(map[int]struct {
			result1 []db.FailedVolume
			result2 error
		})
	}
	fake.getFailedVolumesReturnsOnCall[i] = struct {
		result1 []db.FailedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) GetDestroyingVolumes(workerName string) ([]string, error) {
	fake.getDestroyingVolumesMutex.Lock()
	ret, specificReturn := fake.getDestroyingVolumesReturnsOnCall[len(fake.getDestroyingVolumesArgsForCall)]
	fake.getDestroyingVolumesArgsForCall = append(fake.getDestroyingVolumesArgsForCall, struct {
		workerName string
	}{workerName})
	fake.recordInvocation("GetDestroyingVolumes", []interface{}{workerName})
	fake.getDestroyingVolumesMutex.Unlock()
	if fake.GetDestroyingVolumesStub != nil {
		return fake.GetDestroyingVolumesStub(workerName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDestroyingVolumesReturns.result1, fake.getDestroyingVolumesReturns.result2
}

func (fake *FakeVolumeFactory) GetDestroyingVolumesCallCount() int {
	fake.getDestroyingVolumesMutex.RLock()
	defer fake.getDestroyingVolumesMutex.RUnlock()
	return len(fake.getDestroyingVolumesArgsForCall)
}

func (fake *FakeVolumeFactory) GetDestroyingVolumesArgsForCall(i int) string {
	fake.getDestroyingVolumesMutex.RLock()
	defer fake.getDestroyingVolumesMutex.RUnlock()
	return fake.getDestroyingVolumesArgsForCall[i].workerName
}

func (fake *FakeVolumeFactory) GetDestroyingVolumesReturns(result1 []string, result2 error) {
	fake.GetDestroyingVolumesStub = nil
	fake.getDestroyingVolumesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) GetDestroyingVolumesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.GetDestroyingVolumesStub = nil
	if fake.getDestroyingVolumesReturnsOnCall == nil {
		fake.getDestroyingVolumesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getDestroyingVolumesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindCreatedVolume(handle string) (db.CreatedVolume, bool, error) {
	fake.findCreatedVolumeMutex.Lock()
	ret, specificReturn := fake.findCreatedVolumeReturnsOnCall[len(fake.findCreatedVolumeArgsForCall)]
	fake.findCreatedVolumeArgsForCall = append(fake.findCreatedVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("FindCreatedVolume", []interface{}{handle})
	fake.findCreatedVolumeMutex.Unlock()
	if fake.FindCreatedVolumeStub != nil {
		return fake.FindCreatedVolumeStub(handle)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findCreatedVolumeReturns.result1, fake.findCreatedVolumeReturns.result2, fake.findCreatedVolumeReturns.result3
}

func (fake *FakeVolumeFactory) FindCreatedVolumeCallCount() int {
	fake.findCreatedVolumeMutex.RLock()
	defer fake.findCreatedVolumeMutex.RUnlock()
	return len(fake.findCreatedVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindCreatedVolumeArgsForCall(i int) string {
	fake.findCreatedVolumeMutex.RLock()
	defer fake.findCreatedVolumeMutex.RUnlock()
	return fake.findCreatedVolumeArgsForCall[i].handle
}

func (fake *FakeVolumeFactory) FindCreatedVolumeReturns(result1 db.CreatedVolume, result2 bool, result3 error) {
	fake.FindCreatedVolumeStub = nil
	fake.findCreatedVolumeReturns = struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindCreatedVolumeReturnsOnCall(i int, result1 db.CreatedVolume, result2 bool, result3 error) {
	fake.FindCreatedVolumeStub = nil
	if fake.findCreatedVolumeReturnsOnCall == nil {
		fake.findCreatedVolumeReturnsOnCall = make(map[int]struct {
			result1 db.CreatedVolume
			result2 bool
			result3 error
		})
	}
	fake.findCreatedVolumeReturnsOnCall[i] = struct {
		result1 db.CreatedVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTeamVolumesMutex.RLock()
	defer fake.getTeamVolumesMutex.RUnlock()
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	fake.findBaseResourceTypeVolumeMutex.RLock()
	defer fake.findBaseResourceTypeVolumeMutex.RUnlock()
	fake.createBaseResourceTypeVolumeMutex.RLock()
	defer fake.createBaseResourceTypeVolumeMutex.RUnlock()
	fake.findResourceCacheVolumeMutex.RLock()
	defer fake.findResourceCacheVolumeMutex.RUnlock()
	fake.findTaskCacheVolumeMutex.RLock()
	defer fake.findTaskCacheVolumeMutex.RUnlock()
	fake.createTaskCacheVolumeMutex.RLock()
	defer fake.createTaskCacheVolumeMutex.RUnlock()
	fake.findResourceCertsVolumeMutex.RLock()
	defer fake.findResourceCertsVolumeMutex.RUnlock()
	fake.createResourceCertsVolumeMutex.RLock()
	defer fake.createResourceCertsVolumeMutex.RUnlock()
	fake.findVolumesForContainerMutex.RLock()
	defer fake.findVolumesForContainerMutex.RUnlock()
	fake.getOrphanedVolumesMutex.RLock()
	defer fake.getOrphanedVolumesMutex.RUnlock()
	fake.getFailedVolumesMutex.RLock()
	defer fake.getFailedVolumesMutex.RUnlock()
	fake.getDestroyingVolumesMutex.RLock()
	defer fake.getDestroyingVolumesMutex.RUnlock()
	fake.findCreatedVolumeMutex.RLock()
	defer fake.findCreatedVolumeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolumeFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.VolumeFactory = new(FakeVolumeFactory)
