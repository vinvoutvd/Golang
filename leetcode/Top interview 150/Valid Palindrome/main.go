import (
    "unicode"
    "strings"
)
//проверка строки на палиндром
func isPalindrome(s string) bool {
  //удаляем все специальные символы через strings.Builder
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
      // записываем, переводя в low
			result.WriteRune(unicode.ToLower(r))
		}
	}
  //записываем полученное от strings.Builder
	s = result.String()
    if len(s)==0{
        return true
    }
  // решаем с помощью двух индексов: один в начале другой в конце, если равны смещаем на шаг вправо и влево соответственно 
	i, j := 0, len(s)-1
	for i<=j{
		if s[i]!=s[j]{
			return false
		}
		i++
		j--
	}
	return true
}
