grammar Gramatica;

// ANALIZADOR

ini
    :funcionMetodo EOF
;

funcionMetodo
        :PUBLIC MAIN DPUNTOS ID CABRE PRINCIPAL PARA STRING LLABRE LLACIE PARC DPUNTOS METODO CABRE cuerpoFunciones CCIER CCIER
;

cuerpoFunciones
                :cuerpoFunciones declaraciones
                |declaraciones
;

declaraciones
            :sentenceIf
            |IMPRIMIR PARA listaExpresiones PARC PCOMA
            |variable
;

listaExpresiones
            :listaExpresiones COMA expresion
            |expresion
;

sentenceIf
    :IF DPUNTOS PARA expresion PARC CABRE cuerpoFunciones CCIER sentenceElse
;

sentenceElse
    :ELSE CABRE cuerpoFunciones CCIER
    |ELSE sentenceIf
    |
;

variable
        :DECLARAR identificadores tipo PCOMA
        |DECLARAR identificadores tipo IGUAL expresion PCOMA
        |identificadores IGUAL expresion PCOMA
;

identificadores
            : identificadores COMA identificadores
            | ID
;

tipo
    :INT
    |BOOLEAN
    |DOUBLE
    |CHAR
    |STRING
    |REAL
;

expresion
    :MENOS expresion
    |expresion AND expresion
    |expresion OR expresion

    //ARITMÉTICOS
    |expresion MAS expresion
    |expresion MENOS expresion
    |expresion POR expresion
    |expresion DIVIDIR expresion
    |expresion MOD expresion
    |PARA expresion PARC

    //RELACIONALES
    |expresion IGUALI expresion
    |expresion DIFERENCIA expresion
    |expresion MAYI expresion
    |expresion MENI expresion
    |expresion MAY expresion
    |expresion MEN expresion
    |ID
    |primitivos
;

primitivos
    : ENTERO
    | DECIMAL
    | TRUE
    | FALSE
    | CADENA
;



// LEXEMAS
fragment A                  : [aA] ;
fragment B                  : [bB] ;
fragment C                  : [cC] ;
fragment D                  : [dD] ;
fragment E                  : [eE] ;
fragment F                  : [fF] ;
fragment G                  : [gG] ;
fragment H                  : [hH] ;
fragment I                  : [iI] ;
fragment J                  : [jJ] ;
fragment K                  : [kK] ;
fragment L                  : [lL] ;
fragment M                  : [mM] ;
fragment N                  : [nN] ;
fragment O                  : [oO] ;
fragment P                  : [pP] ;
fragment Q                  : [qQ] ;
fragment R                  : [rR] ;
fragment S                  : [sS] ;
fragment T                  : [tT] ;
fragment U                  : [uU] ;
fragment V                  : [vV] ;
fragment W                  : [wW] ;
fragment X                  : [xX] ;
fragment Y                  : [yY] ;
fragment Z                  : [zZ] ;

//PALABRAS RESERVADAS
IF                      : I F;
ELSE                    : E L S E;
IMPRIMIR                : S E N T E N C I A PUNTO C O N S O L A ;
MAIN                    : M A I N;
PUBLIC                  : P U B L I C O;
CLASS                   : C L A S S;
DECLARAR                : D E C L A R A R;
DOUBLE                  : D O U B L E;

//TIPOS DE DATOS
STRING                  : S T R I N G;
INT                     : I N T;
REAL                    : R E A L;
BOOLEAN                 : B O O L;
PRINCIPAL               : P R I N C I P A L;
METODO                  : M E T O D O;


//ARITMÉTICAS
MOD                     : '%' ;
POR                     : '*' ;
DIVIDIR                 : '/' ;
MAS                     : '+' ;
MENOS                   : '-' ;
IGUAL                   : '=' ;

//LOGICAS
AND                     : '&&' ;
OR                      : '||' ;
NOT                     : '!'  ;

//RELACIONALES
MENI                    : '<='  ;
MEN                     : '<'   ;
MAY                     : '>'   ;
MAYI                    : '>='  ;
IGUALI                  : '=='  ;
DIFERENCIA              : '!='  ;

//EXTRAS
GUIONB                  : '_';
PUNTO                   : '.'   ;
PARA                    : '('   ;
PARC                    : ')'   ;
LLABRE                  : '{'   ;
LLACIE                  : '}'   ;
CABRE                   : '['   ;
CCIER                   : ']'   ;
DPUNTOS                 : ':'   ;
PCOMA                   : ';'   ;
COMA                    : ','   ;

//LITERALES
CADENA                  : ('"' ( '\\'. | '""' | ~('"'| '\\') )* '"');
ENTERO                  : [0-9]+;
DECIMAL                 : [0-9]+ PUNTO [0-9]*| PUNTO [0-9]+;
TRUE                    : T R U E;
FALSE                   : F A L S E;
ID                      : ([a-zA-ZáéíóúÁÉÍÓÚ]+( GUIONB | [0-9A-Za-záéíóúÁÉÍÓÚ])*);

// IGNORED TOKENS
ESPACIOS                       : [ \t\r\n]+    -> skip;



/*


("//".*\r\n)|("//".*\n)|("//".*\r)              return "COMENTARIO";
"/*""/"*([^* /]|[^*]"/"|"*"[^/])*"*"*"* /"        return "COMENTARIO";

%left 'INTERROGACION' 'DPUNTOS'
%left 'OR'
%left 'AND'
%left 'DIFERENCIA' 'IGUALACION'
%left 'MENOR_IGUAL' 'MAYOR_IGUAL' 'MENOR' 'MAYOR'
%left 'MAS' 'MENOS'
%left 'POR' 'DIVIDIR'
%left 'POTENCIA' 'MOD'
%left 'COMA' 'IGUAL'
%left UMENOS
%right 'NOT'
*/