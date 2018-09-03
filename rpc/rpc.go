package rpcdemo

import "errors"

// Service.Method

type DemoService struct{}

type Args struct {
	A,B int
}

func (DemoService) Div(args Args,result *float32) error{
	if args.B == 0{
		return errors.New("division by zero")
	}

	*result = float32(args.A)/float32(args.B)
	return nil
}


