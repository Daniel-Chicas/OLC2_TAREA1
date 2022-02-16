grammar Gramatica;

options{
    tokenVocab = GramaticaLexer;
}

@header {
    import "OLC2_TAREA1/AST/Entornos"
    import "OLC2_TAREA1/AST/Interfaces"
    import "OLC2_TAREA1/AST/Expresion"
    import "OLC2_TAREA1/AST/Funcion"
    import arrayList "github.com/colegno/arraylist"
    import "OLC2_TAREA1/AST/Instruccion"
}



ini returns [*arrayList.List lista]
        :PUBLIC MAIN DPUNTOS ID CABRE PRINCIPAL PARA STRING LLABRE LLACIE PARC DPUNTOS METODO CABRE instrucciones CCIER CCIER {$lista = $instrucciones.l}
;

instrucciones returns [*arrayList.List l]
          @init{
            $l =  arrayList.New()
          }
          : e += declaraciones*  {
                                    listInt := localctx.(*InstruccionesContext).GetE()
                                        for _, e := range listInt {
                                            $l.Add(e.GetInstr())
                                        }
                                  }
;

declaraciones returns [Interfaces.Instruccion instr]
            :sentenceIf                         {$instr = $sentenceIf.instr}
            |IMPRIMIR PARA expresion PARC PCOMA {$instr = Funcion.NewImprimir($expresion.p)}
            |variable                           {$instr = $variable.instr}
;

sentenceIf  returns [Interfaces.Instruccion instr]
    :IF DPUNTOS PARA expresion PARC CABRE instrucciones CCIER sentenceElse  {$instr = Funcion.NewIf($IF.line, localctx.(*SentenceIfContext).Get_IF().GetColumn(), $expresion.p, $instrucciones.l, $sentenceElse.l)}
;

sentenceElse returns[*arrayList.List l]
    @init{$l =  arrayList.New()}
    :ELSE CABRE instrucciones CCIER     {$l = $instrucciones.l}
    |ELSE sentenceIf                    {$l.Add($sentenceIf.instr)}
    |                                   {$l = nil}
;

variable returns[Interfaces.Instruccion instr]
        :DECLARAR identificadores tipo PCOMA                        {$instr = Instruccion.NuevaDeclaracion($DECLARAR.line, localctx.(*VariableContext).Get_DECLARAR().GetColumn(), $identificadores.lista, $tipo.T)}
        |DECLARAR identificadores tipo IGUAL expresion PCOMA        {$instr = Instruccion.NuevaDeclaracionInicio($IGUAL.line, localctx.(*VariableContext).Get_IGUAL().GetColumn(), $identificadores.lista, $tipo.T, $expresion.p)}
        |identificadores IGUAL expresion PCOMA                      {$instr = Instruccion.NuevaDeclaracionInicio($IGUAL.line, localctx.(*VariableContext).Get_IGUAL().GetColumn(), $identificadores.lista, Entornos.NULL, $expresion.p)}
;
//$ID.line, localctx.(*IdentificadoresContext).Get_ID().GetColumn(),
identificadores returns[*arrayList.List lista]
        @init{
            $lista = arrayList.New()
        }
            : sub = identificadores COMA ID                         {
                                                                        $sub.lista.Add(Expresion.NuevoIdentificador($ID.line, localctx.(*IdentificadoresContext).Get_ID().GetColumn(), $ID.text))
                                                                        $lista = $sub.lista
                                                                    }
            | ID                                                    {   $lista.Add(Expresion.NuevoIdentificador($ID.line, localctx.(*IdentificadoresContext).Get_ID().GetColumn(), $ID.text))}
;

tipo returns[Entornos.TipoDato T]
    :INT                                                            {$T =Entornos.INTEGER}
    |BOOLEAN                                                        {$T =Entornos.BOOLEAN}
    |STRING                                                         {$T =Entornos.STRING}
    |REAL                                                           {$T =Entornos.REAL}
;

expresion returns[Interfaces.Expresion p]
    :logicas                                {$p = $logicas.p}
    |relacionales                           {$p = $relacionales.p}
    |aritmeticas                            {$p = $aritmeticas.p}
;

logicas returns[Interfaces.Expresion p]
    : op= NOT  opDe = relacionales                                 {$p = Expresion.NewLogico($op.line, localctx.(*LogicasContext).GetOp().GetColumn(), nil,$op.text,$opDe.p,false)}
    | opIz = relacionales op=( AND | OR ) opDe = relacionales      {$p = Expresion.NewLogico($op.line, localctx.(*LogicasContext).GetOp().GetColumn(), $opIz.p,$op.text,$opDe.p,false)}
    | relacionales                                                 {$p = $relacionales.p}
;

relacionales returns[Interfaces.Expresion p]
    : opIz = relacionales op=( MENI | MAYI | MEN | MAY | IGUALI | DIFERENCIA ) opDe = relacionales      {$p = Expresion.NewRelacional($op.line, localctx.(*RelacionalesContext).GetOp().GetColumn(), $opIz.p,$op.text,$opDe.p,false)}
    | aritmeticas                                                                                       {$p = $aritmeticas.p}
