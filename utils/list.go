package utils

import "container/list"

func Get(l *list.List, index int) *list.Element {
	if nil == l || l.Len() == 0 {
		return nil
	}

	i := 0

	// Front返回链表第一个元素或nil。
	// Next返回链表的后一个元素或者nil。
	// 遍历 list
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		if i == index {
			return iter
		}
		i++
	}

	return nil
}

/**
list.Element 结构。

type Element struct {
    // 元素保管的值
    Value interface{}
    // 内含隐藏或非导出字段
}
 */
func IndexOf(l *list.List, value interface{}) int {
	i := 0
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		if iter.Value == value {
			return i
		}
		i++
	}
	return -1
}

func Remove(l *list.List, value interface{}) {
	var e *list.Element
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		if iter.Value == value {
			e = iter
		}

	}

	if nil != e {
		l.Remove(e)
	}

}

func ToStringArray(l *list.List) []string {
	if nil == l {
		return nil
	}

	values := make([]string, l.Len())

	i := 0
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		// 取值。
		s, _ := iter.Value.(string)
		values[i] = s
		i++
	}

	return values

}
