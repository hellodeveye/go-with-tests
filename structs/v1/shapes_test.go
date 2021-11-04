package shapes

import "testing"

func assertEquals(got, want float64, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	assertEquals(got, want, t)
}
