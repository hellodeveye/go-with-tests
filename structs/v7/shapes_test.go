package shapes

import (
	"fmt"
	"testing"
)

func assertEquals(shape Shape, got, want float64, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got %.2f want %.2f", shape, got, want)
	}
}
func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	assertEquals(rectangle, got, want, t)
}

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{Rectangle{12, 6}, 72.0},
			{Circle{10}, 314.1592653589793},
			{Triangle{12, 6}, 36.0},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			fmt.Println(tt.shape)
			assertEquals(tt.shape, got, tt.want, t)
		}
	})
}
