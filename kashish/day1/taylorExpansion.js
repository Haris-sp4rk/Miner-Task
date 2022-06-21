function taylorExpansion(x,numOfTerms){
    var i, fact=1;
    var sum=0.0;
    
    for(var i=1; i<numOfTerms; i++){
        fact = fact*i;
        sum = sum+(Math.pow(x,i)/fact);
    }

    sum = sum+1;
    console.log(sum)
    ex = document.getElementById('x')
    num = document.getElementById('num')
    if(ex.value != '' && num.value !=''){
        document.querySelector(".output-div").textContent = sum;
        document.querySelector(".error-message").textContent=""
    }
    else{
        document.querySelector(".output-div").textContent = "";
        document.querySelector(".error-message").textContent = "Enter the values of x and num to compute taylor series"
    }
    

}
