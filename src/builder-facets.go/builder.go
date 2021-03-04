package main

import "fmt"

type Person struct {
	// address
	StreetAddress string
	PostCode      string
	City          string

	// job
	CompanyName  string
	Position     string
	AnnualIncome int
}

type PersonBuilder struct {
	person *Person // Pointer to a Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostCode(postalCode string) *PersonAddressBuilder {
	b.person.PostCode = postalCode
	return b
}

func (b *PersonJobBuilder) At(company string) *PersonJobBuilder {
	b.person.CompanyName = company
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	b.person.AnnualIncome = annualIncome
	return b
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {

	pb := NewPersonBuilder()

	pb.
		Lives().
		At("123 London Road").
		In("London").
		WithPostCode("SW12BC").
		Works().
		At("FooBar").
		AsA("Programmer").
		Earning(1)

	person := pb.Build()

	fmt.Println(*person)

}
