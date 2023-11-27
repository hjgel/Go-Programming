package main

import (
	"bufio"
	"fmt"
	"os"
)

// GetStrings reads a string from each line of a file.
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func main() {
	var games map[int]string
	games = make(map[int]string)
	// append
	games[456] = "성기훈"
	games[218] = "박해수"
	games[067] = "강새벽"
	games[001] = "오일남"
	games[199] = "알리"
	games[101] = "아이오아이"

	// fmt.Println(games[067])
	for _, v := range games {
		fmt.Println(v)
	}

	// lines, err := GetStrings("votes.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// counts := make(map[string]int)
	// for _, line := range lines {
	// 	counts[line]++
	// }
	// for name, count := range counts {
	// 	fmt.Printf("Votes for %s: %d\n", name, count)
	// }
}
