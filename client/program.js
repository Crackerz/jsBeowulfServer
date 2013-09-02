{
	evaluate:function(str) {
		switch(str) {
		case 'pi':
			this.funcpi();
			break;
		default:
			log.write(str)
			break;
		}
	},
	exp:1,
	pi:0,
	term:0,
	funcpi:function() {
			this.pi += ((4/(8*this.term+1))-(2/(8*obj.term+4))-(1/(8*obj.term+5))-(1/(8*obj.term+6))) * obj.exp;
			this.exp *= (1/16)
			this.term+=1
			log.write(this.pi)
	}
}
