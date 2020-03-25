// site project site.go
package site

import (
	"fmt"
)

type site struct {
	Name string
	Path string
}

func (s site) String() string { return fmt.Sprintf("Name: %q (%s)", s.Name, s.Path) }
