function [J, grad] = linearRegCostFunction(X, y, theta, lambda)
%LINEARREGCOSTFUNCTION Compute cost and gradient for regularized linear 
%regression with multiple variables
%   [J, grad] = LINEARREGCOSTFUNCTION(X, y, theta, lambda) computes the 
%   cost of using theta as the parameter for linear regression to fit the 
%   data points in X and y. Returns the cost in J and the gradient in grad

% Initialize some useful values
m = length(y); % number of training examples

% You need to return the following variables correctly 
J = 0;
grad = zeros(size(theta));

% ====================== YOUR CODE HERE ======================
% Instructions: Compute the cost and gradient of regularized linear 
%               regression for a particular choice of theta.
%
%               You should set J to the cost and grad to the gradient.
%

% print out params
% X
% y
% theta
% lambda

n = length(theta); % number of parameters

predictions = X * theta;
squaredErrors = (predictions - y) .^ 2;
J = (1/(2*m)) * (sum(squaredErrors) + lambda * sum(theta(2:n) .^ 2));

% gradient
kr = zeros(n, 1);
kr(2:n) = lambda/m;
errors = predictions - y;
grad = (1/m)*sum((X .* errors), 1)' + kr .* theta;

% =========================================================================

grad = grad(:);

end
