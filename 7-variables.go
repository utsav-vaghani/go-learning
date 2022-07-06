package main

/**
	Variables

	Declaration

		1. Without value
			var name string
		
		2. With value
			var name string = "Utsav"
			var name = "Utsav" // GO will consider type of name as `string` automatically
		
		3. Shorthand declaration
			name := "Utsav" // GO will consider type of name as `string` automatically

	Zero values

	- Every variables declared without value, will have zero values

		types		zero-value
		int			0
		float		0.0
		string		""
		bool		false
		interface{}	nil
		pointer		nil
		struct		struct with zero value of fields
**/