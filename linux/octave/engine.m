%Task 2

%Variables

% Crank Angle, Theta
% Bore, B
% Stroke, L
% Compression Ratio, r
% Connecting rod length to crank-arm length a, ratio R=1/a
% Area of Cylinder head, Ach
% Area of Piston, Ap
% Crankshaft Rotational Speed, N [RPM]

%We are looking for Volume (theta),Area (theta), Piston Velocity Up (theta),
%First derivative of volume wrt theta (dV/d(theta))=Vprime,
%First derivative of Area wrt theta, (dA/d(theta)) = Aprime

%Equations

pkg load symbolic
clear 

syms V_d B L
syms qty
syms V_c theta r R
syms A_p A_ch A
syms U_p N qty_1 qty_2 A_prime thetaD pi

V_d = ((pi * (B^2))/4) * L

V_c = V_d / (r - 1)
qty = sqrt((R^2 - sin(theta)^2))
V = V_c * ( 1 + ((r - 1)/2) *( R + 1 - cos(theta) - qty))
A = A_ch + A_p + pi * B * L/2 * ( R + 1 - cos(theta) - qty)
U_p = pi * L * N * sin(theta) * (cos(theta)/qty + 1)

qty_1 = ((r - 1)/2)
qty_2 = (sin(theta) * cos(theta))/(2 * qty)

Vprime = V_c * (( sin(theta) * qty_1) + qty_2 *( r - 1 ))

A_prime = (( pi * B * L)/2) * (sin(theta) + (sin(theta)*cos(theta))/qty)

%Inputs
syms thetaD
thetaD = 15
theta = (thetaD * pi)/sym(180)

B = 14; %[inches];
L = 14; %[inches]
r = 9;
R = 6;
A_ch = 231; %[inches^2]
A_p = 154; %[inches^2]
N = 1500; %[RPM]

printf("V_d = %f\n", eval(V_d))
printf("V_c = %f\n", eval(V_c))

%Keep all inputs the same as before except theta to make plots of:
% V vs theta, Vprime vs theta, A vs theta, Aprime vs theta, Up vs theta

% P_theta=linspace(0, pi, 100)

% Plot 1: V vs theta

theta=linspace(-2*pi, 2*pi);
y=eval(eval(V));
x=eval(theta);
% subplot(2,1,1);
plot(x,y)
%Create an array of Volume values
% from 0 to Vd
xlabel ("Theta");
ylabel ("Volume [in^3]");
title ("Volume versus Crank Angle Relationship");
error("stop")
