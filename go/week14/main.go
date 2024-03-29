package main

import "fmt"

func status(name string) {
	balls := map[string]int{"성기훈": 20, "오일남": 0}
	// ball := balls[name]
	ball, exists := balls[name]
	fmt.Println(ball, exists)
	if !exists {
		fmt.Println(name, "님은 게임에 참여하실 수 없습니다.")
	} else if ball < 1 {
		fmt.Println(name, "님은 ", balls[name], "개로 탈락..탕탕탕")
	} else {
		fmt.Println(name, "님은 게임에서 승리하셨습니다.")
	}
}

func main() {
	status("오일남")
	status("강철")
	status("성기훈")

	// balls := make(map[string]int)
	// // var balls map[string]int
	// fmt.Printf("%#v\n", balls)
	// balls["성기훈"] = 20
	// fmt.Println(balls["성기훈"])
	// fmt.Println(balls["오일남"])

	// games := map[int]string{
	// 	456: "성기훈",
	// 	218: "박해수",
	// 	067: "강새벽",
	// 	001: "오일남",
	// 	199: "알리",
	// 	101: "아이오아이",
	// }
	// // append
	// // games[456] = "성기훈"
	// // games[218] = "박해수"
	// // games[067] = "강새벽"
	// // games[001] = "오일남"
	// // games[199] = "알리"
	// // games[101] = "아이오아이"

	// // fmt.Println(games[067])
	// for _, v := range games {
	// 	fmt.Println(v)
	// }
	// // update
	// games[101] = "장덕수"
	// // delete
	// delete(games, 199)
	// for k, v := range games {
	// 	fmt.Println(k, v)
	// }

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

// GetStrings reads a string from each line of a file.
// func GetStrings(fileName string) ([]string, error) {
// 	var lines []string
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		lines = append(lines, line)
// 	}
// 	err = file.Close()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if scanner.Err() != nil {
// 		return nil, scanner.Err()
// 	}
// 	return lines, nil
// }
