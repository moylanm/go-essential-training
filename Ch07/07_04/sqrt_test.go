package sqrt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	value	 float64
	expected float64
}

func newTestCase(value, expected float64) testCase {
	return testCase{
		value: value,
		expected: expected,
	}
}

func parseTestCaseFile() ([]testCase, error) {
	file, err := os.Open("sqrt_cases.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var testCases []testCase

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringValues := strings.Split(scanner.Text(), ",")
		value, expected, err := parseFloats(stringValues)
		if err != nil {
			return []testCase{}, err 
		}

		testCase := newTestCase(value, expected)
		testCases = append(testCases, testCase)
	}

	if err := scanner.Err(); err != nil {
		return []testCase{}, err
	}

	return testCases, nil
}

func parseFloats(values []string) (float64, float64, error) {
	value, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}

	expected, err := strconv.ParseFloat(values[1], 64)
	if err != nil {
		return 0.0, 0.0, err
	}

	return value, expected, nil
}

func almostEqual(v1, v2 float64) bool {
	return Abs(v1 - v2) <= 0.001
}

func TestMany(t *testing.T) {
	testCases, err := parseTestCaseFile()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatal(err)
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}
