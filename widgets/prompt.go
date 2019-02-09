package widgets

// #cgo CFLAGS: -I${SRCDIR}/../termio
// #cgo LDFLAGS: -L${SRCDIR}/../termio -ltermio
// #include <stdlib.h>
// #include <termio.h>
import "C"
import (
	"unsafe"
)

type Prompt struct {
	ptr *C.Promptwidget
}

func NewPrompt(x, y, width int, title string) Prompt {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	ptr := C.prw_make(
		C.int(x),
		C.int(y),
		C.int(width),
		ctitle,
	)

	return Prompt{ptr}
}

func (p Prompt) Destroy() {
	C.prw_destroy(p.ptr)
}

func (p Prompt) Redraw() {
	C.prw_redraw(p.ptr)
}

func (p Prompt) HandleKey(key int) (string, bool) {
	str := C.prw_handlekey(p.ptr, C.int(key))
	if str == nil {
		return "", false
	}
	return C.GoString(str), true
}

func (p Prompt) ChangeTitle(title string) {
	str := C.CString(title)
	defer C.free(unsafe.Pointer(str))

	C.prw_changetitle(p.ptr, str)
}
