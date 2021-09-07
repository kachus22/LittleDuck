all:
	cd ANTLR && antlr -Dlanguage=Go -o parser LittleDuck.g4
	cd Gocc && gocc LittleDuck.bnf
