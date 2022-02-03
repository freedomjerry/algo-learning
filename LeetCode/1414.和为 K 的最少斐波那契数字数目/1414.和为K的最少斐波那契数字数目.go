func findMinFibonacciNumbers(k int) (ans int) {
    f := []int{1, 1}
    for f[len(f) - 1] < k {
        f = append(f, f[len(f) - 1] + f[len(f) - 2])
    } 
    for i:= len(f) -1; i >= 0 && k >= 0; i-- {
        if k - f[i] >= 0 {
            k = k -f[i]
            ans += 1
        }
    }
    return ans
}