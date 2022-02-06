func max(a int , b int) int {
    if a < b {
        return b
    }else {
        return a
    }
}
func maxArea(height []int) (maxCap int) {
    low := 0
    high := len(height) - 1
    maxCap = 0
    for{
        if(low >= high){
            break
        }
        if height[low] < height[high] {
           maxCap = max(maxCap, height[low] * (high - low))
           low++
        }else {
            maxCap = max(maxCap, height[high] * (high - low))
            high--
        }
    }
    return
}