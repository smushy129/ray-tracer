package renders

import "os"

func outputPPM(ppm string, name string) {
	f, err := os.Create("./renders/" + name)
	defer f.Close()

	if err != nil {
		panic(err)
	}

	f.WriteString(ppm)
}
