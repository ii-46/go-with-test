package iteration

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expected := 15
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestSumAll(t *testing.T) {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5}
	got := SumAll(nums1, nums2)
	want := []int{6, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v given (%v %v)", want, got, nums1, nums2)
	}
}

func TestSumAllTail(t *testing.T) {
	t.Run("make sum of some slice", func(t *testing.T) {
		got := SumAllTail([]int{1, 4}, []int{2, 9})
		want := []int{4, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
