package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("获得 %.2f 预期 %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	got := rectangle.Area()
	// 	want := 72.0

	// 	if got != want {
	// 		t.Errorf("获得 %.2f 预期 %.2f", got, want)
	// 	}
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	got := circle.Area()
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("获得 %.2f 预期 %.2f", got, want)
	// 	}
	// })

	// =============================================
	// 重构 ， 引入 接口
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area() // 接口
	// 	if got != want {
	// 		t.Errorf("获得 %.2f 预期 %。2f", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12, 6}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circles := Circle{10}
	// 	checkArea(t, circles, 314.1592653589793)
	// })

	// =============================================
	// 表格驱动
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 2, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			// tt.shape 结构体中域的值，可以一眼看出被测试的属性
			t.Errorf("%#v 获得 %.2f 预期 %.2f", tt.shape, got, tt.want)
		}
	}
}
