//go:build mage

package main

import (
	"fmt"
	"strings"

	"github.com/bmayfi3ld/advent-of-code/utils/wrapper"
)

func FourA() error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("FourA")()

	// test answer 18
	// matrix := parseInput(testInputFour)
	matrix := parseInput(inputFour)

	found := 0

	for row := range matrix {
		for col := range matrix[row] {
			// check if first spot is X
			if matrix[row][col] != "X" {
				continue
			}

			found += checkRight(row, col, matrix)
			found += checkLeft(row, col, matrix)
			found += checkUp(row, col, matrix)
			found += checkDown(row, col, matrix)
			found += checkDUpLeft(row, col, matrix)
			found += checkDUpRight(row, col, matrix)
			found += checkDDownLeft(row, col, matrix)
			found += checkDDownRight(row, col, matrix)
		}
	}

	fmt.Println(found)

	return nil
}

func parseInput(in string) [][]string {
	out := [][]string{}

	rows := strings.Split(strings.TrimSpace(in), "\n")

	for r, row := range rows {
		out = append(out, []string{})
		for _, point := range row {
			out[r] = append(out[r], string(point))
		}
	}

	return out
}

func checkRight(row, col int, matrix [][]string) int {
	// check if can fit
	if col+3 >= len(matrix[row]) {
		return 0
	}

	if matrix[row][col+1] != "M" {
		return 0
	}

	if matrix[row][col+2] != "A" {
		return 0
	}

	if matrix[row][col+3] != "S" {
		return 0
	}

	return 1
}

func checkLeft(row, col int, matrix [][]string) int {
	// check if can fit
	if col-3 < 0 {
		return 0
	}

	if matrix[row][col-1] != "M" {
		return 0
	}

	if matrix[row][col-2] != "A" {
		return 0
	}

	if matrix[row][col-3] != "S" {
		return 0
	}

	return 1
}

func checkUp(row, col int, matrix [][]string) int {
	// check if can fit
	if row-3 < 0 {
		return 0
	}

	if matrix[row-1][col] != "M" {
		return 0
	}

	if matrix[row-2][col] != "A" {
		return 0
	}

	if matrix[row-3][col] != "S" {
		return 0
	}

	return 1
}

func checkDown(row, col int, matrix [][]string) int {
	// check if can fit
	if row+3 >= len(matrix) {
		return 0
	}

	if matrix[row+1][col] != "M" {
		return 0
	}

	if matrix[row+2][col] != "A" {
		return 0
	}

	if matrix[row+3][col] != "S" {
		return 0
	}

	return 1
}

func checkDUpLeft(row, col int, matrix [][]string) int {
	// check if can fit
	if row-3 < 0 {
		return 0
	}
	if col-3 < 0 {
		return 0
	}

	if matrix[row-1][col-1] != "M" {
		return 0
	}

	if matrix[row-2][col-2] != "A" {
		return 0
	}

	if matrix[row-3][col-3] != "S" {
		return 0
	}

	return 1
}

func checkDUpRight(row, col int, matrix [][]string) int {
	// check if can fit
	if row-3 < 0 {
		return 0
	}
	if col+3 >= len(matrix[row]) {
		return 0
	}

	if matrix[row-1][col+1] != "M" {
		return 0
	}

	if matrix[row-2][col+2] != "A" {
		return 0
	}

	if matrix[row-3][col+3] != "S" {
		return 0
	}

	return 1
}

func checkDDownLeft(row, col int, matrix [][]string) int {
	// check if can fit
	if row+3 >= len(matrix) {
		return 0
	}
	if col-3 < 0 {
		return 0
	}

	if matrix[row+1][col-1] != "M" {
		return 0
	}

	if matrix[row+2][col-2] != "A" {
		return 0
	}

	if matrix[row+3][col-3] != "S" {
		return 0
	}

	return 1
}

func checkDDownRight(row, col int, matrix [][]string) int {
	// check if can fit
	if col+3 >= len(matrix[row]) {
		return 0
	}
	if row+3 >= len(matrix) {
		return 0
	}

	if matrix[row+1][col+1] != "M" {
		return 0
	}

	if matrix[row+2][col+2] != "A" {
		return 0
	}

	if matrix[row+3][col+3] != "S" {
		return 0
	}

	return 1
}

func FourB() error {
	fmt.Println("hello")
	defer wrapper.ProfileFunction("FourA")()

	// test answer 9
	// matrix := parseInput(testInputFour)
	matrix := parseInput(inputFour)

	found := 0

	for row := range matrix {
		for col := range matrix[row] {
			// check if first spot is X
			if matrix[row][col] != "A" {
				continue
			}

			// because it is a consistent box shape, can skip oob checks
			// (probably could have above now that I've slept)
			// and just trim the total area
			// also it is a square
			if row == 0 || col == 0 || row == len(matrix)-1 || col == len(matrix[row])-1 {
				continue
			}



			// need to find 2 for the X
			maybeFound := 0
			maybeFound += checkXDOneOClock(row, col, matrix)
			maybeFound += checkXDFiveOClock(row, col, matrix)
			maybeFound += checkXDSevenOClock(row, col, matrix)
			maybeFound += checkXDElevenOClock(row, col, matrix)

			if maybeFound >= 2 {
				found++
			}
		}
	}

// 	M.S
//  .A.
//  M.S

	fmt.Println(found)

	return nil
}

func checkXDOneOClock(row, col int, matrix [][]string) int {
	if matrix[row-1][col+1] != "M" {
		return 0
	}

	if matrix[row+1][col-1] != "S" {
		return 0
	}

	return 1
}

func checkXDFiveOClock(row, col int, matrix [][]string) int {
	if matrix[row+1][col+1] != "M" {
		return 0
	}
	if matrix[row-1][col-1] != "S" {
		return 0
	}

	return 1
}

func checkXDSevenOClock(row, col int, matrix [][]string) int {
	if matrix[row+1][col-1] != "M" {
		return 0
	}
	if matrix[row-1][col+1] != "S" {
		return 0
	}

	return 1
}

func checkXDElevenOClock(row, col int, matrix [][]string) int {
	if matrix[row-1][col-1] != "M" {
		return 0
	}
	if matrix[row+1][col+1] != "S" {
		return 0
	}

	return 1
}

