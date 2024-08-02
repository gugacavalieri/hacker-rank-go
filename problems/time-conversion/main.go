// This is the main package for this module
package main

import (
	"fmt"
  "strings"
  "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 * Time Complexity: O(1)
 * Space Complexity: O(1)
 */

func timeConversion(s string) string {

  splitString := strings.Split(s, ":")

  hour := splitString[0]
  minute := splitString[1]
  seconds := splitString[2][0:2]
  timeOfDay := splitString[2][2:4]
  hourMilitaryFormat := hour

  /* convert hour to military format */
  if(hour == "12" && timeOfDay == "AM") {
    hourMilitaryFormat = "00"
  } else if(hour == "12" && timeOfDay == "PM") {
    hourMilitaryFormat = hour
  } else if (timeOfDay == "PM") {
    hourInt, _ := strconv.Atoi(hour)
    hourMilitaryFormatInt := (hourInt + 12) % 24
    hourMilitaryFormat = fmt.Sprintf("%d", hourMilitaryFormatInt)
  }

  convertedDate := fmt.Sprintf("%s:%s:%s", hourMilitaryFormat, minute, seconds)
  return convertedDate
}
