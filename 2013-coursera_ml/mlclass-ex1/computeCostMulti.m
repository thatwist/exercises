function J = computeCostMulti(X, y, theta)
%COMPUTECOSTMULTI Compute cost for linear regression with multiple variables
%   J = COMPUTECOSTMULTI(X, y, theta) computes the cost of using theta as the
%   parameter for linear regression to fit the data points in X and y

% Initialize some useful values
m = length(y); % number of training examples

% You need to return the following variables correctly 
J = 0;

% ====================== YOUR CODE HERE ======================
% Instructions: Compute the cost of a particular choice of theta
%               You should set J to the cost.

X       % so X is already prepared for vectorization (has ones as the first column)
y       % it's m dimensional vector
theta   % it's 2 dimensional vector

% plot that shit
plot(X, y, 'b')
hold on

% plot the prediction
predictions = X * theta
plot(X, predictions, 'r')
hold off

squaredErrors = (predictions - y) .^ 2
k = 1/(2*m)

% function result
J = k * sum(squaredErrors)



% =========================================================================

end

% OUTPUT
%
% == Connecting to ml-class ... 
%
% X =
%
%     1.0000    3.4572    1.8594    1.3636
%     1.0000    4.1961    2.0484    1.4312
%     1.0000    4.9350    2.2215    1.4905
%     1.0000    5.6739    2.3820    1.5434
%     1.0000    6.4128    2.5324    1.5913
%     1.0000    7.1517    2.6743    1.6353
%     1.0000    7.8906    2.8090    1.6760
%     1.0000    8.6295    2.9376    1.7139
%     1.0000    9.3684    3.0608    1.7495
%     1.0000   10.1073    3.1792    1.7830
%     1.0000   10.8462    3.2934    1.8148
%     1.0000   11.5851    3.4037    1.8449
%     1.0000   12.3241    3.5106    1.8736
%     1.0000   13.0630    3.6143    1.9011
%     1.0000   13.8019    3.7151    1.9275
%     1.0000   14.5408    3.8132    1.9528
%     1.0000   15.2797    3.9089    1.9771
%     1.0000   16.0186    4.0023    2.0006
%     1.0000   16.7575    4.0936    2.0233
%     1.0000   17.4964    4.1829    2.0452
% 
% y =
% 
%     5.1778
%     6.6755
%     8.4462
%    10.0438
%    11.1175
%    11.5784
%    11.6443
%    11.7325
%    12.2465
%    13.3623
%    14.9358
%    16.5858
%    17.8961
%    18.6286
%    18.8418
%    18.8611
%    19.1117
%    19.8964
%    21.2363
%    22.8612
% 
% theta =
% 
%    0.10000
%    0.20000
%    0.30000
%    0.40000
% 
% predictions =
% 
%    1.8947
%    2.1262
%    2.3496
%    2.5667
%    2.7788
%    2.9868
%    3.1912
%    3.3928
%    3.5917
%    3.7884
%    3.9832
%    4.1761
%    4.3674
%    4.5573
%    4.7459
%    4.9332
%    5.1195
%    5.3046
%    5.4889
%    5.6722
% 
% squaredErrors =
% 
%     10.779
%     20.696
%     37.168
%     55.907
%     69.533
%     73.816
%     71.454
%     69.551
%     74.906
%     91.660
%    119.961
%    154.000
%    183.026
%    198.001
%    198.695
%    193.987
%    195.782
%    212.919
%    247.982
%    295.461
% 
% k =  0.025000
% J =  64.382
% 
% == [ml-class] Submitted Assignment 1 - Part 5 - Computing Cost (for multiple variables)
% == Nice work!
