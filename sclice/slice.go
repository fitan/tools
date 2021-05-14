package sclice

import (
	"github.com/emirpasic/gods/maps/hashmap"
)

type KF func(v interface{}) (interface{}, error)

func L2L(l []interface{}, kf KF) ([]interface{}, error) {
	tmpL := make([]interface{}, 0, len(l))
	for _, v := range l {
		k, err := kf(v)
		if err != nil {
			return nil, err
		}
		tmpL = append(tmpL, k)
	}
	return tmpL, nil
}

// L2M slice []interface to map map[f()]interface
func L2M(l []interface{}, kf KF) (*hashmap.Map, error) {
	m := hashmap.New()
	for _, v := range l {
		tmpK, err := kf(v)
		if err != nil {
			return nil, err
		}
		m.Put(tmpK, v)
	}
	return m, nil
}

// Intersect l1 和 l2 的交集
func Intersect(l1, l2 []interface{}, kf KF) ([]interface{}, error) {
	m1, err := L2M(l1, kf)
	if err != nil {
		return nil, err
	}
	m2, err := L2M(l2, kf)
	if err != nil {
		return nil, err
	}
	l := make([]interface{}, 0, 0)
	for _, k := range m1.Keys() {
		if v, ok := m2.Get(k); ok {
			l = append(l, v)
		}
	}
	return l, nil
}

// Minus l1 和 l2的差集 l1-l2
func Minus(l1, l2 []interface{}, kf KF) ([]interface{}, error) {
	m1, err := L2M(l1, kf)
	if err != nil {
		return nil, err
	}
	m2, err := L2M(l2, kf)
	if err != nil {
		return nil, err
	}
	for _, k := range m2.Keys() {
		m1.Remove(k)
	}
	return m1.Values(), nil
}

// 迭代[]string 传入string的指针  修改会修改slice中的值
func EachStringPtr(l *[]string, f func(ptr *string)) {
	for i, _ := range *l {
		f(&(*l)[i])
	}
}

// 迭代[]string 传入string
func EachString(l []string, f func(s string)) {
	for _, s := range l {
		f(s)
	}
}
