import "strings"

// можно ли собрать s1 из символов s2
// если количество символа в s1 > в s2 то нельзя
func canConstruct(ransomNote string, magazine string) bool {
    for _, letter := range(ransomNote){
        if strings.Count(ransomNote, string(letter)) > strings.Count(magaine, string(letter)){
            return false
        }
    }
    return true
}
