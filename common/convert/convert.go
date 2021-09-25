package convert

import "strconv"

func StringToInt64(in string) (int64, error) {
	tem, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return 0, err
	}
	return tem, nil
}

func Int64ToString(in int64) string {
	return strconv.FormatInt(in, 10)
}

func StringToInt(in string) (int, error) {
	tem, err := strconv.Atoi(in)
	if err != nil {
		return 0, err
	}
	return tem, nil
}

func IntToString(in int) string {
	return strconv.Itoa(in)
}

func StringPtr(in string) *string {
	return &in
}
func Int64Ptr(in int64) *int64 {
	return &in
}
func IntPtr(in int) *int {
	return &in
}
func Int32Ptr(in int32) *int32 {
	return &in
}

func Float64Ptr(in float64) *float64 {
	return &in
}

func BoolPtr(in bool) *bool {
	return &in
}
