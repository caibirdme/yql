grammar Yql;

query: expr;

expr: booleanExpr       #boolExpr
    | expr 'and' expr   #andExpr
    | expr 'or' expr    #orExpr
    | '(' expr ')'      #embbedExpr
    ;

booleanExpr: leftexpr op=('='|'!='|'>'|'<'|'>='|'<=') value
    | leftexpr op=('in'|'!in'|'∩'|'!∩') '(' value (',' value)* ')'
    ;
leftexpr: FIELDNAME (('.' FUNC '()')+)?
    ;
value: STRING | INT | FLOAT | TRUE | FALSE;

TRUE: 'true';
FALSE: 'false';
FUNC: 'count'|'sum'|'avg'|'max'|'min';
FIELDNAME: ([a-zA-Z]|'_')+;
STRING: '\'' .*? '\'';
fragment DIGIT: [0-9];
INT: ('+'|'-')? DIGIT+;
FLOAT: ('+'|'-')? DIGIT+ '.' DIGIT*;
WS: [ \t\r\n]+ -> skip;
