package blocks

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/forestgiant/sliceutil"
)

//- Line
//# Square
//$ Left L
//% Right L
//^ Left Z
//& Right Z
//* Half Cross
var blocks = [7]string{"-", "#", "$", "%", "^", "&", "*"}
var specials = [9]string{"a", "c", "n", "r", "o", "s", "b", "g", "q"}

func GenerateBlock() string {
	rand.Seed(time.Now().UnixNano())
	orientation := rand.Intn(4)

	var block float64 = float64(rand.Intn(100))
	var coef float64 = 100 / 7

	blockType := int(math.Floor(block / coef))
	return blocks[blockType] + strconv.Itoa(orientation)
}

func GenerateSpecials(coordsArray []string) {
	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(5)
	var result []string

	for i := 0; i <= count; i++ {
		itemIndex := rand.Intn(len(coordsArray))

		if !sliceutil.Contains(result, coordsArray[itemIndex]) {
			result = append(result, coordsArray[itemIndex])
		}
	}

	fmt.Println(result)
}
