function binomial() {
    n = 0;
    k = 0;
    var n = prompt("Enter first value: ");
    var k = prompt("Enter second value: ");
    var coeff = 1;
    for (var x = n - k + 1; x <= n; x++) coeff *= x;
    for (x = 1; x <= k; x++) coeff /= x;
    return coeff;
}

//document.getElementById("bino").innerHTML = binomial();
