package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var fon = &cobra.Command{
	Use:   "fon",
	Short: "json flattener",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Validation(args)
	},
}

func Validation(args []string) error {
	Pipe()
	if len(args) == 0 {
		return nil
	}
	return Args(args)
}

func Pipe() error {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var buf []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			buf = append(buf, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		return unmarshal(buf)
	}
	return nil
}

func Args(args []string) error {
	var input []byte
	fileInfo, err := os.Stat(args[0])
	if err == nil {
		if fileInfo.Mode().IsRegular() {
			file, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}
			input = file
		} else {
			input = []byte(args[0])
		}
	} else {
		input = []byte(args[0])
	}
	return unmarshal(input)
}

func unmarshal(input []byte) error {
	jsonObj := make(map[string]interface{})
	err := json.Unmarshal(input, &jsonObj)
	if err != nil {
		return fmt.Errorf("invalid json! %v", err)
	}
	flattenJSON(jsonObj)
	return nil
}

func Execute() {
	if err := fon.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

func flattenJSON(jsonObj map[string]interface{}) map[string]interface{} {
	for k, v := range jsonObj {
		flatten(k, v)
	}
	return nil
}

func flatten(key string, jsonObj interface{}) {
	switch jsonObj.(type) {
	case map[string]interface{}:
		for k, v := range jsonObj.(map[string]interface{}) {
			newkey := fmt.Sprintf("%s.%s", key, k)
			flatten(newkey, v)
		}
	case []interface{}:
		for i, v := range jsonObj.([]interface{}) {
			newkey := fmt.Sprintf("%s[%d]", key, i)
			flatten(newkey, v)
		}
	case int8:
		fmt.Printf("%s: %d", key, jsonObj.(int))
	case int16:
		fmt.Printf("%s: %d", key, jsonObj.(int))
	case int32:
		fmt.Printf("%s: %d", key, jsonObj.(int))
	case int64:
		fmt.Printf("%s: %d", key, jsonObj.(int))
	case float32:
		fmt.Printf("%s: %f\n", key, jsonObj.(float32))
	case float64:
		fmt.Printf("%s: %f\n", key, jsonObj.(float64))
	case []byte:
		fmt.Printf("%s: %s\n", key, string(jsonObj.([]byte)))
	case string:
		fmt.Printf("%s: %s\n", key, jsonObj.(string))
	default:
		fmt.Println(key, jsonObj.(string))
	}
}
