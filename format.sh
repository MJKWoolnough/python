#!/bin/bash

types() {
	for file in ast.go ast_compound.go ast_expression.go ast_statement.go; do
		while read type; do
			echo "$type" "$file";
		done < <(grep "type .* struct {" "$file" | cut -d' ' -f2);
	done | sort;
}

{
	cat <<HEREDOC
package python

// File automatically generated with format.sh.

import "io"
HEREDOC

	while read type file; do
		echo -e "\nfunc (f *$type) printType(w io.Writer, v bool) {";
		echo "	pp := indentPrinter{w}";
		echo;
		echo "	pp.Print(\"$type {\")";
		while read fieldName fieldType; do
			if [ "$fieldType" = "bool" ]; then
				echo;
				echo "	if f.$fieldName || v {";
				echo "		pp.Printf(\"\\n$fieldName: %v\", f.$fieldName)";
				echo "	}";
			elif [ "$fieldType" = "uint" -o "$fieldType" = "int" ]; then
				echo;
				echo "	if f.$fieldName != 0 || v {";
				echo "		pp.Printf(\"\\n$fieldName: %v\", f.$fieldName)";
				echo "	}";
			elif [ "${fieldType:0:2}" = "[]" ]; then
				echo;
				echo "	if f.$fieldName == nil {";
				echo "		pp.Print(\"\\n$fieldName: nil\")";
				echo "	} else if len(f.$fieldName) > 0 {";
				echo "		pp.Print(\"\\n$fieldName: [\")";
				echo;
				echo "		ipp := indentPrinter{&pp}";
				echo;
				echo "		for n, e := range f.$fieldName {";
				echo "			ipp.Printf(\"\n%d: \", n)";
				echo "			e.printType(&ipp, v)";
				echo "		}";
				echo;
				echo "		pp.Print(\"\\n]\")";
				echo "	} else if v {";
				echo "		pp.Print(\"\\n$fieldName: []\")";
				echo "	}";
			elif [ "${fieldType:0:1}" = "[" ]; then
				echo "	pp.Print(\"\\n$fieldName: [\")";
				echo;
				echo "	ipp := indentPrinter{&pp}";
				echo;
				echo "	for n, e := range f.$fieldName {";
				echo "		ipp.Printf(\"\n%d: \", n)";
				echo "		e.printType(&ipp, v)";
				echo "	}";
				echo;
				echo "	pp.Print(\"\n]\")";
			elif [ "${fieldType:0:1}" = "*" ]; then
				echo;
				echo "	if f.$fieldName != nil {";
				echo "		pp.Print(\"\\n$fieldName: \")";
				echo "		f.$fieldName.printType(&pp, v)";
				echo "	} else if v {";
				echo "		pp.Print(\"\\n$fieldName: nil\")";
				echo "	}";
			else
				echo;
				echo "	pp.Print(\"\\n$fieldName: \")";
				echo "	f.$fieldName.printType(&pp, v)";
			fi;
		done < <(sed '/^type '$type' struct {$/,/^}$/!d;//d' "$file");

		echo;
		echo "	io.WriteString(w, \"\n}\")";
		echo "}";
	done < <(types);
} > "format_types.go";

{
	cat <<HEREDOC
package python

// File automatically generated with format.sh.

import "fmt"
HEREDOC

	while read type _; do
		echo -e "\n// Format implements the fmt.Formatter interface";
		echo "func (f $type) Format(s fmt.State, v rune) {";
		echo "	if v == 'v' && s.Flag('#') {";
		echo "		type X = $type";
		echo "		type $type X";
		echo;
		echo "		fmt.Fprintf(s, \"%#v\", $t(f))";
		echo "	} else {";
		echo "		format(&f, s, v)";
		echo "	}";
		echo "}";
	done < <(types);
} > "format_format.go";

{
	cat <<HEREDOC
package python

// File automatically generated with format.sh.

import "fmt"

// Type is an interface satisfied by all python structural types.
type Type interface {
	fmt.Formatter
	pythonType()
}

func (Tokens) pythonType() {}
HEREDOC

	while read type _; do
		echo -e "\nfunc ($type) pythonType() {}";
	done < <(types);
} > "types.go";
