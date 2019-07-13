package iteration

const countRepeat = 5

func Repeat(character string) string {
    var repeated string
    for i := 0; i < countRepeat; i++ {
        repeated += character
    }
    return repeated
}
