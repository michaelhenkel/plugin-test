package types
import(
	"fmt"
)

type Foo struct {
	Name string
}

func (f Foo) Create() string{
	fmt.Println("writing " + f.Name + " to db")
	bar := Bar{
		Name: "barX",
		Prop: "prop1",
	}
	fmt.Println(bar.Create())
	return "Created: " + f.Name
}
func (f Foo) Read() string{
	return "Read: " + f.Name
}
func (f Foo) Update() string{
	return "Update: " + f.Name
}
func (f Foo) Delete() string{
	return "Delete: " + f.Name
}
