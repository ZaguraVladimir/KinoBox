// config project conf.go
package conf

import (
	"log"
	"testing"
)

func TestRead(t *testing.T) {
	log.Println(*Read())
}
