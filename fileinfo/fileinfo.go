package fileinfo

import (
	"fmt"
	"log"
	"os"
)
var (
	fileinfo os.FileInfo
	err      error

)





//this is a function to read files, it takes any kind of string and reads it.
func FileReader(name string) {
	args := os.Args
	// kun filnavn
	if len(args) == 2 {
		//produces error if args length is 2.
		feilmelding := "Du mangler -f"
		panic(feilmelding)
	}
	//attaches the the string "-f" to run func printStats that take args array position 3 as in.
	if len(args) == 3 {
		if args[1] == "-f" {
			printStats(args[2])
		}
	}
}

func printStats(name string) {
	fileinfo, err = os.Lstat(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file name:", fileinfo.Name())
	fmt.Println("Size in bytes", fileinfo.Size())
	bytes := fileinfo.Size()
	//gjør om til float før casting
	kibibytes := (float64) (bytes / 1024) //casting to float64
	mibibytes := (float64)(kibibytes/ 1024)
	gibibytes := (float64) (mibibytes / 1024)
	fmt.Println("KibiBytes:",kibibytes)
	fmt.Println("MibiBytes:",mibibytes)
	fmt.Println("GibiBytes:",gibibytes)
	fmt.Println("Permissions:", fileinfo.Mode())
	fmt.Println("Last modified:", fileinfo.ModTime())
	fmt.Println("Is a directory:", fileinfo.IsDir())
	fmt.Println("Is a regular file:",fileinfo.Mode().IsRegular())
	//mt.Printf("System Interface type: %T\n", fileinfo.Sys())
	//fmt.Printf("System info: %+v\n\n", fileinfo.Sys())
	//hvis blokk device, da er det ikke char device.
	x bool := false
	if fileinfo.Mode()&os.ModeAppend != 0{
		fmt.Println("Is not append only.")

	}else {
		fmt.Println("File is append only")
	}
	if fileinfo.Mode()&os.ModeDevice == 0{
		fmt.Println("File is ", fileinfo.Mode()&os.ModeDevice)
	}else{
		fmt.Println()
	}
}


