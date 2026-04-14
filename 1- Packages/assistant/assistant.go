package assistant

import "fmt"


func Writing(write string){
	fmt.Println("Assistant is saying : ",write)
	nonExportFunc()
	ExportedFunc()
}