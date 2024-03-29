// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package gorest

import (
	"sync"
	"time"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
// 	func TestSomethingThatUsesRepository(t *testing.T) {
//
// 		// make and configure a mocked Repository
// 		mockedRepository := &RepositoryMock{
// 			DeleteFunc: func(s string) error {
// 				panic("mock out the Delete method")
// 			},
// 			FindFunc: func(s string) ([]byte, time.Time, error) {
// 				panic("mock out the Find method")
// 			},
// 			UpdateFunc: func(s string, bytes []byte) error {
// 				panic("mock out the Update method")
// 			},
// 		}
//
// 		// use mockedRepository in code that requires Repository
// 		// and then make assertions.
//
// 	}
type RepositoryMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(s string) error

	// FindFunc mocks the Find method.
	FindFunc func(s string) ([]byte, time.Time, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(s string, bytes []byte) error

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// S is the s argument value.
			S string
		}
		// Find holds details about calls to the Find method.
		Find []struct {
			// S is the s argument value.
			S string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// S is the s argument value.
			S string
			// Bytes is the bytes argument value.
			Bytes []byte
		}
	}
	lockDelete sync.RWMutex
	lockFind   sync.RWMutex
	lockUpdate sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *RepositoryMock) Delete(s string) error {
	if mock.DeleteFunc == nil {
		panic("RepositoryMock.DeleteFunc: method is nil but Repository.Delete was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(s)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedRepository.DeleteCalls())
func (mock *RepositoryMock) DeleteCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Find calls FindFunc.
func (mock *RepositoryMock) Find(s string) ([]byte, time.Time, error) {
	if mock.FindFunc == nil {
		panic("RepositoryMock.FindFunc: method is nil but Repository.Find was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockFind.Lock()
	mock.calls.Find = append(mock.calls.Find, callInfo)
	mock.lockFind.Unlock()
	return mock.FindFunc(s)
}

// FindCalls gets all the calls that were made to Find.
// Check the length with:
//     len(mockedRepository.FindCalls())
func (mock *RepositoryMock) FindCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockFind.RLock()
	calls = mock.calls.Find
	mock.lockFind.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *RepositoryMock) Update(s string, bytes []byte) error {
	if mock.UpdateFunc == nil {
		panic("RepositoryMock.UpdateFunc: method is nil but Repository.Update was just called")
	}
	callInfo := struct {
		S     string
		Bytes []byte
	}{
		S:     s,
		Bytes: bytes,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(s, bytes)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedRepository.UpdateCalls())
func (mock *RepositoryMock) UpdateCalls() []struct {
	S     string
	Bytes []byte
} {
	var calls []struct {
		S     string
		Bytes []byte
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
