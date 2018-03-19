### YQL(Yet another-Query-Language)
[![Build Status](https://www.travis-ci.org/caibirdme/yql.svg?branch=master)](https://www.travis-ci.org/caibirdme/yql)
[![GoDoc](https://godoc.org/github.com/caibirdme/yql?status.svg)](https://godoc.org/github.com/caibirdme/yql)


YQL is very similar with the `where` part of sql. You can see it as another sql which also support comparison between two sets. YQL have little concepts thus you can use it well short after reading the examples.Though it's designed for rule engine, it can be widely used in your code logic.

### Install
`go get github.com/caibirdme/yql`

### Exmaple
See more examples in the `yql_test.go` and godoc.

``` go
	rawYQL := `name='deen' and age>=23 and (hobby in ('soccer', 'swim') or score>90))`
	result, _ := yql.Match(rawYQL, map[string]interface{}{
		"name":  "deen",
		"age":   int64(23),
		"hobby": "basketball",
		"score": int64(100),
	})
	fmt.Println(result)
	rawYQL = `score ∩ (7,1,9,5,3)`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int64{3, 100, 200},
	})
	fmt.Println(result)
	rawYQL = `score in (7,1,9,5,3)`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int64{3, 5, 2},
	})
	fmt.Println(result)
	rawYQL = `score.sum() > 10`
	result, _ = yql.Match(rawYQL, map[string]interface{}{
		"score": []int{1, 2, 3, 4, 5},
	})
	fmt.Println(result)
	//Output:
	//true
	//true
	//false
	//true
```

And In most cases, you can use `Rule` to cache the AST and then use `Match` to get the result, which could avoid hundreds of thousands of repeated parsing process.

```go
	rawYQL := `name='deen' and age>=23 and (hobby in ('soccer', 'swim') or score>90)`
	ruler,_ := yql.Rule(rawYQL)

	result, _ := ruler.Match(map[string]interface{}{
		"name":  "deen",
		"age":   23,
		"hobby": "basketball",
		"score": int64(100),
	})
	fmt.Println(result)
	result, _ = ruler.Match(map[string]interface{}{
		"name":  "deen",
		"age":   23,
		"hobby": "basketball",
		"score": int64(90),
	})
	fmt.Println(result)
	//Output:
	//true
	//false
```

Though the to be matched data is the type of `map[string]interface{}`, there're only 5 types supported:
* int
* int64
* float64
* string
* bool

### Helpers
In `score.sum() > 10`, `sum` is a helper function which adds up all the numbers in score, which also means the type of score must be one of the []int,[]int64 or []float64.

This repo is in the early stage, so now there are just a few helpers, feel free to create an issue about your needs. Supported helpers are listed below:
* sum: ...
* count: return the length of a slice or 1 if not a slice
* avg: return the average number of a slice(`float64(total)/float64(len(slice))`)
* max: return the maximum number in a slice
* min: return the minimum number in a slice

### Usage scenario
Obviously, it's easy to use in rule engine.
```go
var handlers = map[int]func(map[string]interface{}){
	1: sendEmail,
	2: sendMessage,
	3: alertBoss,
}

data := resolvePostParamsFromRequest(request)
rules := getRulesFromDB(sql)

for _,rule := range rules {
	if success,_ := yql.Match(rule.YQL, data); success {
		handler := handlers[rule.ID]
		handler(data)
		break
	}
}
```

Also, it can be used in your daily work, which could significantly reduce the deeply embebbed `if else` statements:
```go
func isVIP(user User) bool {
	rule := fmt.Sprintf("monthly_vip=true and now<%s or eternal_vip=1 or ab_test!=false", user.ExpireTime)
	ok,_ := yql.Match(rule, map[string]interface{}{
		"monthly_vip": user.IsMonthlyVIP,
		"now": time.Now().Unix(),
		"eternal_vip": user.EternalFlag,
		"ab_test": isABTestMatched(user),
	})
	return ok
}
```

Even, you can use `json.Marshal` to generate the map[string]interface{} if you don't want to write it manually. Make sure the structure tag should be same as the name in rawYQL.

### Syntax
See [grammar file](./internal/grammar/Yql.g4)

### Compatibility promise
The API `Match`is stable now. Its grammar won't change any more, and what I only will do next is to optimize, speed up and add more helpers if needed.


### Further Trial
Though it's kinder difficult to create a robust new Go compiler, there're still some interesting things could do. For example, bringing lambda function in Go which maybe look like:
```go
var scores = []int{1,2,3,4,5,6,7,8,9,10}
newSlice := yql.Filter(`(v) => v % 2 == 0`).Map(`(v) => v*v`).Call(scores).Interface()
//[]int{4,16,36,64,100}
```
If the lambda function won't change all time, it can be cached like opcode, which is as fast as the compiled code. And in most cases, who care?(pythoner?)

It's not easy but interesting, isn't it? Welcome to join me, open some issues and talk about your ideas with me. Maybe one day it can become a pre-compile tool like [babel](http://babeljs.io/) in javascript.  

#### Attention
`Lambda expression` now is in its very early stage, **DO NOT USE IT IN PRODUCTION**.

You can take a quick preview in [test case](/lambda/lambda_test.go)

```go
type Student struct {
	Age  int
	Name string
}

var students = []Student{
	Student{
		Name: "deen",
		Age:  24,
	},
	Student{
		Name: "bob",
		Age:  22,
	},
	Student{
		Name: "alice",
		Age:  23,
	},
	Student{
		Name: "tom",
		Age:  25,
	},
	Student{
		Name: "jerry",
		Age:  20,
	},
}

t = yql.Filter(`(v) => v.Age > 23 || v.Name == "alice"`).Call(students).Interface()
res,_ := t.([]Student)
// res: Student{"deen",24} Student{"alice", 23} Student{"tom", 25}
```

Chainable
```go
dst := []int{1, 2, 3, 4, 5, 6, 7}
r := Filter(`(v) => v > 3 && v <= 7`).Map(`(v) =>  v << 2`).Filter(`(v) => v % 8 == 0`).Call(dst)
s, err := r.Interface()
ass := assert.New(t)
ass.NoError(err)
ass.Equal([]int{16, 24}, s)
```