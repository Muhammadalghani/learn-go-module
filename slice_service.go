package learngomodule

import "fmt"

func PrintSliceOfString( slc []string){
	for _, v := range slc {
		fmt.Println(v)
	}
}