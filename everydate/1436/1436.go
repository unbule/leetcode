package main

func main() {

}

func destCity(paths [][]string) string {
	city := map[string]bool{}
	for _, path := range paths {
		city[path[0]] = true
	}
	for _, path := range paths {
		if !city[path[1]] {
			return path[1]
		}
	}
	return ""
}
