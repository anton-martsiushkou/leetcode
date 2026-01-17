import pytest
from encode_and_decode_strings import encode, decode


def test_example_1():
    """Test case from example 1 - basic strings."""
    strs = ["Hello", "World"]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_example_2():
    """Test case from example 2 - empty string."""
    strs = [""]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_example_3():
    """Test case from example 3 - empty array."""
    strs = []
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_example_4():
    """Test case from example 4 - strings with special characters."""
    strs = ["a#b", "c", "de#f"]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_single_string():
    """Test with a single string."""
    strs = ["test"]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_multiple_empty_strings():
    """Test with multiple empty strings."""
    strs = ["", "", ""]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_strings_with_spaces():
    """Test with strings containing spaces."""
    strs = ["hello world", "foo bar"]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_strings_with_numbers():
    """Test with strings containing numbers and special chars."""
    strs = ["123", "456#789"]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_long_strings():
    """Test with long strings."""
    strs = ["this is a very long string with many characters", "another long one"]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_unicode_characters():
    """Test with unicode characters and emojis."""
    strs = ["Hello 世界", "🚀 emoji"]
    encoded = encode(strs)
    decoded = decode(encoded)
    assert decoded == strs


def test_encode_format():
    """Test that encode produces the expected format."""
    strs = ["Hello", "World"]
    encoded = encode(strs)
    assert encoded == "5#Hello5#World"


def test_decode_specific_string():
    """Test decode with a specific encoded string."""
    encoded = "5#Hello5#World"
    decoded = decode(encoded)
    assert decoded == ["Hello", "World"]


def test_encode_empty_string():
    """Test encoding a list with an empty string."""
    strs = [""]
    encoded = encode(strs)
    assert encoded == "0#"


def test_decode_empty_encoded():
    """Test decoding an empty string."""
    encoded = ""
    decoded = decode(encoded)
    assert decoded == []


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
