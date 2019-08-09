# Pico y Placa Exercise

This repository contains the Pico y Placa exercise given as a part of the [Stack
Builders'][1] recruitment process.

The implementation is done using [Golang][2].

## Implementation details

 - The code is written in a procedural fashion.
 - Unit tests are written for the appropriate sections of the code.
 - The only external library used is the extension for Golang's testing library
   called [check.v1][3]

 ## How to build the program

 - The repository comes with `Makefile` which has appropriate targets to build /
   test the program.
 - The `Makefile` has a `help` target that briefly explains what each of the
   target does.

## How to run the program

- The program is a command line program and the user is expected to run it in
  a command line environment by invoking the executable name
- Example outputs
  - Invoking without any parameters
    ```
    $ ./picoyplaca
    Usage: picoyplaca LICENSE-PLATE DATE TIME
    Check if a LICENSE-PLATE is allowed within Quito City for a given DATE and TIME,
    Pico y Placa policy implemented in Quito.
    Example: picoyplaca AAB-0123 2019-05-20 09:31:00
  
    Interpretation of parameters:
         LICENSE-PLATE   This is the license plate number as described in the legal
                         system of Ecuador, only Standard Plates are implemented as of now.
                         For example: ABC-0123
  
         DATE            The date to check for if a vehicle is allowed in the city.
                         The input expects ISO 8601 standard of the form YYYY-MM-DD.
                         For example: 2019-05-20
  
         TIME            The time to check for if a vehicle is allowed in the city.
                         The input expects ISO 8601 standard of the form HH:MM:SS.
                         For example: 09:31:00
    ```
  - Invoking with parameters
    ```
    $ ./picoyplaca ABC-0123 2019-08-06 08:31:00
    ABC-0123 is NOT allowed on 2019-08-06 at 08:31:00
    $
    ```
    ```
    $ ./picoyplaca ABC-0123 2019-08-09 08:31:00
    ABC-0123 is allowed on 2019-08-09 at 08:31:00
    $
    ```

License
-------

BSD 2-clause. See LICENSE.

[1]: https://www.stackbuilders.com/
[2]: https://golang.org
[3]: https://gopkg.in/check.v1
