// +build linux

package basestation

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var device = "/dev/ttyAMA0"
var baud = uint32(syscall.B19200)

type Serial struct {
	File    *os.File
	Scanner *bufio.Scanner
}

func SerialOpen() *Serial {
	// open device
	file, err := os.OpenFile(device, syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)
	if err != nil {
		panic(err)
	}

	// perform termios syscall to set mode correctly
	fd := file.Fd()
	termios := syscall.Termios{
		Iflag:  syscall.IGNPAR,                                                // ignore framing / parity errors
		Cflag:  syscall.CS8 | syscall.CREAD | syscall.CLOCAL | syscall.B19200, // 8 bit characters, rx mode, ignore modem lines, set baud rate
		Cc:     [32]uint8{syscall.VMIN: 1},                                    // return data when there is at least 1 character
		Ispeed: syscall.B19200,                                                // speeds
		Ospeed: syscall.B19200,
	}

	_, _, errno := syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TCSETS),
		uintptr(unsafe.Pointer(&termios)),
		0,
		0,
		0,
	)
	if errno != 0 {
		panic(fmt.Sprintf("Error setting serial mode: %v", errno))
	}

	if err = syscall.SetNonblock(int(fd), false); err != nil {
		panic(err)
	}

	// open buffered reader
	scanner := bufio.NewScanner(file)
	scanner.Split(scanAnsiLines)

	return &Serial{File: file, Scanner: scanner}
}

func (s *Serial) Scan() bool {
	// blocks until we have read a line
	return s.Scanner.Scan()
}

func (s *Serial) Text() string {
	return s.Scanner.Text()
}

func (s *Serial) Close() {
	s.File.Close()
}

//----------------------------------------------------------------------------

// scan for ANSI escape sequence line breaks (\x1BE)
func scanAnsiLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	escapeIndex := bytes.IndexByte(data, 27)
	// check whether we have an ANSI escape character, and more data following
	if escapeIndex >= 0 && len(data) > escapeIndex+1 {
		// check whether the escape sequence is CNL (cursor next line)
		if data[escapeIndex+1] == 69 {
			// We have a full line
			return escapeIndex + 2, data[0:escapeIndex], nil
		}
	}

	// Request more data.
	return 0, nil, nil
}
