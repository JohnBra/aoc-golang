package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

type QItem struct {
	r, c int32
	seq  string
}

type Move struct {
	r, c, k int32 // new r, new c, directional rune
}

type Mem struct {
	path string
	len  int
}

func getSequences(keypad [][]rune) map[[2]rune][]string {
	pos := map[rune][2]int32{}

	for r, keys := range keypad {
		for c, key := range keys {
			if key != '_' {
				pos[key] = [2]int32{int32(r), int32(c)}
			}
		}
	}

	seqs := map[[2]rune][]string{}

	for a := range pos {
		for b := range pos {
			sKey := [2]rune{a, b}
			if a == b {
				seqs[sKey] = []string{"A"}
				continue
			}

			possibilities := []string{}
			p := pos[a]
			q := ds.NewDeque([]QItem{{p[0], p[1], ""}})
			optimal := math.MaxInt

		outer:
			for q.Len() > 0 {
				cur, _ := q.PopFront()
				moves := []Move{
					{cur.r - 1, cur.c, '^'},
					{cur.r + 1, cur.c, 'v'},
					{cur.r, cur.c - 1, '<'},
					{cur.r, cur.c + 1, '>'},
				}
				for _, n := range moves {
					if n.r < 0 || n.c < 0 || n.r == int32(len(keypad)) || n.c == int32(len(keypad[0])) {
						continue
					}

					if keypad[n.r][n.c] == '_' {
						continue
					}

					if keypad[n.r][n.c] == b {
						if optimal < len(cur.seq)+1 {
							break outer
						}

						optimal = len(cur.seq) + 1
						possibilities = append(possibilities, cur.seq+string(n.k)+"A")
					} else {
						q.PushBack(QItem{n.r, n.c, cur.seq + string(n.k)})
					}
				}
			}
			seqs[sKey] = possibilities
		}
	}

	return seqs
}

func getPathOptions(seqs map[[2]rune][]string, code string) []string {
	options := [][]string{}
	for _, t := range utils.ZipMerge([]rune("A"+code), []rune(code)) {
		options = append(options, seqs[[2]rune{t[0], t[1]}])
	}

	return sliceProduct(options...)
}

func complexity(code string, len int, re *regexp.Regexp) int {
	snum := re.FindAllString(string(code), -1)
	num, err := strconv.Atoi(snum[0])
	utils.Check(err)

	return num * len
}

// creates cartesian product of input and returns
// joined paths of robot as slice of strings
func sliceProduct(args ...[]string) []string {
	pools := args
	npools := len(pools)
	indices := make([]int, npools)
	result := make([]string, npools)

	for i := range result {
		if len(pools[i]) == 0 {
			return nil
		}
		result[i] = pools[i][0]
	}

	results := [][]string{result}

	for {
		i := npools - 1
		for ; i >= 0; i -= 1 {
			pool := pools[i]
			indices[i] += 1

			if indices[i] == len(pool) {
				indices[i] = 0
				result[i] = pool[0]
			} else {
				result[i] = pool[indices[i]]
				break
			}

		}

		if i < 0 {
			res := []string{}
			for _, s := range results {
				res = append(res, strings.Join(s, ""))
			}
			return res
		}

		newresult := make([]string, npools)
		copy(newresult, result)
		results = append(results, newresult)
	}
}

func pathLenAtDepth(
	memo map[Mem]int,
	dlens map[[2]rune]int,
	dseqs map[[2]rune][]string,
	path string,
	depth int,
) int {
	if _, ok := memo[Mem{path, depth}]; ok {
		return memo[Mem{path, depth}]
	}

	subpaths := utils.ZipMerge([]rune("A"+path), []rune(path))

	if depth == 1 {
		len := 0
		for _, subpath := range subpaths {
			len += dlens[[2]rune{subpath[0], subpath[1]}]
		}
		return len
	}

	minLen := 0
	for _, subpath := range subpaths {
		seqMinLen := math.MaxInt
		for _, seq := range dseqs[[2]rune{subpath[0], subpath[1]}] {
			seqMinLen = utils.Min(seqMinLen, pathLenAtDepth(memo, dlens, dseqs, seq, depth-1))
		}
		minLen += seqMinLen
	}
	memo[Mem{path, depth}] = minLen

	return minLen
}

func solve(input []string, nseqs, dseqs map[[2]rune][]string) (int, int) {
	p1, p2 := 0, 0
	re := regexp.MustCompile(`\d+`)

	dlens := map[[2]rune]int{}
	for rtuple, seqs := range dseqs {
		dlens[rtuple] = len(seqs[0])
	}

	memo := map[Mem]int{}

	for _, code := range input {
		paths := getPathOptions(nseqs, code)
		optimalP1, optimalP2 := math.MaxInt, math.MaxInt

		for _, path := range paths {
			optimalP1 = utils.Min(optimalP1, pathLenAtDepth(memo, dlens, dseqs, path, 2))
			optimalP2 = utils.Min(optimalP2, pathLenAtDepth(memo, dlens, dseqs, path, 25))
		}

		p1 += complexity(code, optimalP1, re)
		p2 += complexity(code, optimalP2, re)
	}

	return p1, p2
}

func parseInput(filepath string) ([]string, error) {
	codes := []string{}

	f, err := os.Open(filepath)
	if err != nil {
		return codes, err
	}
	defer f.Close()

	b := new(strings.Builder)
	io.Copy(b, f)

	codes = strings.Split(strings.TrimSuffix(b.String(), "\n"), "\n")
	return codes, nil
}

func main() {
	input, err := parseInput(utils.GetPuzzleInputSrc())
	utils.Check(err)

	npad := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{'_', '0', 'A'},
	}

	dpad := [][]rune{
		{'_', '^', 'A'},
		{'<', 'v', '>'},
	}

	nseqs := getSequences(npad)
	dseqs := getSequences(dpad)

	partOneRes, partTwoRes := solve(input, nseqs, dseqs)
	fmt.Println("Part one res", partOneRes)
	fmt.Println("Part two res", partTwoRes)
}
