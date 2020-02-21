pkg load symbolic;
clear;
clc;


function result = sumarray(arraySize)
    syms n;
    y(n) = (1/((2*n + 1)*(2*n + 2)));
    n = linspace (0, arraySize, (arraySize + 1));
    sum(y(n));
    result = eval(ans);
endfunction

disp("This snippet evaluates a function by passing an array:")
disp("    f(n) = (1/((2*n + 1)*(2*n + 2)))");
disp("\nThen it does an iterative approach.")
disp("\nIt first attempts to pass an array to the function.")

input ("\nFirst, try with n[1:50]:");
ans1 = feval("sumarray", 50);

disp("Answer: "), disp(ans1);
disp("ln(2):  "), disp(log(2));

%  ----------
input ("Press Enter to continue: ")
clear;
clc;
disp ("Next, try with n[1:500]:");
disp ("This will likely fail with an overrun:");
input ("Press Enter to continue: ");
% warning("on");

ans1 = feval("sumarray", 500);
disp("Answer: "), disp(ans1);
disp("ln(2):  "), disp(log(2));

%  ----------
input ("Press Enter to continue: ")
clear;
clc;
disp ("NAN is an overrun. Let's try with iterative method:");
input ("Press Enter to continue: ");
syms n;
y(n) = (1/((2*n + 1)*(2*n + 2)));
sum = 0;
for i = 0:500
    sum = sum + eval(y(i));
    fprintf(".");
endfor
fprintf("\n")

disp("Answer: "), disp(sum)
disp("ln(2):  "), disp(log(2));
