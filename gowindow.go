package gowindow

import (
	"regexp"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32                   = windows.NewLazyDLL("user32.dll")
	procEnumWindows          = user32.NewProc("EnumWindows")
	procGetWindowTextW       = user32.NewProc("GetWindowTextW")
	procGetWindowTextLengthW = user32.NewProc("GetWindowTextLengthW")
)

func enumWindows(lpEnumFunc uintptr, lParam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, lpEnumFunc, lParam, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r1, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r1)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

//Returns a slice of window titles matching the passed regexp
func FindWindow(rTitle *regexp.Regexp) (s []string) {

	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 255)
		_, err := getWindowText(h, &b[0], int32(len(b)))
		if err != nil {
			return 1
		}
		if rTitle.MatchString(syscall.UTF16ToString(b)) {
			s = append(s, syscall.UTF16ToString(b))
			return 1
		}
		return 1
	})
	enumWindows(cb, 0)

	return
}
