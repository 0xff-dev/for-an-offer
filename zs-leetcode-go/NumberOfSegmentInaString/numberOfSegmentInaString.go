package NumberOfSegmentInaString


func countSegments(s string) int {
    if len(s) == 0 { return 0 }
    cnt := 0
    needAddOne := false
    for _, char := range s {
        if char == ' ' {
            if needAddOne { cnt++ }
            needAddOne = false
            continue
        }
        needAddOne = true
    }
    if needAddOne { cnt++ }
    return cnt
}
