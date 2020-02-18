disp("FLux Demo")
disp("Demonstration of some Octave functions:")
input ("Press Enter for next example.")

pkg load symbolic
syms x
y=(5*x.^2 - 8*x + 5)
int(y)
y=(-6*x.^3 + 9*x.^2 + 4*x - 3)
int(y)

input ("Press Enter for next example.")
clc
disp("The error is from using floating point variables.")
y=(x.^(3/2) + 2*x + 3)



input ("Press Enter for next example.")
clc
disp("We can override this with:")
y=(x.^(sym(3)/2) + 2*x + 3)
int(y)
