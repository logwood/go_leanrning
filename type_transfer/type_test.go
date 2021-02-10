package type_transfer

import (
	"testing"
)

func TestType(t *testing.T) {
	var var1 int = 7
	t.Logf("%T->%v\n", var1, var1)
}
