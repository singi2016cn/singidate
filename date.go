package singidate

import (
	"time"
)

// timeFormat 格式化标准字符串
const timeFormat = "2006-01-02 15:04:05"

// zone 时区
const zone = "CST"

// zoneOffsetUnit 时区偏移单位，秒
const zoneOffsetUnit = 3600

//Timestamp2DateTime 时间戳转时间格式
func Timestamp2DateTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(timeFormat)
}

//DateTime2TimestampWithZoneOffset 时间格式转时间戳,根据偏移量计算对应时区
func DateTime2TimestampWithZoneOffset(dateTime string, zoneOffset int) int64 {
	var cstEastZone8 = time.FixedZone(zone, zoneOffset*zoneOffsetUnit)
	t, _ := time.ParseInLocation(timeFormat, dateTime, cstEastZone8)
	return t.Unix()
}

//DateTime2TimestampWithoutZone 时间格式转时间戳,0时区
func DateTime2TimestampWithoutZone(dateTime string) int64 {
	t, _ := time.Parse(timeFormat, dateTime)
	return t.Unix()
}

//DateTime2TimestampWithShanghai 时间格式转时间戳,上海时区
func DateTime2TimestampWithShanghai(dateTime string) int64 {
	return DateTime2TimestampWithZoneOffset(dateTime, 8)
}
