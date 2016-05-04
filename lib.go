package villa2

import "fmt"

type S struct {
	D int
	Q string
}

func (s *S) Do() {
	fmt.Println("S.Do")
}
