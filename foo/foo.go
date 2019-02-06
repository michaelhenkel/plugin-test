package main
import (
	"fmt"
	"github.com/michaelhenkel/intentresourceinterface"
)

type Foo struct {
	Name string
}

func (i Foo) Create() {
	fmt.Println("\t\t Creating Foo:", i.Name)
}

func GetIntentResource(data map[string]interface{}) (i intentresourceinterface.IntentResourceInterface, err error) {
	fmt.Println("foo data: ", data["name"])
	fooName := data["name"]
	i = Foo{
		Name: fooName.(string),
	}
	fmt.Printf("[plugin GetFilter] Returning interface: %T %v\n", i, i)
	return
}
