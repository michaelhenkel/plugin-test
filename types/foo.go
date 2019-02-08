package types
import(
	"fmt"
	"math/rand"
)

var jobs chan int
type Foo struct {
	Name string
}

func (f Foo) Create(jobs chan<- int, out <-chan string) string{
	fmt.Println("writing " + f.Name + " to db")
	bar := Bar{
		Name: "barX",
		Prop: "prop1",
	}
	fmt.Println(bar.Create())
	randInt := rand.Intn(100)
	fmt.Println("sending to writer: ",randInt)
	jobs <- randInt
	fmt.Println(<-out)
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
