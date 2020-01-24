go build
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify left standard --align=right
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify right standard --align=left
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify hello shadow --align=center
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify "1 Two 4" shadow --align=justify
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify 23/32 standard --align=right
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify ABCabc123 thinkertoy --align=right
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify "\"!#$%&" thinkertoy --align=center
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify "23Hello World!" standard --align=left
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify "HELLO there HOW are YOU?!" thinkertoy --align=justify
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify "a -> A b -> B c -> C" shadow --align=right
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify abcd shadow --align=right
echo ------------------------------------------------------------------------------------------------------------------------------------
./ascii-fustify ola standard --align=center
echo ------------------------------------------------------------------------------------------------------------------------------------
rm ascii-fustify