;

aritmeticas returns[Interfaces.Expresion p]
    : op='-' opDe = aritmeticas                                 {$p = Expresion.NewAritmetica($op.line, localctx.(*AritmeticasContext).GetOp().GetColumn(),  nil,$op.text,$opDe.p,true)}
    | opIz = aritmeticas op=MOD opDe = aritmeticas              {$p = Expresion.NewAritmetica($op.line, localctx.(*AritmeticasContext).GetOp().GetColumn(), $opIz.p,$op.text,$opDe.p,false)}
    | opIz = aritmeticas op=(POR|DIVIDIR) opDe = aritmeticas    {$p = Expresion.NewAritmetica($op.line, localctx.(*AritmeticasContext).GetOp().GetColumn(), $opIz.p,$op.text,$opDe.p,false)}
    | opIz = aritmeticas op=(MAS|MENOS) opDe = aritmeticas      {$p = Expresion.NewAritmetica($op.line, localctx.(*AritmeticasContext).GetOp().GetColumn(), $opIz.p,$op.text,$opDe.p,false)}
    | primitivos                                                {$p = $primitivos.p}
    | PARA expresion PARC                                       {$p = $expresion.p}
;

primitivos returns[Interfaces.Expresion p]
    : ENTERO                                                {
                                                                num, err := strconv.Atoi($ENTERO.text)
                                                                if err != nil{
                                                                    fmt.Println(err)
                                                                }
                                                                $p = Expresion.NewPrimitivo($ENTERO.line, localctx.(*PrimitivosContext).Get_ENTERO().GetColumn(), num, Entornos.INTEGER)
                                                            }
    | DECIMAL                                               {
                                                                num, err := strconv.ParseFloat($DECIMAL.text, 64)
                                                                if err != nil{
                                                                    fmt.Println(err)
                                                                }
                                                                $p = Expresion.NewPrimitivo($DECIMAL.line, localctx.(*PrimitivosContext).Get_DECIMAL().GetColumn(), num, Entornos.REAL)
                                                            }
    | TRUE                                                  {
                                                                $p = Expresion.NewPrimitivo($TRUE.line, localctx.(*PrimitivosContext).Get_TRUE().GetColumn(), $TRUE.text, Entornos.BOOLEAN)
                                                            }
    | FALSE                                                 {
                                                                $p = Expresion.NewPrimitivo($FALSE.line, localctx.(*PrimitivosContext).Get_FALSE().GetColumn(), $FALSE.text, Entornos.BOOLEAN)
                                                            }
    | CADENA                                                {
                                                                str:= $CADENA.text[1:len($CADENA.text)-1]
                                                                $p = Expresion.NewPrimitivo($CADENA.line, localctx.(*PrimitivosContext).Get_CADENA().GetColumn(), str,Entornos.STRING)
                                                            }
    | ID                                                    {$p = Expresion.NewAcceso($ID.line, localctx.(*PrimitivosContext).Get_ID().GetColumn(), $ID.text)}
;


//TOKENS

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
fragment ESC_SEQ            : '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ');

//PALABRAS RESERVADAS
IF                      : I F;
ELSE                    : E L S E;
IMPRIMIR                : S E N T E N C I A PUNTO C O N S O L A ;
MAIN                    : M A I N;
PUBLIC                  : P U B L I C O;
CLASS                   : C L A S S;
DECLARAR                : D E C L A R A R;

//TIPOS DE DATOS
STRING                  : S T R I N G;
INT                     : I N T;
REAL                    : R E A L;
BOOLEAN                 : B O O L;
PRINCIPAL               : P R I N C I P A L;
METODO                  : M E T O D O;

//LITERALES
CADENA                  : ('"' ( '\\'. | '""' | ~('"'| '\\') )* '"');
ENTERO                  : [0-9]+;
DECIMAL                 : [0-9]+ PUNTO [0-9]*| PUNTO [0-9]+;
TRUE                    : T R U E;
FALSE                   : F A L S E;
ID                      : ([a-zA-ZáéíóúÁÉÍÓÚ]+( GUIONB | [0-9A-Za-záéíóúÁÉÍÓÚ])*);

//RELACIONALES
MENI                    : '<='  ;
MAYI                    : '>='  ;
MEN                     : '<'   ;
MAY                     : '>'   ;
IGUALI                  : '=='  ;
DIFERENCIA              : '!='  ;

//LOGICAS
AND                     : '&&' ;
OR                      : '||' ;
NOT                     : '!'  ;

//ARITMÉTICAS
MOD                     : '%' ;
POR                     : '*' ;
DIVIDIR                 : '/' ;
MAS                     : '+' ;
MENOS                   : '-' ;
IGUAL                   : '=' ;

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

// IGNORED TOKENS
ESPACIOS                       : [ \t\r\n]+    -> skip;
COMENTARIO_MUL: '(*' (~[/])+ '*)' -> skip;
COMENTARIO_LIN: '//'(~[\n])+ -> skip;