//+build !jsoniter-sloppy

package jsoniter

import "fmt"

func (iter *Iterator) skipNumber() {
	if !iter.trySkipNumber() {
		iter.unreadByte()
		iter.ReadFloat32()
	}
}

func (iter *Iterator) trySkipNumber() bool {
	dotFound := false
	for i := iter.head; i < iter.tail; i++ {
		c := iter.buf[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if dotFound {
				iter.ReportError("validateNumber", `more than one dot found in number`)
				return true // already failed
			} else {
				dotFound = true
			}
		default:
			switch c {
			case ',', ']', '}', ' ', '\t', '\n', '\r':
				iter.head = i
				return true // must be valid
			}
			return false // may be invalid
		}
	}
	return false
}

func (iter *Iterator) skipString() {
	if !iter.trySkipString() {
		iter.unreadByte()
		iter.ReadString()
	}
}

func (iter *Iterator) trySkipString() bool {
	for i := iter.head; i < iter.tail; i++ {
		c := iter.buf[i]
		if c == '"' {
			iter.head = i + 1
			return true // valid
		} else if c == '\\' {
			return false
		} else if c < ' ' {
			iter.ReportError("ReadString",
				fmt.Sprintf(`invalid control character found: %d`, c))
			return true // already failed
		}
	}
	return false
}

func (iter *Iterator) skipObject() {
}

func (iter *Iterator) skipArray() {
}