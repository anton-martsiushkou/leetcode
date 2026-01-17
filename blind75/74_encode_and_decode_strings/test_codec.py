import pytest
from codec import Codec


def test_example_1():
    """Test case from example 1."""
    codec = Codec()
    strs = ["Hello", "World"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_example_2_empty_string():
    """Test case from example 2 - empty string."""
    codec = Codec()
    strs = [""]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_empty_list():
    """Test with empty list."""
    codec = Codec()
    strs = []
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_single_string():
    """Test with single string."""
    codec = Codec()
    strs = ["test"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_string_with_delimiter():
    """Test strings containing the delimiter character."""
    codec = Codec()
    strs = ["#", "##", "a#b"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_string_with_numbers():
    """Test strings with numbers."""
    codec = Codec()
    strs = ["123", "456"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_mixed_strings():
    """Test mixed strings including empty ones."""
    codec = Codec()
    strs = ["abc", "", "def", "ghij"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_special_characters():
    """Test strings with special characters."""
    codec = Codec()
    strs = ["!@#$%", "^&*()", "{}[]"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_long_strings():
    """Test with long strings."""
    codec = Codec()
    strs = ["This is a long string with many words", "Another long string"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_unicode_characters():
    """Test with unicode characters."""
    codec = Codec()
    strs = ["Hello 世界", "🎉🎊"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_empty_string_in_middle():
    """Test with empty string in the middle."""
    codec = Codec()
    strs = ["before", "", "after"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_multiple_empty_strings():
    """Test with multiple empty strings."""
    codec = Codec()
    strs = ["", "", ""]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


def test_strings_that_look_like_encoding():
    """Test strings that look like the encoding format."""
    codec = Codec()
    strs = ["5#hello", "10#world"]
    encoded = codec.encode(strs)
    decoded = codec.decode(encoded)
    assert decoded == strs


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
