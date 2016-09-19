package gocrc48

import (
	"testing"
)

func TestCRC48(t *testing.T) {
	str := `29234:C 19 Sep 20:21:52.777 * DB saved on disk
29234:C 19 Sep 20:21:52.778 * RDB: 8 MB of memory used by copy-on-write
4868:M 19 Sep 20:21:52.844 * Background saving terminated with success
4868:M 19 Sep 20:36:53.036 * 1 changes in 900 seconds. Saving...
4868:M 19 Sep 20:36:53.038 * Background saving started by pid 29383
29706:C 19 Sep 21:06:55.030 * RDB: 8 MB of memory used by copy-on-write
4868:M 19 Sep 21:06:55.106 * Background saving terminated with success`
	expected := int64(0x6b11055dbb63)
	crc := CRC48([]byte(str))
	if crc != expected {
		t.Errorf(" CRC48(%s) failed! expected [0x%x], but got [0x%x]", str, expected, crc)
	}
}
