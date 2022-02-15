// Generated from C:/Users/Daniel Chicas/Desktop/PRIMER_SEMESTRE_2022/COMPI2/LABORATORIO/TAREAS_LAB/OLC2_TAREA1/Backend/Gramatica\GramaticaLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GramaticaLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		IF=1, ELSE=2, IMPRIMIR=3, MAIN=4, PUBLIC=5, CLASS=6, DECLARAR=7, STRING=8, 
		INT=9, REAL=10, BOOLEAN=11, PRINCIPAL=12, METODO=13, CADENA=14, ENTERO=15, 
		DECIMAL=16, TRUE=17, FALSE=18, ID=19, MENI=20, MAYI=21, MEN=22, MAY=23, 
		IGUALI=24, DIFERENCIA=25, AND=26, OR=27, NOT=28, MOD=29, POR=30, DIVIDIR=31, 
		MAS=32, MENOS=33, IGUAL=34, GUIONB=35, PUNTO=36, PARA=37, PARC=38, LLABRE=39, 
		LLACIE=40, CABRE=41, CCIER=42, DPUNTOS=43, PCOMA=44, COMA=45, ESPACIOS=46;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", 
			"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "ESC_SEQ", 
			"IF", "ELSE", "IMPRIMIR", "MAIN", "PUBLIC", "CLASS", "DECLARAR", "STRING", 
			"INT", "REAL", "BOOLEAN", "PRINCIPAL", "METODO", "CADENA", "ENTERO", 
			"DECIMAL", "TRUE", "FALSE", "ID", "MENI", "MAYI", "MEN", "MAY", "IGUALI", 
			"DIFERENCIA", "AND", "OR", "NOT", "MOD", "POR", "DIVIDIR", "MAS", "MENOS", 
			"IGUAL", "GUIONB", "PUNTO", "PARA", "PARC", "LLABRE", "LLACIE", "CABRE", 
			"CCIER", "DPUNTOS", "PCOMA", "COMA", "ESPACIOS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "'<='", "'>='", "'<'", 
			"'>'", "'=='", "'!='", "'&&'", "'||'", "'!'", "'%'", "'*'", "'/'", "'+'", 
			"'-'", "'='", "'_'", "'.'", "'('", "')'", "'{'", "'}'", "'['", "']'", 
			"':'", "';'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "IF", "ELSE", "IMPRIMIR", "MAIN", "PUBLIC", "CLASS", "DECLARAR", 
			"STRING", "INT", "REAL", "BOOLEAN", "PRINCIPAL", "METODO", "CADENA", 
			"ENTERO", "DECIMAL", "TRUE", "FALSE", "ID", "MENI", "MAYI", "MEN", "MAY", 
			"IGUALI", "DIFERENCIA", "AND", "OR", "NOT", "MOD", "POR", "DIVIDIR", 
			"MAS", "MENOS", "IGUAL", "GUIONB", "PUNTO", "PARA", "PARC", "LLABRE", 
			"LLACIE", "CABRE", "CCIER", "DPUNTOS", "PCOMA", "COMA", "ESPACIOS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public GramaticaLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "GramaticaLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\60\u01a6\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t"+
		"=\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4"+
		"I\tI\4J\tJ\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t"+
		"\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3"+
		"\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3"+
		"\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\34\3\35\3\35\3\35\3"+
		"\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3"+
		"\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3#\3#\3#\3$\3"+
		"$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3(\3("+
		"\3(\3(\3(\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\7*\u012f"+
		"\n*\f*\16*\u0132\13*\3*\3*\3+\6+\u0137\n+\r+\16+\u0138\3,\6,\u013c\n,"+
		"\r,\16,\u013d\3,\3,\7,\u0142\n,\f,\16,\u0145\13,\3,\3,\6,\u0149\n,\r,"+
		"\16,\u014a\5,\u014d\n,\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3/\6/\u015b\n"+
		"/\r/\16/\u015c\3/\3/\7/\u0161\n/\f/\16/\u0164\13/\3\60\3\60\3\60\3\61"+
		"\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\64\3\65\3\65\3\65\3\66\3\66"+
		"\3\66\3\67\3\67\3\67\38\38\39\39\3:\3:\3;\3;\3<\3<\3=\3=\3>\3>\3?\3?\3"+
		"@\3@\3A\3A\3B\3B\3C\3C\3D\3D\3E\3E\3F\3F\3G\3G\3H\3H\3I\3I\3J\6J\u01a1"+
		"\nJ\rJ\16J\u01a2\3J\3J\2\2K\3\2\5\2\7\2\t\2\13\2\r\2\17\2\21\2\23\2\25"+
		"\2\27\2\31\2\33\2\35\2\37\2!\2#\2%\2\'\2)\2+\2-\2/\2\61\2\63\2\65\2\67"+
		"\29\3;\4=\5?\6A\7C\bE\tG\nI\13K\fM\rO\16Q\17S\20U\21W\22Y\23[\24]\25_"+
		"\26a\27c\30e\31g\32i\33k\34m\35o\36q\37s u!w\"y#{$}%\177&\u0081\'\u0083"+
		"(\u0085)\u0087*\u0089+\u008b,\u008d-\u008f.\u0091/\u0093\60\3\2\"\4\2"+
		"CCcc\4\2DDdd\4\2EEee\4\2FFff\4\2GGgg\4\2HHhh\4\2IIii\4\2JJjj\4\2KKkk\4"+
		"\2LLll\4\2MMmm\4\2NNnn\4\2OOoo\4\2PPpp\4\2QQqq\4\2RRrr\4\2SSss\4\2TTt"+
		"t\4\2UUuu\4\2VVvv\4\2WWww\4\2XXxx\4\2YYyy\4\2ZZzz\4\2[[{{\4\2\\\\||\t"+
		"\2\"#%%--/\60<<BB]_\4\2$$^^\3\2\62;\16\2C\\c|\u00a3\u00a3\u00ab\u00ab"+
		"\u00af\u00af\u00b5\u00b5\u00bc\u00bc\u00c5\u00c5\u0163\u0163\u201e\u201e"+
		"\u2032\u2032\uffff\uffff\17\2\62;C\\c|\u00a3\u00a3\u00ab\u00ab\u00af\u00af"+
		"\u00b5\u00b5\u00bc\u00bc\u00c5\u00c5\u0163\u0163\u201e\u201e\u2032\u2032"+
		"\uffff\uffff\5\2\13\f\17\17\"\"\2\u0196\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2"+
		"\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2"+
		"K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3"+
		"\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2"+
		"\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2"+
		"q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3"+
		"\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2"+
		"\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2\2\2\u008f"+
		"\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2\2\3\u0095\3\2\2\2\5\u0097\3\2\2"+
		"\2\7\u0099\3\2\2\2\t\u009b\3\2\2\2\13\u009d\3\2\2\2\r\u009f\3\2\2\2\17"+
		"\u00a1\3\2\2\2\21\u00a3\3\2\2\2\23\u00a5\3\2\2\2\25\u00a7\3\2\2\2\27\u00a9"+
		"\3\2\2\2\31\u00ab\3\2\2\2\33\u00ad\3\2\2\2\35\u00af\3\2\2\2\37\u00b1\3"+
		"\2\2\2!\u00b3\3\2\2\2#\u00b5\3\2\2\2%\u00b7\3\2\2\2\'\u00b9\3\2\2\2)\u00bb"+
		"\3\2\2\2+\u00bd\3\2\2\2-\u00bf\3\2\2\2/\u00c1\3\2\2\2\61\u00c3\3\2\2\2"+
		"\63\u00c5\3\2\2\2\65\u00c7\3\2\2\2\67\u00c9\3\2\2\29\u00cc\3\2\2\2;\u00cf"+
		"\3\2\2\2=\u00d4\3\2\2\2?\u00e6\3\2\2\2A\u00eb\3\2\2\2C\u00f3\3\2\2\2E"+
		"\u00f9\3\2\2\2G\u0102\3\2\2\2I\u0109\3\2\2\2K\u010d\3\2\2\2M\u0112\3\2"+
		"\2\2O\u0117\3\2\2\2Q\u0121\3\2\2\2S\u0128\3\2\2\2U\u0136\3\2\2\2W\u014c"+
		"\3\2\2\2Y\u014e\3\2\2\2[\u0153\3\2\2\2]\u015a\3\2\2\2_\u0165\3\2\2\2a"+
		"\u0168\3\2\2\2c\u016b\3\2\2\2e\u016d\3\2\2\2g\u016f\3\2\2\2i\u0172\3\2"+
		"\2\2k\u0175\3\2\2\2m\u0178\3\2\2\2o\u017b\3\2\2\2q\u017d\3\2\2\2s\u017f"+
		"\3\2\2\2u\u0181\3\2\2\2w\u0183\3\2\2\2y\u0185\3\2\2\2{\u0187\3\2\2\2}"+
		"\u0189\3\2\2\2\177\u018b\3\2\2\2\u0081\u018d\3\2\2\2\u0083\u018f\3\2\2"+
		"\2\u0085\u0191\3\2\2\2\u0087\u0193\3\2\2\2\u0089\u0195\3\2\2\2\u008b\u0197"+
		"\3\2\2\2\u008d\u0199\3\2\2\2\u008f\u019b\3\2\2\2\u0091\u019d\3\2\2\2\u0093"+
		"\u01a0\3\2\2\2\u0095\u0096\t\2\2\2\u0096\4\3\2\2\2\u0097\u0098\t\3\2\2"+
		"\u0098\6\3\2\2\2\u0099\u009a\t\4\2\2\u009a\b\3\2\2\2\u009b\u009c\t\5\2"+
		"\2\u009c\n\3\2\2\2\u009d\u009e\t\6\2\2\u009e\f\3\2\2\2\u009f\u00a0\t\7"+
		"\2\2\u00a0\16\3\2\2\2\u00a1\u00a2\t\b\2\2\u00a2\20\3\2\2\2\u00a3\u00a4"+
		"\t\t\2\2\u00a4\22\3\2\2\2\u00a5\u00a6\t\n\2\2\u00a6\24\3\2\2\2\u00a7\u00a8"+
		"\t\13\2\2\u00a8\26\3\2\2\2\u00a9\u00aa\t\f\2\2\u00aa\30\3\2\2\2\u00ab"+
		"\u00ac\t\r\2\2\u00ac\32\3\2\2\2\u00ad\u00ae\t\16\2\2\u00ae\34\3\2\2\2"+
		"\u00af\u00b0\t\17\2\2\u00b0\36\3\2\2\2\u00b1\u00b2\t\20\2\2\u00b2 \3\2"+
		"\2\2\u00b3\u00b4\t\21\2\2\u00b4\"\3\2\2\2\u00b5\u00b6\t\22\2\2\u00b6$"+
		"\3\2\2\2\u00b7\u00b8\t\23\2\2\u00b8&\3\2\2\2\u00b9\u00ba\t\24\2\2\u00ba"+
		"(\3\2\2\2\u00bb\u00bc\t\25\2\2\u00bc*\3\2\2\2\u00bd\u00be\t\26\2\2\u00be"+
		",\3\2\2\2\u00bf\u00c0\t\27\2\2\u00c0.\3\2\2\2\u00c1\u00c2\t\30\2\2\u00c2"+
		"\60\3\2\2\2\u00c3\u00c4\t\31\2\2\u00c4\62\3\2\2\2\u00c5\u00c6\t\32\2\2"+
		"\u00c6\64\3\2\2\2\u00c7\u00c8\t\33\2\2\u00c8\66\3\2\2\2\u00c9\u00ca\7"+
		"^\2\2\u00ca\u00cb\t\34\2\2\u00cb8\3\2\2\2\u00cc\u00cd\5\23\n\2\u00cd\u00ce"+
		"\5\r\7\2\u00ce:\3\2\2\2\u00cf\u00d0\5\13\6\2\u00d0\u00d1\5\31\r\2\u00d1"+
		"\u00d2\5\'\24\2\u00d2\u00d3\5\13\6\2\u00d3<\3\2\2\2\u00d4\u00d5\5\'\24"+
		"\2\u00d5\u00d6\5\13\6\2\u00d6\u00d7\5\35\17\2\u00d7\u00d8\5)\25\2\u00d8"+
		"\u00d9\5\13\6\2\u00d9\u00da\5\35\17\2\u00da\u00db\5\7\4\2\u00db\u00dc"+
		"\5\23\n\2\u00dc\u00dd\5\3\2\2\u00dd\u00de\5\177@\2\u00de\u00df\5\7\4\2"+
		"\u00df\u00e0\5\37\20\2\u00e0\u00e1\5\35\17\2\u00e1\u00e2\5\'\24\2\u00e2"+
		"\u00e3\5\37\20\2\u00e3\u00e4\5\31\r\2\u00e4\u00e5\5\3\2\2\u00e5>\3\2\2"+
		"\2\u00e6\u00e7\5\33\16\2\u00e7\u00e8\5\3\2\2\u00e8\u00e9\5\23\n\2\u00e9"+
		"\u00ea\5\35\17\2\u00ea@\3\2\2\2\u00eb\u00ec\5!\21\2\u00ec\u00ed\5+\26"+
		"\2\u00ed\u00ee\5\5\3\2\u00ee\u00ef\5\31\r\2\u00ef\u00f0\5\23\n\2\u00f0"+
		"\u00f1\5\7\4\2\u00f1\u00f2\5\37\20\2\u00f2B\3\2\2\2\u00f3\u00f4\5\7\4"+
		"\2\u00f4\u00f5\5\31\r\2\u00f5\u00f6\5\3\2\2\u00f6\u00f7\5\'\24\2\u00f7"+
		"\u00f8\5\'\24\2\u00f8D\3\2\2\2\u00f9\u00fa\5\t\5\2\u00fa\u00fb\5\13\6"+
		"\2\u00fb\u00fc\5\7\4\2\u00fc\u00fd\5\31\r\2\u00fd\u00fe\5\3\2\2\u00fe"+
		"\u00ff\5%\23\2\u00ff\u0100\5\3\2\2\u0100\u0101\5%\23\2\u0101F\3\2\2\2"+
		"\u0102\u0103\5\'\24\2\u0103\u0104\5)\25\2\u0104\u0105\5%\23\2\u0105\u0106"+
		"\5\23\n\2\u0106\u0107\5\35\17\2\u0107\u0108\5\17\b\2\u0108H\3\2\2\2\u0109"+
		"\u010a\5\23\n\2\u010a\u010b\5\35\17\2\u010b\u010c\5)\25\2\u010cJ\3\2\2"+
		"\2\u010d\u010e\5%\23\2\u010e\u010f\5\13\6\2\u010f\u0110\5\3\2\2\u0110"+
		"\u0111\5\31\r\2\u0111L\3\2\2\2\u0112\u0113\5\5\3\2\u0113\u0114\5\37\20"+
		"\2\u0114\u0115\5\37\20\2\u0115\u0116\5\31\r\2\u0116N\3\2\2\2\u0117\u0118"+
		"\5!\21\2\u0118\u0119\5%\23\2\u0119\u011a\5\23\n\2\u011a\u011b\5\35\17"+
		"\2\u011b\u011c\5\7\4\2\u011c\u011d\5\23\n\2\u011d\u011e\5!\21\2\u011e"+
		"\u011f\5\3\2\2\u011f\u0120\5\31\r\2\u0120P\3\2\2\2\u0121\u0122\5\33\16"+
		"\2\u0122\u0123\5\13\6\2\u0123\u0124\5)\25\2\u0124\u0125\5\37\20\2\u0125"+
		"\u0126\5\t\5\2\u0126\u0127\5\37\20\2\u0127R\3\2\2\2\u0128\u0130\7$\2\2"+
		"\u0129\u012a\7^\2\2\u012a\u012f\13\2\2\2\u012b\u012c\7$\2\2\u012c\u012f"+
		"\7$\2\2\u012d\u012f\n\35\2\2\u012e\u0129\3\2\2\2\u012e\u012b\3\2\2\2\u012e"+
		"\u012d\3\2\2\2\u012f\u0132\3\2\2\2\u0130\u012e\3\2\2\2\u0130\u0131\3\2"+
		"\2\2\u0131\u0133\3\2\2\2\u0132\u0130\3\2\2\2\u0133\u0134\7$\2\2\u0134"+
		"T\3\2\2\2\u0135\u0137\t\36\2\2\u0136\u0135\3\2\2\2\u0137\u0138\3\2\2\2"+
		"\u0138\u0136\3\2\2\2\u0138\u0139\3\2\2\2\u0139V\3\2\2\2\u013a\u013c\t"+
		"\36\2\2\u013b\u013a\3\2\2\2\u013c\u013d\3\2\2\2\u013d\u013b\3\2\2\2\u013d"+
		"\u013e\3\2\2\2\u013e\u013f\3\2\2\2\u013f\u0143\5\177@\2\u0140\u0142\t"+
		"\36\2\2\u0141\u0140\3\2\2\2\u0142\u0145\3\2\2\2\u0143\u0141\3\2\2\2\u0143"+
		"\u0144\3\2\2\2\u0144\u014d\3\2\2\2\u0145\u0143\3\2\2\2\u0146\u0148\5\177"+
		"@\2\u0147\u0149\t\36\2\2\u0148\u0147\3\2\2\2\u0149\u014a\3\2\2\2\u014a"+
		"\u0148\3\2\2\2\u014a\u014b\3\2\2\2\u014b\u014d\3\2\2\2\u014c\u013b\3\2"+
		"\2\2\u014c\u0146\3\2\2\2\u014dX\3\2\2\2\u014e\u014f\5)\25\2\u014f\u0150"+
		"\5%\23\2\u0150\u0151\5+\26\2\u0151\u0152\5\13\6\2\u0152Z\3\2\2\2\u0153"+
		"\u0154\5\r\7\2\u0154\u0155\5\3\2\2\u0155\u0156\5\31\r\2\u0156\u0157\5"+
		"\'\24\2\u0157\u0158\5\13\6\2\u0158\\\3\2\2\2\u0159\u015b\t\37\2\2\u015a"+
		"\u0159\3\2\2\2\u015b\u015c\3\2\2\2\u015c\u015a\3\2\2\2\u015c\u015d\3\2"+
		"\2\2\u015d\u0162\3\2\2\2\u015e\u0161\5}?\2\u015f\u0161\t \2\2\u0160\u015e"+
		"\3\2\2\2\u0160\u015f\3\2\2\2\u0161\u0164\3\2\2\2\u0162\u0160\3\2\2\2\u0162"+
		"\u0163\3\2\2\2\u0163^\3\2\2\2\u0164\u0162\3\2\2\2\u0165\u0166\7>\2\2\u0166"+
		"\u0167\7?\2\2\u0167`\3\2\2\2\u0168\u0169\7@\2\2\u0169\u016a\7?\2\2\u016a"+
		"b\3\2\2\2\u016b\u016c\7>\2\2\u016cd\3\2\2\2\u016d\u016e\7@\2\2\u016ef"+
		"\3\2\2\2\u016f\u0170\7?\2\2\u0170\u0171\7?\2\2\u0171h\3\2\2\2\u0172\u0173"+
		"\7#\2\2\u0173\u0174\7?\2\2\u0174j\3\2\2\2\u0175\u0176\7(\2\2\u0176\u0177"+
		"\7(\2\2\u0177l\3\2\2\2\u0178\u0179\7~\2\2\u0179\u017a\7~\2\2\u017an\3"+
		"\2\2\2\u017b\u017c\7#\2\2\u017cp\3\2\2\2\u017d\u017e\7\'\2\2\u017er\3"+
		"\2\2\2\u017f\u0180\7,\2\2\u0180t\3\2\2\2\u0181\u0182\7\61\2\2\u0182v\3"+
		"\2\2\2\u0183\u0184\7-\2\2\u0184x\3\2\2\2\u0185\u0186\7/\2\2\u0186z\3\2"+
		"\2\2\u0187\u0188\7?\2\2\u0188|\3\2\2\2\u0189\u018a\7a\2\2\u018a~\3\2\2"+
		"\2\u018b\u018c\7\60\2\2\u018c\u0080\3\2\2\2\u018d\u018e\7*\2\2\u018e\u0082"+
		"\3\2\2\2\u018f\u0190\7+\2\2\u0190\u0084\3\2\2\2\u0191\u0192\7}\2\2\u0192"+
		"\u0086\3\2\2\2\u0193\u0194\7\177\2\2\u0194\u0088\3\2\2\2\u0195\u0196\7"+
		"]\2\2\u0196\u008a\3\2\2\2\u0197\u0198\7_\2\2\u0198\u008c\3\2\2\2\u0199"+
		"\u019a\7<\2\2\u019a\u008e\3\2\2\2\u019b\u019c\7=\2\2\u019c\u0090\3\2\2"+
		"\2\u019d\u019e\7.\2\2\u019e\u0092\3\2\2\2\u019f\u01a1\t!\2\2\u01a0\u019f"+
		"\3\2\2\2\u01a1\u01a2\3\2\2\2\u01a2\u01a0\3\2\2\2\u01a2\u01a3\3\2\2\2\u01a3"+
		"\u01a4\3\2\2\2\u01a4\u01a5\bJ\2\2\u01a5\u0094\3\2\2\2\16\2\u012e\u0130"+
		"\u0138\u013d\u0143\u014a\u014c\u015c\u0160\u0162\u01a2\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}