const testInputFour = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`


const inputFour = `
SSSMMSAMXSSSSSSMSSSSMAMSMMSMSMXSASMMMMAMXAXMAXXSSMSSSMSMMSXMAXXMAXSAMXMXMAXXMAAMMMMMAASXMSAMXMASMMSMSMSXXMSMSXAXMSMMSXASXSMMSMMXMMMMXMXAMMSX
XAAAASMSAAAASAAASAAXSAAAMMXAAAASAMXAMMSSSSSMXSAXSAAAASMAAXASMSMAMXAMASXMMSMSMMXSMAASAMXAXXXMAMXSXMSMSASMXAAAMMMAMXAASMXSAAMASASAMSAMAMXMSAMS
MSMMMSAMMMMMMSMMMMMMSMSMSAASMXMMAMSXSAAAXAAAAMAMMMMSMMXMMSXMAAAAXMXSAAAMAXXAAAASXSMXMASMMMSSSMAXAXSAMAMAMSMSMAMSMMMMMAXMMMMXSAMAMSASMSAAMXMA
XAAMXMAMXXAAMASXXXXAXMXXAAMXSAXMAMAAMMMXMSMMMMMMXMXMASXMMXMMSMSMSMAXMAMMASMSMMXSAMMSMAAXAAAAAMSSSMMAMMMAMXAMMXMAAAAAXSMMSXMMMXMXMSAMAMMMMAXM
MXMAXMAMMMSXSASMMAMMSMSAMMMASXMSXMMSMAMSXAXXAAAAASXAXAAAASAAXXXAXMASXSXAXXAAXXAMASAAMSMSXMMSXMAMASMMMAMAMMSMMXSSMMXMXMAXSASAAAMAMMSMMMSASXSX
AAMASMSSMAMMMAMMMASAAAMAXXMAMSASAMMMMMSAMXMSSSSSSSMMMSSMMAMXMMMMMSXSAMASMMSMMMXSAMMSMXMAXXMMMMASAMXMSXSASXMAMXXAXSASMSSMXASAMASXSAAXAASAXAAX
SASAXXMAMAXAMXMXSASMSMSXMXMAMAXSAMMSSMAXXMMAAMAMXMAMMAMXMMMMSMAMAXSMAMAXXAXAXXMMASXMXMAMMMMAXSASAXXMAASAMMSAMXMMMMAXASMAMMMMSXMAMXMSMXSASAMM
XXMASMMSMASMMSMAMASAMXXMSMMSXMXMAMXAAMAMXXMXMSMMASAXXXMAMAAAASMMXMXSXMASMXMASXXSXMAMXMMSSMSAMMASMSSMAXMAMASAMMSMSMAMSMMMXMAMXAAAXAMXMASASAMX
MAMAMAAXMXSAAAMASMMAAMSAXMAXAXXSMMMSSMAMMMMAMAAXAXXSSSSMSSSSMMXXMAMMXMASAXMASAAAMMAMMSAAAMXMSMMMMAAMMMMAMXSAMAAAMXXXMASMSMMMSMMXXASXMMMXMASX
AMMAXMMMSMSMMMSASXSXMAASMMMXAMXAAAMAMXAMMASASXSMXSSMMAAXMAAAXMMSMMXAAMMXMMXAXMXMAXXAAMMSMMAXMMAAMMXMAMSSSXMASXMSMMXMSSMXAAAXSMAXSSMMMASASXMX
SMXAMSMMXAXAXXMXSASXMXMAMASMASMAXSMASMMXSASASAXAAMAAMSMMMMMMAAAMAAMMMSAMSSSSSXSXMXXMSSMMMSSSSSSSSSSMMMAMXXXXMAXAMXAMMXASXSMXMMMXMAAMSASXXAAM
MMMSMMAMMSMMMSMXMXMASMXMSAMMAMAXSMXXMAAAMAMXMMMMMMSMMXSSXMASMMXSMMSAAXMAXAAAXASASMXXMXMAAXAAAAXAAAAASMMSXSSMMSSSSSMSAMXXAXMASAMASXMMXMSXXMMM
MXAXASAMXAAAAAAXMAMMMMXXMAXMSSSXSAASMMMSSSMMSXMASXMAXAMMXMASMMMSXAXMMMSSMMMMMMMXAXSXMASMMMMMMMMMMMMXMASAMSAXSAAXXAMMMSMMMMSASASASAMSAMXMASAS
MMMSMSAMSSSMSMMSSSXSASXMMMSAMAXAMXMSMMAXAMXAXMSAAMSSMMMSASXSMSAMMMSAMXAAAXMAXXMSMMMASAMAAXSXXSXXSXXSXXMAMSAMMMXMSMMAAAXASAMXMMMXSMMXAMXAMSAS
XSASXXAMXXMXMXMMAMASXSAMAAMAMMMXMXMSAMXMAMMSSMXSMMAMAAXSASAMAMASMAMMMMSXMMSXMSAMXASXMMSSMMMMMMAMXMAMXMSMMMAMAXSAAASMSSSMSXMMSMXMMASXMMSSXMAM
AMAXAMSMMAMXMAMMAMAMXMAMMSSSSXSMSXASAMXXASAMAMMMXMAXMMMMAMXMAXMMMASMSAMAXXXMXMAXSMSSXAXMAMAAAXSXAMAMAXAXAXAAXMASMXMAAAXASMMAAXAXSAMASAMXAMMM
SMAMSMMASXMASAXSXMXSMSMMMMAAXASAAMXMAXXSMAXSAMASMMMMASAMXMASXXSASMSXMASMMMMAXSXMXXSXMSMSMSSMSAMXXSASMSASMSSMSAMAMMSMXXMSMAMMSSSXMXXAMAXSMMMS
XMAXXASAMXSXXAMSXSAAMAMAAMSMMMMMMXXSXMMXAMXSASMSAASXMMXSAMXXMASASMXAMAMAMSASMXXAXMXXAAXAAMAAAMSAMSAXXMAMAAAAXXXAXASASMSXSXMSMMMAMXMMSXMMXAAA
SMSMSMMAXMXSMSXSAMXMMAXSXXAMMMAMXSAMMSXXSXASXMXSMMMAAAMSAMXMXMMAMAMSMMXAMXXXAAXSXASXSSSMSMMSMXMAMMXMSSMMMMMMMSSSMMSAMSAAMSMAASXXMASAAMSSSMSM
AAAXAXSMMSAMXXAMXMAMXMMMXSASMSMXMMXSASAMXMASXMASMMSSMMASXSAXSAMXMAMXAMSSSMXSXMAMXSMAAXAAXMAMXXSSMMXMAAAXSAMAXMAMAMMAMXMXMASXSMAXMAMMXMMAMAAX
MSMSXMSXAMAMAMMXAMASAMAMXXAMMAMXMMMMASASXMAMAMASAMXMASXMASAMXAMXMXXXAMAMAMASMSAAMXMMMMMMMMMSMXAAXAXMAXMMSASMMAMSSMMXMAMASXSMXXMMMSXSAAAMMMMS
XMXMXMXMXMMMXSMSXXASMXAMSMSMSXSAAAXMMMAMXAASXMMSMMAMMSAXAMMMXSMMXSMSSMMSAMMSASASASXXXXAXAAMAMSSMMSXMMSSXSAMXAAMMAXAXSXSMXASAXXXXMXAMXSMXMSAM
XAAMAMAASMSSMAMASMAMXSSSMAXAXAXXXMXMSMXMXSMSXSAMASXMASAMXXMSMXAMAMXMXAXXXXXMXMXMXMMMSSMMSSSMMXAXAMASAAXMMMMMSMSSMMMXSASXMAMMMSSMSMSMAXAXXMAS
SSXMASXSAMXAXMMAXMASMMMAMXMMMMMSMMXSAMMSAXAMMMAMASMMASAMXSMXASAMASMSMSMSASXSAMAMXSAAAAAMAAMXXSXMASAMMSMMSSMMAMAAMAMAMXMXMXXAAAAAAAAMMSAMXSAM
XAXSAMXXXSSSMXMSSMSAASMSMXMMASASAAMSASAXASAMXSXMASASMMAAAAXMMMASAXXAAAXSAMXSASMSMSMSSSMMMSMMXMASMMXSXXAAAAASAMMSMMMASMMXMXSMMSSMMSMSXSAMAMAS
MAMMSAMXAAAAMXAXXMXMMMXMXAMSMSASMMMSAMMSASXMASXMAMAMXSMMASAMXMMXAXXMMMXMXMMSMMAAMXMXMAMXAAAAAMMMAAAMAMMMMXMMASXMAXSAMXSASXXSAMAAAXXXASAMXSAM
MMMMMXXMAXXAMSAMXXXXASAMMXXAAMAMAXAMXAMMMMAMASMMSMMAMMASAAMMASXMSMXXASAMAMXXMASXMASAMAMSXSMSMSAMMMSSMXSXSXMMMMAMAMAAAMXAMMXMSSMMSXSMMMXSMMAS
SASASMXMASMSMXMXMMMSMXASAMSMMMSSSMXSMXSAMXMMXXAAMXSMASAMXMXSASAAAAAXMSASXXAAMAXMSXSMSAXSAXAMXSXSXAMAMSMASAMSSMMMMMSXMXMSMMSMASXAMMMAMXAXAXXM
MAMASAAMAXAAXAXAMXXXSMSMXXAXAXMAMAXXAMSMMSXSMSMMSXXAMMMSASXMMSMSMXSAMXXMAXSMMASAMXMAMXSMAMAMAXAXMXSAMAXAMAMAAAXSAMXXMMMMAAXMASMXMASAMMSSMSSS
MAMAMMXMSSSMMMSMSMAAMXXSXSASXSMMMSSMSMMAXXAAAAMAXMMSMAMMMAMXAXMMXAAAMMAMXAXASASMSAMAMXSMAMXMMMXMMXXAXMAXSXMXSMMXAXXXXAASMMSMASAMSAXAXAMXXAAA
SMMSAMXSAAAXSXAAMXMXMAMAXMAAAMMXAXMAAMSSMMSMSMMMMAAASXMASMAMMSASMSSMMMMMMASAMXSASXSASAMXAXSAMXMASMSXMXSXSASXAASMMMSMXXMXXAXMAXMMMMMSMMXSMMSM
AAAMAMSMMXMAAXMXMAMAMASMSXSMSMAMXSSSXMAAAAAAAMASMMSMSXSXXXSMXAAMXMXMAMSAMXMXSMMMMASAMMSSMSMMSMSASXSAMXXAMAXSAMXAAAAMMSSMMSXMASXMASAXAXSAAXMM
MMMSSMSASXSAMXSASASASASXAAMXXMAMXMAAMMSSSMMSMMASAAXMXMMSMSMSSMXMXXAMMSSXXMSAMXAAMMMAMXXXMAMAAAMASAMAMASAMAMXXSSSMSXSAAAMAAXMXAXSAXMSSMMSMMSX
MXMXAAXAMAXSXXSASASMSXSMMSMAMMSMXMASMXXMMAMXXMMSMMMMSAAAMSXAMMMSXSMSXMXSAAMASMSSMMSAMXMXSASMMMMMMMSXMASXMASMMAMAAXAMMSSMMSAMXSMMSXMAAAAXXASM
SAMSMMMSSMSAMAMMMASASAMXAXXXXAAAXASMMXMAMSAMXSAMXMAXMMMSSMMXSXAMAXMAMMAXMXMAMMAXAXAAXSAMSASAMXSXAXAXMXMXMAAMMASMMMXMMAXMXXAMAAAAAXMASMMMMSSM
SAMXMXAAAXMAMSXSXXSAMMMMXSASMSMSSMMMMAMXAMMSMMASMSMSAAXMAMASXMAMSASMMMAMXMMXMMAMXMSAMASAMMMXXAXAXMMXSAMMMMMXSASAAXSAMXSMASAMSSMMSSSMMAXSXAXA
SMMAAMMSSMSAMXAAXMMAMMAMAMAMAXXAAMAAXAXMASAMASXMAXMAXMXMAMXXSAAAMXMAMXAXXSMMXMXXAMXSMSMMMASAMSSMMMXMSMMAXASXXAXXMMXMASXMXSAMMAMAXMXAMMMAMMSM
XASMMMAXAMSASMMMMXSSMSASAMAMAMMSSMSXSAASAMXSMXAMSAMXASXSASXMASMXXXMAMXSSMXAMSMMSASAXXXAXSMSMXXAMMXAXMXSXSASAMSSMMSMMMMMSMSAMMAMSSMSSMMMMSMAM
SMXAAMMSMASXMXMASMMMAXASASMMASXXMAXAAMMMASAAXXMMAAMAMAMSAXASXXMSSXSASXMAMMAMMAMXAXXMSMMMAMSXXMAMXSMSSXAAMXMAMAAMSAMXAAMSASXMXXMXAAAAXXAXMASX
AXSSMSMAXAMASMMAXAASXMXMMMXSAMXXMAMMMMXSAMASXMXSMXMMSMMMSMAXAAXAAASAMXXAMMAXMAMMXMAMSAASAMSASMMMMXAAMMMMMAMAMMXMSAMXXXXMXMASXMMSMMMXMSMMSAMX
MMXAAXMMMASAMAMSSSMSAMXAXAXMASXXMAMXAXMAMXMMMAAXXAXXMXSAMMAMXSMMSMMASXSXMXMSSXXAAMAMMSMXMSMAMMXSMXMAMXAAXMSSMXSASAMSXSAAMMAMASMMXSXAASAMSAMS
MMAMMMMXSAMAMSAMAMXSMMSXMMMSAMXMASXSASMAMSXAAMSMSMSSXAMSMMMSMAXMAXMAXMAXSAMXMMMSMSAMMAMSXAMAMSMSMASXMSMSXMAXAAAMMAMAASXMMMAXSMAMASXMSMAMMAMM
SMSMAXMAMASXMMAMXMASMASXXAAMASAMXXMASMMAMAMXXSAMXAAXMXMAXMAAXSMSASXSSXSASASXXAXAASAMSAMASXSASXAMSASAAXAAXMASMXMXXSMMMMMXSSSSXSSMAMXMAMAMSMMX
AAAXXXMXSMMMSSXMAXMXMASXSMSSXMASMXXXAXMXXAAMXXMASMSMSASXSMSSXMAMMSAAAAXXSMMXSMSMMMAMSASMSXXXMMSMMXMMMMXMSMASXMXSAMASXMMAAXMAMAMMMXAMSSMMAMMX
MSMXMSMXSAASMMMXXXXXMMSAMAMXMXMAMMSMSMSSSMSAMXXAMXXASASXAAAMAMAMAMMXMMMMMXAMXAAAASAMXAMXMMMMSAMAMSSSMMAXAMMSMAMMMXAMAAMMMSMMMSMAAXSMMAXSAXSA
MAAASAMXSMMSAMMMSAMXMAMAMXMASASAXMAAAAAAAXXMAXMMSAMMMMXXAMXSMSAMXSXSMMAAAMXXXMMSMSMMMXMAXXXAAXSAMMAAXSXSXSXMASXAMMASMMMMSMASAMMASAXASAMSMMAX
MMSMSASAXXXSAMAAAASXMASAMAMAXAXAMSMSMSMMMMSXSMSAMMSXSXMSMSXSXSAMXAAXASMSSSSXASAMXMMSMSMSSXMASMSMMMMMMAAMAXAMSASAMAASXMAMASAMXSXXMAMAMMXXXAMX
MXAXSAMXMAMXAMMMMXMMSMSASXSSSSMSAMAMAXXMSASXXAMASXXAXMAAAXAXXXAMSSSMAMAAAAXSAMXMAMAAAAAMMAMAXXXAMMAAMMXMASXMMMXSMSXMMSMSAMXXAMMAMAMXMXMAMSXM
XSAMMXMASXXSAMXXMMSAXASMMMAXAXAXMMAXXXAXMASAMMMMXMMSMASMSMXMMMAMXAAMAMMMMMMSXMSXSXSXSMSMSMMMSMSSMSSSSXMMAMAAAXXXXMAAMAMXXMASAAAASAMASAMAMAAS
MSXMAXXXXMAMAMXSAAXMMMMMAMXMMMSMXSSSSMSMMXMXMMAMAMAXAMXAAASMSAMSMMMXMXAAAXXMAASAXAMXAAXAAXMXAXMAAAAAMMAMXSSSMSSMASXMMASXMAXASXSMSASASXMAXSAM
ASAMXMSAMAXMSMASMSMXAAAMXSMSXAAMXMAAXAXXSAMXXXAXSMMSSSMXMXAAXXSAAAAASMSMSASXMMMAMAMSMMMSMSXSXSSMMMSMAMSMMAAAXXASXMMMSAMAAMSMMMAXSXMXSAXMXXAX
SMAMAMSASXSAXMXMMXMXSSSSXXAXMXSMAMMMMXMAXASXXSSSXAXAAAAASAMXMSMMSMSXSAAXMASXAAXAMAMASMAAAMAXAMAMXAXAMXAAAMXMMSMMMMAMMMSSMMXAASAMXXMMXMASMSXM
ASMMXMSAMAMASXMXSMSXAAXAMMXMAAAMSMXMASMSSMMXXMMAMSMMMMSASAXAXSAAMMMMMMMMMXSMSMSXSXSASMSMMMAMMMAAMMSAMSMSXSAAAXXMASASAMXMMAXXMMXMAXSAAXXMASXM
MAXXSAMXMXMAMMAXXAMMMMMAXXMASMSMMAXSSXXAAAMMMMMXMXAAXXAAMAMXMMMMSXAAAAAXXXSAXXMXXMMXMMASAMMXSAMXSAMAMSAAAMXMMSMSASMSASAMXMSSMSASAMXSMSAMAMAM
XXMXMAXSAMMSSXSXMSMSMSSXXSMXMAMAXSMMXSMSMMMAAXSSMXSMSMSSSMSXMAMASXSSSMSASXMMMAXMMMMSMSASMSASMMMXMASXMMMMSMSMAAMMMSMMXSXMXXAAASXMAAAMXSAMSSXM
XMXAXSMMSAMXMAAMAAAAAMAMMASAMXMMMMAXXMXXASXSMSAXMAXXXAAXAAAASMSMSAMAMXMMSAAASAAASAAAXMASAMMMAAXAMXAMAASXMMSMSSSMASAMASMSMSMMMMASAMXMAMXMAAAM
ASMMMAXXAMMAMXMSSMSMSMAMMAMXMXAXASMMSSMSMMAAXMMMMMSMMMMSMMMAMAAASXMAAXMASMMMSMMMMMMXMMXMXMXXSMSXSXMXSMSXMASAAMAMAXAMAMXAAAXSAMXSMMAMASAMXXXM
XMAMSMMMMMSXSXXXAAAXAMXMASXMMSMSMSMAAAAAAMMMMSXSXAAAASAXXXSXMSMMMASXSAXASASASMSASMSSMMMMAMAXMAMASMXMMSXXMMMMMSSMMXSMSSSMSMSSXMASASMSMSASMSMS
MMMMAAAAXASASAAXMASXXXASXMAAMAAAASMMSMXAMAMXAAAXMSSSMSMSSXMAAXAAXXMAXAMASMMASASXSAAXAAAMMMMXMAMXMASASAMXSAMXAAMASMAAAXAAAXMXAMASXMASASASXAAA
AAASMSMSMASAMMSXXXMAMMAMAXMMSMSMSMAMMMMMSAXSAMXMXMAMAXMAXAMSMSXMSXMXMASASXMMMAMAMMMMSSSXASXMSAMXXXMXSAMSAMXMMMMAMMMMMSSSMSMSSMMSAXAMXMAMMMSM
MSXSAXMXXAMAMXMAMMSXMASMXMXAXAXMAMAMXSAMMMMAMXMMASAMXMMSSMMAAAMXXMASMMMASXAAMAMAMAAMAXXXAMAASASMSMMXSAMXMASMAAMMMMASXAAAXXAAAAASMMXSXMMMAMAM
MXAMXMSMMXMXMAMAMAAAAAXAXXMSMMMSXSXSAMMSSXSAMXMXASASMXMMAXSMMMMSAXAAMXMAMMSASXSMSXSMASMMSMMMMASAAAXMXXMAXXXMMMSAASASMMSMASMXSMMMXSASAMXXXSAM
XMAMAMXAXMSASAMAMXXSMMMASXAMAXXAMXXMAMXXAXSXSMXMASMMAAXMMMXAAAASXMMXMAXXXMMXMXXAXXMAMMAAXMAXMAXMSSMASMMXSMXXMAXXMMASXAMXMXAXMMSMAMAMAMAAXMXS
SSMMAAXSAAAXMAASMSMXASXAAMXMAMMASXMMSXXMAAMMSMXAXXASXMSAASXSMMMSAXAASXSSMXMASAMXMMSASXMMSSSXMASMMMXSXAAMMMSASMSXXMAMMMSSXMSMSAAMXMAMXMMSXAAX
XAASMSMXASMSMXMAAAAMAMMXMSAMXMSAMAMAMMXMSMMASAMMMSAMAAXXASAMMXXSAMMXMAAXAAXXMAMXAAXXXAMXXAMMMMXMAMXXSMMSAXXAMAAMSMSAMXAXXAXAMSMMAMAXAAMAMXMX
XSMMSAMXXXXAAAMMSMMMAMASXSASXXMAMXMAXAAXAAMXSAMAXMMMMMMMMMXMSMMSAMXSMMMMSXSAAAMMSSMSSXMMMMMMAMASMSAMASMMMSMSMMMSAAAMMMAXSMMSMMMSMSASXSMSAMXM
MAMXMMSMMMSMSXXMAMXSXSAMMMMMXMASMMSSSMMXSXMAXAMXMAAXASMXSAAMAAASAMAMAXSXAXMMMMMAMAXMAMXAXSASASMSXMASAMSAMXAAMMMXMSMSMMSXAMAMASAAXMASAAAAMXAA
AMMSSMAAAAAXMXXSMXXAMMASMMAXXXAMXMAXXASAXASMMMMXSSMSASAAMSSSMMMSSMAMAMSMMSSMASMMSSMXMMMSMAASASAMXMMMXSMMMMSMSSXMAXAXXAMSSMASXMMSSMMMXMMMSSSS
XAAXASXMMSASMXMMSMMMMSMMASASXMSSSMSSMMMAMAAXAAXXMAAXAMMMMAMXXMAMAMSSSXXAMAAMASXMAXMAAAAMAMAMAMMMXSAMXMASXAAASASAAXMXMAXMMXAMMXAAAAXAAXMAMAXA
MMSXMMMXMMMMMAMAAXAMAMASMMXXAAAAAAAAXSMSMSMSSMSMSMMMMXMSMSSSSMMSMMMAMMSSMMXMMSAMASXSSMSSSMXMXMAAXXASASMMMSSMMAMMMMMASXSASMXSAMMXSMMSMXMAMSMM
MAXAXSMMAAAASXMSXSMSASMXAXXXMMMXMMMXMAAMAMMXMXAMAMMSMAAAAMAMXAMXMXMAMXMASAMXMSMMAXXAXAXAXXMASXMSMSAMAXSAAAAAMAMASASASMSAMAAAMXXMAMXMAMSAMXSX
MASMMAAAMSXMMAMXASASASXSSSSXSASASASXAMXMASXAXMASMSAAMXXMXMAMSMMMMMXSMASAMXSAAXMASXMMMMMMMXMMXASAAMMMXXSMMSMMSSMXSASASAMMMMMSMMSAMXAMAMMASXMS
MASMSSSMXMMSSXMMXMAMAMMAAAAASASMSAMSSXXSSMMMSAXSAMXSXMMSASAMAAAMASMMMAMASASMSMSMXMASASAMAMSASMMMXMSMSASMAXMMMXSXMAMMMMMAMAMXAMMMMMMSMMSMMASX
MAMXAAXMAXAXMMMXXMAMMMMMMMMMMMXXMXMSXMAMAXASAMSMAMMMXXASAMASMSXSASAAMMSXMAXAAAAXAAXXAXASAAMASXAXSMAAMAMMAMAAMMASMXMAMASASASXSMAAXAMAMMAXSXMX
ASAMMSMMMSMSAAAAXXXXXXMASAXXXXAMAAXMAXSSSMMSAMXMAMAAMMAMMSMMMAXMASXMSXMMMAMSMSMMMXSXMSAMXMMAMMAMAMMSMAMMAMSMSSMXMASASASASASAMSSSSMSASXMMSXSX
MMAMXXMAMAASMMSSMMSMSSSXSASMMMMASMMMAMMAMMAMMMXSSSMAXXXAMAXAMXMMMMAXSAMXMAXAAXAMSAAAMXXMMSMAXXXXXXXMXXXMSMXAAMXMAMSMXXMXMXMAMAMAMASXMASMSASM
XXMMXXSASMXMXMAMAAAAAXXXSAXAXAAXMXSMMXMAMMAMAMAMXAMAMSXXSMSXSSSSXSXMXAMSXSMMMMAAXAMSMAAMASMSSMSSXMSSMMSMMAMMMXAMAMXMMMSXMXMSMSMSMMMSMXMAMAMA
AASAMXSASXMSXMASMMMMMXMAMAMSSMSMSASASXMMXSMSSXAASMMSMMAMAXAASXXSAAXMSSMMAMXMXSMMXSXXMXMMASMAAAASAMXAAAAAMSMSMSXSASXSAAAMSAXMAMXMAMASXAMAMMMS
SMSAMXMAMXXASMAMAXMASXSAMSMXAMAAMASAMAXMMXMAAMSMMXAMASXMAMMXMAAMMMSXSXAMAMAMAXAXAXXSMSSMASMSXMASXSSSMMSAMXMAAAMSAMXSMSSXSASMAMMSAMAXAASASAAM
XXXAMXMAMAMMXXXXXMMAXMXAXMMXMAMMMAMASAMXAXMXMAXSMMMSMMAASMMAMMMMAXSXXSMMMXAMSSMMMSASAASAMXXMAMAMMXAXMMMXMASMXMXMMMMMMMMMSXMMASXSAMXSMXSASMSS
SSSSMAMASMAXASAXSAMASMMXMAXMXXXXSSSMMMSMMSSMXAXAXMAAAXXMXAMAXAASMSSMXXSASMMAXMXSAAAMMMSXSSMSSMAXAMMMMAMAMAMAASASASMSMSAAXMASXMASMMMSXAMMMMAM
AAAXSXSAMXXMAXAASAMXSXAXSXMXAMXXXAAAXAAAXAAAMSSMXMSSXMXXXSSMSSMXSAXAMMSXSAMSXXASXMAMXAMAMAXAXXMMSMSAMAMMMSMXMXAXAMAAAXMMSXAAASMMAMAMMXMAAMAS
MMMMXXMASMXSMMXMSAMXXMXMXAAMMASMSSMMMSSSMXSSMMAXXMAMXASXAXAAAMASMSMSAASXSXMXSMMXXXAMAAMAMMMMSAMXAASMSMSMXXAASMSMSMMMMMMAMMMSMMASMMAMASASXMAS
AXAMMXMAMMASASMXSAAXXMSASMSXSASAAXAAXXAAXAMXASXMSMAMMMAMXXMMMSSMAAMAMMSMMMSMAMAMMSMSSXSASXAAAXMMMXMAAAAXAMMMMAXAXAXSASMASAXAASAMXXAMMXMMSMAS
XSMSMAMSXMMSAMAAMSXXXXAXSAAAMAMMMSSMSMSMMSSSMMSAAMSSSMSAASMXAXAMAMXAXMSAMAAAAMASMAMAXXSAXAMXSMSASAMSMSMMSXSAMAMMMMMMAXMMMXSSXMMMXSMMMAMAMMXM
XSAAAMXAAMAMXMMMMXMMMSMSMXMXMXMXAXAAXXXAAXAAMAXXMSXAMAMSMAAMSSMMMXXXSMXAMSSSXSAAXAMMMXMMMSMMXMMAXXXAMXXAAAXMAMAMAAAMXMAAMXMXMASAAAAXSSMSXMAM
SMSMSXMSXSAXMAMMMMMAAMMAMAXMAMXMASMMMXXMAMMMMMSMXMMMMXMAMMXMAAASXXSAMXSMMAMMXXMXSMSXSAMXAXASXSMAMSSSSXMAMSSMSXSXSSSXMXXASXMASXMMMSSMAXSXASMS
AMAMXAMAAXMMMSMMAASMSSSXSASAXAMXAXAAXMAMXSXAXAAMAMXSAAMMMXMMMSMMAAMAMAAAMAXXAXMASAAAMAMMMXXMAASAMXAMMMMSMAAAMAMAAMAASMSMMXSXMMMXXAAMAMMSAMXM
AMAMMAMMSMAXMAASMMSXXAMASASMSSMMSMSMSXXXMAXMMSMSMSMSSSSMMAMXMAMMXMAXMMSSMSMMMSMAMMMMMASAXAAMMMXSAMXMMAAAMMMMMAMMMMSMMAAAXASMSAMAMMMSAMXMMMAS
SSSSSSMAMMMSMSMMASMMSAMXMAMXMAAXMAXAMXSXMAXSAMXSAAAMAMXAXMSMMMSMMSMXSXAMAXMAXXMXSSXSSXSAXSSMXMAXSAMXSMSXSAMXMXMASXMMXSSSMASASMSAAXXAMAAMASAX
AAAAAASXMAAAXXAMMMMAXAMXMXMSSSMMMXMAMAMAMAMSASAMMMSMMMSSMSAXAXXAAAMAXMMMAMSSSMSMXMAMMXMAMXAAXMAXAMMAMXMAMAMAMXMAMAAMMMAXMAMXMAXXMMAAXSMSASXS
MMMMSMMXSMXMSXXXAAMMSXMAMAMXAMMAXSSXMXMAMXMMAMMMAAXXAAMAXSASXSMMXSMSXSXMMMAMAMXMAMAMSXSAAMMMMMSSMSMASAMAMASMSAMMMXMMAMMMMXSXMAMMSMXSAXAMMMMA
XAAMXMMXMASXMMMSMSSMAASASMXMAMXSAXAAMSMMSXXAAMAMMSSSMMSMMMAMAMXSMMAXASXMXMXMXXASMSXMXASXSXSAMXXAAAMXMASXSXMASMSMASMSSSMAAXXMMMSAAAAMASMMMMAM
SSSMAASAMAMAAAAXAAAMAMSAMAASAMAMMAMSMSAXAMXSSSXSXAXXAAAAAMAMXXAXAMAMXMAMSMSMSSXSMAMSMXMAMXSMSAMMXMSXXXMAMXMAMAXMAMAAAMXMSMMAAAMMMMMMAMXAAXSA
AMAMSMSASMSXMMSMASMMMXMAMSMXAMXASMXAASMMASMAMXMAXASMMSSSXSXSXSXSAMXSXMAMAAAAMAMXMSXAAXSAMXXAAXAXAXMMSAMAMXMXSMSMMSMXMAAMAAAMMXSAMAMMMMMXMAAX
MSAMMXSXMXSAXXXAXXMASMSAMASMSMSAMAXMMSXAAMMAMAXMXXXAAMXXMMAMAAAXMMAAMMXMMXMMMMAXSAXXSAMXMMMXMAMSXSAASASXSAMMAXXAXSXSXMASMMSASASXSASAMMSMXMXM
XMXMAXMAMAMAMASMSASMMXSASAXXAASASMXXAXAMXSXMSXSASMSSSSSMSMAMMMMMMMXSXASMMMXXAXXMAMASMMMAAAXMXSAAMMMMSAMMSASAMSMAMXAMXMASXAXAMASMSASMSAAAAMAM
SMSSMMSSMMSMMXAAMAMAAASXMXSSMMMAAXSMMSXMAXAMAMXAAAAAMXAAXMAMAAXAASAXMAXAASMSSSMMMXXXAASXSSMSAXMMMAMXMAAXSMMMAAAAXMXMAMSMMSMXMAMAMAMAMXMXXXAM
AAAAAMAMXMAAXXMAMSMXMASXSSMAASMMMXMAMXAMSXMMAAMXMSMXMXMMMMSXSSMXAMXSXAMXMMAAAMAAXXMSMMSXXXAMXXMXMASXSSMMSAMXSXSMSMMXAXAAXAASMSSXMAMXMSSMSSSS
MMMXSMSSSSMSSSXAAAMXSXMASAMAMMXMSMSXMXSXAASASXSXXXMASMXSXAMXMXMMAXXAMXMMAMMMXMSMSMAMMASMMMMMXASAMMSMMAXAMAMMMXMMXMASXSMAMMSMAAMMSSSMXAAXMAXA
XASAMMXMAAAMXMASMMSAAAMXSAMSXSXMAASMMMXMSMMAMASMXAXAMAAMMMSASAXMMMMMSAAXXMSSMXAAAMAMMASMSAAMMMSASASASMMXMMMAXAAAASMXMAXMXXAMXMAAAAAAMSSXMMMX
MSMASAXMSMMMAMAXSXMASXMMMXMXASMSMSMAAAXAAXMAMMMAXMMSMMAMAXSAMSMMAAXXAAMSAMAAXXXSMSMSMMSASXSSSXSAMXSAMXSXASMMSMMSMSAXSMMASMMMAXMMMSMMMAMXMXMX
SAMXMMMMAMMSMMSMAMSXMMXMASMMAMAXXAMSXSSSMSMXXMASXMAMASAMXMMXMAMXSMSAASMAMMSMMMMAAAMMAMMAMAMAXAMXMXMAXAMSXMAXAAAXMMXMAAAAXAASXSMSAMMXASXMASMS
SXSASAAMASAXXAXAMAXASAXMAMXMAMSMSMAXMAMAXSAXXSAMXMASXXAMXXXAMASXMASMMMMXMXMAAAAMMMMSAMMMMMMAMMMAMXSAMMXMSSXMSMMXSAMMSMMMSSMSAAXMAXAMXMAXAAXA
SASASXSSXSAMMMMSXSMXMASMSSMMAXMAMXAMASXMMMMSMMASMAAXAMASXMSXSASAMMMXAXXMMASXMMSXAAXMSMXMAMSXSAXAMMASMXASAMSAMAMAMASXMAAAMAAMMMMMAMMSAXXMXMAM
MAMAMMMAXMXXAASXXXAAXAMXMAMSSSMXAMXSXXAXMAAAXSAMXAMXMXAMAMAMMMSMMAAMSSMASASAMXXMMSMMXXXMAMXASASXMMAXMXXMAXXMSAMXSAMASMMSSMMMXMAMSSMAMAMSSSMX
MXMXMAMMMMASMXXSAXSMSAMMMAMMMAMSMSASMSXMSMSXXAMXAXSXSMMXAMASAMMMSMSMAAXAMASXXXMASAAMMMMSMSMMMMMMAMSMMXSSSMMMSXMXMAMXMAAAAASXMAMAMAXMAXMAAAXM
MSSSXXXXAXAXXXMMSMAXSAASMSSSMAMAAMAMXSAAXXAMSSMMSXMASXSSXSXSMSAMMAMMMMSXSMMXSAMXMSSMAAAAAXAAASMMSMAASXMAXAMAMXMAXAXSMMMMSMMMSSMSSSMXMXSMSMMA
MMAAASMSXSMSMSMAMSMMMAMXAAXMXMMMMMAMMMMSSMXXAAAAXAMAMAAMAMAMMSASMAMXXAMXXAXMMMSXMAMXSMSMSMSMMMAAMSSSMAMAMXMAXAAMXMXMAAAXMMAAAXAAAMAAMAMXAMXS
MMMMMMAAASXAAAMAMAMAMAMMMMXSASXSAMASAAAAMAMMSSMMSXMSMSMXAMAMXMXMXXSXMASMSSMSAAXXMAXXMMXAMAMMSSMMMXMAXSMMSASAMSXSASXSAMSSSSSSMMSMXMSAMMSXMSAM
SSSSSMXMMMAMSMSMMASASMSXMAMSAXASAXAXMMMXMAXXAXAXAAXXAAASXSMSXMASMXMASXAAAXASMSSMSMSMSMMXMAMXAAMAMMSSMXMASAMAXMASMXAASXMAAAAMXAXMAXMAMASAMXAM
AAAAXSSMMAXMXMAXMASMSMSAMAXMXMXMMMSSMXSASXSSMSMMSSMMMMXMXAXAXSASMAAAXMASMMXMAXMAAXAAAMSMSSSSMSMAMMAMXXMASAMMMMAMXMMMMMMMMMMMMASXXSSXMASAMMSS
MMMMMMSAASMMXMAMXAMXSXSASMSSMMSASXAAMASASXAAMXMAMMMSAMAMSSMSMMXSXXSAMAMXXMMMMXMSMSMSMSAAXMAXAXXXXMASXMMAMXMXSMMMSMXAXAMMSXMMMSSMXXXAMAXAMXAX
XSXSASMSMXAXMMASMAMXSASMMAAXAXSAMMXSMASMMMSMMXMMXAASMSMXAXAXSMAMAXMXASXMXXAAMMMAASAMXMMSMMSMMMAMSMMSAMXMSAMAMAAAMAMXMASAMAMXMAMXMXSSMASMMMMS
XAASMSAXXMXMMMASXSMMMXMAMMMMMMMMMMXMMXXAMXMAMSSSSMMSMAMMMMMMAMXSXSASXXASMSSXSAXMSMAMXSAMAAAASMSMXAASAMAMSXMAXSMSMSSSMAMMSMMAMMSASMAAAAXAAAAX
XMMMMMAMMSMMXMAXAMMXXASMSAAXAAXAAMXMASXSMASAAMAAXSAXMAMAXAMSMXMAMMAMMXAXAAXMXXMXXXXMMMASMXSMMXMASMMSMMSMMMSXMXMXAMAXMAMXAASASXSASMSMMMMXMMMX
MXAAAMAMAAASXSSMAMMSSMMXSXSSSSXSSSMXMAAXSXSMXMMMMMSSSSSSXMXMXSXXXMAMASMMSMXXXASMASMAXMMMMAMAMXMMMAAMMAXAXAAMSMSMAMMMMASMSXSASXMXMXXXXMAMSSSM
ASMSMSSMXMMMAAASXMSAXAAMXAMMAAAXAMMASMMMSMXMSAAMMAXMMXXAASMMASAMMXMMXXXAAAMSMMAMAMXAMXMAMAXMMXSSSXMSMMSSMMXMAAASAMXAMMAMXAMXMASASMAMSMMSAAMA
MSAMAXMASXXMMMMSMMMASMMSMSMAMMMMAMSASXMAXMASMSSSMSSMSMMSMMAMASAMAASMSXMMMXMAASXMASXXAMSAMXSXMAMXMMXXAMXAMXASMSMSMMMXSMAAMXMASAMMAMAAMAMMMMMX
MMAMMMMMMMAAAMXXAMMMMXAXAMAAMXXMXMMASAMMSSXSXAXXXMAAXXAAMSAMMMMMSASAMXMASASMXMMMXSAMAXMXMXAXMXSAMSMSXMXXMSXSAMXSAMXAMMMMMSSMMASMSSMSMMMAXXXX
XSAMXSXSASXSMSXSAMXSSMMMMMMMSXAMMXMASAMAMMMMMMMSMSMMMMSMXAASMMXAMXXAAXAXSASMSAMSAMXSMSMMMSMXXAMXSAAMMSAMXSXMASASASMXSAMAXAAXSAMMMAMMAMMSSMAS
XSAXAXMMASMAAXASXXMASMXAXAAXXXAMXXSXSMMMSAAAAAAAAXAMXXAMAMXAAXMXSAMXSSMMMMMMSAMAAXAXAAAAAAXMMXMMSMSMAAXAXSAXAMASAMXAMASXXMSMMMSMXAMSAMMXAMXA
ASAMSSMMMMMMMMMXMXMAMXSXSSSSMSSMMASAMXSASMSMXMXMMMMSMSMSMXMSXMMASAMAAAAMSXMXSAMSSMSSSSSMMSASXMSAXAMMMMXSXSXMASAMXMAMSMMXSXXMXXMAXSMSMSXMAMXX
MMAMMAXSSXMXSXMAMAMASMSAAAAAXAAASAMAMAMASMMMSSMSMMMAMXMAMAXXAMMASAMSSSMMSAMXMAMXXMMAMXAXXMAMAAMAMXMSAAMXMXMAXMASXSXMAAXAAAMAMMMAMXAMAAXSSMMS
XSAMSMMMAASAMAMSSMSXSAMMMMMMMSSMMMSAMASXMASAMXAAASXSMMSASXMSAMAMMXMAAAXASAMXMXMSXSMAMSMMXMAMMMMASXASMMAMMSXMASXMMMASMSMXMAMAAAMSSMAMMMMAAMAM
MSAXAXAMSMMASAMMAXXMMXXMASXXAXXXXMSMMXSMMXMAMMMMSMAMMASXSAAXMMSSMMMMSMMXMAMSMSAMAASAMAAAASXSMMAAAMAMAMSAAMMSAMMMMSAMAMMMMXSXSSXMAMXMMMMMSMMS
XSXSSSXMMMSMMXMSAMXXAAXMSMSMXSAMXMMMSMMASAXAMXXSXMAMMASXSMMMSAAAMXSMAMSMSMMAAMXMAMSMSSSSMSMSAAMASMASXMXMXXAAMAAAAMASAMSAMXAAXMASMMMXSXSAAASX
AMAMAXMASAXAAAXAAMMMMXMXAAXMASAMMXAAAMXMSXMMXAXMAXSXMASXSASAMMSMMAXSAXXAAXMMMMMXMMXMAMXMASASMMMXXXASXMSASMSMSSSSXMAMAXSSSXMAXAXASAAASAMXSXMX
MMAMMMXXMXSMMMSSMMAAXXSMMSMMXSXMMXMSSSSXSMASMXMASXMAMXSAMAMSSMXAMXXMMSMSMSMSSXXXMASMMSSMAMXMASXXMMMSAXMAMMXXXMAMXMSSMMXAXXXMASMMSMMSMAMAMMMM
AXASXXMMMASXAXMAASXMSMSAAXXMMSMSSMMXAAAAXAXMASAXMASAMSMXMSMMSAXAMXXAASXMASAAMSMXMAXMAAXMAXMSMMMMMAMMMMMMMASMSXSASMAAASMMSMMMAXAAMMMXMAMXXAAA
MSMSMSAAXAMMMMMSMMAAXAMXMXAMAMAXMASXMMMSMMSAMMSMXMMAMXAMMMAMXMSSMASMMSAMAMMMMAMAXMXMMSSSMXMXXAAAXSMMXSAAXMMAAAXASMMSMMAAAAAMMSMMMAXMSSSMSSMS
XAAMAMMXMMSASMXMXSMMMSMSSSSMSSSMSAMAASAMXAXAXMAMASMMMMXMASAMXAAAMMXMASAMAXAAXAMMMMASXMMAAASMSMSXXAASASMMMAMSMMMMMAAXXSMMSMXMMXMASMSMAAAAAXAM
SMSMSMSXXXMAMAAMAXAXXMAMAAXAXAMAMMMSMMMSMSSMMSASXSAMASASASAXXMSMMXAAMMXXMXSXSMSMASASAMXMSMSAAAAXMMMMASAMSMMMAMAAMXMXAXXMXXSMSAMXAMMMASMMMMMM
AAAXAXAXXMASMSMMMSMMMMAMMMMXMAMSMMAXAMASXMAXASASMSAMMSASASMMSXMASMMSMMAASAXASAAXAMASMMAXAAMMMSASMSMMMMXMAXASASXSMAXXXMMXMAXXXXMAMMMMAXXMSSSM
MSMSSSMMXAMXAAAXXAMAXSXSXMAXSAAXAMXXAMXSASAMMSXMASMMMMAMAMAASASAMAXMAMXSMASAMSMMSXMMAMXMMMMXMMAAAAXSMSMSMSMSASAMXXSXMAXMMSSMMSMXSAMSMSSXAAAM
XMAXAAXAMSMSSSMMSMMSMSXXMAXXMMSSSMSSXMSSMMMSAMAMAMAAMMMMXMMXSMMMSMMSAMXXMXMXMXXAAAXSXMAXAASAXMAMXMMMAAMAAMXMMMAMXMAXAMXXAAAAXSAAMMXAAAAMMSMM
SMMMXXMMXAAAXAMXAAXMXSXMASXSXXAAXAAAAMAMAAAMAMAMXXSMSSXSMSXMXXAAAAXSAXMSMSAMXXMASMMAASXSSSSXMMSXMSAMSMSMXMAXASAMAAXXMXSMMSSMMMMMSXSMSMMXMAMS
AAAMSSMMSMSMSMMSSMMMASAMAXMMMMMSMMSSMMASMMSSSMSMSMASXMAMSAXMASMSSSMSAMSAMAMMSMMMAASMMSAMXAMAXSXMASAXAAAMMSMSXSASXSMMSMMAXAAAXSAXSMSMXXXXSASM
SSMSAAAAXAMXXMXAMASMAMAMAXAAAAAAAMMAMXASAAAAAAAAAXMASMAMAMXMAMMAAAMMAMSASXSAAAXXMMMAAXAXMAMSMMAMAXMSMSMSAAMSASAMAAXAAASXMASXMASMSAMXSMSAMXAM
AAMMXSMASAMXMXMASXMASMXMMSSSSMSSSMSXMMASMMMSMMMSMSAMXMASMXSASMMMSMMSMMSXMAMMSSMAXXSMMMMXSASAASXMXSASAMXMMSSMAMAMSXMSSMMMSAMXMAMXMXMASAMXMSMX
`