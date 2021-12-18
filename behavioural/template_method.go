package behavioural

import "fmt"

/*

Summary:
Template Method is used to define the skeleteon of the algorithm
as a sequence of operations.
How the operation is performed is left to the concrete implementations.

Example: SurgeryWorkflowTemplate method
Let's perform a surgery workflow for a patient using this. Surgery's
CareProvider interface defines the workflow. The template method executes
that workflow for different concrete implementations: shalby, practo surgeries

Benefit:
The skeleton of the algorithm which performs certain work
in some sequence is decoupled with actual implementations. This is very
useful way to implement workflows and also help in testing as it works on
an interface.
*/

// CareProvider defines what makes a care provider for surgery
type CareProvider interface {
	// Book appointment to consult a doctor
	BookConsult()
	// Consult can lead to surgery or not
	Consult()
	// BookSurgery book the surgery
	BookSurgery()
	// Operate is used to perform the surgery
	Operate()
	// Recover is used to recover from surgery
	Recover()
}

// SurgeryWorkflowTemplate is the template method that does a sequence of work to
// get the surgery done
func SurgeryWorkflowTemplate(provider CareProvider) {
	provider.BookConsult()
	provider.Consult()
	provider.BookSurgery()
	provider.Operate()
	provider.Recover()
}

type shalbyHospital struct{}

func NewShalbyHospital() CareProvider  { return &shalbyHospital{} }
func (s *shalbyHospital) BookConsult() { fmt.Printf("shalby book consult ") }
func (s *shalbyHospital) Consult()     { fmt.Printf("shalby consult ") }
func (s *shalbyHospital) BookSurgery() { fmt.Printf("shalby book surgery ") }
func (s *shalbyHospital) Operate()     { fmt.Printf("shalby operate ") }
func (s *shalbyHospital) Recover()     { fmt.Printf("shalby recover") }

type practoCareSurgery struct{}

func NewPractoCareSurgery() CareProvider  { return &practoCareSurgery{} }
func (s *practoCareSurgery) BookConsult() { fmt.Printf("practo book consult ") }
func (s *practoCareSurgery) Consult()     { fmt.Printf("practo consult ") }
func (s *practoCareSurgery) BookSurgery() { fmt.Printf("practo book surgery ") }
func (s *practoCareSurgery) Operate()     { fmt.Printf("practo operate ") }
func (s *practoCareSurgery) Recover()     { fmt.Printf("practo recover") }
