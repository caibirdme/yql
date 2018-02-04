/*
Package yql implements sql-like grammar for rule engines.

yql support compare operators like:

	=
	!=
	>
	>=
	<
	<=
	in
	!in
	∩ // intersection (not the letter n)
	!∩

below are legal yqls:

	a=1 and b>2 and (c<=3 or d in ('boy','girl')) or e ∩ (1,2,3)
	a=boy and b=1 or c in ('boy','girl')
	// boy will be interpreted as string if the type of a is string
	// 1 will be interpreted according to b's type

for more details,see yql_test.go
*/
package yql
