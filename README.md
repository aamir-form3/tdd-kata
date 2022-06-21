## TDD Kata

### Bowling
This kata is taken from Uncle Bob's book

### HTTP Server

### Stack
Inspired by Uncle Bob's kata for stack

### Roman Numerals
Convert arabic numbers to roman

### Intellij template 

```go
type $suite_name$TestSuite struct {
	suite.Suite
}

func Test$suite_name$Suite(t *testing.T) {
	suite.Run(t, new($suite_name$TestSuite))
}

func (s *$suite_name$TestSuite) SetupTest() {
}

func (s *$suite_name$TestSuite) $END$() {
}
```

```go
func (s *$VAR$)Test$NAME$() {
    $END$
}
```