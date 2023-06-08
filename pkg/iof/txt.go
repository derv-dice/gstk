package iof

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"os"
	"strings"
)

// ReadFileLineByLine - Прочитать текстовый файл построчно в []string
// Флаг unique=true отбросит повторяющиеся строки
func ReadFileLineByLine(filename string, unique bool) (lines []string, err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return
	}

	defer func() {
		_ = file.Close()
	}()

	s := bufio.NewScanner(file)

	var uniqMap map[string]bool
	if unique {
		uniqMap = map[string]bool{}
	}

	for s.Scan() {
		if unique {
			if strings.TrimSpace(s.Text()) == "" {
				continue
			}

			hash := hashFromString(s.Text())

			if uniqMap[hash] {
				continue
			}

			uniqMap[hash] = true
		}

		lines = append(lines, s.Text())
	}

	err = s.Err()
	return
}

func hashFromString(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
