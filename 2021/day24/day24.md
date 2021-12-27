# Day 24

## Puzzle Input

| Digit A  |   w |   x |    y |    z | Digit B  |   w |       x |    y |                z | Digit C  |   w |                x |    y |                               z | Digit D  |   w |                               x |    y |                              z | Digit E  | Digit F  | Digit G  | Digit H   | Digit I  | Digit J   | Digit K  | Digit L  | Digit M  | Digit O  |
| -------- | --: | --: | ---: | ---: | -------- | --: | ------: | ---: | ---------------: | -------- | --: | ---------------: | ---: | ------------------------------: | -------- | --: | ------------------------------: | ---: | -----------------------------: | -------- | -------- | -------- | --------- | -------- | --------- | -------- | -------- | -------- | -------- |
| inp w    |   A |     |      |    0 | inp w    |   B |         |      |                  | inp w    |   C |                  |      |                                 | inp w    |   D |                                 |      |                                | inp w    | inp w    | inp w    | inp w     | inp w    | inp w     | inp w    | inp w    | inp w    | inp w    |
| mul x 0  |     |   0 |      |      | mul x 0  |     |       0 |      |                  | mul x 0  |     |                0 |      |                                 | mul x 0  |     |                               0 |      |                                | mul x 0  | mul x 0  | mul x 0  | mul x 0   | mul x 0  | mul x 0   | mul x 0  | mul x 0  | mul x 0  | mul x 0  |
| add x z  |     |   0 |      |      | add x z  |     |    A+15 |      |                  | add x z  |     | 26.(A+15) + B+12 |      |                                 | add x z  |     | 26.26.(A+15) + 26.(B+12) + C+15 |      |                                | add x z  | add x z  | add x z  | add x z   | add x z  | add x z   | add x z  | add x z  | add x z  | add x z  |
| mod x 26 |     |   0 |      |      | mod x 26 |     |    A+15 |      |                  | mod x 26 |     |             B+12 |      |                                 | mod x 26 |     |                            C+15 |      |                                | mod x 26 | mod x 26 | mod x 26 | mod x 26  | mod x 26 | mod x 26  | mod x 26 | mod x 26 | mod x 26 | mod x 26 |
| div z 1  |     |     |      |    0 | div z 1  |     |         |      |             A+15 | div z 1  |     |                  |      |                26.(A+15) + B+12 | div z 26 |     |                                 |      |               26.(A+15) + B+12 | div z 26 | div z 1  | div z 26 | div z 26  | div z 1  | div z 26  | div z 1  | div z 1  | div z 26 | div z 26 |
| add x 12 |     |  12 |      |      | add x 14 |     | A+15+14 |      |                  | add x 11 |     |        B+12 + 11 |      |                                 | add x -9 |     |                       C+ 15 - 9 |      |                                | add x -7 | add x 11 | add x -1 | add x -16 | add x 11 | add x -15 | add x 10 | add x 12 | add x -4 | add x 0  |
| eql x w  |     |   0 |      |      | eql x w  |     |       0 |      |                  | eql x w  |     |                0 |      |                                 | eql x w  |     |                               0 |      |                                | eql x w  | eql x w  | eql x w  | eql x w   | eql x w  | eql x w   | eql x w  | eql x w  | eql x w  | eql x w  |
| eql x 0  |     |   1 |      |      | eql x 0  |     |       1 |      |                  | eql x 0  |     |                1 |      |                                 | eql x 0  |     |                               1 |      |                                | eql x 0  | eql x 0  | eql x 0  | eql x 0   | eql x 0  | eql x 0   | eql x 0  | eql x 0  | eql x 0  | eql x 0  |
| mul y 0  |     |     |    0 |      | mul y 0  |     |         |    0 |                  | mul y 0  |     |                  |    0 |                                 | mul y 0  |     |                                 |    0 |                                | mul y 0  | mul y 0  | mul y 0  | mul y 0   | mul y 0  | mul y 0   | mul y 0  | mul y 0  | mul y 0  | mul y 0  |
| add y 25 |     |     |   25 |      | add y 25 |     |         |   25 |                  | add y 25 |     |                  |   25 |                                 | add y 25 |     |                                 |   25 |                                | add y 25 | add y 25 | add y 25 | add y 25  | add y 25 | add y 25  | add y 25 | add y 25 | add y 25 | add y 25 |
| mul y x  |     |     |   25 |      | mul y x  |     |         |   25 |                  | mul y x  |     |                  |   25 |                                 | mul y x  |     |                                 |   25 |                                | mul y x  | mul y x  | mul y x  | mul y x   | mul y x  | mul y x   | mul y x  | mul y x  | mul y x  | mul y x  |
| add y 1  |     |     |   26 |      | add y 1  |     |         |   26 |                  | add y 1  |     |                  |   26 |                                 | add y 1  |     |                                 |   26 |                                | add y 1  | add y 1  | add y 1  | add y 1   | add y 1  | add y 1   | add y 1  | add y 1  | add y 1  | add y 1  |
| mul z y  |     |     |      |    0 | mul z y  |     |         |      |        26.(A+15) | mul z y  |     |                  |      |        26.26.(A+15) + 26.(B+12) | mul z y  |     |                                 |      |       26.26.(A+15) + 26.(B+12) | mul z y  | mul z y  | mul z y  | mul z y   | mul z y  | mul z y   | mul z y  | mul z y  | mul z y  | mul z y  |
| mul y 0  |     |     |    0 |      | mul y 0  |     |         |    0 |                  | mul y 0  |     |                  |    0 |                                 | mul y 0  |     |                                 |    0 |                                | mul y 0  | mul y 0  | mul y 0  | mul y 0   | mul y 0  | mul y 0   | mul y 0  | mul y 0  | mul y 0  | mul y 0  |
| add y w  |     |     |    A |      | add y w  |     |         |    B |                  | add y w  |     |                  |    C |                                 | add y w  |     |                                 |    D |                                | add y w  | add y w  | add y w  | add y w   | add y w  | add y w   | add y w  | add y w  | add y w  | add y w  |
| add y 15 |     |     | A+15 |      | add y 12 |     |         | B+12 |                  | add y 15 |     |                  | C+15 |                                 | add y 12 |     |                                 | D+12 |                                | add y 15 | add y 2  | add y 11 | add y 15  | add y 10 | add y 2   | add y 0  | add y 0  | add y 15 | add y 15 |
| mul y x  |     |     | A+15 |      | mul y x  |     |         | B+12 |                  | mul y x  |     |                  | C+15 |                                 | mul y x  |     |                                 | D+12 |                                | mul y x  | mul y x  | mul y x  | mul y x   | mul y x  | mul y x   | mul y x  | mul y x  | mul y x  | mul y x  |
| add z y  |     |     |      | A+15 | add z y  |     |         |      | 26.(A+15) + B+12 | add z y  |     |                  |      | 26.26.(A+15) + 26.(B+12) + C+15 | add z y  |     |                                 |      | 26.26.(A+15) + 26.(B+12) +D+12 | add z y  | add z y  | add z y  | add z y   | add z y  | add z y   | add z y  | add z y  | add z y  | add z y  |

