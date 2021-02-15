package humanity

import {
	"fmt"
	"SoftwareGoDay1/data"
	"io/ioutil"
	"strconv"
}

type Human struct {
	Name string
	Age int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	humAge, err := strconv.Atoi(csv[1])
	if err != nil
		return nil, fmt.Errorf("Bad parameter, try int")
	hum := Human{csv[0], humAge, csv[2]}
	return &hum, err
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	myData, err := data.ReadFile(path)
	if err != nil
		return nil,err
	lenData := len(myData)
	hum := make([]*Human, lenData)
	for i := 0; i < lenData; i++ {
		line, err := data.LineToCSV(myData[i])
		if err != nil {
			nil, err
		}
		hum[i], err = NewHumanFromCSV(line)
		fmt.Println(line)
		if err != nil {
			return nil, err
		}
	}
	return hum, err
}

func (h *Human) String() string {
	return fmt.Sprintf("%s, %d, %s\n", h.Name, h.Age, h.Country)
}