package code

import (
	"fmt"
	"testing"
)

func Test_FindString(t *testing.T) {
	html := `{"I":"5333","V":"马经理"},`
	Linkman := FindString(`{"I":"5333","V":"(?P<value>[^"]+)"}`, html, "value")
	fmt.Println(Linkman)
}
