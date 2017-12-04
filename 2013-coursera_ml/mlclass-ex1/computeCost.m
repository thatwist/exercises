function J = computeCost(X, y, theta)
%COMPUTECOST Compute cost for linear regression
%   J = COMPUTECOST(X, y, theta) computes the cost of using theta as the
%   parameter for linear regression to fit the data points in X and y

% Initialize some useful values
m = length(y); % number of training examples

% You need to return the following variables correctly 
J = 0;

% ====================== YOUR CODE HERE ======================
% Instructions: Compute the cost of a particular choice of theta
%               You should set J to the cost.

% print out params
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



% Submitting output
% 
% == Connecting to ml-class ... 
%
% X =
% 
%     1.0000    3.4572
%     1.0000    4.1961
%     1.0000    4.9350
%     1.0000    5.6739
%     1.0000    6.4128
%     1.0000    7.1517
%     1.0000    7.8906
%     1.0000    8.6295
%     1.0000    9.3684
%     1.0000   10.1073
%     1.0000   10.8462
%     1.0000   11.5851
%     1.0000   12.3241
%     1.0000   13.0630
%     1.0000   13.8019
%     1.0000   14.5408
%     1.0000   15.2797
%     1.0000   16.0186
%     1.0000   16.7575
%     1.0000   17.4964
% 
% y =
% 
%     3.3480
%     4.5439
%     5.9972
%     7.3354
%     8.2459
%     8.6391
%     8.6955
%     8.7709
%     9.2115
%    10.1728
%    11.5389
%    12.9826
%    14.1363
%    14.7836
%    14.9724
%    14.9895
%    15.2115 
%    15.9079
%    17.1010
%    18.5538
% 
% theta =
% 
%    0.50000
%   -0.50000
% 
% predictions =
% 
%   -1.2286
%   -1.5980
%   -1.9675
%   -2.3370
%   -2.7064
%   -3.0759
%   -3.4453
%   -3.8148
%   -4.1842
%   -4.5537
%   -4.9231
%   -5.2926
%   -5.6620
%   -6.0315
%   -6.4009
%   -6.7704
%   -7.1398
%   -7.5093
%   -7.8787
%   -8.2482
% 
% squaredErrors =
% 
%     20.946
%     37.724
%     63.437
%     93.555
%    119.953
%    137.241
%    147.398
%    158.399
%    179.445
%    216.870
%    270.999
%    333.983
%    391.974
%    433.269
%    456.819
%    473.493
%    499.581
%    548.365
%    623.986
%    718.346
% 
% k =  0.025000
% J =  148.14
% 
% == [ml-class] Submitted Assignment 1 - Part 2 - Computing Cost (for one variable)
% == Nice work!
