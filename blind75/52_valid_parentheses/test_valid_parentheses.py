import pytest
from valid_parentheses import is_valid


def test_example_1():
    """Test case from example 1."""
    s = "()"
    result = is_valid(s)
    assert result is True


def test_example_2():
    """Test case from example 2."""
    s = "()[]{}"
    result = is_valid(s)
    assert result is True


def test_example_3():
    """Test case from example 3."""
    s = "(]"
    result = is_valid(s)
    assert result is False


def test_example_4():
    """Test case from example 4."""
    s = "([])"
    result = is_valid(s)
    assert result is True


def test_nested_brackets():
    """Test with nested brackets."""
    s = "{[()]}"
    result = is_valid(s)
    assert result is True


def test_mismatched_brackets():
    """Test with mismatched brackets."""
    s = "([)]"
    result = is_valid(s)
    assert result is False


def test_only_opening_brackets():
    """Test with only opening brackets."""
    s = "(((("
    result = is_valid(s)
    assert result is False


def test_only_closing_brackets():
    """Test with only closing brackets."""
    s = "))))"
    result = is_valid(s)
    assert result is False


def test_empty_string():
    """Test with empty string."""
    s = ""
    result = is_valid(s)
    assert result is True


def test_single_opening_bracket():
    """Test with single opening bracket."""
    s = "("
    result = is_valid(s)
    assert result is False


def test_single_closing_bracket():
    """Test with single closing bracket."""
    s = ")"
    result = is_valid(s)
    assert result is False


def test_complex_valid():
    """Test with complex valid pattern."""
    s = "{[()()]}(){}"
    result = is_valid(s)
    assert result is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
