package main

import (
	"aoc_2021/lib/set"
	"strings"
)

func part1(paths map[string][]string) int {
	routes := findRoutes(paths, []string{}, "start", set.New[string]())
	//for _, r := range routes {
	//	fmt.Printf("%s\n", r)
	//}
	return len(routes)
}

func findRoutes(paths map[string][]string, route []string, current string, visited *set.Set[string]) []string {
	//fmt.Printf("%s\n", strings.Join(route, " - "))
	route = append(route, current)
	if current == "end" {
		// found a route!
		return []string{strings.Join(route, ",")}
	} else if isSmall(current) && visited.Has(current) {
		return []string{}
	}
	var routes []string
	for _, fork := range paths[current] {
		v := visited.Clone()
		v.Add(current)
		routes = append(routes, findRoutes(paths, route[:], fork, v)...)
	}
	return routes
}

func isSmall(s string) bool {
	return strings.ToLower(s) == s
}

func part2(paths map[string][]string) int {
	routes := findRoutes2(paths, []string{}, "start", false, set.New[string]())
	//for _, r := range routes {
	//	fmt.Printf("%s\n", r)
	//}
	return len(routes)
}

func findRoutes2(paths map[string][]string, route []string, current string, smallTwice bool, visited *set.Set[string]) []string {
	//fmt.Printf("%s\n", strings.Join(route, " - "))
	route = append(route, current)
	if current == "end" {
		return []string{strings.Join(route, ",")}
	} else if isSmall(current) && visited.Has(current) {
		if smallTwice || current == "start" {
			return []string{}
		} else {
			smallTwice = true
		}
	}
	var routes []string
	for _, fork := range paths[current] {
		v := visited.Clone()
		v.Add(current)
		routes = append(routes, findRoutes2(paths, route[:], fork, smallTwice, v)...)
	}
	return routes
}

func Parse(lines []string) map[string][]string {
	paths := make(map[string][]string)
	for _, line := range lines {
		from, to, _ := strings.Cut(line, "-")
		paths[from] = append(paths[from], to)
		paths[to] = append(paths[to], from)
	}
	return paths
}
