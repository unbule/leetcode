package main

func main() {

}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	area1 := (ax2 - ax1) * (ay2 - ay1)
	area2 := (bx2 - bx1) * (by2 - by1)
	x := min(ax2, bx2) - max(ax1, bx1)
	y := min(ay2, by2) - max(ay1, by1)
	over := max(x, 0) * max(y, 0)
	return area1 + area2 - over
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
