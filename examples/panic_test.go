package examples

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/aghosn/gasm/wasi"
	"github.com/aghosn/gasm/wasm"
	"github.com/stretchr/testify/require"
)

func Test_panic(t *testing.T) {
	buf, err := ioutil.ReadFile("wasm/panic.wasm")
	require.NoError(t, err)

	mod, err := wasm.DecodeModule(bytes.NewBuffer(buf))
	require.NoError(t, err)

	vm, err := wasm.NewVM(mod, wasi.New().Modules())
	require.NoError(t, err)

	defer func() {
		err := recover()
		require.Equal(t, "unreachable", err)
	}()
	vm.ExecExportedFunction("cause_panic")
}
