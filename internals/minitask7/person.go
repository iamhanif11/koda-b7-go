package minitask7

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

//print
func (p Person) Print() string {
	return fmt.Sprintf("Nama: %s, Alamat: %s, Telepon: %s", p.Name, p.Address, p.Phone)
}

// greet
func (p *Person) Greet() {
	fmt.Printf("Hai %s salam kenal\n", p.Name)
}

// setter
func (p *Person) SetName(newName string) {
	p.Name = newName

}
