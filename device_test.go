package main

//import (
//	"context"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"github.com/stretchr/testify/require"
//)
//
//var deviceFile = "/dev/ttyUSB0"
//
//func MockedProcessor(s string) bool {
//	return false
//}
//
//func TestOpen(t *testing.T) {
//
//	d, err := OpenDevice(deviceFile)
//	if err == nil {
//		defer d.Close()
//	}
//
//	require.NoError(t, err)
//}
//
//func TestOpen_DeviceNotExist(t *testing.T) {
//	_, err := OpenDevice(deviceFile)
//
//	require.Error(t, err)
//}
//
//func TestRead_DeviceNotOpened(t *testing.T) {
//	d := Device{}
//	err := d.Process(context.Background(), MockedProcessor)
//
//	require.Error(t, err)
//}
//
//func TestRead_DeviceClosed(t *testing.T) {
//
//	d, _ := OpenDevice(deviceFile)
//	d.Close()
//
//	err := d.Process(context.Background(), MockedProcessor)
//
//	require.Error(t, err)
//}
//
//func TestRead(t *testing.T) {
//	lines := []string{
//		"some line to read process",
//		"some other line to read and process",
//	}
//
//	m := mock.Mock{}
//	m.On("func1", lines[0]).Once()
//	m.On("func1", lines[1]).Once()
//
//	d, _ := OpenDevice(deviceFile)
//	counter := 0
//	err := d.Process(context.Background(), func(s string) bool {
//		counter++
//		m.Called(s)
//		return counter == 2
//	})
//	assert.NoError(t, err)
//	m.AssertNumberOfCalls(t, "func1", 2)
//
//}
//
//func TestRead_stopProcessing(t *testing.T) {
//	l1 := "some line to read and process"
//	l2 := "some other line to read and process"
//	l3 := "this line shouldn't be processed"
//
//
//	m := mock.Mock{}
//	m.On("func1", l1).Once()
//	m.On("func1", l2).Once()
//
//	d, _ := OpenDevice(deviceFile)
//	counter := 0
//	d.Process(context.Background(), func(s string) bool {
//		counter++
//		m.Called(s)
//		return counter == 2
//	})
//
//	m.AssertExpectations(t)
//	m.AssertNotCalled(t, "func1", l3)
//}
