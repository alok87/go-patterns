package creational

/*

Summary:
Fluent Builder helps in creating a complex object in steps using
method chaining technique.

Example:
Create an Employee object which has many user specified parameters.
With this approach we construct the object easily in one line (readability):
emp = builder.Called("Bezos").Company("Amazon").Address("USA").Build()

Benefit: Readability
When the Constructor is a a bad choice to construct an object?
When the constructor depends on huge number of user specified parameters,
using the fluent method chaining simplies the object creation in steps.
*/

type Employee struct {
	Name    string
	Company string
	Address string
}

type EmployeeBuilder struct {
	name, address, company string
}

func NewEmployeeBuilder() *EmployeeBuilder {
	return &EmployeeBuilder{}
}

func (c *EmployeeBuilder) WithName(name string) *EmployeeBuilder {
	c.name = name
	return c
}
func (c *EmployeeBuilder) WithCompany(company string) *EmployeeBuilder {
	c.company = company
	return c
}
func (c *EmployeeBuilder) WithAddress(address string) *EmployeeBuilder {
	c.address = address
	return c
}

func (c *EmployeeBuilder) Build() Employee {
	return Employee{
		Name:    c.name,
		Company: c.company,
		Address: c.address,
	}
}
