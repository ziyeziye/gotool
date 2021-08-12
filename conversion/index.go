package conversion

import "strconv"

// Conversion 类型转换结构定义
type Conversion struct {
}

// StrToInt 字符串转int类型，当存在defaultNum时，出现异常会返回设置的默认值，若不出现异常正常返回
func (c Conversion) StrToInt(strNum string, defaultNum ...int) int {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		} else {
			return 0
		}
	}
	return num
}

// StrToInt64 StrNum需要转换的字符串
//defaultNum默认值
//String类型转int64
func (c Conversion) StrToInt64(strNum string, defaultNum ...int64) int64 {
	num, err := strconv.ParseInt(strNum, 10, 64)
	if err != nil {
		if len(defaultNum) > 0 {
			return defaultNum[0]
		} else {
			return 0
		}
	}
	return num
}

// StrToFloat32 String转float32位
//strFloat 需要转的字符串
//defaultFloat 默认值，若没有出现转换异常直接返回0.0
func (c Conversion) StrToFloat32(strFloat string, defaultFloat ...float32) float32 {
	float, err := strconv.ParseFloat(strFloat, 32)
	if err != nil {
		if len(defaultFloat) > 0 {
			return defaultFloat[0]
		} else {
			return 0.0
		}
	}
	return float32(float)
}

// StrToFloat64 String转float64位
//strFloat 需要转的字符串
//defaultFloat 默认值，若没有出现转换异常直接返回0.0
func (c Conversion) StrToFloat64(strFloat string, defaultFloat ...float64) float64 {
	float, err := strconv.ParseFloat(strFloat, 36)
	if err != nil {
		if len(defaultFloat) > 0 {
			return defaultFloat[0]
		} else {
			return 0.0
		}
	}
	return float
}

// IntToStr int类型转字符串
//num int类型的数据，不区分int64或者int类型
func (c Conversion) IntToStr(num interface{}) string {
	var str string
	switch num.(type) {
	case int:
		str = strconv.Itoa(num.(int))
	case int64:
		str = strconv.FormatInt(num.(int64), 10)
	}
	return str
}
