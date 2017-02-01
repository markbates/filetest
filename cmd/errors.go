package cmd

import "sync"

var errs []error
var moot = &sync.Mutex{}

func Add(err error) error {
	moot.Lock()
	defer moot.Unlock()
	errs = append(errs, err)
	if failFast {
		return err
	}
	return nil
}
