package structs

import "testing"

func TestPerimeter(t *testing.T) {
	ractangle := Rectangle{10.0, 10.0}
	got := ractangle.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12.0, 60.0}, 72.0},
		{"circle", Circle{10}, 314.1592653589793},
		{"triangle", Triangle{12, 6}, 36.0},
	}

	for _, s := range areaTests {
		t.Run(s.name, func(t *testing.T) {
			checkArea(t, s.shape, s.want)
		})
	}
}
