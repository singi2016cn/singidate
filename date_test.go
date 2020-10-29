package singidate

import (
	"testing"
)

func TestTimestamp2DateTime(t *testing.T) {
	actualVal := Timestamp2DateTime(1586501854)
	exceptVal := "2020-04-10 14:57:34"
	if actualVal != exceptVal {
		t.Errorf(`%s != %s`, actualVal, exceptVal)
	}
}

func TestDateTime2TimestampWithShanghai(t *testing.T) {
	actualVal := DateTime2TimestampWithShanghai("2020-10-27 00:00:00")
	exceptVal := int64(1603728000)
	if actualVal != exceptVal {
		t.Errorf(`%d != %d`, actualVal, exceptVal)
	}
}
