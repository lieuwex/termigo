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

// Log is a logging widget, copying an instance doesn't actually copy stuff.
type Log struct {
	ptr *C.Logwidget
}

// NewLog creates a new log widget instance at x,y with the given width, height,
// and title. If timestamps is true, a timestamp is prepended to every entry.
// Destroy has to be called when done with the widget.
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

// Destroy the current log widget.
func (l Log) Destroy() {
	C.lgw_destroy(l.ptr)
}

// Redraw the current widget onto the screen.
func (l Log) Redraw() {
	C.lgw_redraw(l.ptr)
}

// Add adds the given line to the logging widget.
func (l Log) Add(line string) {
	str := C.CString(line)
	defer C.free(unsafe.Pointer(str))

	C.lgw_add(l.ptr, str)
}

// Addf applies the format and arguments onto fmt, and adds the result to the
// logging widget.
func (l Log) Addf(format string, a ...interface{}) {
	l.Add(fmt.Sprintf(format, a...))
}

// Clear removes all entries from the widget.
func (l Log) Clear() {
	C.lgw_clear(l.ptr)
}

// ChangeTitle changes the title of the widget.
func (l Log) ChangeTitle(title string) {
	str := C.CString(title)
	defer C.free(unsafe.Pointer(str))

	C.lgw_changetitle(l.ptr, str)
}
