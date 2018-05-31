package tensorcv

import (
	"encoding/json"
	"io/ioutil"
)

// LoadLabels takes a JSON file which contains the 1000 classes from ImageNet and returns a map
// from class index to class description, e.g. 1 => "goldfish, Carassiu auratus"
func LoadLabels(jsonPath string) (map[int]string, error) {

	labelMap := make(map[int]string)

	bytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &labelMap)
	if err != nil {
		return nil, err
	}

	return labelMap, nil
}
