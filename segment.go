package segment

import (
	"fmt"
	"sort"
	"strings"
)

type Segment[T comparable] struct {
	start, end int64
	value      T
}

type Segments[T comparable] []Segment[T]

func (s Segment[T]) String() string {
	return fmt.Sprintf("{%d~%d:%v}", s.start, s.end, s.value)
}

func (s Segment[T]) Start() int64 { return s.start }
func (s Segment[T]) End() int64   { return s.end }
func (s Segment[T]) Value() T     { return s.value }

func (ss Segments[T]) String() string {
	var output []string
	for _, s := range ss {
		output = append(output, s.String())
	}
	return strings.Join(output, ", ")
}

func New[T comparable](start, end int64, value T) (Segment[T], error) {
	if end < start {
		return Segment[T]{}, fmt.Errorf("end < start: nil segment returned")
	}
	return Segment[T]{start, end, value}, nil
}

func Must[T comparable](start, end int64, value T) Segment[T] {
	s, err := New(start, end, value)
	if err != nil {
		panic(err)
	}
	return s
}

//Cover 把segment覆盖在segments中
func Cover[T comparable](ss Segments[T], in Segment[T]) Segments[T] {
	result := make(Segments[T], 0, len(ss))
	var done bool
	idx := sort.Search(len(ss), func(i int) bool {
		return ss[i].end >= in.start
	})
	result = append(result, ss[:idx]...)
	for i, v := range ss[idx:] {
		if in.start <= v.end && !done {
			if in.start > v.start {
				result = append(result, Must(v.start, in.start-1, v.value))
			}
			result = append(result, Must(in.start, in.end, in.value))
			done = true
		}
		if in.end < v.end {
			result = append(result, Must(in.end+1, v.end, v.value))
			result = append(result, ss[idx+i+1:]...)
			break
		}
	}
	return result
}

//Merge 合并segments(如果value相同)
func Merge[T comparable](ss ...Segment[T]) Segments[T] {
	if len(ss) < 1 {
		return ss
	}
	ssSorted := append(Segments[T]{}, ss...)
	sort.Slice(ssSorted, func(i, j int) bool { return ssSorted[i].start < ssSorted[j].start })
	var output Segments[T]
	last := ssSorted[0]
	for _, v := range ssSorted[1:] {
		if v.value == last.value && v.start == last.end+1 {
			last.end = v.end
			continue
		}
		output = append(output, last)
		last = v
	}
	output = append(output, last)
	return output
}

//continuity  s在v前面返回-1 在后面返回1
func (s Segment[T]) continuity(v Segment[T]) int {
	if s.end == v.start-1 {
		return -1
	}
	if s.start == v.end+1 {
		return 1
	}
	return 0
}

//Continuity 判断ss是否连续
func Continuity[T comparable](ss ...Segment[T]) bool {
	if len(ss) < 2 {
		return true
	}
	last := ss[0]
	for _, v := range ss[1:] {
		if 1 != v.continuity(last) {
			return false
		}
		last = v
	}
	return true
}
