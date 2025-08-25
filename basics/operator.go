package main

import "fmt"

/*
Arithmetic operators:
Operator	Name	Description	Example	Try it
+	Addition	Adds together two values	x + y
-	Subtraction	Subtracts one value from another	x - y
*	Multiplication	Multiplies two values	x * y
/	Division	Divides one value by another	x / y
%	Modulus	Returns the division remainder	x % y
++	Increment	Increases the value of a variable by 1	x++
--	Decrement	Decreases the value of a variable by 1	x--

comparison operators:
Operator	Name	Example	Try it
==	Equal to	x == y
!=	Not equal	x != y
>	Greater than	x > y
<	Less than	x < y
>=	Greater than or equal to	x >= y
<=	Less than or equal to	x <= y

Logical operators:
Operator	Name	Description	Example	Try it
&& 	Logical and	Returns true if both statements are true	x < 5 &&  x < 10
|| 	Logical or	Returns true if one of the statements is true	x < 5 || x < 4
!	Logical not	Reverse the result, returns false if the result is true	!(x < 5 && x < 10)

Bitwise Operators:
Operator	Name	Description	Example	Try it
& 	AND	Sets each bit to 1 if both bits are 1	x & y
|	OR	Sets each bit to 1 if one of two bits is 1	x | y
 ^	XOR	Sets each bit to 1 if only one of two bits is 1	x ^ b
<<	Zero fill left shift	Shift left by pushing zeros in from the right	x << 2
>>	Signed right shift	Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off	x >> 2
*/

func mainOperator() {
	var x = 27
	// 27: 0001 1011, 3: 0000 0011
	fmt.Println(x & 3)  // 3
	fmt.Println(x | 3)  // 27
	fmt.Println(x ^ 3)  //24
	fmt.Println(x >> 3) // 3
	fmt.Println(x << 3) // 216
}
