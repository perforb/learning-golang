package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Temp file name: /var/folders/mp/mfphy1md3y5dwc184k9xxmd00000gn/T/sample2586306116
// Temp dir name: /var/folders/mp/mfphy1md3y5dwc184k9xxmd00000gn/T/sampledir3336164410
func main() {
	f, err := os.CreateTemp("", "sample")
	check(err)
	fmt.Println("Temp file name:", f.Name())
	defer func(name string) {
		_ = os.Remove(name)
	}(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
