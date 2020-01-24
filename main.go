package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	tag := flag.Arg(0)
	path := flag.Arg(1)
	p, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Println(err)
	}
	conf, err := readOnConfig(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", conf)
	for _, v := range conf.Images {
		v.NewTag = tag
	}
	yaml, err := yaml.Marshal(&conf)
	fmt.Println(string(yaml))
	fmt.Println(err)
	err = ioutil.WriteFile(p, yaml, 0644)
	fmt.Println(err)
}

func readOnConfig(fileBuffer []byte) (*Kustomization, error) {
	data := Kustomization{}
	err := yaml.Unmarshal(fileBuffer, &data)
	if err != nil {
		fmt.Println(err)
		return &data, err
	}
	return &data, nil
}
