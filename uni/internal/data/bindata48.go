package data

import(
	"os"
	"time"
)


func dictUnidictBwBytes() ([]byte, error) {
	return _dictUnidictBw, nil
}

func dictUnidictBw() (*asset, error) {
	bytes, err := dictUnidictBwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600786977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
