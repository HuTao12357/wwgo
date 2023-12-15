package utils

// 实现有序map
type OrderMap struct {
	key   []string
	value map[string]interface{}
}

func NewOrderMap() *OrderMap {
	return &OrderMap{key: make([]string, 0), value: make(map[string]interface{})}
}
func (o *OrderMap) SetMap(key string, value interface{}) {
	o.value[key] = value
	o.key = append(o.key, key)
}
func (o *OrderMap) GetMap(key string) interface{} {
	return o.value[key]
}
func removeDuplicates(arr []string) []string {
	//可以使用一个 map 来记录切片中的元素是否已经出现过，遍历切片时，如果元素尚未出现在 map 中，则将其添加到新的切片中
	noMap := make(map[string]bool)
	res := make([]string, 0)
	for _, v := range arr {
		if noMap[v] == false {
			noMap[v] = true
			res = append(res, v)
		}
	}
	return res
}
