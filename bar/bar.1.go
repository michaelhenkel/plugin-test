package main
import (
	"fmt"
	"github.com/michaelhenkel/plugin-test/types"
	"github.com/michaelhenkel/plugin-test/intentresourceinterface"
)

func GetIntentResource(data map[string]interface{}) (i intentresourceinterface.IntentResourceInterface, err error) {
	fmt.Println("bar data: ", data)
	barName := data["name"]
	barProp := data["properties"]
	i = types.Bar{
		Name: barName.(string),
		Prop: barProp.(string),
	}
	fmt.Printf("[plugin GetFilter] Returning interface: %T %v\n", i, i)
	return
}
