// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/config"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/algorithm"
	"github.com/concourse/atc/scheduler"
)

type FakeJobServiceDB struct {
	GetJobStub        func(job string) (db.SavedJob, error)
	getJobMutex       sync.RWMutex
	getJobArgsForCall []struct {
		job string
	}
	getJobReturns struct {
		result1 db.SavedJob
		result2 error
	}
	GetRunningBuildsBySerialGroupStub        func(jobName string, serialGroups []string) ([]db.BuildDB, error)
	getRunningBuildsBySerialGroupMutex       sync.RWMutex
	getRunningBuildsBySerialGroupArgsForCall []struct {
		jobName      string
		serialGroups []string
	}
	getRunningBuildsBySerialGroupReturns struct {
		result1 []db.BuildDB
		result2 error
	}
	GetNextPendingBuildBySerialGroupStub        func(jobName string, serialGroups []string) (db.Build, bool, error)
	getNextPendingBuildBySerialGroupMutex       sync.RWMutex
	getNextPendingBuildBySerialGroupArgsForCall []struct {
		jobName      string
		serialGroups []string
	}
	getNextPendingBuildBySerialGroupReturns struct {
		result1 db.Build
		result2 bool
		result3 error
	}
	UpdateBuildPreparationStub        func(prep db.BuildPreparation) error
	updateBuildPreparationMutex       sync.RWMutex
	updateBuildPreparationArgsForCall []struct {
		prep db.BuildPreparation
	}
	updateBuildPreparationReturns struct {
		result1 error
	}
	IsPausedStub        func() (bool, error)
	isPausedMutex       sync.RWMutex
	isPausedArgsForCall []struct{}
	isPausedReturns     struct {
		result1 bool
		result2 error
	}
	LoadVersionsDBStub        func() (*algorithm.VersionsDB, error)
	loadVersionsDBMutex       sync.RWMutex
	loadVersionsDBArgsForCall []struct{}
	loadVersionsDBReturns     struct {
		result1 *algorithm.VersionsDB
		result2 error
	}
	GetNextInputVersionsStub        func(versions *algorithm.VersionsDB, job string, inputs []config.JobInput) ([]db.BuildInput, bool, db.MissingInputReasons, error)
	getNextInputVersionsMutex       sync.RWMutex
	getNextInputVersionsArgsForCall []struct {
		versions *algorithm.VersionsDB
		job      string
		inputs   []config.JobInput
	}
	getNextInputVersionsReturns struct {
		result1 []db.BuildInput
		result2 bool
		result3 db.MissingInputReasons
		result4 error
	}
	UseInputsForBuildStub        func(buildID int, inputs []db.BuildInput) error
	useInputsForBuildMutex       sync.RWMutex
	useInputsForBuildArgsForCall []struct {
		buildID int
		inputs  []db.BuildInput
	}
	useInputsForBuildReturns struct {
		result1 error
	}
}

func (fake *FakeJobServiceDB) GetJob(job string) (db.SavedJob, error) {
	fake.getJobMutex.Lock()
	fake.getJobArgsForCall = append(fake.getJobArgsForCall, struct {
		job string
	}{job})
	fake.getJobMutex.Unlock()
	if fake.GetJobStub != nil {
		return fake.GetJobStub(job)
	} else {
		return fake.getJobReturns.result1, fake.getJobReturns.result2
	}
}

func (fake *FakeJobServiceDB) GetJobCallCount() int {
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	return len(fake.getJobArgsForCall)
}

func (fake *FakeJobServiceDB) GetJobArgsForCall(i int) string {
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	return fake.getJobArgsForCall[i].job
}

func (fake *FakeJobServiceDB) GetJobReturns(result1 db.SavedJob, result2 error) {
	fake.GetJobStub = nil
	fake.getJobReturns = struct {
		result1 db.SavedJob
		result2 error
	}{result1, result2}
}

func (fake *FakeJobServiceDB) GetRunningBuildsBySerialGroup(jobName string, serialGroups []string) ([]db.BuildDB, error) {
	fake.getRunningBuildsBySerialGroupMutex.Lock()
	fake.getRunningBuildsBySerialGroupArgsForCall = append(fake.getRunningBuildsBySerialGroupArgsForCall, struct {
		jobName      string
		serialGroups []string
	}{jobName, serialGroups})
	fake.getRunningBuildsBySerialGroupMutex.Unlock()
	if fake.GetRunningBuildsBySerialGroupStub != nil {
		return fake.GetRunningBuildsBySerialGroupStub(jobName, serialGroups)
	} else {
		return fake.getRunningBuildsBySerialGroupReturns.result1, fake.getRunningBuildsBySerialGroupReturns.result2
	}
}

