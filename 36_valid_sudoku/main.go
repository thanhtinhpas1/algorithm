package main

import "fmt"

func isValidSudoku_brute_force(board [][]byte) bool {
	// each row must contain digits 1-9, not repetition
	// each column must contain digits 1-9, not repetition
	// each sub boxes contains digits 1-9, not repetition

	rowLen := len(board)
	columnLen := len(board[0])

	// step 1: validate row
	for i := 0; i < rowLen; i++ {
		hash := make(map[byte]bool)

		for j := 0; j < columnLen; j++ {
			selectByte := board[i][j]

			if selectByte == byte('.') {
				continue
			}

			if !isValidByte(selectByte) {
				return false
			}

			if _, exists := hash[selectByte]; exists {
				return false
			}

			hash[board[i][j]] = true
		}
	}

	// step 2: validate column
	for i := 0; i < columnLen; i++ {
		hash := make(map[byte]bool)

		for j := 0; j < rowLen; j++ {
			selectByte := board[j][i]

			if selectByte == byte('.') {
				continue
			}

			if !isValidByte(selectByte) {
				return false
			}

			if _, exists := hash[selectByte]; exists {
				return false
			}

			hash[selectByte] = true
		}
	}

	// step 3: validate sub boxes
	for i := 0; i < rowLen-2; i += 3 {

		for j := 0; j < columnLen-2; j += 3 {
			hash := make(map[byte]bool)

			for k := 0; k < 3; k++ {
				for e := 0; e < 3; e++ {
					selectByte := board[i+k][j+e]

					if selectByte == byte('.') {
						continue
					}

					if !isValidByte(selectByte) {
						return false
					}

					if _, exists := hash[selectByte]; exists {
						return false
					}

					hash[selectByte] = true
				}
			}
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	// row:{row}:byte
	// col:{col}:byte
	// sub:{row}:{col}:byte
	hash := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == byte('.') {
				continue
			}

			selectByte := board[i][j]
			if !isValidByte(selectByte) {
				return false
			}

			if hash[fmt.Sprintf("row:%v:%v", i, selectByte)] || hash[fmt.Sprintf("col:%v:%v", j, selectByte)] || hash[fmt.Sprintf("sub:%v:%v:%v", selectByte, i/3, j/3)] {
				return false
			}

			hash[fmt.Sprintf("row:%v:%v", i, selectByte)] = true
			hash[fmt.Sprintf("col:%v:%v", j, selectByte)] = true
			hash[fmt.Sprintf("sub:%v:%v", i/3, j/3)] = true
		}
	}

	return true
}

func isValidByte(b byte) bool {
	return (b >= byte('1') && b <= byte('9'))
}

func main() {
	board := [][]byte{
		[]byte{byte('5'), byte('3'), byte('.'), byte('.'), byte('7'), byte('.'), byte('.'), byte('.'), byte('.')},
		[]byte{byte('6'), byte('.'), byte('.'), byte('1'), byte('9'), byte('5'), byte('.'), byte('.'), byte('.')},
		[]byte{byte('.'), byte('9'), byte('8'), byte('.'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.')},
		[]byte{byte('8'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('3')},
		[]byte{byte('4'), byte('.'), byte('.'), byte('8'), byte('.'), byte('3'), byte('.'), byte('.'), byte('1')},
		[]byte{byte('7'), byte('.'), byte('.'), byte('.'), byte('2'), byte('.'), byte('.'), byte('.'), byte('6')},
		[]byte{byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('.'), byte('2'), byte('8'), byte('.')},
		[]byte{byte('.'), byte('.'), byte('.'), byte('4'), byte('1'), byte('9'), byte('.'), byte('.'), byte('5')},
		[]byte{byte('.'), byte('.'), byte('.'), byte('.'), byte('8'), byte('.'), byte('.'), byte('7'), byte('9')},
	}

	fmt.Println(isValidSudoku(board))
}
