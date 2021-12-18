package behavioural

import (
	"testing"

	capture "github.com/alok87/go-capture"
	"github.com/stretchr/testify/assert"
)

func TestTempalteMethod(t *testing.T) {
	type test struct {
		provider CareProvider
		expected string
	}

	tests := []test{
		{
			provider: NewShalbyHospital(),
			expected: "shalby book consult shalby consult shalby book surgery shalby operate shalby recover",
		},
		{
			provider: NewPractoCareSurgery(),
			expected: "practo book consult practo consult practo book surgery practo operate practo recover",
		},
	}

	for _, tc := range tests {
		tc = tc
		out := capture.Output(func() {
			SurgeryWorkflowTemplate(tc.provider)
		})
		assert.Equal(t, tc.expected, out)
	}
}
