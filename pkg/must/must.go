package must

import (
	"errors"
	"fmt"
)

func Must[E any](s E, err error) E {
	if err != nil {
		panic(errors.Join(fmt.Errorf("error while musting %T", s), err))
	}
	return s
}
