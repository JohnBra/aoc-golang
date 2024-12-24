package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

type Gate struct {
	w1, w2, op string
}

func getBitVal(wires map[string]bool, gates map[string]Gate, bit string) bool {
	if v, ok := wires[bit]; ok {
		return v
	}

	g := gates[bit]
	switch g.op {
	case "AND":
		return getBitVal(wires, gates, g.w1) && getBitVal(wires, gates, g.w2)
	case "OR":
		return getBitVal(wires, gates, g.w1) || getBitVal(wires, gates, g.w2)
	case "XOR":
		return getBitVal(wires, gates, g.w1) != getBitVal(wires, gates, g.w2)
	}

	panic(fmt.Errorf("logical operator not found for %v", g))
}

func partOne(wires map[string]bool, gates map[string]Gate) int {
	res := 0
	bits := []string{}

	for k := range gates {
		if strings.HasPrefix(k, "z") {
			bits = append(bits, k)
		}
	}

	slices.Sort(bits)

	for i, bit := range bits {
		if getBitVal(wires, gates, bit) {
			res += (1 << i)
		} else {
			res += (0 << i)
		}
	}

	return res
}

func parseInput(filepath string) (map[string]bool, map[string]Gate, error) {
	wires := map[string]bool{}
	gates := map[string]Gate{}

	f, err := os.Open(filepath)
	if err != nil {
		return wires, gates, err
	}
	defer f.Close()

	b := new(strings.Builder)
	io.Copy(b, f)

	strs := strings.Split(strings.TrimSuffix(b.String(), "\n"), "\n\n")
	re := regexp.MustCompile(`([a-z]\d{2})|(\s\d+)`)
	for _, v := range strings.Split(strs[0], "\n") {
		w := re.FindAllString(v, -1)

		val, err := strconv.ParseBool(strings.TrimSpace(w[1]))
		if err != nil {
			return wires, gates, err
		}
		wires[w[0]] = val
	}

	for _, v := range strings.Split(strs[1], "\n") {
		g := strings.Split(v, " ")
		gates[g[4]] = Gate{w1: g[0], w2: g[2], op: g[1]}
	}

	return wires, gates, nil
}

func main() {
	w, g, err := parseInput(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := partOne(w, g)
	fmt.Println("Part one res", partOneRes)
}
