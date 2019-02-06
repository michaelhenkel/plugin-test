package main
import (
	"fmt"
	"github.com/michaelhenkel/intentresourceinterface"
)

type Bar struct {
	Name	string
	Prop	string
	Ref	string
}

func (i Bar) Create() {
	fmt.Println("\t\t Creating Bar:", i)
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
