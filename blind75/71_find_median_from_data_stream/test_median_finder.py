import pytest
from median_finder import MedianFinder


def test_example_1():
    """Test case from example 1."""
    mf = MedianFinder()
    mf.addNum(1)
    mf.addNum(2)
    assert abs(mf.findMedian() - 1.5) < 1e-9

    mf.addNum(3)
    assert abs(mf.findMedian() - 2.0) < 1e-9


def test_single_element():
    """Test with single element."""
    mf = MedianFinder()
    mf.addNum(5)
    assert abs(mf.findMedian() - 5.0) < 1e-9


def test_negative_numbers():
    """Test with negative numbers."""
    mf = MedianFinder()
    mf.addNum(-1)
    mf.addNum(-2)
    mf.addNum(-3)
    assert abs(mf.findMedian() - (-2.0)) < 1e-9


def test_mixed_numbers():
    """Test with mixed positive and negative numbers."""
    mf = MedianFinder()
    mf.addNum(5)
    mf.addNum(15)
    mf.addNum(1)
    mf.addNum(3)
    assert abs(mf.findMedian() - 4.0) < 1e-9  # (3 + 5) / 2


def test_many_elements():
    """Test with many elements."""
    mf = MedianFinder()
    nums = [12, 10, 13, 11, 5, 15, 1, 11, 6, 12]
    for num in nums:
        mf.addNum(num)
    assert abs(mf.findMedian() - 11.0) < 1e-9  # (11 + 11) / 2


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
