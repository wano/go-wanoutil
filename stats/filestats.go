package stats

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
)

func GetFileSizeAndMd5Sum(filepath string) (fileSizeByte int64, md5hash string, err error) {
	const filechunk = 8192
	file, err := os.Open(filepath)
	if err != nil {
		return 0, "", err
	}
	defer func() {
		_ = file.Close()
	}()
	// calculate the file size
	info, err := file.Stat()
	if err != nil {
		return 0, "", err
	}
	fileSize := info.Size()
	blocks := uint64(math.Ceil(float64(fileSize) / float64(filechunk)))
	hash := md5.New()
	for i := uint64(0); i < blocks; i++ {
		blockSize := int(math.Min(filechunk, float64(fileSize-int64(i*filechunk))))
		buf := make([]byte, blockSize)
		_  , err = file.Read(buf)
		if err != nil {
			return 0 , "" , err
		}

		_ ,  err = io.WriteString(hash, string(buf)) // append into the hash
		if err != nil {
			return 0 , "" , err
		}
	}
	checkSumString := fmt.Sprintf("%x", hash.Sum(nil))

	return fileSize, checkSumString, nil
}