package bitcharacters

func isOneBitCharacter(bits []int) bool {
    flag := false
    aux(0, bits, &flag)
    return flag
}

//11, 10 ,0 三者
func aux(index int, bits []int, flag *bool) {
    if index == len(bits)-1 && bits[index] == 0 {
        *flag = true
        return
    }
    if index > len(bits)-1 {
        return
    }
    if bits[index] == 0 {
        aux(index+1, bits, flag)
        return
    }
    aux(index+2, bits, flag)
}
