[![](https://pkg.go.dev/badge/github.com/lixiangzhong/segment)](https://pkg.go.dev/github.com/lixiangzhong/segment)

# Segment

> 操作数字区段

```go
package main

import (
	"fmt"

	"github.com/lixiangzhong/segment"
)

func main() {
	s1 := segment.Must(1, 10, "value1")
	s2 := segment.Must(11, 20, "value1")
	s3 := segment.Must(21, 25, "value3")

	ss := segment.Merge(s1, s2, s3)
	fmt.Println(ss)
	//{1~20:value1}, {21~25:value3}

	fmt.Println(segment.Continuity(s1, s2, s3))
	//true

	fmt.Println(segment.Cover(ss, segment.Must(5, 15, "valueCover")))
	//{1~4:value1}, {5~15:valueCover}, {16~20:value1}, {21~25:value3}
}

```

> 更多例子 example/

> 参考 github.com/google/segment/segment
