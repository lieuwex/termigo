package termigo

// #cgo CFLAGS: -I${SRCDIR}/termio
// #cgo LDFLAGS: -L${SRCDIR}/termio -ltermio
// #include <stdlib.h>
// #include <termio.h>
// #include "shared.h"
// void cgo_setstyle(int fg, int bg, bool bold, bool ul) {
// 	struct Style s = { fg, bg, bold, ul };
// 	setstyle(&s);
// }
// void initworld() {
// 	installredrawhandler(redrawhandler);
// }
import "C"
import (
	"errors"
	"fmt"
	"io"
	"unsafe"
)

// Style contains styling information for the text on the terminal.
type Style struct {
	Foreground, Background int
	Bold, Underlined       bool
}

// Init initialises the screen and world.
func Init() {
	C.initworld()
	C.initscreen()
}

// End ends the screen and world.
func End() {
	C.endscreen()
}

// InitKeyboard starts catching keyboard events.
func InitKeyboard(nosigkeys bool) {
	C.initkeyboard(C.bool(nosigkeys))
}

// EndKeyboard stops catching keyboard events.
func EndKeyboard() {
	C.endkeyboard()
}

// HandleCLs enables or disables ^L handling.
func HandleCLs(install bool) {
	C.installCLhandler(C.bool(install))
}

// OnRedraw calls the given fn when the screen is redrawn.
func OnRedraw(fn func(fullRedraw bool)) {
	// REVIEW: why is this actually through termio when we could just do it
	// ourselves?
	redrawHandlers = append(redrawHandlers, fn)
}

// ClearScreen clears the screen.
func ClearScreen() {
	C.clearscreen()
}

// Putc puts the given char onto the screen.
func Putc(c int) {
	C.tputc(C.int(c))
}

// Print prints the given str onto the screen.
func Print(str string) int {
	cstr := C.CString(str)
	res := C.tprint(cstr)
	C.free(unsafe.Pointer(cstr))
	return int(res)
}

// Printf formats the given format string using the given parameters through the
// fmt package, then prints the result onto the screen.
func Printf(format string, a ...interface{}) int {
	return Print(fmt.Sprintf(format, a...))
}

// GetTerminalSize returns the width and height of the terminal.
func GetTerminalSize() (width, height int) {
	size := C.gettermsize()
	return int(size.w), int(size.h)
}

// SetStyle sets the styling of the text on the terminal.
func SetStyle(style Style) {
	C.cgo_setstyle(
		C.int(style.Foreground),
		C.int(style.Background),
		C.bool(style.Bold),
		C.bool(style.Underlined),
	)
}

// SetForegound sets the foreground color.
func SetForegound(color int) {
	C.setfg(C.int(color))
}

// SetForegound sets the background color.
func SetBackground(color int) {
	C.setbg(C.int(color))
}

// SetForegound sets whether or not text should be bold.
func SetBold(val bool) {
	C.setbold(C.bool(val))
}

// SetForegound sets whether or not text should be underlined.
func SetUnderlined(val bool) {
	C.setul(C.bool(val))
}

// FillRectangle fills a rectangle on screen starting at x,y with the given
// width,height using the given char.
func FillRectangle(x, y, width, height, char int) {
	C.fillrect(
		C.int(x),
		C.int(y),
		C.int(width),
		C.int(height),
		C.int(color),
	)
}

// Redraw redraws the screen using diffing.
func Redraw() {
	C.redraw()
}

// Redraw redraws the screen, ignoring diffing.
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

// GetBufferChar returns the char in the buffer at the given x,y.
func GetBufferChar(x, y int) int {
	val := C.getbufferchar(C.int(x), C.int(y))
	return int(val)
}

// MoveTo moves the cursor to the given x,y.
func MoveTo(x, y int) {
	C.moveto(C.int(x), C.int(y))
}
func PushCursor() {
	C.pushcursor()
}
func PopCursor() {
	C.popcursor()
}

// Bel makes a bel sound
func Bel() {
	C.bel()
}

// CursorVisible sets whether or not the cursor should be visible in the
// terminal emulator. Takes effect immediately.
func CursorVisible(visible bool) {
	C.cursorvisible(C.bool(visible))
}

// GetKey waits and returns a key or error.
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

// GetLine waits and returns an entered line or error.
func GetLine() (string, bool) {
	str := C.tgetline()
	if str == nil {
		return "", false
	}
	return C.GoString(str), true
}
