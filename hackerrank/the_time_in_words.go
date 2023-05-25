// https://www.hackerrank.com/challenges/the-time-in-words/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

func getWordRepOfNums(num int32) string {
	if num < 0 || num > 50 {
		panic("invalid number")
	}

	switch num {
	case 0:
		return ""
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "quarter"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	case 30:
		return "half"
	default:
		postfix := getWordRepOfNums(num % 10)
		if postfix != "" {
			postfix = " " + postfix
		}
		var prefix string
		switch {
		case num >= 20 && num < 30:
			prefix = "twenty"
		case num > 30 && num < 40:
			prefix = "thirty"
		case num >= 40 && num < 50:
			prefix = "forty"
		case num >= 50 && num < 60:
			prefix = "fifty"
		default:
			panic("should not reach here")
		}
		return prefix + postfix
	}
}

func getMinuteStr(min int32) string {
	if min == 1 {
		return " minute"
	} else if min == 15 || min == 30 || min == 45 {
		return ""
	} else {
		return " minutes"
	}
}

func timeInWords(h int32, m int32) string {
	if h < 1 || h > 12 || m < 0 || m >= 60 {
		panic("Invalid args")
	}

	if m == 0 {
		return getWordRepOfNums(h) + " o' clock"
	} else if m <= 30 {
		return getWordRepOfNums(m) + getMinuteStr(m) + " past " + getWordRepOfNums(h)
	} else {
		return getWordRepOfNums(60-m) + getMinuteStr(m) + " to " + getWordRepOfNums(h+1)
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)
	fmt.Println(result)
	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
