gopa
====

Gopa enables simple and safe access to parameters

Installation
============

Get it by issuing:

    go get github.com/Blacktremolo/gopa

Use it by:

    import "github.com/Blacktremolo/gopa"

Usage overview
==============

Gopa offers a few commands to REQUIRE a parameter of a specific type and on a specific position.
Requiering means that if the param type does not match or it is non-existent the program will exit.
Gopa also offers commands to check if param of any type exists on a certain position,
so the number of parameters can be flexible.
The return number and fail message can simply be set in a public module variable.

