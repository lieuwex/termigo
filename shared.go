package termigo

// #include <stdbool.h>
// #include "shared.h"
import "C"

var redrawHandlers []func(fullRedraw bool)

//export redrawhandler
func redrawhandler(b C.bool) {
	for _, fn := range redrawHandlers {
		fn(bool(b))
	}
}
