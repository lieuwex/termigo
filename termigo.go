package termigo

// #cgo CFLAGS: -I${SRCDIR}/termio
// #cgo LDFLAGS: -L${SRCDIR}/termio -ltermio
// #include <stdlib.h>
// #include <termio.h>
// void cgo_setstyle(int fg, int bg, bool bold, bool ul) {
// 	struct Style s = { fg, bg, bold, ul };
// 	setstyle(&s);
// }
import "C"
import (
	"errors"
	"fmt"
	"io"
	"unsafe"
)

type Style struct {
	Foreground, Background int
	Bold, Underlined       bool
}

func InitScreen() {
	C.initscreen()
}
func EndScreen() {
	C.endscreen()
}

func InitKeyboard(nosigkeys bool) {
	C.initkeyboard(C.bool(nosigkeys))
}
func EndKeyboard() {
	C.endkeyboard()
}

func InstallCLHandler(install bool) {
	C.installCLhandler(C.bool(install))
}
func InstallRedrawHandler() {
	// TODO
}

func ClearScreen() {
	C.clearscreen()
}

func Putc(c int) {
	C.tputc(C.int(c))
}
func Print(str string) int {
	cstr := C.CString(str)
	res := C.tprint(cstr)
	C.free(unsafe.Pointer(cstr))
	return int(res)
}
func Printf(format string, a ...interface{}) int {
	return Print(fmt.Sprintf(format, a...))
}

func GetTerminalSize() (width, height int) {
	size := C.gettermsize()
	return int(size.w), int(size.h)
}
func SetStyle(style Style) {
	C.cgo_setstyle(
		C.int(style.Foreground),
		C.int(style.Background),
		C.bool(style.Bold),
		C.bool(style.Underlined),
	)
}
func SetForegound(color int) {
	C.setfg(C.int(color))
}
func SetBackground(color int) {
	C.setbg(C.int(color))
}
func SetBold(val bool) {
	C.setbold(C.bool(val))
}
func SetUnderlined(val bool) {
	C.setul(C.bool(val))
}
func FillRectangle(x, y, width, height, color int) {
	C.fillrect(
		C.int(x),
		C.int(y),
		C.int(width),
		C.int(height),
		C.int(color),
	)
}
func Redraw() {
	C.redraw()
}
func RedrawFull() {
	C.redrawfull()
}

func ScrollTerminal(x, y, width, height, amount int) {
	C.scrollterm(
		C.int(x),
		C.int(y),
		C.int(width),
		C.int(height),
		C.int(amount),
	)
}

func GetBufferChar(x, y int) int {
	val := C.getbufferchar(C.int(x), C.int(y))
	return int(val)
}

func MoveTo(x, y int) {
	C.moveto(C.int(x), C.int(y))
}
func PushCursor() {
	C.pushcursor()
}
func PopCursor() {
	C.popcursor()
}

func Bel() {
	C.bel()
}
func CursorVisible(visible bool) {
	C.cursorvisible(C.bool(visible))
}

func GetKey() (Key, error) {
	res := C.tgetkey()
	switch res {
	case -1:
		return 0, io.EOF
	case -2:
		return 0, errors.New("unknown error") // TODO
	default:
		return Key(res), nil
	}
}

func GetLine() (string, bool) {
	str := C.tgetline()
	if str == nil {
		return "", false
	}
	return C.GoString(str), true
}
