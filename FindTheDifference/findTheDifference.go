package FindTheDifference

func findTheDifference(s string, t string) byte {
    var start uint8 = 0
    for _, item := range s {
        start ^= uint8(item)
    }
    for _, item := range t {
        start ^= uint8(item)
    }

    return byte(start)
}
