package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	var file string
	flag.StringVar(&file, "file", "", "the yaml input file to convert to JSON")

	flag.Parse()

	if flag.NArg() == 0 {
		printErr("Wrong command line, it s missing the command argument")
		printErr("\nPlease use")
		printErr(" mzr <command> -arg value...")
		printErr("\nAvailable command are:")
		printErr(" tojson -file <yaml file>")
		os.Exit(1)
	}

	cmd := flag.Arg(0)

	if cmd == "tojson" {
		err := convertToJSON(file)
		if err != nil {
			printErr("The program failed to convert the file %q to JSON", file)
			printErr("\nPlease check the following error message before proceeding further:")
			printErr("\n%v", err)
			os.Exit(1)
		}
	} else {
		printErr("Wrong command line, %q is not a command this program provides", cmd)
		printErr("\nAvailable command are:")
		printErr(" tojson -file <yaml file>")
		os.Exit(1)
	}

	os.Exit(0)
}

func printErr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

type employee struct {
	Name   string   `yaml:"name"`
	Job    string   `yaml:"job"`
	Skills []string `yaml:"skills"`
}

func convertToJSON(input string) error {

	content, err := ioutil.ReadFile(input)
	if err != nil {
		return fmt.Errorf("The file could not be read, the reason is: %v", err.Error())
	}

	m := make([]map[string]employee, 0)
	err = yaml.Unmarshal(content, &m)
	if err != nil {
		return fmt.Errorf("The file could not be parsed as YAML, the reason is: %v", err.Error())
	}

	err = json.NewEncoder(os.Stdout).Encode(m)

	if err != nil {
		return fmt.Errorf("The program failed to encode the decoced YAML to JSON, the reason is: %v", err.Error())
	}

	return nil
}
