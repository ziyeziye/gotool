package test

import (
	"fmt"
	"github.com/ziyeziye/gotool"
	"testing"
)

func TestPage(t *testing.T) {
	paginator := gotool.PageUtils.Paginator(5, 20, 500)
	fmt.Println(paginator)
}
