package pkg

import "strconv"

func ToUint64(str string) uint64 {
    id, err := strconv.Atoi(str)
    if err != nil {
        return 0
    }
    return uint64(id)
}
func ToUint32(str string) uint32 {
    id, err := strconv.Atoi(str)
    if err != nil {
        return 0
    }
    return uint32(id)
}

