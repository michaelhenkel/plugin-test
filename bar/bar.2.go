package main
import (
	"fmt"
	"github.com/michaelhenkel/plugin-test/intentresourceinterface"
	"github.com/michaelhenkel/plugin-test/types"
)

type Bar types.Barv2

func (i Bar) Create() string {
	fmt.Println("\t\t Creating Bar:", i)
	return i.Name
}

func GetIntentResource(data map[string]interface{}) (i intentresourceinterface.IntentResourceInterface, err error) {
	fmt.Println("bar data: ", data)
	barName := data["name"]
	barProp := data["properties"]
	barRef := data["references"]
	i = Bar{
		Name: barName.(string),
		Prop: barProp.(string),
		Ref: barRef.(string),
	}
	fmt.Printf("[plugin GetFilter] Returning interface: %T %v\n", i, i)
	return
}
