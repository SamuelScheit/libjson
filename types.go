package libjson

// json type
type t_json int32

type token struct {
	Type t_json
	// only populated for number and string
	Start int
	End   int
}

var empty = token{Type: t_eof}

const (
	t_string       t_json = iota + 1 // anything between ""
	t_number                         // floating point, hex, etc
	t_true                           // true
	t_false                          // false
	t_null                           // null
	t_left_curly                     // {
	t_right_curly                    // }
	t_left_braket                    // [
	t_right_braket                   // ]
	t_comma                          // ,
	t_colon                          // :
	t_eof                            // for any non structure characters outside of strings and numbers
)

var tokennames = map[t_json]string{
	t_string:       "string",
	t_number:       "number",
	t_true:         "true",
	t_false:        "false",
	t_null:         "null",
	t_left_curly:   "{",
	t_right_curly:  "}",
	t_left_braket:  "[",
	t_right_braket: "]",
	t_comma:        ",",
	t_colon:        ":",
	t_eof:          "EOF",
}
