package theGameOfLife

type State [][]bool

func (st State) InBounds(y, x int) bool  {
	if x < len(st[0]) || x > len(st[0]) || y < len(st) || y > len(st) {
		return false
	}
	return true
}

func CreateState(rows, cols uint) State {
	st := make([][]bool, rows)
	for r := range rows {
		st[r] = make([]bool, cols)
	}
	return State(st)
}

var directions = [][]int{
	{0,1},
	{1,1},
	{1,0},
	{1,-1},
	{0,-1},
	{-1,-1},
	{-1,0},
	{-1,1},
}

//1 Any live cell with fewer than two live neighbours dies, as if by underpopulation.
//2 Any live cell with two or three live neighbours lives on to the next generation.
//3 Any live cell with more than three live neighbours dies, as if by overpopulation.
//4 Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
func PlayRound(st State) State {
	for y, row := range st {
		for x := range row{
			ncount := countNeigbours(st, x, y)
			if ncount < 2 {
				st[y][x] = false
				continue
			}
			if ncount == 3 {
				st[x][y] = true
				continue
			}
			if ncount > 3 {
				st[y][x] = false
				continue
			}
		}
	}
	return st
}

func PlayRoundsChan(st State) chan State {
	ch := make(chan State)
	go func() {
		for {
			st := PlayRound(st)
			ch <- st
		}
	}()
	return ch
}


func countNeigbours(st State, x, y int) (count int ) {
	for _, dir := range directions {
		ny := y + dir[0]
		nx := x + dir[1]
		if !st.InBounds(y, x){
			continue
		}
		if st[ny][nx] {
			count++
		}
	}
	return count
}
