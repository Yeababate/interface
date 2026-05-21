package main

import (
	"strconv"
)

func countRepeats(input string, start int, unit string) int{
	count := 0 
	for {
		nextStart := start + count*len(unit)
		nextEnd := nextStart + len(unit)
		if nextEnd > len(input) {
			break
		}
		if input[nextStart:nextEnd] == unit {
			count++
		} else {
			break
		}
	}
	return count

}
func encoder(input string) (string , bool){
	output := ""
	for _, v:= range input{
		if v == ']' || v == '['{
			return output, false
		}
	}
	i := 0
	for i < len(input) {
		unit1 := string(input[i])
		count1 := countRepeats(input, i, unit1)

		unit := unit1
		
		count := count1

		if i+1 < len(input) {
			unit2 := input[i:i+2]
			count2 := countRepeats(input, i, unit2)
			if count2 > count1{
				unit = unit2
				count = count2
			}
		}
		if count == 1 {
			output += unit
		}else {
			output += "[" + strconv.Itoa(count) + " " + unit + "]"
		}
		i += count * len(unit)
	}

	return output, true
}
