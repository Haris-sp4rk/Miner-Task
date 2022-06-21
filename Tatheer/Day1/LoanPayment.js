var a;
var p= prompt('Enter your loan amount : ');
var r = prompt('Enter nterest rate per period : ');
var n = prompt('Enter total number of periods :  ');
a=p*(r*(1+r)^n) / ( (1+r)^n -1 );
alert("Payment amount per period : " + a);