z is the list of digits of a base 26 number.

The first bit with X will take the last digit held in z and add a certain amount (and compare with the current digit?)

Z will either be the same or have that last digit removed. (peek or pop)

Each digit starts as 1-9

| Sames?   | Digit A  | Digit B  | Digit C  | Digit D  | Digit E  | Digit F  | Digit G  | Digit H   | Digit I  | Digit J   | Digit K  | Digit L  | Digit M  | Digit O  |
| -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | --------- | -------- | --------- | -------- | -------- | -------- | -------- |
| inp w    |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| mul x 0  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| add x z  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| mod x 26 |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
|          | div z 1  | div z 1  | div z 1  | div z 26 | div z 26 | div z 1  | div z 26 | div z 26  | div z 1  | div z 26  | div z 1  | div z 1  | div z 26 | div z 26 |
|          | add x 12 | add x 14 | add x 11 | add x -9 | add x -7 | add x 11 | add x -1 | add x -16 | add x 11 | add x -15 | add x 10 | add x 12 | add x -4 | add x 0  |
| eql x w  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| eql x 0  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| mul y 0  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| add y 25 |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| mul y x  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| add y 1  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| mul z y  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| mul y 0  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| add y w  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
|          | add y 15 | add y 12 | add y 15 | add y 12 | add y 15 | add y 2  | add y 11 | add y 15  | add y 10 | add y 2   | add y 0  | add y 0  | add y 15 | add y 15 |
| mul y x  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |
| add z y  |          |          |          |          |          |          |          |           |          |           |          |          |          |          |

Working backward what are the restrictions on the last digit?

div z 1 is push
div z 26 is pop?
always paired with a subtraction which means the if statement is irrelevant for push, only matters for pop

z.pop() is always > 9, so the condition is only met when subtracting from z.pop.
This means to reduce the stack we have to have the pops as pushes.

z.push(A + 15)

z.push(B + 12)

z.push(C + 15)

if D != z.pop() - 9 : D + 12

if E != z.pop() - 7 : E + 15

z.push(F + 2)

if G != z.pop() - 1 : G + 11

if H != z.pop() - 16 : H + 15

z.push(I + 10)

if J != z.pop() - 15 : J + 2

z.push(K + 0)

z.push(L + 0)

if M != z.pop() - 4 : M + 15

if O != z.pop() - 0 : O + 15

### Rules:

D = C + 15 - 9
E = B + 12 - 7
G = F + 2 - 1
H = A + 15 - 16
J = I + 10 - 15
M = L + 0 - 4
O = K + 0 - 0

D = C + 6
E = B + 5
G = F + 1
H = A - 1
J = I - 5
M = L - 4
O = K

A = H + 1
B = E - 5
C = D - 6
F = G - 1
I = J + 5
K = O
L = M + 4

Maximum:

| Digit A | Digit B | Digit C | Digit D | Digit E | Digit F | Digit G | Digit H | Digit I | Digit J | Digit K | Digit L | Digit M | Digit O |
| ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- |
| 9       | 4       | 3       | 9       | 9       | 8       | 9       | 8       | 9       | 4       | 9       | 9       | 5       | 9       |

94399898949959

Minimum:

| Digit A | Digit B | Digit C | Digit D | Digit E | Digit F | Digit G | Digit H | Digit I | Digit J | Digit K | Digit L | Digit M | Digit O |
| ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- | ------- |
| 2       | 1       | 1       | 7       | 6       | 1       | 2       | 1       | 6       | 1       | 1       | 5       | 1       | 1       |

21176121611511
