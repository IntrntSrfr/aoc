package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	passports := getPassports(true)
	start := time.Now()
	validatePassports(passports)
	fmt.Println("\nTime taken:", time.Since(start))
}

func validatePassports(passports []map[string]string) {
	good1 := 0
	good2 := 0
	for _, pass := range passports {
		if !validPassport(pass) {
			continue
		}

		good1++

		if !validBYR(pass["byr"]) {
			continue
		}
		if !validIYR(pass["iyr"]) {
			continue
		}
		if !validEYR(pass["eyr"]) {
			continue
		}
		if !validHGT(pass["hgt"]) {
			continue
		}
		if !validHCL(pass["hcl"]) {
			continue
		}
		if !validECL(pass["ecl"]) {
			continue
		}
		if !validPID(pass["pid"]) {
			continue
		}

		good2++
	}
	fmt.Println(fmt.Sprintf("Valid passports\n\nPart 1: %v valid passports\n\nPart 2: %v valid passports", good1, good2))
}

func getPassports(chung bool) []map[string]string {

	var data []byte
	if chung {
		data, _ = ioutil.ReadFile("./chungus.txt")
	} else {
		data, _ = ioutil.ReadFile("./input.txt")
	}

	lines := strings.Split(string(data), "\n\n")

	var pps []map[string]string

	for _, line := range lines {
		m := make(map[string]string)
		for _, part := range strings.Fields(line) {
			kv := strings.Split(part, ":")
			m[kv[0]] = kv[1]
		}
		pps = append(pps, m)
	}
	return pps
}

func validPassport(pass map[string]string) bool {
	_, cidFound := pass["cid"]
	return cidFound && len(pass) == 8 || !cidFound && len(pass) == 7
}
func validBYR(inp string) bool {
	year, err := strconv.Atoi(inp)
	return err == nil && year >= 1920 && year <= 2002
}
func validIYR(inp string) bool {
	year, err := strconv.Atoi(inp)
	return err == nil && year >= 2010 && year <= 2020
}
func validEYR(inp string) bool {
	year, err := strconv.Atoi(inp)
	return err == nil && year >= 2020 && year <= 2030
}
func validHGT(inp string) bool {
	if len(inp) < 4 {
		return false
	}
	val, units := inp[:len(inp)-2], inp[len(inp)-2:]
	n, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if units == "cm" {
		return 150 <= n && n <= 193
	} else if units == "in" {
		return 59 <= n && n <= 76
	} else {
		return false
	}
}

var (
	hclReg = regexp.MustCompile("^#[0-9a-f]{6}$")
	eclReg = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	pidReg = regexp.MustCompile("^\\d{9}$")
)

func validHCL(inp string) bool {
	return hclReg.MatchString(inp)
}
func validECL(inp string) bool {
	return eclReg.MatchString(inp)
}
func validPID(inp string) bool {
	return pidReg.MatchString(inp)
}
