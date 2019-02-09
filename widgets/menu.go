package widgets

/*

// #cgo CFLAGS: -I${SRCDIR}/../termio
// #cgo LDFLAGS: -L${SRCDIR}/../termio -ltermio
// #include <stdlib.h>
// #include <termio.h>
// Menuitem cgo_menuitem(char *text, int hotkey, void (*func)(int index)) {
// 	Menuitem res = {text, hotkey, func};
// 	return res;
// }
//
// Menudata *cgo_menudata(int nitems, Menuitem *items) {
// 	Menudata *res = malloc(sizeof(Menudata));
//
// 	res->nitems = nitems;
// 	res->items = items;
//
// 	return res;
// }
import "C"
import "unsafe"

type MenuKey int

const (
	MenuHandled MenuKey = iota
	MenuIgnored
	MenuQuit
	MenuCalled
)

type MenuItem struct {
	text   string
	hotkey int
	fn     func(int)
}

type Menu struct {
	ptr  *C.Menuwidget
	data *C.Menudata
}

func NewMenu(x, y int, items []MenuItem) Menu {
	nitems := len(items)
	itemsRaw := make([]C.Menuitem, nitems)
	for i, item := range items {
		itemsRaw[i] = C.cgo_menuitem(
			item.text,
			C.int(hotkey),
			item.fn,
		)
	}
	itemsptr := unsafe.Pointer(&items[0])
	menudata := C.cgo_menudata(nitems, (*C.Menuitem)(itemsptr))

	ptr := C.menu_make(
		C.int(x),
		C.int(y),
		menudata,
	)

	return Menu{ptr, menudata}
}

func (m Menu) Destroy() {
	C.menu_destroy(m.ptr)
	C.free(unsafe.Pointer(m.data))
}

func (m Menu) Redraw() {
	C.menu_redraw(m.ptr)
}

func (m Menu) HandleKey(key int) MenuKey {
	res := C.menu_handlekey(m.ptr, C.int(key))
	return MenuKey(res)
}
*/
