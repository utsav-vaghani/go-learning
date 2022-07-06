package main

/**
	Types

	- GO has following basic(primitive/built-in) types.
		1. Numbers
			int  int8  int16  int32  int64
			uint uint8 uint16 uint32 uint64 uintptr
			float32 float64

		2. Boolean
			bool

		3. String
			string
			for character: rune // alias for int32

		4. Byte
			byte // alias for uint8

		5. interface{}
			It can hold any kind(type) of value.

	- Apart from the basic types, one can write own type with keyword `type`.
		type <type-name> <original-type>
			here <original-type> can be any basic type or [ struct or interface will be discussed later].

		ex.
			type MyFloat float64

			now we can use this type to declare any variable.
			ex.
				var num MyFloat = 1.3

	- On custom type we can define functions which is called methods.
		ex. on our MyFloat type 

		func (m MyFloat) converToInt() int {
			return int(m) // here m will be `1.3` and returned `1`
		}
**/
