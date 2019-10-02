CXX = go build
EXEC = main
OBJECTS = main.go producer.go consumer.go helpers.go

${EXEC}: ${OBJECTS}
	${CXX} ${OBJECTS}

-include ${DEPENDS}

.PHONY: clean

clean:
	rm ${EXEC}
