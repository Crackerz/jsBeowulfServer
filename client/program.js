{
	evaluate:function(str) {
		digit = parseInt(str,10);
		if(isNaN(digit)) {
			log.write('Invalid input: '+str);
		} else {
			log.write(this.spigot(digit));
		}
	},
	spigot:function(digit) {
		len = Math.floor(10*digit/3)+1;
		A = new Array(len);

		for(var i=0;i<len;i++) A[i]=2;
		var finalDigit = 0;
		var nines = 0;
		var predigit = 0;

		for(i=1;i<digit+1;i++) {
			q = 0;
			for(j=len;j>0;j--) {
				x = 10*A[j-1] + q*j;
				x = Math.floor(x);
				A[j-1] = this.mod(x,(2*j-1));
				A[j-1] = Math.floor(A[j-1]);
				q = x / ((2*j) -1);
				q = Math.floor(q);
			}

			A[0] = Math.floor(this.mod(q,10));
			A[0] = 0
			q = Math.floor(q/10);

			if(q==9) {
				nines++;
			} else if (q==10) {
				finalDigit = predigit+1;
				for(j=0;j<nines;j++) {
					finalDigit = 0;
				}
				predigit=0;
				nines=0;
			} else {
				finalDigit = predigit;
				predigit=q;
				for(j=0;j<nines;j++) {
					finalDigit = 9;
				}
				nines=0;
			}
		}
		return finalDigit;
	},
	mod:function(a,b) {
		return (a/b-Math.floor(a/b))*b;
	}
}
