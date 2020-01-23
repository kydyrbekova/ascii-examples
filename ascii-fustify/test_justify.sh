echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go left standard --align=right
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go right standard --align=left
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go hello shadow --align=center
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go "1 Two 4" shadow --align=justify
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go 23/32 standard --align=right
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go ABCabc123 thinkertoy --align=right
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go "\"!#$%&" thinkertoy --align=center
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go "23Hello World!" standard --align=left
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go "HELLO there HOW are YOU?!" thinkertoy --align=justify
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go "a -> A b -> B c -> C" shadow --align=right
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go abcd shadow --align=right
echo -----------------------------------------------------------------------------------------------------------------------------------------------
go run justify.go ola standard --align=center
echo -----------------------------------------------------------------------------------------------------------------------------------------------