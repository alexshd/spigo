// test the flow package
package flow

import (
	"fmt"
	"github.com/adrianco/spigo/archaius"
	"github.com/adrianco/spigo/gotocol"
	"testing"
)

func TestFlow(t *testing.T) {
	archaius.Conf.Collect = true
	r1 := gotocol.NewTrace()
	r2 := gotocol.NewTrace()
	r3 := gotocol.NewTrace()
	archaius.Conf.Arch = "test"
	Update(r1, "one")
	Update(r2, "two")
	Update(r3, "three")
	Update(r1.NewSpan(), "une")
	Update(r2.NewSpan(), "deux")
	Update(r3.NewSpan(), "trois")
	fmt.Println(flowmap)
	End(r1)
	fmt.Println("Walk")
	Walk(flowmap)
	Shutdown()
}