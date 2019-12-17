package renders

import "os"

var rootDir = "./"
var renderDir = "./renders/"

func outputPPM(ppm string, name string) {
	f, err := os.Create(rootDir + name)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	f.WriteString(ppm)
}
