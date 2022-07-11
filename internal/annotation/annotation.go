package annotation

import "fmt"

type Annotation string

func (t Annotation) String() string {
	return string(t)
}

func (t Annotation) Tag() string {
	return fmt.Sprintf(`name:"%s"`, t.String())
}
