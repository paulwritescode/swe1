package swe1_3

func ReverseString(inputParameter string) string {
    runesSlice := []rune(inputParameter)

    for i, j := 0, len(runesSlice)-1; i < j; i,j = i+1, j-1{
        runesSlice[i], runesSlice[j] = runesSlice[j],runesSlice[i]
    }

    return string(runesSlice)
}