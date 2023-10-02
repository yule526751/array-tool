package array

type IntUintFloat interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func In[T IntUintFloat | string](val T, arr []T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// 数组转map
func ToMap[T any](arr []T) map[any]struct{} {
	m := make(map[any]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	return m
}

// 获取数组中最大值
func Max[T IntUintFloat](arr []T) (n T) {
	for k := range arr {
		if n < arr[k] {
			n = arr[k]
		}
	}
	return
}

// 获取数组中最小值
func Min[T IntUintFloat](arr []T) (n T) {
	for _, v := range arr {
		if n > v {
			n = v
		}
	}
	return
}

// 数组去重
func RemoveRepeated[T IntUintFloat | string](arr []T) []T {
	result := make([]T, 0, len(arr))
	temp := map[T]struct{}{}
	for _, item := range arr {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// 去除0值
func ArrRemoveZero[T IntUintFloat](arr []T) []T {
	result := make([]T, 0, len(arr))
	for _, item := range arr {
		if item != 0 {
			result = append(result, item)
		}
	}
	return result
}

// 去除空字符串
func RemoveEmpty(arr []string) []string {
	result := make([]string, 0, len(arr))
	for _, item := range arr {
		if item != "" {
			result = append(result, item)
		}
	}
	return result
}
