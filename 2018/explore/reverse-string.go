package main

func reverseString(s string) string {
    var (
        i = 0
        j = len(s)-1
        ns = []byte(s)
    )
    for i < j {
        ns[i], ns[j] = ns[j], ns[i]
        i++
        j--
    }
    return string(ns)
}

func main() {
}
