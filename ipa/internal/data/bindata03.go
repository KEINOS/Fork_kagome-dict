package data

import(
	"os"
	"time"
)


func dictIpadictAdBytes() ([]byte, error) {
	return _dictIpadictAd, nil
}

func dictIpadictAd() (*asset, error) {
	bytes, err := dictIpadictAdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/ipadict.ad", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1600787007, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
