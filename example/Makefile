CXX = go build
EXEC = example
OBJECTS = main.go producer.go consumer.go helpers.go

example: ${OBJECTS}
	${CXX} -o ${EXEC}

.PHONY: clean

clean:
	rm ${EXEC}
