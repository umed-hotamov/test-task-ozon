package urlshortener

const (
  BASE             = 62
  UPPERCASE_OFFSET = 26
  DIGIT_OFFSET     = 52
)
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func mapToAlphabet(index int) byte {
  return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"[index]
}

func Encode(num int) string {
  var shortenedUrl string

  for num != 0 {
    ch := mapToAlphabet(num % BASE)
    shortenedUrl += string(ch)
    num /= BASE
  }
  
  return reverse(shortenedUrl)
}

func Decode(shortenedUrl string) int {
  var id int 

  for _, r := range shortenedUrl {
    if 'a' <= r && r < 'z' {
      id = id * BASE + int(r) - 'a'
    } else if 'A' <= r && r <= 'Z' {
      id = id * BASE +  int(r) - 'A' + UPPERCASE_OFFSET
    } else if '0' <= r && r <= '9' {
      id = id * BASE + int(r) - '0' + DIGIT_OFFSET
    }
  }

  return id
}
