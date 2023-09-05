package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/vjeantet/jodaTime"
)

func TestJodaTime(t *testing.T) {
	date := jodaTime.Format("YYMMdd", time.Now())
	fmt.Println(date)
}
