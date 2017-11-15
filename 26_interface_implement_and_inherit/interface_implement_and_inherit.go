package main

import "fmt"

//------------------------------------------------------------
// Runnable Interface
//------------------------------------------------------------
type Runnable interface {
	Init()
	Run()
}

//------------------------------------------------------------
// BaseObject Type
//------------------------------------------------------------
type BaseObject struct {
}

func (b *BaseObject) Init() {
	fmt.Println("Initing BaseObject")
}

func (b *BaseObject) Run() {
	fmt.Println("Running BaseObject")
}

//------------------------------------------------------------
// InheritedObject Type
//------------------------------------------------------------
type InheritedObject struct {
	*BaseObject // Embedded anonymous type
}

func (i *InheritedObject) Init() {
	fmt.Println("Initing InheritedObject")
}

//------------------------------------------------------------
// BaseObject2 Type
// Implement Runnable interface
//------------------------------------------------------------
type BaseObject2 struct {
}

func (b BaseObject2) Init() {
	fmt.Println("Initing BaseObject2")
}

func (b BaseObject2) Run() {
	fmt.Println("Running BaseObject2")
}

//------------------------------------------------------------
// Run Runnable
//------------------------------------------------------------
func Run(r Runnable) {
	r.Run()
}

//------------------------------------------------------------
// Main Function
//------------------------------------------------------------
func main() {
	var bo BaseObject
	//Run(bo) //Error: BaseObject does not implement Runnable (Init method has pointer receiver)
	Run(&bo)

	pbo := new(BaseObject)
	Run(pbo)

	var bo2 BaseObject2
	Run(bo2)  // BaseObject2 has implemented Runnable interface
	Run(&bo2) // Pointer to BaseObject2 is also valid

	piho := new(InheritedObject)
	Run(piho) //Use the method Run defined by BaseObject
}
