default: c

c:
	@echo "Compiling ${f}..."
	@clang++ -Wall -pipe -O2 -std=c++2a \
	-Wall -Wextra -pedantic -Wshadow -Wformat=2 -Wfloat-equal -Wconversion \
	-Wshift-overflow -Wcast-qual -Wcast-align \
	-Wno-unused-result -Wno-sign-conversion \
	${f} -o $$(basename -s .cpp ${f}) \
	&& echo "Compiled without errors"

new:
	@cp templates/template.cpp ${f} \
		&& echo "New file \"${f}\" created from template"

clean:
	@rm *.cpp \
		&& echo "All *.cpp files are removed from the workspace"
