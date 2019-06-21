/*
BITONIC POINT

Given an array. The task is to find the bitonic point of the array.
The bitonic point in an array is the index before which all the numbers
are in increasing order and after which, all are in decreasing order.
*/

package main

import (
    "fmt"
)

func bitonic(a []int, n int) int {
    l := 1
    r := n - 2
    for (l <= r) {
        m := (l + r) / 2
        if (a[m] > a[m - 1] && a[m + 1] < a[m]) {
            return m
        }
        if (a[m] < a[m + 1]) {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return -1
}

func main() {
    var n int
    fmt.Scan(&n)
    var a []int
    for i := 0; i < n; i++ {
        var temp int
        fmt.Scan(&temp)
        a = append(a, temp)
    }
    point := bitonic(a, n)
    if (point == -1) {
        fmt.Print("There is no bitonic point")
    } else {
        fmt.Print("The bitonic point is ", point)
    }
}

/*
INPUT :
n = 6
a = [1 4 8 4 2 1]
OUTPUT :
The bitonic point is 2
*/
