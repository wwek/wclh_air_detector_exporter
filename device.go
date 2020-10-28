// 空气检测仪设备
package main

import (
	"bufio"
	"github.com/pkg/errors"
	"log"
	"os"

	"context"
	"sync"

	"github.com/spf13/afero"
	"github.com/tarm/serial"
)

var (
	// AppFs is an abstraction of the file system
	// to allow mocking in tests.
	AppFs = afero.NewOsFs()
)

const (

	// ReceivePrefixStar
	// *--------------------------------------- Data Output Format Definition ----------------------------------------*
	ReceivePrefixStar = "*"

	// ReceivePrefixField
	// TEMP HUMI CH_PM1.0 CH_PM2.5 CH_PM10 US_PM1.0 US_PM2.5 US_PM10 >0.3um >0.5um >1.0um >2.5um >5.0um >10um HCHO TVOC
	ReceivePrefixField = "TEMP"

	// ResetCmd is resetting the homeduino device to set it in an expected state
	ResetCmd = "RESET"

	// ResetCmdResponse is the response to ResetCmd homeduiono
	ResetCmdResponse = "ready"
)

// ProcessorFunc are implementations for consuming line by line of the device
type ProcessorFunc func(s string) bool

// Device represents the device file of an tty
// connected to the USB port.
type Device struct {
	afero.File
	sync.Mutex
	open bool
}

func SetupDevice(name string) {
	log.Println("Setting up serial port to use 9600 baud")
	c := &serial.Config{Name: name, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	s.Close()
	log.Println("Serial port prepared")
}

// OpenDevice opens the named device file for reading.
func OpenDevice(name string) (*Device, error) {
	file, err := AppFs.OpenFile(name, os.O_RDWR, 0644)

	if err != nil {
		return nil, errors.Wrapf(err, "Failed to open '%v'", name)
	}
	log.Printf("Device '%v' opened", file.Name())

	d := &Device{
		File:  file,
		open:  true,
		Mutex: sync.Mutex{},
	}

	return d, nil
}

// Process 循环的从设备文件中读取下一行，并且
// 对它应用ProcessorFunc上下文用于停止读取
// 在使用ReadProcess之前，需要通过“OpenDevice”打开设备文件。
func (d *Device) Process(ctx context.Context, handle ProcessorFunc) error {
	d.Lock()
	defer d.Unlock()
	if !d.open {
		return errors.New("File already closed")
	}

	scanner := bufio.NewScanner(d)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("Line Scanned:", line)

		stop := handle(line)
		if stop {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrapf(err, "Error while scanning device file '%s'", d.Name())
	}

	return nil
}

// Close closes the opened device file.
func (d *Device) Close() error {
	d.Lock()
	if !d.open {
		d.Unlock()
		return errors.New("File already closed")
	}
	log.Printf("Closing '%v'", d.Name())
	d.open = false
	d.Unlock()
	return d.Close()
}
