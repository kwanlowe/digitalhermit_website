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

% We are looking for Volume (theta), Area (theta), Piston Velocity Up (theta),
% First derivative of volume wrt theta (dV/d(theta)) = Vprime
% First derivative of Area wrt theta   (dA/d(theta)) = Aprime

% Equations

pkg load symbolic
clear 

syms V_d B L
syms qty
syms V_c theta r R
syms A_p A_ch A
syms U_p N qty_1 qty_2 A_prime thetaD pi

V_d = ((pi * (B^2))/4) * L

V_c = V_d / (r - 1)
qty = sqrt((R^2 - sind(theta)^2))
V = V_c * ( 1 + ((r - 1)/2) *( R + 1 - cosd(theta) - qty))
A = A_ch + A_p + pi * B * L/2 * ( R + 1 - cosd(theta) - qty)
U_p = pi * L * N * sind(theta) * (cosd(theta)/qty + 1)

qty_1 = ((r - 1)/2)
qty_2 = (sind(theta) * cosd(theta))/(2 * qty)

Vprime = V_c * (( sind(theta) * qty_1) + qty_2 *( r - 1 ))

A_prime = (( pi * B * L)/2) * (sind(theta) + (sind(theta)*cosd(theta))/qty)

%Inputs

theta = 15

B = 14; %[inches];
L = 14; %[inches]
r = 9;
R = 6;
A_ch = 231; %[inches^2]
A_p = 154; %[inches^2]
N = 1500; %[RPM]

printf("V_d = %f\n", eval(V_d));
printf("V_c = %f\n", eval(V_c));

% Keep all inputs the same as before except theta to make plots of:
%   V vs theta
%   Vprime vs theta
%   A vs theta
%   Aprime vs theta
%   Up vs theta

theta=linspace(-180, 180);

% Plot 1: V vs theta
y_1=eval(V);
subplot(3,2,1);
plot(theta,y_1)
grid on;
axis([ -180, 180, 0,1]);
axis ("auto y");
xlabel ("Theta");
ylabel ("Volume [in^3]");
title ("Volume versus Crank Angle");

%Create an array of Volume values
% from 0 to Vd

V_prime=diff(V);
y_2=eval(V_prime);
subplot(3,2,2);
plot(theta, y_2)
grid on;
axis ([-180, 180, -1,1]);
axis ("auto y");
xlabel ("Theta");
ylabel ("V'");
title ("dV versus Crank Angle");


subplot(3,2,3);
y_3=eval(A);
plot(theta, y_3)
grid on
axis ([-180, 180, -1,1]);
axis ("auto y");
xlabel ("Theta");
ylabel ("Area");
title ("Area versus Crank Angle");

subplot(3,2,4);
y_4=eval(A_prime);
plot(theta, y_4)
grid on;
axis ([-180, 180, -1, 1]);
axis("auto y");
xlabel ("Theta");
ylabel ("dA");
title ("dA versus Crank Angle");

subplot (3,2,5);
y_5=eval(U_p);
plot(theta,y_5);
grid on;
axis ([-180, 180, -1, 1]);
axis("auto y");
xlabel ("Theta");
ylabel ("U_p");
title ("U_p versus Crank Angle");

% Save Images
print -dpng diagram_engine_degrees.png;