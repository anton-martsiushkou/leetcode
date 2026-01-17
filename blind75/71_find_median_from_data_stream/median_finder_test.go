package median_finder

import (
	"math"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(1)
		mf.AddNum(2)
		got := mf.FindMedian()
		want := 1.5
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("FindMedian() = %v, want %v", got, want)
		}

		mf.AddNum(3)
		got = mf.FindMedian()
		want = 2.0
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("FindMedian() = %v, want %v", got, want)
		}
	})

	t.Run("single element", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(5)
		got := mf.FindMedian()
		want := 5.0
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("FindMedian() = %v, want %v", got, want)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(-1)
		mf.AddNum(-2)
		mf.AddNum(-3)
		got := mf.FindMedian()
		want := -2.0
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("FindMedian() = %v, want %v", got, want)
		}
	})

	t.Run("mixed numbers", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(5)
		mf.AddNum(15)
		mf.AddNum(1)
		mf.AddNum(3)
		got := mf.FindMedian()
		want := 4.0 // (3 + 5) / 2
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("FindMedian() = %v, want %v", got, want)
		}
	})

	t.Run("many elements", func(t *testing.T) {
		mf := Constructor()
		nums := []int{12, 10, 13, 11, 5, 15, 1, 11, 6, 12}
		for _, num := range nums {
			mf.AddNum(num)
		}
		got := mf.FindMedian()
		want := 11.0 // (11 + 11) / 2
		if math.Abs(got-want) > 1e-9 {
			t.Errorf("FindMedian() = %v, want %v", got, want)
		}
	})
}
