package pdf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"immotep/backend/services/pdf"
)

func TestPDF(t *testing.T) {
	pdf.Test = true
	p := pdf.NewPDF()

	t.Run("GetWidth", func(t *testing.T) {
		assert.InDelta(t, 190, p.GetWidth(), 0.001)
	})

	t.Run("NoTestNeeded", func(_ *testing.T) {
		p.Ln(10)
		p.AddCenteredTitle("Test", pdf.H1)
		p.AddTitle("Test", pdf.H2)
		p.AddText("Test")
		p.Add2Texts("Test", "Test")
		p.AddMultiLineText("Test")
		p.AddLine()
		p.AddImages(nil)
	})

	t.Run("Output", func(t *testing.T) {
		output, err := p.Output()
		require.NoError(t, err)
		assert.NotEmpty(t, output)
	})
}
