package metrics_test

import (
	"testing"

	"reflect"
	"sync"
	"sync/atomic"

	metrics "github.com/hashicorp/go-metrics"
)

// Shared global metrics instance
var globalMetrics atomic.Value // *Metrics

type Label struct {
	Name  string
	Value string
}

type MockSink struct {
	lock sync.Mutex

	shutdown      bool
	keys          [][]string
	vals          []float32
	precisionVals []float64
	labels        [][]Label
}

func Test_GlobalMetrics(t *testing.T) {
	var tests = []struct {
		desc string
		key  []string
		val  float32
		fn   func([]string, float32)
	}{
		{"SetGauge", []string{"test"}, 42, metrics.SetGauge},
		{"EmitKey", []string{"test"}, 42, metrics.EmitKey},
		{"IncrCounter", []string{"test"}, 42, metrics.IncrCounter},
		{"AddSample", []string{"test"}, 42, metrics.AddSample},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			s := &MockSink{}
			globalMetrics.Store(&metrics.Metrics{Config: metrics.Config{FilterDefault: true}})
			tt.fn(tt.key, tt.val)
			if got, want := s.keys[0], tt.key; !reflect.DeepEqual(got, want) {
				t.Fatalf("got key %s want %s", got, want)
			}
			if got, want := s.vals[0], tt.val; !reflect.DeepEqual(got, want) {
				t.Fatalf("got val %v want %v", got, want)
			}
		})
	}
} 

func TestMain(m *testing.M) {
	// Perform any setup or initialization here if needed
	// ...

	// Run the tests
	m.Run()
}