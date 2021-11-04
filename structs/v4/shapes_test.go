package shapes

import "testing"

func assertEquals(got, want float64, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0
	assertEquals(got, want, t)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0
		assertEquals(got, want, t)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		assertEquals(got, want, t)
	})
}