func (fake *FakeJobServiceDB) GetRunningBuildsBySerialGroupCallCount() int {
	fake.getRunningBuildsBySerialGroupMutex.RLock()
	defer fake.getRunningBuildsBySerialGroupMutex.RUnlock()
	return len(fake.getRunningBuildsBySerialGroupArgsForCall)
}

func (fake *FakeJobServiceDB) GetRunningBuildsBySerialGroupArgsForCall(i int) (string, []string) {
	fake.getRunningBuildsBySerialGroupMutex.RLock()
	defer fake.getRunningBuildsBySerialGroupMutex.RUnlock()
	return fake.getRunningBuildsBySerialGroupArgsForCall[i].jobName, fake.getRunningBuildsBySerialGroupArgsForCall[i].serialGroups
}

func (fake *FakeJobServiceDB) GetRunningBuildsBySerialGroupReturns(result1 []db.BuildDB, result2 error) {
	fake.GetRunningBuildsBySerialGroupStub = nil
	fake.getRunningBuildsBySerialGroupReturns = struct {
		result1 []db.BuildDB
		result2 error
	}{result1, result2}
}

func (fake *FakeJobServiceDB) GetNextPendingBuildBySerialGroup(jobName string, serialGroups []string) (db.Build, bool, error) {
	fake.getNextPendingBuildBySerialGroupMutex.Lock()
	fake.getNextPendingBuildBySerialGroupArgsForCall = append(fake.getNextPendingBuildBySerialGroupArgsForCall, struct {
		jobName      string
		serialGroups []string
	}{jobName, serialGroups})
	fake.getNextPendingBuildBySerialGroupMutex.Unlock()
	if fake.GetNextPendingBuildBySerialGroupStub != nil {
		return fake.GetNextPendingBuildBySerialGroupStub(jobName, serialGroups)
	} else {
		return fake.getNextPendingBuildBySerialGroupReturns.result1, fake.getNextPendingBuildBySerialGroupReturns.result2, fake.getNextPendingBuildBySerialGroupReturns.result3
	}
}

func (fake *FakeJobServiceDB) GetNextPendingBuildBySerialGroupCallCount() int {
	fake.getNextPendingBuildBySerialGroupMutex.RLock()
	defer fake.getNextPendingBuildBySerialGroupMutex.RUnlock()
	return len(fake.getNextPendingBuildBySerialGroupArgsForCall)
}

func (fake *FakeJobServiceDB) GetNextPendingBuildBySerialGroupArgsForCall(i int) (string, []string) {
	fake.getNextPendingBuildBySerialGroupMutex.RLock()
	defer fake.getNextPendingBuildBySerialGroupMutex.RUnlock()
	return fake.getNextPendingBuildBySerialGroupArgsForCall[i].jobName, fake.getNextPendingBuildBySerialGroupArgsForCall[i].serialGroups
}

