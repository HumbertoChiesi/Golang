package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		ret := Rectangle{10, 15}
		expectedArea := float64(150)
		areaReceived := ret.area()
		if expectedArea != areaReceived {
			t.Errorf("Expected area: %0.2f   Area received: %0.2f",
				expectedArea,
				areaReceived,
			)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		cir := Circle{5}
		expectedArea := float64(math.Pi * (cir.radius * cir.radius))
		areaReceived := cir.area()
		if expectedArea != areaReceived {
			t.Errorf("Expected area: %0.2f   Area received: %0.2f",
				expectedArea,
				areaReceived,
			)
		}
	})
}
