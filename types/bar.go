package types

type Bar struct {
	Name string
	Prop string
}

func (b Bar) Create() string{
	return "Created: " + b.Name
}
func (b Bar) Read() string{
	return "Read: " + b.Name
}
func (b Bar) Update() string{
	return "Update: " + b.Name
}
func (b Bar) Delete() string{
	return "Delete: " + b.Name
}
