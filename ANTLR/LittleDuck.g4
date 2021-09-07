grammar LittleDuck;

/*
 Lexer Rules
 */

fragment NON_DIGIT: [a-zA-Z_];
fragment DIGIT: [0-9];
fragment NON_ZERO_DIGIT: [1-9];
fragment HEXADECIMAL_DIGIT: [0-9a-fA-F];
fragment HEXQUAD: HEXADECIMAL_DIGIT HEXADECIMAL_DIGIT HEXADECIMAL_DIGIT HEXADECIMAL_DIGIT;
fragment UNIVERSAL_CHARACTER_NAME
    :   '\\u' HEXQUAD
	|   '\\U' HEXQUAD HEXQUAD;
fragment ID_NON_DIGIT
    :   NON_DIGIT
    |   UNIVERSAL_CHARACTER_NAME;
fragment INT_PART: DIGIT+;
fragment FRACTION: '.' DIGIT+;
fragment EXPONENT: [eE] [+-]? DIGIT+;
fragment POINT_FLOAT
    :   INT_PART? FRACTION
    |   INT_PART '.';
fragment EXPONENT_FLOAT: ( INT_PART | POINT_FLOAT ) EXPONENT;

// STRING : '"'(ESC|.)*? '"';
// fragment ESC: '\\"' | '\\\\';

// Words
PROGRAM: 'programa';
VAR: 'var';
INT: 'int';
FLOAT: 'float';
PRINT: 'print';
IF: 'if';
ELSE: 'else';

// Symbols
COMMA: ',';
COLON: ':';
SEMICOLON: ';';
LEFT_CURLY: '{';
RIGHT_CURLY: '}';
LEFT_ROUND: '(';
RIGHT_ROUND: ')';

STAR: '*';
PLUS: '+';
MINUS: '-';
DIV: '/';
EQUAL: '=';

LESS_THAN: '<';
GREATER_THAN: '>';
NOT_EQUAL: '<>';

// Complex
LET_ID: ID_NON_DIGIT (ID_NON_DIGIT | DIGIT)*;
LET_STRING: '"' ('\\' ["\\] | ~["\\\r\n])* '"';
LET_INT
    :   '0'
    |   NON_ZERO_DIGIT (DIGIT)*;
LET_FLOAT
    :   POINT_FLOAT
    |   EXPONENT_FLOAT;

// Ignore
WHITESPACE: [ \n\t\r]+ -> skip;

/*
 Parser Rules
 */
programa: PROGRAM LET_ID SEMICOLON programa_vars bloque;
programa_vars
    :   vars
    |   empty;

vars: VAR vars_id vars_more;
vars_id: LET_ID vars_id_next COLON tipo SEMICOLON;
vars_id_next
    :   COMMA LET_ID vars_id_next
    |   empty;
vars_more
    :   vars_id vars_more
    |   empty;

bloque: LEFT_CURLY bloque_more RIGHT_CURLY;
bloque_more
    :   estatuto bloque_more
    |   empty;

tipo
    :   INT 
    |   FLOAT;

estatuto
    :   asignacion
    |   condicion
    |   escritura;

asignacion: LET_ID EQUAL expresion SEMICOLON;

condicion: IF LEFT_ROUND expresion RIGHT_ROUND bloque condicion_else SEMICOLON;
condicion_else
    :   ELSE bloque
    |   empty;

escritura: PRINT LEFT_ROUND escritura_params RIGHT_ROUND SEMICOLON;
escritura_params
    :   LET_STRING escritura_params_next
    |   expresion escritura_params_next;
escritura_params_next
    :   COMMA escritura_params
    |   empty;

expresion: exp expresion_next;
expresion_next
    :   GREATER_THAN exp
    |   LESS_THAN exp
    |   NOT_EQUAL exp
    |   empty;

exp: termino exp_next;
exp_next
    :   PLUS exp
    |   MINUS exp
    |   empty;

termino: factor termino_next;
termino_next
    :   STAR termino
    |   DIV termino
    |   empty;

factor
    :   LEFT_ROUND expresion RIGHT_ROUND
    |   factor_sign cte;
factor_sign
    :   PLUS
    |   MINUS
    |   empty;

cte
    :   LET_ID
    |   LET_INT
    |   LET_FLOAT;

empty: ;