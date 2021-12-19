package behavioural

import "fmt"

/*

Summary:
Chain of Responsilbility pattern is used when a work is passed to a chain
of handlers until one of the processes it. The handlers can break the chain
and return as well.

Example:
Alert escalation policy follows a chain. If the alert is resolved by the first
one in chain, its done, otherwise next is tried. Finally the escaltion timesout.

Benefit:
- The client is decopupled from a series of task as it chains only to one
and then follows the chain to execute many.
- Design pattern to add series of tasks as chains of tasks.
- Easier to add new task, easier to write unit test for each task.
*/

type EscalationHandler interface {
	Handle(alert string)
}

type escalation struct {
	user      string
	canHandle []string
	next      EscalationHandler
}

func NewEscalation(
	u string,
	c []string,
	h EscalationHandler,
) EscalationHandler {
	return &escalation{user: u, canHandle: c, next: h}
}

func (e *escalation) Handle(alert string) {
	for _, a := range e.canHandle {
		if a == alert {
			fmt.Printf("resolved by %s", e.user)
			return
		}
	}

	if e.next == nil {
		fmt.Printf("timedout")
		return
	}

	e.next.Handle(alert)
}
