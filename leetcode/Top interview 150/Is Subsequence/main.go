func isSubsequence(s string, t string) bool {
    count := 0
    if len(s)==0{
        return true
    }
    if len(t)==0{
        return false
    }
    for i :=range(t){
        if t[i] == s[count]{
            count++
        }
        if count==len(s){
            return true
        }
    }
    return false
}
