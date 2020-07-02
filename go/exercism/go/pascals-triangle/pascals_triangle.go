package pascal

// Triangle generates Pascal's triangle and returns it as
// a 2-D slice.
func Triangle(depth int) [][]int {
	triangle := make([][]int, 0)

	triangle = append(triangle, []int{1})
	for i := 1; i < depth; i++ {
		row := []int{1}
		for j := 1; j < len(triangle[i-1]); j++ {
			row = append(row, triangle[i-1][j-1]+triangle[i-1][j])
		}
		row = append(row, 1)
		triangle = append(triangle, row)
	}
	return triangle
}
