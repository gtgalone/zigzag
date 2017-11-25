// 메인 패키지 정의 실행모듈
package main

// 의존 주입
import (
	"fmt"
	"regexp"
	"strconv"
	"math"
)

// 메인 함수 정의
func main() {

	// 변수 정의
	foo := "ab2v9bc13j5jf4jv21"
	
	// 정규표현식 정의
	re := regexp.MustCompile("[0-9]+")

	// foo 변수에서 숫자만 추출하여 bar 리스트에 넣음
	bar := re.FindAllString(foo, -1)

	// 결과를 담을 빈 리스트 생성
	var result = []int{}

	/* 
		bar list에서 형변환 후 홀수인 것만 제곱을 한 뒤에 result 리스트에 넣고 아래 쪽에 정의해준 리스트 내부의 값을
		모두 합하는 sum 변수에 인자로 전달하여 출력
	*/
	for _, x := range bar {
		y, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		if y % 2 != 0 {
			result = append(result, int(math.Pow(float64(y), 2)))
		}
	}
	fmt.Println(sum(result))
}

// 리스트를 인자로 받아서 내부에 있는 값을 모두 더하여 리턴
func sum(array[]int) int {
	sum := 0

	for _, x := range array {
		sum += x
	}

	return sum
}