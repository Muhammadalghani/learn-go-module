package learngomodule

import "fmt"

func PrintSliceOfString( slc []string){
	for _, v := range slc {
		fmt.Println(v)
	}
}

// git tag v1.0.0
// git push origin v1.0.0 

// go mod tidy -> untuk mengimpor nya 
// atau bisa juga dengan 
// go get link_githubnya

/*
1 create repo di github (github.com/Muhammadalghani/learn-go-module)
2. inisiasi go mod dengan domain github repository
3. push
4 step setelah push :
	- create tagging 
		- git tag v1.0.0
		- git push origin v1.0.0
5. consume package 
	- go get -u github.com/Muhammadalghani/learn-go-module
*/ 