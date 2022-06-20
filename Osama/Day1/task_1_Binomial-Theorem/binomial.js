const prompt = require("prompt-sync")({
    fake_val : "OPTIONAL CONFIG VALUES HERE",
});

function binomial(n, k) {
   var coeff = 1;
   for (var x = n-k+1; x <= n; x++) coeff *= x;
   for (x = 1; x <= k; x++) coeff /= x;
   return coeff;
}

var x = prompt("Enter first value: ");
var y = prompt("Enter second value: ");

console.log(binomial(x,y));