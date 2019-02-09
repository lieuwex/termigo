package widgets

// #cgo CFLAGS: -I${SRCDIR}/../termio
// #cgo LDFLAGS: -L${SRCDIR}/../termio -ltermio
// #include <stdlib.h>
// #include <termio.h>
import "C"
import (
	"unsafe"
)

// Prompt is an interactive prompt on screen.
type Prompt struct {
	ptr *C.Promptwidget
}

// NewPrompt creates a new prompt at x,y with the given width and title.
// Destroy has to be called when done with the widget.
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

// Destroy destroys the current prompt.
func (p Prompt) Destroy() {
	C.prw_destroy(p.ptr)
}

// Redraw redraws the prompt on screen.
func (p Prompt) Redraw() {
	C.prw_redraw(p.ptr)
}

// HandleKey should be provided with the key given by `getkey`, it returns a
// string with the prompt input or an error.
func (p Prompt) HandleKey(key int) (string, bool) {
	str := C.prw_handlekey(p.ptr, C.int(key))
	if str == nil {
		return "", false
	}
	return C.GoString(str), true
}

// ChangeTitle changes the title of the prompt.
func (p Prompt) ChangeTitle(title string) {
	str := C.CString(title)
	defer C.free(unsafe.Pointer(str))

	C.prw_changetitle(p.ptr, str)
}
