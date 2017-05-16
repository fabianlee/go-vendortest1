package main

import (
	"fmt"
	"github.com/fabianlee/go-myutil/myutil"
)

func main() {
	fmt.Println("Version: ",myutil.GetBranch())
}
