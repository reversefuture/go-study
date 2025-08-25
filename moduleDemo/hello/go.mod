module example.com/hello

go 1.25.0

// redirect package to local
replace example.com/greetings => ../greetings

// reference local module with pseudo-version number 
require example.com/greetings v0.0.0-00010101000000-000000000000
