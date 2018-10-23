package structs

import "strconv"

type Counter struct {
	Source     [][]string
	stopChanel chan stopStruct
	counters int
}

func (c *Counter) Count() [][]string{
	c.stopChanel = make(chan stopStruct)
	c.counters = 0

	for hIndex := 0; hIndex < len(c.Source); hIndex++ {
		for vIndex := 0; vIndex < len(c.Source); vIndex++ {
			if c.Source[hIndex][vIndex] != "X" {
				c.counters++
				go c.countOne(hIndex, vIndex)
			}
		}
	}

	for {
		select{
		case result := <- c.stopChanel:
			c.Source[result.Coordinates[0]][result.Coordinates[1]] = result.Result
			c.counters--
			if c.counters == 0{
				return c.Source
			}
		}
	}
}

func (c *Counter) countOne(x int, y int) {
	answer := stopStruct{Coordinates: [2]int{x, y}}
	temp := 0
	for hIndex := x - 1; hIndex <= x+1; hIndex++ {
		for vIndex := y - 1; vIndex <= y+1; vIndex++ {
			temp += c.getValue(hIndex,vIndex,)
		}
	}
	answer.Result = strconv.Itoa(temp)
	c.stopChanel <- answer
}

func (c *Counter) getValue(x int, y int) int {

	if (x < 0) || (x == len(c.Source)) || (y < 0) || (y == len(c.Source)) {
		return 0
	}

	return boolToInt(c.Source[x][y] == "X")
}

type stopStruct struct {
	Coordinates [2]int
	Result      string
}

func boolToInt(exp bool) int{
	if exp{
		return 1
	}else{
		return 0
	}
}