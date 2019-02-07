package main
import (
	"fmt"
	"github.com/michaelhenkel/plugin-test/types"
	"github.com/michaelhenkel/plugin-test/intentresourceinterface"
)

func GetIntentResource(data map[string]interface{}) (i intentresourceinterface.IntentResourceInterface, err error) {
	fmt.Println("foo data: ", data["name"])
	fooName := data["name"]
	i = types.Foo{
		Name: fooName.(string),
	}
	fmt.Printf("[plugin GetFilter] Returning interface: %T %v\n", i, i)
	return
}
