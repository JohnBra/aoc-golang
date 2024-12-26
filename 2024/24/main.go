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
	w1, w2, op, out string
}

func getFrozen(reverseLookup map[[3]string]string, key [3]string) (string, bool) {
	if val, ok := reverseLookup[key]; ok {
		return val, ok
	}

	val, ok := reverseLookup[[3]string{key[1], key[0], key[2]}]
	return val, ok
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

func getGateLookups(gates []Gate) (map[string]Gate, map[[3]string]string) {
	lookup, reverseLookup := map[string]Gate{}, map[[3]string]string{}

	for _, g := range gates {
		key := [3]string{g.w1, g.w2, g.op}
		reverseLookup[key] = g.out
		lookup[g.out] = g
	}

	return lookup, reverseLookup
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

func swap(a, b string, pairs *[][2]string, gates *[]Gate) {
	*pairs = append(*pairs, [2]string{a, b})
	for i, gate := range *gates {
		if gate.out == a || gate.out == b {
			if a != gate.out {
				(*gates)[i] = Gate{gate.w1, gate.w2, gate.op, a}
			} else {
				(*gates)[i] = Gate{gate.w1, gate.w2, gate.op, b}
			}
		}
	}
}

func processBitPosition(
	xi, yi, zi string,
	lookup map[string]Gate,
	reverseLookup map[[3]string]string,
	carry string,
	pairs *[][2]string,
	gates *[]Gate,
) (bool, string) {
	bit, _ := getFrozen(reverseLookup, [3]string{xi, yi, "XOR"})
	adder, ok := getFrozen(reverseLookup, [3]string{bit, carry, "XOR"})
	ncarry := carry

	if ok {
		c1, _ := getFrozen(reverseLookup, [3]string{xi, yi, "AND"})
		c2, _ := getFrozen(reverseLookup, [3]string{bit, carry, "AND"})
		ncarry, _ = getFrozen(reverseLookup, [3]string{c1, c2, "OR"})
	} else {
		gate := lookup[zi]
		if gate.w1 != ncarry {
			swap(bit, gate.w1, pairs, gates)
		} else {
			swap(bit, gate.w2, pairs, gates)
		}
		return false, ncarry
	}

	if adder != zi {
		swap(adder, zi, pairs, gates)
		return false, ncarry
	}
	return true, ncarry
}

func processGateSwaps(
	zs int,
	lookup map[string]Gate,
	reverseLookup map[[3]string]string,
	pairs *[][2]string,
	gates *[]Gate,
) {
	carry, _ := getFrozen(reverseLookup, [3]string{"x00", "y00", "AND"})

	for i := 1; i < zs; i++ {
		xi := fmt.Sprintf("x%02d", i)
		yi := fmt.Sprintf("y%02d", i)
		zi := fmt.Sprintf("z%02d", i)

		ok := false
		ok, carry = processBitPosition(xi, yi, zi, lookup, reverseLookup, carry, pairs, gates)

		if !ok {
			return
		}
	}
}

func partTwo(gates []Gate) string {
	pairs := &[][2]string{}
	ngates := &[]Gate{}
	zs := 0

	for _, gate := range gates {
		*ngates = append(*ngates, gate)
		if strings.HasPrefix(gate.out, "z") {
			zs += 1
		}
	}

	for len(*pairs) < 4 {
		lookup, reverseLookup := getGateLookups(*ngates)
		processGateSwaps(zs, lookup, reverseLookup, pairs, ngates)
	}

	res := []string{}
	for _, p := range *pairs {
		res = append(res, p[0], p[1])
	}
	slices.Sort(res)

	return strings.Join(res[:], ",")
}

func parseInput(filepath string) (map[string]bool, []Gate, error) {
	wires := map[string]bool{}
	gates := []Gate{}

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
		gates = append(gates, Gate{w1: g[0], w2: g[2], op: g[1], out: g[4]})
	}

	return wires, gates, nil
}

func main() {
	wires, gates, err := parseInput(utils.GetPuzzleInputSrc())
	utils.Check(err)
	lookup, _ := getGateLookups(gates)

	partOneRes := partOne(wires, lookup)
	fmt.Println("Part one res", partOneRes)

	partTwoRes := partTwo(gates)
	fmt.Println("Part two res", partTwoRes)
}
