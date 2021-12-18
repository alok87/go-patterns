package creational

import (
	"bytes"
	"encoding/gob"
)

/*

Summary:
Prototype pattern is used to create copies of objects.
Self clone: Actual objects have the responsiblity to make their own clone.
The object to be cloned exposes Clone() methods.

DeepCopy:
DeepCopy() is used in Golang to copy objects with actual data
without referrence copy. Two approaches to deepcopy:
1. DeepCopy using function: Use function to copy referrences as fresh object.
Downside is we need to make many functions for a large nested objects.
2. DeepCopy through searialization: Use serialization. Serialize an object
and deserialize it another object. Saves developer from writing a lot of code.

Example:
The autoscaler controller wants to work on a copy of the object, make changes
to it and then call the API to update the object status. We do not want to
create a new object as it is expensive. DeeCopy is used.

1. Complex objects can require nested copies and many deep copies, it helps
to use this pattern.
2. Objects private variables can only be performed by the same object. So self
clone is also a necessity.
3. Copied object are better than creating objects from scratch. This saves
us from heavy computations which may be required when created from scratch.

*/

type CopyMode int

const (
	Serialization CopyMode = iota
	Function
)

type AutoScaler struct {
	Spec   *AutoScalerSpec
	Status *AutoScalerStatus

	Mode CopyMode
}

type AutoScalerSpec struct {
	Min    *int
	Max    *int
	Target *float32
}

type AutoScalerStatus struct {
	Current   int
	Desired   int
	Available int
}

func NewAutoScaler(min, max int, target float32) *AutoScaler {
	return &AutoScaler{
		Spec: &AutoScalerSpec{
			Min:    &min,
			Max:    &max,
			Target: &target,
		},
		Status: &AutoScalerStatus{},
		Mode:   Serialization,
	}
}

func DeepCopyWithSerialization(a *AutoScaler) AutoScaler {
	buf := bytes.Buffer{}

	encode := gob.NewEncoder(&buf) // buf implements io.Writer's Write()
	encode.Encode(a)               // error not handled for simplicity

	var copy AutoScaler

	decoder := gob.NewDecoder(&buf) // buf implements io.Reader's Read()
	decoder.Decode(&copy)

	return copy
}

func (a *AutoScaler) UpdateStatus() AutoScaler {
	var copy AutoScaler
	switch a.Mode {
	case Serialization:
		copy = DeepCopyWithSerialization(a)
		// case Function: not implemented as serialization is better
	}

	copy.Status.Current = 27
	copy.Status.Desired = 27
	copy.Status.Available = 27

	return copy
}
