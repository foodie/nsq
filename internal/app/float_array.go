package app

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

//设置浮点数数组
type FloatArray []float64

//把用,划分的字符串转换成float array
func (a *FloatArray) Set(param string) error {
	for _, s := range strings.Split(param, ",") {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Fatalf("Could not parse: %s", s)
			return nil
		}
		*a = append(*a, v)
	}
	sort.Sort(*a)
	return nil
}

//便于排序
func (a FloatArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FloatArray) Less(i, j int) bool { return a[i] > a[j] }
func (a FloatArray) Len() int           { return len(a) }

//打印浮点数数组
func (a *FloatArray) String() string {
	var s []string
	for _, v := range *a {
		s = append(s, fmt.Sprintf("%f", v))
	}
	return strings.Join(s, ",")
}
