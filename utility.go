package utility

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Clear line cursor is currently on to render new line of len(string) < len(previous_string).
func placeholder() {
	var argument = string
	fmt.Printf("\r\033[2K%s", argument)

}

// ByteCount functions returns string representation of byte sizes.
func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

/*
Example output of ByteCount.

int64 passed		ByteCountSI		ByteCountIEC
<= 999				999 B			999 B
>= 1000				1.0 kB			1000 B
1023				1.0 kB			1023 B
1024				1.0 kB			1.0 KiB
987, 654, 321		987.7 MB		941.9 MiB
math.MaxInt64		9.2 EB			8.0 EiB
*/

// Convert number range ie. 3-7 [3, 4, 5, 6, 7]
func convertEntry(str []string) []int {
	var rtn []int
	for _, v := range str {
		if strings.Contains(v, "-") {
			expand := strings.Split(v, "-")
			begin, _ := strconv.Atoi(expand[0])
			end, _ := strconv.Atoi(expand[1])

			for i := begin; i <= end; {
				rtn = append(rtn, i)
				i++
			}
		} else {
			i, _ := strconv.Atoi(v)
			rtn = append(rtn, i)
		}
	}
	return rtn
}

// End

// Walk path and increment size.
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
