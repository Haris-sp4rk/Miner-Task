function taylorExpansion(){
    var i, fact=1;
    var sum=0.0;
    const x = prompt("Enter value of x: ");
    const num = prompt("Enter number of terms: ");

    for(var i=1; i<num; i++){
        fact = fact*i;
        sum = sum+(Math.pow(x,i)/fact);
}

    sum = sum+1;
    // alert(sum);
    return sum;

}
