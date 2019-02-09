package termigo

// Key is a keycode or a keycode+ modifier.
//
// For special control keys:
// Add (+) the relevant character: ascii '@' (64) till '_' (95)
// E.g. ^C is KEY_CTRL + 'C'
// Note that alt-[ is interpreted as the start of an escape sequence, and not normally readable
type Key int

const (
	KeyTab       Key = 9
	KeyLF        Key = 10
	KeyCR        Key = 13
	KeyEsc       Key = 27
	KeyBackspace Key = 127

	KeyRight    Key = 1001
	KeyUp       Key = 1002
	KeyLeft     Key = 1003
	KeyDown     Key = 1004
	KeyPageup   Key = 1021
	KeyPagedown Key = 1022
	KeyDelete   Key = 1100
	KeyShifttab Key = 1200

	KeyCtrl      Key = -64
	KeyShift     Key = 1000000
	KeyAlt       Key = 120000
	KeyCtrlAlt   Key = KeyCtrl + KeyAlt
	KeyCtrlShift Key = KeyCtrl + KeyShift
)
