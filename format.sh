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
HEREDOC

	while read type file; do
		echo -e "\nfunc (f *$type) printType(w writer, v bool) {";
		echo "	pp := w.Indent()";
		echo;
		echo "	pp.WriteString(\"$type {\")";
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
				echo "		pp.WriteString(\"\\n$fieldName: nil\")";
				echo "	} else if len(f.$fieldName) > 0 {";
				echo "		pp.WriteString(\"\\n$fieldName: [\")";
				echo;
				echo "		ipp := pp.Indent()";
				echo;
				echo "		for n, e := range f.$fieldName {";
				echo "			ipp.Printf(\"\n%d: \", n)";
				echo "			e.printType(ipp, v)";
				echo "		}";
				echo;
				echo "		pp.WriteString(\"\\n]\")";
				echo "	} else if v {";
				echo "		pp.WriteString(\"\\n$fieldName: []\")";
				echo "	}";
			elif [ "${fieldType:0:1}" = "[" ]; then
				echo "	pp.WriteString(\"\\n$fieldName: [\")";
				echo;
				echo "	ipp := pp.Indent()";
				echo;
				echo "	for n, e := range f.$fieldName {";
				echo "		ipp.Printf(\"\n%d: \", n)";
				echo "		e.printType(ipp, v)";
				echo "	}";
				echo;
				echo "	pp.WriteString(\"\n]\")";
			elif [ "${fieldType:0:1}" = "*" ]; then
				echo;
				echo "	if f.$fieldName != nil {";
				echo "		pp.WriteString(\"\\n$fieldName: \")";
				echo "		f.$fieldName.printType(pp, v)";
				echo "	} else if v {";
				echo "		pp.WriteString(\"\\n$fieldName: nil\")";
				echo "	}";
			else
				echo;
				echo "	pp.WriteString(\"\\n$fieldName: \")";
				echo "	f.$fieldName.printType(pp, v)";
			fi;
		done < <(sed '/^type '$type' struct {$/,/^}$/!d;//d' "$file");

		echo;
		echo "	w.WriteString(\"\n}\")";
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
