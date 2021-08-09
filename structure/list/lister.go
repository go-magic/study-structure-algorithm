package list

type Lister interface {
	Append(interface{})
	Remove(interface{})
	RemoveAt(int)
	Show(int) interface{}
}
