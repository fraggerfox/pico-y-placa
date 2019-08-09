# Notes and Implementation details on the exercise

## Number plates in Ecuador

In Ecuador number plates have a unique method of numbering. Standard plates have
a pattern `ABC-0123` which is consistent. The alphabets used have unique meaning
to indicate the province and the type of the vehicle it is etc. A full reference
of this can be found in the Wikipedia [entry][1] for vehicle registration plates
of Ecuador.

## Pico y Placa rule

This is an interesting rule applied in various parts of the world to reduce
traffic congestion in cities where the density of traffic can explode during
peak hours. Even though this rule is not unique to [Ecuador][2], this
implementation will only cover the rules in Ecuador.

A rough translation of the relevant sections of the Wikipedia entry is given
below.

Peak hours are considered as follows:

|  |Start Time|End Time|
|--|--|--|
|Morning| 07:00:00 | 09:30:00 |
|Evening| 16:00:00 | 19:30:00 |

Day of week and the disallowed last digits during the peak hours:

|Day|Last digit|
|--|--|
|Monday|1 & 2|
|Tuesday|3 & 4|
|Wednesday|5 & 6|
|Thursday|7 & 8|
|Friday|9 & 0|

The above rules do not apply to the weekends / holidays.

## Implementation notes

- The program is quite basic and uses a command line approach for user
  interaction and outputs.
- The program is written in such a way that it is modular enough to be
  exported as an API (if needed) with a bit of minimal changes.
- License plate validation can use a bit more things like
  - Checking for the type of license plate depending on the patters.
  - Checking for holidays (other than weekends) where the rules are not applied.
  - Checking if the given number plate is emergency vehicle.
  - Better validation by checking the length of digits and alphabets.

[1]: https://en.wikipedia.org/wiki/Vehicle_registration_plates_of_Ecuador
[2]: https://es.wikipedia.org/wiki/Pico_y_placa#Quito,_Ecuador
