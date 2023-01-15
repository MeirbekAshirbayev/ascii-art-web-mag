package main

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"io"
	"log"
	"os"
)

const (
	errstr1 = "File was modified!"
	errstr2 = "Check the text for ascii char!"
	errstr3 = "Check the banner type!"
)

func AsciiForWeb(text, banner string) (string, error) {
	if !CheckStringForAscii(text) {
		return "", errors.New(errstr2)
	}

	asciiMap, err := MakeAsciiMap(banner + ".txt")
	if err != nil {
		return "", errors.New(errstr1)
	}

	str := Printing(text, asciiMap)
	return str, nil
}

func ConverAscii(str string, m map[rune][]string) string {
	str1 := ""
	for i := 0; i < 8; i++ {
		for _, v := range str {
			s1 := m[v]
			str1 += s1[i]
		}
		if i == 7 {
			continue
		}
		str1 += "\n"
	}

	return str1
}

func CheckStringForAscii(s string) bool {
	for _, v := range s {
		if !(v >= ' ' && v <= '~') {
			if v != '\n' && v != '\r' {
				return false
			}
		}
	}
	return true
}

func PrepareForPrinting(str1 string) []string {
	w := []rune(str1)
	str := ""
	s := []string{}
	b := false
	for i := 0; i < len(w); i++ {
		if b {
			b = false
			continue
		}

		if i != len(w)-1 && w[i] == '\\' && w[i+1] == 'n' {
			b = true

			s = append(s, str, "\n")
			str = ""
		} else if w[i] == '\n' || (w[i] == '\r' && w[i+1] == '\n') {
			s = append(s, str, "\n")
			str = ""
		} else {
			str += string(w[i])
		}
	}

	s = append(s, str)

	return s
}

func Printing(s string, m map[rune][]string) string {
	str := ""
	slice := PrepareForPrinting(s)
	for _, v := range slice {
		if v == "\n" {
			str += "\n"
		} else if v == "" {
			continue
		} else {
			str += ConverAscii(v, m)
		}
	}
	if len(slice) > 1 && slice[len(slice)-1] == "" {
		return str
	}
	str += "\n"
	return str
}

func MakeAsciiMap(str string) (map[rune][]string, error) {
	file, err := os.Open("./web/" + str)
	if err != nil {
		return nil, errors.New(errstr2)
	}
	strSlice := []string{}
	fileScaner := bufio.NewScanner(file)
	fileScaner.Split(bufio.ScanLines)

	for fileScaner.Scan() {
		text := fileScaner.Text()
		strSlice = append(strSlice, text)
	}

	if err = file.Close(); err != nil {
		return nil, errors.New("incorrect txt file")
	}
	var ind rune
	ind = 32
	checkermap := make(map[string][]byte)
	checkermap["shadow.txt"] = []byte{38, 185, 77, 11, 19, 75, 119, 233, 253, 35, 224, 54, 11, 253, 129, 116, 15, 128, 251, 127, 101, 65, 209, 216, 197, 216, 94, 115, 238, 85, 15, 115}
	checkermap["standard.txt"] = []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}
	checkermap["thinkertoy.txt"] = []byte{100, 40, 94, 73, 96, 209, 153, 244, 129, 147, 35, 196, 220, 99, 25, 186, 52, 241, 240, 221, 157, 161, 77, 7, 17, 19, 69, 245, 215, 108, 63, 163}
	hashStr := CheckForHash("./web/" + str)
	if string(hashStr) != string(checkermap[str]) {
		return nil, errors.New(errstr1)
	}
	s := []string{}
	asciiMap := make(map[rune][]string)

	for _, v := range strSlice {
		if v == "" && len(s) != 0 {
			asciiMap[ind] = s
			ind++
			s = []string{}
		} else {
			if v == "" {
				continue
			}
			s = append(s, v)
		}
	}
	asciiMap[ind] = s
	return asciiMap, nil
}

func CheckForHash(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return h.Sum(nil)
}
