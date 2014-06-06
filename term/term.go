package term

/*
#cgo LDFLAGS: -ltermcap
#include <stdlib.h>
#include <unistd.h>
#include <termios.h>
#include <curses.h>
#include <term.h>

int put_int(int c)
{
  return write(1, &c, 1);
}

int tgetent_() {
  if (tgetent(NULL, getenv("TERM")) == -1)
    return 1;
  return 0;
}

int set_cap(char *str)
{
  char  *area;

  if ((area = tgetstr(str, NULL)) == NULL)
    return 1;
  tputs(area, 3, &put_int);
  return 0;
}

*/
import "C"

import (
	"errors"
	"unsafe"
)

// See man (5) terminfo
func SetCap(cpblt string) error {
	cp := C.CString(cpblt)
	defer C.free(unsafe.Pointer(cp))

	if err := C.set_cap(cp); err == 1 {
		return errors.New("Capability failed")
	}
	return nil
}

// See man (3) curs_termcap
func TGetEnt() error {
	if err := C.tgetent_(); err != 1 {
		if err == 0 {
			errors.New("TGetEnt error: There is no such entry.")
		} else if err == -1 {
			errors.New("TGetEnt error: Terminfo database cannot be found.")
		}
	}
	return nil
}
