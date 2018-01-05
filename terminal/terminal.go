package terminal

import (
	"errors"
	"syscall"
	"unsafe"
)

type terminal struct {
	originalTermios *syscall.Termios
	fd              uintptr
}

func getTerminal() (*terminal, error) {
	fd := uintptr(syscall.Stdin)
	originalTermios := &syscall.Termios{}
	if _, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd,
		syscall.TIOCGETA,
		uintptr(unsafe.Pointer(originalTermios)),
	); err != 0 {
		return nil, errors.New("could not load Termios")
	}
	return &terminal{
		originalTermios: originalTermios,
		fd:              fd,
	}, nil
}

func (t *terminal) updateTerminal(input *syscall.Termios) error {
	if _, _, err := syscall.Syscall(
		syscall.SYS_IOCTL,
		t.fd,
		uintptr(syscall.TIOCSETA),
		uintptr(unsafe.Pointer(input)),
	); err != 0 {
		return errors.New("could not update Termios")
	}
	return nil
}

func (t *terminal) restoreTerminal() error {
	return t.updateTerminal(t.originalTermios)
}

func Echo(input string) error {
	t, err := getTerminal()
	if err != nil {
		return err
	}
	defer t.restoreTerminal()

	temp := *t.originalTermios
	temp.Lflag &^= syscall.ECHO
	t.updateTerminal(&temp)

	for _, c := range []byte(input) {
		if _, _, err := syscall.Syscall(
			syscall.SYS_IOCTL,
			t.fd,
			syscall.TIOCSTI,
			uintptr(unsafe.Pointer(&c)),
		); err != 0 {
			return errors.New("could not write to terminal")
		}
	}
	return nil
}

func Run(input string) error {
	return Echo(input + "\n")
}
