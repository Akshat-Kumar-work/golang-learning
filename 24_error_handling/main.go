package errorhandling

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
)

func main() {
	// in go errors are value not exceptions
	type error interface {
		Error() string
	}
	//Any type that implements Error() string is an error.
	type MyError struct {
		Code int
	}

	// func (e MyError) Error() string {
	// 	return fmt.Sprintf("error code %d",e.Code)
	// }
	err := sql.ErrNoRows
	if err != nil {
		return //fmt.Errorf("something %w", sql.ErrNoRows)
	}

	//Go errors can form a chain (via wrapping with %w).
	//fmt.Errorf -> use to add context over error
	//it wrap error and create a stack
	// %w is used to wrap errors
	//high-level error
	//└── wrapped error
	//	└── root cause

	//err2 := fmt.Errorf("failed to load %w", sql.ErrNoRows)

	//errors.Is
	//check error identity
	//works even through wrapped errors
	//“Is this error (or any wrapped error) equal to THAT error?”
	//It checks identity, not structure.
	//// ✅ CORRECT: sql.ErrNoRows is a sentinel(predefined) error
	// if errors.Is(err, sql.ErrNoRows) {
	// 	return nil, //appErrors.NewUserNotFoundError()
	// }
	// ✅ CORRECT: io.EOF is a sentinel error
	// example......=> if err := c.ShouldBindJSON(&request); err != nil && !errors.Is(err, io.EOF) {
	// 	// Handle non-EOF errors
	// }
	// if error.Is(err, sql.ErrNoRows) {
	// }

	//errors.As
	//“Does this error (or any wrapped error) have THIS TYPE?”
	//It checks type, not equality.
	//example=>...... errors.As(err,validationErrCustom) for custom errors checking

	var pathErr *fs.PathError
	if errors.As(err, &pathErr) {
		fmt.Println(pathErr.Path)
	}

	//immediate program disruption
	// crash program
	panic("something unexpected happened")

}
