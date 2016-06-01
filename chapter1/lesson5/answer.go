package lesson5

import (
	"bytes"
    "fmt"
)

func countCompression(text string) int {
    last := text[0]
    size := 0
    count := 1
    for i := 1 ; i < len(text); i++ {
        if text[i] == last {
            count++
        } else {
            last = text[i]
            size += 1 + len(fmt.Sprintf("%d", count))
            count = 1
        }
    }
    size += 1 + len(fmt.Sprintf("%d", count))
    return size
}

func Compress(text string) string {
	size := countCompression(text)
    if size >= len(text) {
        return text
    }
    
    var buf bytes.Buffer
    last := text[0]
    count := 1
    for i := 1 ; i < len(text); i++ {
        if text[i] == last {
            count++
        } else {
            buf.WriteByte(last)
            buf.WriteString(fmt.Sprintf("%d", count))
            last = text[i]
            count = 1
        }
    }
    buf.WriteByte(last)
    buf.WriteString(fmt.Sprintf("%d", count))        
	return buf.String()
}
