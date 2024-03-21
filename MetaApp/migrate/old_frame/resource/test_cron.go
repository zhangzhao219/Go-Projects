package resource

import (
	"context"
	"fmt"
	"time"
)

func TestUpdate1(ctx context.Context) error {
	fmt.Printf("test1 execute, time %v\n", time.Now())
	return nil
}

func TestUpdate2(ctx context.Context) error {
	fmt.Printf("test2 execute, time %v\n", time.Now())
	return nil
}
