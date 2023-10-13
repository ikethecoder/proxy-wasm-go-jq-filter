package main

import (
	"reflect"
	"testing"

	"github.com/ikethecoder/proxy-wasm-go-filter-template/config"
	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/proxytest"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func Test_myPluginHTTPContext_OnPluginStart(t *testing.T) {
	vmTest(t, func(t *testing.T, vm types.VMContext) {
		opt := proxytest.NewEmulatorOption().WithVMContext(vm)
		host, reset := proxytest.NewHostEmulator(opt)
		defer reset()

		// Call OnPluginStart.
		require.Equal(t, types.OnPluginStartStatusOK, host.StartPlugin())

		// Check Envoy logs.
		logs := host.GetInfoLogs()
		require.Contains(t, logs, "OnPluginStart from Go!")
	})
}

func Test_myPluginHTTPContext_OnHttpResponseBody(t *testing.T) {
	type fields struct {
		conf               *config.Config
		DefaultHttpContext types.DefaultHttpContext
	}
	type args struct {
		in0 int
		in1 bool
	}

	tests := []struct {
		name   string
		body   []byte
		fields fields
		args   args
		want   types.Action
	}{
		{
			name: "simple",
			body: []byte(`
				[
					{
						"Id": "1",
						"Priority": "Urgent"
					},
					{
						"Id": "2",
						"Priority": "Low"
					}
				]
			`),
			fields: fields{
				conf: &config.Config{Query: string(`.[] | select(.Priority == "Urgent") | .Id`)},
			},
			args: args{in0: 0, in1: false},
			want: types.ActionContinue,
		},
	}

	for _, tt := range tests {

		vmTest(t, func(t *testing.T, vm types.VMContext) {
			opt := proxytest.NewEmulatorOption().WithVMContext(vm)
			host, reset := proxytest.NewHostEmulator(opt)
			defer reset()

			// Call OnPluginStart.
			//require.Equal(t, types.OnPluginStartStatusOK, host.StartPlugin())

			contextId := host.InitializeHttpContext()

			action := host.CallOnResponseBody(contextId, tt.body, true)

			if !reflect.DeepEqual(action, types.ActionContinue) {
				t.Errorf("myPluginHTTPContext.OnHttpResponseBody() = %v, want %v", action, types.ActionContinue)
			}

			m := &myPluginHTTPContext{
				conf:               tt.fields.conf,
				DefaultHttpContext: tt.fields.DefaultHttpContext,
			}
			if got := m.OnHttpResponseBody(tt.args.in0, tt.args.in1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("myPluginHTTPContext.OnHttpResponseBody() = %v, want %v", got, tt.want)
			}

		})

	}
}

func vmTest(t *testing.T, f func(*testing.T, types.VMContext)) {
	t.Helper()

	t.Run("go", func(t *testing.T) {
		f(t, &myPluginVMContext{})
	})
}
