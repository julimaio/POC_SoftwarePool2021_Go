package data

import {
	"fmt"
	"io/ioutil"
	"errors"
	"strings"
}

func LineToCSV(line string) ([]string, error) {
	myData := sring.Split(line, ",")
	if len(myData) != 3
		return nil, fmt.Errorf("Bad format")
	return myData, nil
}

func ReadFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil;
		return nil, err
	myData := strings.Split(string(data), "\n")
	return myData, err
}
