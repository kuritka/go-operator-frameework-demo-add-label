package memcached


import (
	"context"
	"github.com/example-inc/memcached-operator/pkg/apis/cache/v1alpha1"
	"sync"


	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// stateSetter sets all GSLB properties once in case they are not defined
type stateSetter struct {
	predefinedStrategy v1alpha1.MemcachedSpec
	strategyOnce       sync.Once
	applyOnce          sync.Once
}
// If you have no serious reason newStateSetter() within Init() func
func newStateSetter()  *stateSetter {
	stateSetter := new(stateSetter)
	stateSetter.predefinedStrategy = v1alpha1.MemcachedSpec{Interval: 30, Threshold: 20, Size: 2}
	return stateSetter
}

// Sets once yaml variables to predefined - non default values
func (s *stateSetter) SetPredefinedStrategy(spec *v1alpha1.MemcachedSpec) *stateSetter{
	s.strategyOnce.Do(func(){
		if spec.Interval== 0 {
			spec.Interval  = s.predefinedStrategy.Interval
		}
		if spec.Threshold  == 0 {
			spec.Threshold  = s.predefinedStrategy.Threshold
		}
	})
	return s
}

// Applies values once
func (s *stateSetter) ApplyOnce(client client.Client, ctx context.Context,gslb runtime.Object) error {
	var err error
	s.applyOnce.Do(func(){
		err = client.Update(ctx, gslb)
	})
	return err
}

