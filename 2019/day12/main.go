package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	m     = make(map[string]int)
	s     = 0
	found = false
)

func main() {

	inp := getMoons()

	for !found {
		inp = step(inp)
		s++
	}
}

func step(moons []Moon) []Moon {
	nm := make([]Moon, len(moons))

	for i, m := range moons {
		nms := m
		nm[i] = nms
	}

	for i := 0; i < len(nm); i++ {
		for j := 0; j < len(nm); j++ {
			// fix x velocity
			if nm[i].X < nm[j].X {
				nm[i].velX++
			} else if nm[i].X > nm[j].X {
				nm[i].velX--
			}

			// fix y velocity
			if nm[i].Y < nm[j].Y {
				nm[i].velY++
			} else if nm[i].Y > nm[j].Y {
				nm[i].velY--
			}

			// fix z velocity
			if nm[i].Z < nm[j].Z {
				nm[i].velZ++
			} else if nm[i].Z > nm[j].Z {
				nm[i].velZ--
			}
		}
	}

	for i := 0; i < len(nm); i++ {
		nm[i].X += nm[i].velX
		nm[i].Y += nm[i].velY
		nm[i].Z += nm[i].velZ
	}

	if s == 999 {
		ans := 0
		for i := 0; i < len(nm); i++ {
			nm[i].pot += int(math.Abs(float64(nm[i].X)) + math.Abs(float64(nm[i].Y)) + math.Abs(float64(nm[i].Z)))
			nm[i].kin += int(math.Abs(float64(nm[i].velX)) + math.Abs(float64(nm[i].velY)) + math.Abs(float64(nm[i].velZ)))
			ans += nm[i].pot * nm[i].kin
		}
		fmt.Println("part 1:", ans)
	}

	if nm[0].velX == 0 && nm[0].velY == 0 && nm[0].velZ == 0 && nm[1].velX == 0 && nm[1].velY == 0 && nm[1].velZ == 0 && nm[2].velX == 0 && nm[2].velY == 0 && nm[2].velZ == 0 && nm[3].velX == 0 && nm[3].velY == 0 && nm[3].velZ == 0 {
		fmt.Println(s)
	}

	k := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v",
		nm[0].X, nm[0].Y, nm[0].Z, nm[0].velX, nm[0].velY, nm[0].velZ,
		nm[1].X, nm[1].Y, nm[1].Z, nm[1].velX, nm[1].velY, nm[1].velZ,
		nm[2].X, nm[2].Y, nm[2].Z, nm[2].velX, nm[2].velY, nm[2].velZ,
		nm[3].X, nm[3].Y, nm[3].Z, nm[3].velX, nm[3].velY, nm[3].velZ,
	)

	if !found {
		_, ok := m[k]
		if ok {
			fmt.Println("part 2:", s)
			found = true
		}
	}

	if s == 0 {
		m[k] = s
	}

	return nm
}

func getMoons() []Moon {
	f, _ := os.Open("./input.txt")
	defer f.Close()

	ms := []Moon{}
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		s := strings.Trim(scan.Text(), "<>")
		s2 := strings.Split(s, ", ")
		vars := make(map[string]int)
		for _, w := range s2 {
			s3 := strings.Split(w, "=")
			n, _ := strconv.Atoi(s3[1])
			vars[s3[0]] = n
		}
		m := Moon{X: vars["x"], Y: vars["y"], Z: vars["z"]}
		ms = append(ms, m)
	}
	return ms
}

type Moon struct {
	X    int
	Y    int
	Z    int
	velX int
	velY int
	velZ int
	pot  int
	kin  int
}
