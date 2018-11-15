package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	IDENTIFIER = "IDENTIFIER"
	LINEBREAK = "LINEBREAK"
	QUOTE = "'"
	QUOTEQUOTE = "\""
	BACKSLASH = "\\"
	//Punctuators
	LBRACE = "{"
	RBRACE = "}"
	LPAREN = "("
	RPAREN = ")"
	LBRACK = "["
	RBRACK = "]"
	DOT = "."
	SEMICOLON = ";"
	COMMA = ","
	LT = "<"
	GT = ">"
	LTASSIGN = "<="
	GTASSIGN = ">="
	ASSIGNASSIGN = "=="
	ASSIGNASSIGNASSIGN = "==="
	BANGASSIGN = "!="
	BANGASSIGNASSIGN = "!=="
	PLUS = "+"
	MINUS = "-"
	STAR = "*"
	PERCENT = "%"
	PLUSPLUS = "++"
	MINUSMINUS = "--"
	LTLT = "<<"
	GTGT = ">>"
	GTGTGT = ">>>"
	AND = "&"
	OR = "|"
	CARAT = "^"
	BANG = "!"
	TILDE = "~"
	ANDAND = "&&"
	OROR = "||"
	QUESTION = "?"
	COLON = ":"
	ASSIGN = "="
	PLUSASSIGN = "+="
	MINUSASSIGN = "-="
	STARASSIGN = "*="
	PERCENTASSIGN = "%="
	LTLTASSIGN = "<<="
	GTGTASSIGN = ">>="
	GTGTGTASSIGN = ">>>="
	ANDASSIGN = "&="
	ORASSIGN = "|="
	CARATASSIGN = "^="
	//DivPunctuators
	DIV = "/"
	DIVASSIGN = "/="
	//Keywords
	BREAK = "BREAK"
	CASE = "CASE"
	CATCH = "CATCH"
	CONTINUE = "CONTINUE"
	DEFAULT = "DEFAULT"
	DELETE = "DELETE"
	DO = "DO"
	ELSE = "ELSE"
	FINALLY = "FINALLY"
	FOR = "FOR"
	FUNCTION = "FUNCTION"
	IF = "IF"
	IN = "IN"
	INSTANCEOF = "INSTANCEOF"
	NEW = "NEW"
	RETURN = "RETURN"
	SWITCH = "SWITCH"
	THIS = "THIS"
	THROW = "THROW"
	TRY = "TRY"
	TYPEOF = "TYPEOF"
	VAR = "VAR"
	VOID = "VOID"
	WHILE = "WHILE"
	WITH = "WITH"
	//Future reserved words
	ABSTRACT = "ABSTRACT"
	BOOLEAN = "BOOLEAN"
	BYTE = "BYTE"
	CHAR = "CHAR"
	CLASS = "CLASS"
	CONST = "CONST"
	DEBUGGER = "DEBUGGER"
	DOUBLE = "DOUBLE"
	ENUM = "ENUM"
	EXPORT = "EXPORT"
	EXTENDS = "EXTENDS"
	FINAL = "FINAL"
	FLOAT = "FLOAT"
	GOTO = "GOTO"
	IMPLEMENTS = "IMPLEMENTS"
	IMPORT = "IMPORT"
	INT = "INT"
	INTERFACE = "INTERFACE"
	LONG = "LONG"
	NATIVE = "NATIVE"
	PACKAGE = "PACKAGE"
	PRIVATE = "PRIVATE"
	PROTECTED = "PROTECTED"
	PUBLIC = "PUBLIC"
	SHORT = "SHORT"
	STATIC = "STATIC"
	SUPER = "SUPER"
	SYNCHRONIZED = "SYNCHRONIZED"
	THROWS = "THROWS"
	TRANSIENT = "TRANSIENT"
	VOLATILE = "VOLATILE"
	//Literals 
	TRUE = "TRUE"
	FALSE = "FALSE"
	NULL = "NULL"
	UNDEFINED = "UNDEFINED"
	NUMBER = "NUMBER"
	STRING = "STRING"
)

var keywords = map[string]TokenType {
	//Keywords
	"break": BREAK,
	"case": CASE, 
	"catch": CATCH,
	"continue": CONTINUE,
	"default": DEFAULT,
	"delete": DELETE,
	"do": DO,
	"else": ELSE,
	"finally": FINALLY,
	"for": FOR,
	"function": FUNCTION,
	"if": IF,
	"in": IN,
	"instanceof": INSTANCEOF,
	"new": NEW,
	"return": RETURN,
	"switch": SWITCH,
	"this": THIS,
	"throw": THROW,
	"try": TRY,
	"typeof": TYPEOF,
	"var": VAR,
	"void": VOID,
	"while": WHILE,
	"with": WITH,
	//Future reserved words
	"abstract": ABSTRACT,
	"boolean": BOOLEAN,
	"byte": BYTE,
	"char": CHAR,
	"class": CLASS,
	"const": CONST,
	"debugger": DEBUGGER,
	"double": DOUBLE,
	"enum": ENUM,
	"export": EXPORT,
	"extends": EXTENDS,
	"final": FINAL,
	"float": FLOAT,
	"goto": GOTO,
	"implements": IMPLEMENTS,
	"import": IMPORT,
	"int": INT,
	"interface": INTERFACE,
	"long": LONG,
	"native": NATIVE,
	"package": PACKAGE,
	"private": PRIVATE,
	"protected": PROTECTED,
	"public": PUBLIC,
	"short": SHORT,
	"static": STATIC,
	"super": SUPER,
	"synchronized": SYNCHRONIZED,
	"throws": THROWS,
	"transient": TRANSIENT,
	"volatile": VOLATILE,
	//Literals 
	"true": TRUE,
	"false": FALSE,
	"null": NULL,
	"undefined": UNDEFINED,
}

func FindIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token 
	}
	return IDENTIFIER
}
type Token struct {
	Type TokenType
	Literal string
}
