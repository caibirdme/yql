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
	result := yql.Match(rawYQL, map[string]interface{}{
		"name":  "deen",
		"age":   int64(23),
		"hobby": "basketball",
		"score": int64(100),
	})
	fmt.Println(result)
	rawYQL = `score ∩ (7,1,9,5,3)`
	result = yql.Match(rawYQL, map[string]interface{}{
		"score": []int64{3, 100, 200},
	})
	fmt.Println(result)
	rawYQL = `score in (7,1,9,5,3)`
	result = yql.Match(rawYQL, map[string]interface{}{
		"score": []int64{3, 5, 2},
	})
	fmt.Println(result)
	//Output:
	//true
	//true
	//false
```

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
```
grammar Yql;

query: expr;

expr: booleanExpr       #boolExpr
    | expr 'and' expr   #andExpr
    | expr 'or' expr    #orExpr
    | '(' expr ')'      #embbedExpr
    ;

booleanExpr: FIELDNAME op='=' value
    | FIELDNAME op='!=' value
    | FIELDNAME op='>' value
    | FIELDNAME op='<' value
    | FIELDNAME op='>=' value
    | FIELDNAME op='<=' value
    | FIELDNAME op='in' '(' value (',' value)* ')'
    | FIELDNAME op='!in' '(' value (',' value)* ')'
    | FIELDNAME op='∩' '(' value (',' value)* ')'
    | FIELDNAME op='!∩' '(' value (',' value)* ')'
    ;

value: STRING | INT | FLOAT | TRUE | FALSE;

TRUE: 'true';
FALSE: 'false';
FIELDNAME: [a-zA-Z]+;
STRING: '\'' .*? '\'';
fragment DIGIT: [0-9];
INT: ('+'|'-')? DIGIT+;
FLOAT: ('+'|'-')? DIGIT+ '.' DIGIT*;
WS: [ \t\r\n]+ -> skip;

```