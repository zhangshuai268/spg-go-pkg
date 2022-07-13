package spg_go_pkg

// ParamMap 用于包内参数传递
type ParamMap map[string]interface{}

// Set 设置参数
func (pm ParamMap) Set(key string, value interface{}) ParamMap {
	pm[key] = value
	return pm
}
