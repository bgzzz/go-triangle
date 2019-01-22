# go-triangle
This is the utility that determines the type of a triangle based on the input lengths of triangle sides.

The types of the triangle that can be determined:

- equilateral
- isosceles
- scalene

## How to compile

```
$ make build
```

## How to run


```
$ ./bin/go-triangle 10 10 10
``` 

Example of output is:

```
Current triangle is equilateral
```

## How to run unit tests

```
make test-unit
```

## In case you do not have infrastructure for GoLang

### How to build docker 

```
make docker-build
```

## How to run unit tests in docker 

```
make docker-run-unit
``` 