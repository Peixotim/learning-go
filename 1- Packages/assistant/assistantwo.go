package assistant

import "fmt"

//First capital letter does not export
func nonExportFunc(){
	fmt.Println("Is function not exported")
}

//First capital letter exports outside the package
func ExportedFunc(){
	fmt.Println("Is function exported")
}