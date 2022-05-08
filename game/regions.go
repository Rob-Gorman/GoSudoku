package game

type regionType [9]map[int]int // { index: value }
type regionFunc func(int) int

var emptyBoard [81]int
var boardSize int = len(emptyBoard)

// empty regions as prebuilt structures to build out actual instances
var Rows regionType = buildFullRegion(emptyBoard, rowAlg)
var Cols regionType = buildFullRegion(emptyBoard, colAlg)
var Boxes regionType = buildFullRegion(emptyBoard, boxAlg)

var regionsAlgs = [3]regionFunc{rowAlg, colAlg, boxAlg}
var regions = [3]regionType{Rows, Cols, Boxes}

func ValidNumbers(board [81]int, index int) (result []int) {
	targetRegions := regionsIncludingIndex(index)
	fillIndexRegions(board, targetRegions)
	filledValues := valuesInRegions(targetRegions)

	for i := 1; i <= 9; i++ {
		if !filledValues[i] {
			result = append(result, i)
		}
	}

	return result
}

func regionsIncludingIndex(index int) [3]map[int]int {
	var result [3]map[int]int
	for i := 0; i < len(regions); i++ {
		regionClass := regions[i]
		regionInstance := regionsAlgs[i](index)
		result[i] = regionClass[regionInstance]
	}
	return result
}

func fillIndexRegions(board [81]int, emptyRegions [3]map[int]int) {
	for _, reg := range emptyRegions {
		for k, _ := range reg {
			reg[k] = board[k]
		}
	}
}

func valuesInRegions(regs [3]map[int]int) map[int]bool {
	var result = make(map[int]bool)
	for _, region := range regs {
		for _, v := range region {
			result[v] = true
		}
	}
	return result
}

func ValidateBoard(board [81]int) bool {
	for _, regionClass := range buildAll(board) {
		for _, regionInstance := range regionClass {
			var values map[int]bool
			for index, value := range regionInstance {
				if values[index] && value != 0 {
					return false
				}
				values[index] = true
			}
		}
	}
	return true
}

func buildAll(board [81]int) (result [3]regionType) {
	// builds out entire region mapping for board
	for i := 0; i < len(regionsAlgs); i++ {
		result[i] = buildFullRegion(board, regionsAlgs[i])
	}

	return result
}

func buildFullRegion(board [81]int, fn regionFunc) (result regionType) {
	// initialize each map for region data structure
	for regNum := 0; regNum < 9; regNum++ {
		result[regNum] = make(map[int]int)
	}

	// generate map keys as appropriate indicies of board structure
	for i := 0; i < len(board); i++ {
		regNum := fn(i)
		(result[regNum])[i] = board[i]
	}

	return result
}

func rowAlg(index int) int {
	return index / 9
}

func colAlg(index int) int {
	return index % 9
}

func boxAlg(index int) int {
	row := rowAlg(index)
	col := colAlg(index)
	box := (row/3)*3 + (col / 3)
	return box
}