package skyhookcomp

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"testing"
)

func TestOnce(t *testing.T) {
	s := NewStore("localhost:6380")

	for i := 0; i < 1; i++ {
		v1, _ := s.Get(context.TODO(), fmt.Sprintf("id:%v", i))
		time.Sleep(time.Millisecond)
		v2, err := s.Get(context.TODO(), fmt.Sprintf("id:%v", i))
		if err != nil {
			t.Errorf("err: %v", err)
		}

		if !reflect.DeepEqual(v1, v2) {
			t.Errorf("v1: %#v, v2: %#v", v1, v2)
		}
	}
	// just do this so that the logs are always shown
	t.Fail()
}
