package helpers

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
    if condition {
        return a
    }
    return b
}

// Here are some test cases to show how you can use it

// func TestIfThenElse(t *testing.T) {
//     assert.Equal(t, IfThenElse(1 == 1, "Yes", false), "Yes")
//     assert.Equal(t, IfThenElse(1 != 1, nil, 1), 1)
//     assert.Equal(t, IfThenElse(1 < 2, nil, "No"), nil)
// }

// For fun, I wrote more useful utility functions such as:

// IfThen(1 == 1, "Yes") // "Yes"
// IfThen(1 != 1, "Woo") // nil
// IfThen(1 < 2, "Less") // "Less"

// IfThenElse(1 == 1, "Yes", false) // "Yes"
// IfThenElse(1 != 1, nil, 1)       // 1
// IfThenElse(1 < 2, nil, "No")     // nil

// DefaultIfNil(nil, nil)  // nil
// DefaultIfNil(nil, "")   // ""
// DefaultIfNil("A", "B")  // "A"
// DefaultIfNil(true, "B") // true
// DefaultIfNil(1, false)  // 1

// FirstNonNil(nil, nil)                // nil
// FirstNonNil(nil, "")                 // ""
// FirstNonNil("A", "B")                // "A"
// FirstNonNil(true, "B")               // true
// FirstNonNil(1, false)                // 1
// FirstNonNil(nil, nil, nil, 10)       // 10
// FirstNonNil(nil, nil, nil, nil, nil) // nil
// FirstNonNil()  