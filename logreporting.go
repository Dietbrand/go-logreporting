package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readLog(s string) (fp *os.File) {
	fp, err := os.Open(s)
	check(err)
	return fp
}

func checkHttpCode(i int) (string, error) {
	httpCodes := [...]int{200, 301, 302, 403, 404, 500, 503}
	for _, n := range httpCodes {
		if n == i {
			return " " + strconv.Itoa(i) + " ", nil
		}
	}
	return "", errors.New("Not a valid HTTP code")
}

func createReport() (fp *os.File) {
	t := time.Now()
	s := "/tmp/report" + string(os.Args[2]) + t.Format("02012006") + ".txt"
	fp, err := os.Create(s)
	check(err)
	return fp
}

func main() {
	//Step 0: define variables
	//var httpCodeOK, reportOK string

	//Step 1. read argument 1, the logfile
	logFile := readLog(string(os.Args[1]))

	//Step 2: read argument 2, the HTTP status code
	httpCodeInput, err := strconv.Atoi(os.Args[2])
	check(err)
	httpCode, err := checkHttpCode(httpCodeInput)
	check(err)

	//Step 3: Create report file
	//Function: createReport()
	reportFile := createReport()

	//Step 4: Read each line from file (arg1) and compare (arg2)
	scanner := bufio.NewScanner(logFile)
	newLine := []byte{'\n'}
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), httpCode) {
			reportFile.WriteString(scanner.Text())
			reportFile.Write(newLine)
		}
	}

	//Step 4: output info to terminal
	fmt.Printf("Opening log,: %s\n", logFile.Name())
	fmt.Printf("Validating HTTP code %s\n", httpCode)
	fmt.Printf("Writing reportfile: %s:\n", reportFile.Name())

	//Step 6: Close files
	defer logFile.Close()
	defer reportFile.Close()
}
