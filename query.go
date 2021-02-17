package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"

	"github.com/machship-mm/go-graphql-client/ident"
)

func constructQuery(v interface{}, variables map[string]interface{}, name string) string {
	query := query(v)
	if len(variables) > 0 {
		return "query " + name + "(" + queryArguments(variables, false) + ")" + query
	}

	if name != "" {
		return "query " + name + query
	}
	return query
}

func constructMutation(v interface{}, variables map[string]interface{}, name string) string {
	query := query(v)
	if len(variables) > 0 {
		return "mutation " + name + "(" + queryArguments(variables, true) + ")" + query
	}
	if name != "" {
		return "mutation " + name + query
	}
	return "mutation" + query
}

func constructSubscription(v interface{}, variables map[string]interface{}, name string) string {
	query := query(v)
	if len(variables) > 0 {
		return "subscription " + name + "(" + queryArguments(variables, false) + ")" + query
	}
	if name != "" {
		return "subscription " + name + query
	}
	return "subscription" + query
}

// queryArguments constructs a minified arguments string for variables.
//
// E.g., map[string]interface{}{"a": Int(123), "b": NewBoolean(true)} -> "$a:Int!$b:Boolean".
func queryArguments(variables map[string]interface{}, isMutation bool) string {
	// Sort keys in order to produce deterministic output for testing purposes.
	// TODO: If tests can be made to work with non-deterministic output, then no need to sort.
	keys := make([]string, 0, len(variables))
	for k := range variables {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, k := range keys {
		io.WriteString(&buf, "$")
		io.WriteString(&buf, k)
		io.WriteString(&buf, ":")
		writeArgumentType(&buf, reflect.TypeOf(variables[k]), true, isMutation)
		// Don't insert a comma here.
		// Commas in GraphQL are insignificant, and we want minified output.
		// See https://facebook.github.io/graphql/October2016/#sec-Insignificant-Commas.
	}
	return buf.String()
}

// writeArgumentType writes a minified GraphQL type for t to w.
// value indicates whether t is a value (required) type or pointer (optional) type.
// If value is true, then "!" is written at the end of t.
func writeArgumentType(w io.Writer, t reflect.Type, value, isMutation bool) {
	if t.Kind() == reflect.Ptr && !isMutation {
		// Pointer is an optional type, so no "!" at the end of the pointer's underlying type.
		writeArgumentType(w, t.Elem(), false, isMutation)
		return
	} else if t.Kind() == reflect.Ptr {
		writeArgumentType(w, t.Elem(), true, isMutation)
		return
	}

	switch t.Kind() {
	case reflect.Slice, reflect.Array:
		// List. E.g., "[Int]".
		io.WriteString(w, "[")
		writeArgumentType(w, t.Elem(), true, isMutation)
		io.WriteString(w, "]")
	default:
		// Named type. E.g., "Int".
		name := t.Name()
		switch name {
		// case "string":
		// 	name = "ID"
		case "GqlBool":
			name = "Boolean"
		case "GqlFloat64":
			name = "Float!"
		case "GqlInt64":
			name = "Int!"
		case "GqlString":
			name = "String!"
		case "GqlTime":
			name = "DateTime!"
		}
		io.WriteString(w, name)
	}

	if value {
		// Value is a required type, so add "!" to the end.
		io.WriteString(w, "!")
	}
}

// query uses writeQuery to recursively construct
// a minified query string from the provided struct v.
//
// E.g., struct{Foo Int, BarBaz *Boolean} -> "{foo,barBaz}".
func query(v interface{}) string {
	var buf bytes.Buffer
	seen := make(map[string]struct{})
	writeQuery(&buf, reflect.TypeOf(v), false, "")
	fmt.Println(seen)
	return buf.String()
}

// writeQuery writes a minified query for t to w.
// If inline is true, the struct fields of t are inlined into parent struct.
// Seen is used to stop infinite loops
func writeQuery(w io.Writer, t reflect.Type, inline bool, inverseName string) {
	switch t.Kind() {
	case reflect.Ptr, reflect.Slice:
		writeQuery(w, t.Elem(), false, inverseName)
	case reflect.Struct:
		// If the type implements json.Unmarshaler, it's a scalar. Don't expand it.
		if reflect.PtrTo(t).Implements(jsonUnmarshaler) {
			return
		}
		if !inline {
			io.WriteString(w, "{")
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			value, ok := f.Tag.Lookup("graphql")
			if value == "-" {
				//skip (this is 'omit')
				continue
			} else if value == "" {
				value = ident.ParseMixedCaps(f.Name).ToLowerCamelCase()
			}
			if inverseName == value {
				continue //Don't allow recursion
			}
			thisInverseName, _ := f.Tag.Lookup("hasInverse")
			// if thisInverseName != "" {
			// 	fmt.Println()
			// }
			if i != 0 {
				io.WriteString(w, ",")
			}
			inlineField := f.Anonymous && !ok
			if !inlineField {
				io.WriteString(w, value)
			}
			writeQuery(w, f.Type, inlineField, thisInverseName)
		}
		if !inline {
			io.WriteString(w, "}")
		}
	}
}

var jsonUnmarshaler = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
