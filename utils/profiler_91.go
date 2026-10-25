package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.7" )

func itaTsWQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Laqj1tygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rt3kbZ8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAMLpHqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PttXNASPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aCZl8FnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dM65liRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVkWE9LKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zk13Dg8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQoxFeOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCJqWZorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXCI5epTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M4prBzWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VqDaRZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcNOoczDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M0AHNYIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzrEWjSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ox4bDvo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zw1YsZrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ntl9jpnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJcBSjdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDHxM0T2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVyw4uQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmjCGiPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43YjoEckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRsrAuwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GT0MRxj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLgtLiZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzRr0ftkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcIcLbeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EkdtVheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBAQ3fdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRsX7qOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjaPpzmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oh1TS6YCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AqgHJijcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdlGJnSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4F8h8tYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVrazbq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0GjUAwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFwngiJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6fHBxELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2qxO0fLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNiTioEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InbLcPXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUU5DHvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJqLwfy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqEisCC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQmGxsfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c00AbwFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iADOLf06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wm5Pwb3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLuwopaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcLsnk3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yrn6uZoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caJDX9fjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3D7pBLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrK6AjYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUd3UoXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6wN8AOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fS1iHuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOjlxov3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDYWIrMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4ZWwCzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FdiJYkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfZ98UJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1UYpavCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REkZZUqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQ3rEoGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oyn536prWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMfR73lzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYU87DgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HvgHae3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrSWUZHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AdcbzYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htdbXo7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMGf6Z5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2ZM2Z7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8zwBGD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iz1l2gNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDAMJ5rkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n68tYLodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2G0xrhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODFKeCLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltULFdBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9GKc2gjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TI370OQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2mRjKN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIgO8vs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUK42gkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEygAkNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sMMcyj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dflNCONVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLx1pYC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DkSdY5S3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZka4VmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGOjo0WLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CU8XZp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cun7OzuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJvfqCAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOSS5kKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBE9SXkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDatl2CnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6SQQzVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8k3ePjznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XK97LcqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mv2LUGDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWUsvxICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhcv81IwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfgbisQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j55ucAVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t015O9v8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnqW1af9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUmugtwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LER5PZCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWEOPDxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onnG6EZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mV1hMfTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yl73fkULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Z3r3lVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ccJOIOsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5T4davtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpmTpUV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyLyCAe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Az4Ul75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8cNis0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkTO026DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0wk7G4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inVGabPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5yTNjlTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svruhbHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTBRgML4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJoE1HrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvGhzCJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toAF1TfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1u032mXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g45Hdw00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEmtOSZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func di13dStQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfpSl5giWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1QeUuaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGBE1MHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyTIqRaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcAChHeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DooTRnGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYndmLcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ENg2wAj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsHhTqdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hs4fqbnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QfTtdJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CuyIRTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cz6KzduiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZ0WEITfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uO7mxrlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3PRVl5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A62ujWHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLU4WRcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3niBzKyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdqeqUXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPepaX9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zsb1BvHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2LXdpG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDLtw1owWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1eyhMpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgsTyai1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEQ9uuQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMIK3vi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func id5CnilxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bek2Z9U1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBOn6tjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nHGjf8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5L8fUK4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlMLvgPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BUiTjjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1jSjyk1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ox2PiU4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bouwY3fqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUFgE90pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUycrOhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ta0SJAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M4bgpeS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtuJf4SbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkgcnkOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QY8pRbXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fInXV90iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMZiYpgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgzwed4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goEqUbm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0O3dokMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCMZQRfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYVhQbYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIrqfSd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aD35Mz1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6BGQ1PdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYeysVgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func quCeaMRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TlyTpcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1moh1V4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCQSLAclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYhWZRnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sQfYV3DhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgMl1Dc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a215ennWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xby63PYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NN74fW8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSz6EiF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVlNb9t1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFrLJHRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncSGxzgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mT5CwE5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Uv79BoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YsqIjyINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEedBBcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nduroz5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRYwJboUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LkIPejxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClanU73fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWUalNLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFrmOLaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YVqmaTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08UrGvRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IuWJCWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkHp7mCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovUsFfF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPrKqaIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdDjgOGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hguGIpG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cl9TKwIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QFSBqHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xqo3ksS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0oLjgCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztWTjuTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuNOVZBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bafaIqojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjwHHEgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qSr6A4WnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSVcSD31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7HKl7EaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YsedFEW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfkiAHJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Nb4WrcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8f7Tm7HQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFzoQ8ahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYKCeUKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fIZzSzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMDr6hqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEIrfxkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3eJnNUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etfN1ms1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvvpKDtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAlDMEREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKiEtTq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inyBhHByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ChG6FiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVIJIjIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jy9xrVj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Ej9xTrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8GdryqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xw3r2WxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMzvWk0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUCKd6z3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBLiqWoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UpVE8wmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrNbXj9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfl4sCdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjsAQS6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VFTg8DtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0TobzJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umpzLA6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nkb2CP29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCinhFdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcPfx1IMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmcrRUhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7bMTmCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func koNKHaxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzpZMLMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiou8x0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xmb8jI03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uxhoc5pDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SzZhO7vXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aFMDQU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rcsg1x8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AD3hUaeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Uwrf8QxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POg5kBlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2YnBUraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func netlEbcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rsn9dlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kwz75MhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qxut5CwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJgbPfc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMAxzN9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIOuWsIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPMEcktaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dujT0ywcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BeL75WB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHVURM0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbdkT6iHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fskrGKPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAS6YjviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWtaX2nfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afBgHKEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ug662eO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEV7J7k2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcAuDH2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xezLr7dWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weBNwHhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOpFabfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuMZFu3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnMJYrGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XhfgX9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ya1ioV4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9efgwYsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7haXze6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clWiSaUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JbmnYmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGckn2pvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9Y0BX2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qR1149Y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXwbOIAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWdl0K1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RGFbq3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvlmIezUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgEZCCAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U55d0jXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAiUP7igWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T90pGMYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWNlpBgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXUabKtaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vzGaxWNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmnGcPq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BJKPDb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqSpkqHrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4NuyGUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZE9GJZHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acBeyVjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mouIHkQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3KWfyFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mq60G1A2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mqy9NAzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zaeld0ZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFyqBGtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8SXLgpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1X90BoLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQdAE9z7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwPe6DH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ohIr2LrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RiC7M7KTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKXTeyK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGjeLav6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgaVvaO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usHUjdTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7knsVFvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0evfzIiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VwQ3WVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVP3nrnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1ZI2HNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rb5mTE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEOKEmDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wwNr2goWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaNUixPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uawwvIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQwoSv6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLslFxhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sr7XPxpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkXxdXfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67nz3lQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0UYZ8SLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrTXD6XiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5yG0P6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMmCaEShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AoX5zUaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmw20tamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFZUP5CkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chj2v0ZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dPb7fjG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcO79TYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGZ44MmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QK6x3NcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ettY6VaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZPRzs5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oTa7oRC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOuuaEKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkxGmA2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOO8XcXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWaL5cyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgxUVxJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gF3QKJdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7n30800Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHU4cQ2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6Az1vNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxzYeeabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fO5nGdQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdxmoqIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZHdz96kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipSIzmA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhdNZq3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8n7xliSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WD2rRvFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LaZp6rZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElOWxjzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcK3VZJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUp3vODUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUekeCw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uTCmZZW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X99590vUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZ95WLgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOCZxwM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKEh4dj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fE8CHQpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duTk3EptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXGYTlmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VVDEe01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3wAE6ubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aByVomHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbUeiDJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfVSziD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c63tmVofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xuuwqbt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3I6v7wBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHWI3NSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5CB8m30wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eq7jVTIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTrXvqGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zm5NhuLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxInXpTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9NTXQHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nrIKuDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JsCbyTNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hildkmI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rF2ae7KhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YmUxSZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAFsqCR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKWTpuTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3pRy19pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwMfDuJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tgdHnyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJDG63RZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JvfZGmzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFRuDrGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHBR2CH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nj7DEzBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yswZPNCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlIUgr0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZnHVKwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeGsDdC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vlb4rqlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXVEZoCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zd63maUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faSuh2QwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rHx8zqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjXBaPfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUveBEqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6Wp3KMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MxaAj3ZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jHvQ3oRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nYQtQ95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iWGL70bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTbvY0KgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I1LY0bRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0apJVaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwlMfA7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUrfjG0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSQXUZj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYzh88EzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1E2i3MwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qztQRHTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgVYuV82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDGrJWLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InS7BSz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxPJtPSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cdm4Ht8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNSO8gNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7lPfVWpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uj08iExdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUA7meFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgabfhUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PeDK1MPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Up9To1c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFRxIZudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkpz9FakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7w9qcx0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTRfmm9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEe16IJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7JAjmoaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmPVGJxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func noqIEW1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eE0837jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWEnUHVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdWA8dY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2L7koSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkJJ0aO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQxuoAVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5cGxHpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7efO6qiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlH8riuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iq8sGWsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nFNcpNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRw65EJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYo3uUT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iE5cAIOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAuUlBx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJSeOHzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nG8x8nVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T33oPCJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uaBd1UYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZwzJQwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRMG6JjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMJXOx8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOu5t0CSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2nmW3ZTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBrikN4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OESSIJLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0xazQ7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKseXkfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6CbU0qZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8LLIwxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdbmZO8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpCjqd6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAWTUd5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpeN3L6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJrEFs9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5PThdp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGlyUHoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjHfyWBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MR94cfJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVlAt0Z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1H2xrraWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKHVUJWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGjCcb7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xb17d5IzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pG2dcXouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pq7dScF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAVT5vTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfLGDjwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vHndgdV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgwwz2UMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PNxadAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cE7UPSAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IznmPYqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LKkceWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Axn9uR1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36rsLT3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fo777QacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U61KacQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8t6xMmzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBjFRxYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CubfqZd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IiLhiYs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWTh8dknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdlwSI2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfQb6L4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fuid0gSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxFDnlobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KwLV6AeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRwq2p4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wIbeum81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qGos3VyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TINd6ApeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B9cED2fNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWf0M2kDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CK2ektWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVVGO4mVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qHVO6CuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYDCFIYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7TSV4wQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfkSporSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zacsbrX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GnlHYTVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2axX9eQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYY9tZtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKMygYPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIT3VffcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFlz7lbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5tZt9kdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNINjdoqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dK2KgC1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRRId2HGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dom6wzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kv3xKQjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InOsfOXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szTMEqyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xx69YTqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFQvXiA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLXqlY7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6bZHajo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdiZADKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3y6BR22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPOsMkkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfG7aUVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nL7yauY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0JJR1MLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqj5xwldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FOnlv46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1EoaGzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bmlWDXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFeqtwr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5E1gZpDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgiP1Sm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRv2hGqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MH6wosQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WabcTmOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsnZP3kfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func poqbsLd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSxImjCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEQy7gCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3D7mmg7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJ5l0BExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rL1feNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Ke5o86oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyETT327Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDTNvsV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CXIgJ75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHODN6h8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7mfXKTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekQNGAjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmCMMp4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwUGK2zmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kq4Op5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tMX1MrvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGJLD1uGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WnKU9ArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSfEsR8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaRjES0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYwM6O1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkBY76YuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hv6NydXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNs4ltm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bhxeaKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMNw7ZGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxy4APRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6HpSkbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrR77v6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIspsmivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJTsiSGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqwbDR9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9cTd5iwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQVXWJgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z98qBdFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XF7i9F9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPqbggxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioLpWyLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BfVKFiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxqzbjNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOOqsdECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3InkB0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qp63uWNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqnamY92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mY0bzpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gha5TAbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElN7efMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84vFddDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MagdkLChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzLdCY91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLdOa1qtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rkd5q2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KL5WGeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hV2e4TbCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSgHkoLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dWRSivlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFAs8jUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1ZjbCjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmQmmYYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFt2JnbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wj4GI9uVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsxuDSLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JD7UqG10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TaTMYYv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X26Nfm5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gj8IFvcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrCIEcimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhgcmiBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pz2CeCXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzGxPcHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uEfDaNrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhQzJqFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7YaaDO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TOMy9GTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8g0IL2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biUWg26vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGPp4lGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtJHc4lmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9mfa3MQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e13DA2psWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpkU5hwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfE4RidWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 831Mnr3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YC7TgOkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lzd50zWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vf68BVBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QPLYYL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDDmWW1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4eJ2WedYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9i9RndcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlB0VV4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zZ3qCr2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TD7Js1rlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQRWUZ3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRtT9a1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2LgVuP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBngjCfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RcFu8I71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nO9rxZF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWP5zbY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZWNQhSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpFSx182Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOM8Ldg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3hTTKz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igjm33R3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDiphGhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLDkIUQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sp7lguYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nkc90IoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KfRvgx1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvt5aD4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gn1eNW7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVlnPyNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsfCdG3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HvROecwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umbrb3ITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpbhBAU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3iEHIjzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jATq4vAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvyNgqMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rqnjR0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wuNHtnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8DEfh9tfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pShUF86zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umnIVCm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1jdHi8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbVcJZjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gyGVd6DMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHgyYB7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iChAlrjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Xi9r1hxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KO5Uv5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNzsaiHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ibf8ZYAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRhOjTGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0aK8C4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRIC12VHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLmbHki8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5K4z3KxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9k6dQqFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itPJ51J4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTiHHcenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90lSyd8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfPirfLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhKJyTIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0C7yHxm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dRgwfbSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zo23nw9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5tT6Mo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6HFvSfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHRfmcrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypLEnVXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qMx7F6lLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SYcVEP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qR29x1zqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y44R3w7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fYPbVSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzWFjc1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eq4vXCthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYQ7MMViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xk2IiEqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddiRSKIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKuehATeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkLcsOafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func doBtD3xgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1oVDrsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EorKAFm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yAsVYgoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdhiZNvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7R1suuYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 33okn4AaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ov9eZJ5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHWxdQGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXMH6OwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9IdOkN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvVQ3YzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNIce2hvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFwnxSgSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEiGuF6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8mhpFDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAkblQzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flUfakmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfpC8juyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8qPTkjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HOQKUoGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usW6q3SeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoSJtC3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfkquALOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9m5vi34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zGYfw2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNeGIZNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWhCx6iiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZ8zgEcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTj0MRvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zThyfa8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZG6cKmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xMpSCSTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5U37VIUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkXvFWUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yygwMQhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