func (fake *FakeJobServiceDB) GetNextPendingBuildBySerialGroupReturns(result1 db.Build, result2 bool, result3 error) {
	fake.GetNextPendingBuildBySerialGroupStub = nil
	fake.getNextPendingBuildBySerialGroupReturns = struct {
		result1 db.Build
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeJobServiceDB) UpdateBuildPreparation(prep db.BuildPreparation) error {
	fake.updateBuildPreparationMutex.Lock()
	fake.updateBuildPreparationArgsForCall = append(fake.updateBuildPreparationArgsForCall, struct {
		prep db.BuildPreparation
	}{prep})
	fake.updateBuildPreparationMutex.Unlock()
	if fake.UpdateBuildPreparationStub != nil {
		return fake.UpdateBuildPreparationStub(prep)
	} else {
		return fake.updateBuildPreparationReturns.result1
	}
}

func (fake *FakeJobServiceDB) UpdateBuildPreparationCallCount() int {
	fake.updateBuildPreparationMutex.RLock()
	defer fake.updateBuildPreparationMutex.RUnlock()
	return len(fake.updateBuildPreparationArgsForCall)
}

func (fake *FakeJobServiceDB) UpdateBuildPreparationArgsForCall(i int) db.BuildPreparation {
	fake.updateBuildPreparationMutex.RLock()
	defer fake.updateBuildPreparationMutex.RUnlock()
	return fake.updateBuildPreparationArgsForCall[i].prep
}

func (fake *FakeJobServiceDB) UpdateBuildPreparationReturns(result1 error) {
	fake.UpdateBuildPreparationStub = nil
	fake.updateBuildPreparationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeJobServiceDB) IsPaused() (bool, error) {
	fake.isPausedMutex.Lock()
	fake.isPausedArgsForCall = append(fake.isPausedArgsForCall, struct{}{})
	fake.isPausedMutex.Unlock()
	if fake.IsPausedStub != nil {
		return fake.IsPausedStub()
	} else {
		return fake.isPausedReturns.result1, fake.isPausedReturns.result2
	}
}

func (fake *FakeJobServiceDB) IsPausedCallCount() int {
	fake.isPausedMutex.RLock()
	defer fake.isPausedMutex.RUnlock()
	return len(fake.isPausedArgsForCall)
}

func (fake *FakeJobServiceDB) IsPausedReturns(result1 bool, result2 error) {
	fake.IsPausedStub = nil
	fake.isPausedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeJobServiceDB) LoadVersionsDB() (*algorithm.VersionsDB, error) {
	fake.loadVersionsDBMutex.Lock()
	fake.loadVersionsDBArgsForCall = append(fake.loadVersionsDBArgsForCall, struct{}{})
	fake.loadVersionsDBMutex.Unlock()
	if fake.LoadVersionsDBStub != nil {
		return fake.LoadVersionsDBStub()
	} else {
		return fake.loadVersionsDBReturns.result1, fake.loadVersionsDBReturns.result2
	}
}

func (fake *FakeJobServiceDB) LoadVersionsDBCallCount() int {
	fake.loadVersionsDBMutex.RLock()
	defer fake.loadVersionsDBMutex.RUnlock()
	return len(fake.loadVersionsDBArgsForCall)
}

func (fake *FakeJobServiceDB) LoadVersionsDBReturns(result1 *algorithm.VersionsDB, result2 error) {
	fake.LoadVersionsDBStub = nil
	fake.loadVersionsDBReturns = struct {
		result1 *algorithm.VersionsDB
		result2 error
	}{result1, result2}
}

func (fake *FakeJobServiceDB) GetNextInputVersions(versions *algorithm.VersionsDB, job string, inputs []config.JobInput) ([]db.BuildInput, bool, db.MissingInputReasons, error) {
	fake.getNextInputVersionsMutex.Lock()
	fake.getNextInputVersionsArgsForCall = append(fake.getNextInputVersionsArgsForCall, struct {
		versions *algorithm.VersionsDB
		job      string
		inputs   []config.JobInput
	}{versions, job, inputs})
	fake.getNextInputVersionsMutex.Unlock()
	if fake.GetNextInputVersionsStub != nil {
		return fake.GetNextInputVersionsStub(versions, job, inputs)
	} else {
		return fake.getNextInputVersionsReturns.result1, fake.getNextInputVersionsReturns.result2, fake.getNextInputVersionsReturns.result3, fake.getNextInputVersionsReturns.result4
	}
}

func (fake *FakeJobServiceDB) GetNextInputVersionsCallCount() int {
	fake.getNextInputVersionsMutex.RLock()
	defer fake.getNextInputVersionsMutex.RUnlock()
	return len(fake.getNextInputVersionsArgsForCall)
}

func (fake *FakeJobServiceDB) GetNextInputVersionsArgsForCall(i int) (*algorithm.VersionsDB, string, []config.JobInput) {
	fake.getNextInputVersionsMutex.RLock()
	defer fake.getNextInputVersionsMutex.RUnlock()
	return fake.getNextInputVersionsArgsForCall[i].versions, fake.getNextInputVersionsArgsForCall[i].job, fake.getNextInputVersionsArgsForCall[i].inputs
}

func (fake *FakeJobServiceDB) GetNextInputVersionsReturns(result1 []db.BuildInput, result2 bool, result3 db.MissingInputReasons, result4 error) {
	fake.GetNextInputVersionsStub = nil
	fake.getNextInputVersionsReturns = struct {
		result1 []db.BuildInput
		result2 bool
		result3 db.MissingInputReasons
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeJobServiceDB) UseInputsForBuild(buildID int, inputs []db.BuildInput) error {
	fake.useInputsForBuildMutex.Lock()
	fake.useInputsForBuildArgsForCall = append(fake.useInputsForBuildArgsForCall, struct {
		buildID int
		inputs  []db.BuildInput
	}{buildID, inputs})
	fake.useInputsForBuildMutex.Unlock()
	if fake.UseInputsForBuildStub != nil {
		return fake.UseInputsForBuildStub(buildID, inputs)
	} else {
		return fake.useInputsForBuildReturns.result1
	}
}

func (fake *FakeJobServiceDB) UseInputsForBuildCallCount() int {
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return len(fake.useInputsForBuildArgsForCall)
}

func (fake *FakeJobServiceDB) UseInputsForBuildArgsForCall(i int) (int, []db.BuildInput) {
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return fake.useInputsForBuildArgsForCall[i].buildID, fake.useInputsForBuildArgsForCall[i].inputs
}

func (fake *FakeJobServiceDB) UseInputsForBuildReturns(result1 error) {
	fake.UseInputsForBuildStub = nil
	fake.useInputsForBuildReturns = struct {
		result1 error
	}{result1}
}

var _ scheduler.JobServiceDB = new(FakeJobServiceDB)
