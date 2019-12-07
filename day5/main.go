package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp := getList()

	pos := 0
	for pos < len(inp) {
		s := "00000"
		s = s[len(inp[pos]):] + inp[pos]
		s = s[:5]

		if s[3:] == "99" {
			break
		}

		pm1 := s[2]
		pm2 := s[1]

		switch s[3:] {
		case "01", "02", "07", "08":
			// 3 parameters

			p1, _ := strconv.Atoi(inp[pos+1])
			p2, _ := strconv.Atoi(inp[pos+2])
			p3, _ := strconv.Atoi(inp[pos+3])

			pos1, pos2 := 0, 0

			if pm1 == '1' {
				pos1 = p1
			} else {
				pos1, _ = strconv.Atoi(inp[p1])
			}
			if pm2 == '1' {
				pos2 = p2
			} else {
				pos2, _ = strconv.Atoi(inp[p2])
			}

			if s[3:] == "01" {
				inp[p3] = strconv.Itoa(pos1 + pos2)
			} else if s[3:] == "02" {
				inp[p3] = strconv.Itoa(pos1 * pos2)
			} else if s[3:] == "07" {
				if pos1 < pos2 {
					inp[p3] = "1"
				} else {
					inp[p3] = "0"
				}
			} else if s[3:] == "08" {
				if pos1 == pos2 {
					inp[p3] = "1"
				} else {
					inp[p3] = "0"
				}
			}
		case "05", "06":
			// 2 parameters

			p1, _ := strconv.Atoi(inp[pos+1])
			p2, _ := strconv.Atoi(inp[pos+2])

			pos1 := 0
			pos2 := 0

			if pm1 == '1' {
				pos1 = p1
			} else {
				pos1, _ = strconv.Atoi(inp[p1])
			}
			if pm2 == '1' {
				pos2 = p2
			} else {
				pos2, _ = strconv.Atoi(inp[p2])
			}

			if s[3:] == "05" {
				if pos1 != 0 {
					pos = pos2
					continue
				}
			} else if s[3:] == "06" {
				if pos1 == 0 {
					pos = pos2
					continue
				}
			}
			pos += 3
			continue
		case "03", "04":
			// 1 parameter

			p1, _ := strconv.Atoi(inp[pos+1])
			pos1 := 0
			if pm1 == '1' {
				pos1 = p1
			} else {
				pos1, _ = strconv.Atoi(inp[p1])
			}

			if s[3:] == "03" {

				input := ""
				fmt.Scan(&input)
				inp[p1] = input
			} else if s[3:] == "04" {
				fmt.Println(pos1)
			}
			pos += 2
			continue
		default:
			break
		}
		pos += 4
	}
}

func getList() []string {

	f, _ := os.Open("./input.txt")
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	v := strings.Split(string(b), ",")

	return v
}
