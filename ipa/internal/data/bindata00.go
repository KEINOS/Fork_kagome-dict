package data

import(
	"os"
	"time"
)


func dictIpadictAaBytes() ([]byte, error) {
	return _dictIpadictAa, nil
}

func dictIpadictAa() (*asset, error) {
	bytes, err := dictIpadictAaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.aa", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600787007, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
