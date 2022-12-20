package main

import "log"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool
func firstBadVersion(n int) int {
	left := 0
	rigth := n
	for {
		if left > rigth {
			break
		}
		mid := (left + rigth) / 2
		if isBadVersion(mid) {
			rigth = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	log.Println(firstBadVersion(5))
}
