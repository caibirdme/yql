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
INT: DIGIT+;
FLOAT: DIGIT+ '.' DIGIT*;
WS: [ \t\r\n]+ -> skip;
