// +build !linux

package basestation

import (
	"fmt"

	t "github.com/lucaspiller/meteorpi/src/sensors/types"
)

func Start(data chan *t.Measurement) {
	fmt.Println("Basestation not available on this platform")
}
