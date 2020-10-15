package rpc

import "errors"

var (
	ErrNilQueue = errors.New("eboolkiq: queue must not empty")
)

func (x *ListReq) Validate() error {
	if x == nil {
		return ErrNilRequest
	}
	return nil
}

func (x *GetReq) Validate() error {
	if x == nil {
		return ErrNilRequest
	}
	return nil
}

func (x *CreateReq) Validate() error {
	if x == nil {
		return ErrNilRequest
	}

	if err := x.Queue.Validate(); err != nil {
		return err
	}

	return nil
}

func (x *DeleteReq) Validate() error {
	if x == nil {
		return ErrNilRequest
	}
	return nil
}

func (x *UpdateReq) Validate() error {
	if x == nil {
		return ErrNilRequest
	}
	if x.Queue == nil {
		return ErrNilQueue
	}
	return nil
}

func (x *FlushReq) Validate() error {
	if x == nil {
		return ErrNilRequest
	}
	return nil
}

func (x *CountJobReq) Validate() error {
	if x == nil {
		return ErrNilRequest
	}
	return nil
}
