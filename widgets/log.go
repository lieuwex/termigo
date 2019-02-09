package widgets

// #cgo CFLAGS: -I${SRCDIR}/../termio
// #cgo LDFLAGS: -L${SRCDIR}/../termio -ltermio
// #include <stdlib.h>
// #include <termio.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Log struct {
	ptr *C.Logwidget
}

func NewLog(x, y, width, height int, title string, timestamps bool) Log {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	ptr := C.lgw_make(
		C.int(x),
		C.int(y),
		C.int(width),
		C.int(height),
		ctitle,
		C.bool(timestamps),
	)

	return Log{ptr}
}

func (l Log) Destroy() {
	C.lgw_destroy(l.ptr)
}

func (l Log) Redraw() {
	C.lgw_redraw(l.ptr)
}

func (l Log) Add(line string) {
	str := C.CString(line)
	defer C.free(unsafe.Pointer(str))

	C.lgw_add(l.ptr, str)
}

func (l Log) Addf(format string, a ...interface{}) {
	l.Add(fmt.Sprintf(format, a...))
}

func (l Log) Clear() {
	C.lgw_clear(l.ptr)
}

func (l Log) ChangeTitle(title string) {
	str := C.CString(title)
	defer C.free(unsafe.Pointer(str))

	C.lgw_changetitle(l.ptr, str)
}
