const leap = new Vue({     
    el: '#stack-stack',  
    data: {
        num: '',
        messages: []
    },
    methods: {
        init(){
            this.messages = []
            this.hanoi(this.num,'左','右','中')
            this.messages.push("完了")
        },
        hanoi(n,from_stack,to_stack,tmp){
            if (n == 1){
                console.log("L14:n=" + n + ",from_stack=" + from_stack +",to_stack=" + to_stack + ",tmp=" +tmp )
                this.messages.push(n + " を " + from_stack + " から " + to_stack + " へ移動")
            }else{
                console.log("L17:n=" + n + ",from_stack=" + from_stack +",to_stack=" + to_stack + ",tmp=" +tmp )
                this.hanoi(n-1,from_stack,tmp,to_stack)
                this.messages.push(n + " を " + from_stack + " から " + to_stack + " へ移動")
                console.log("L20:n=" + n + ",from_stack=" + from_stack +",to_stack=" + to_stack + ",tmp=" +tmp )
                this.hanoi(n-1,tmp,to_stack,from_stack)
            }
        },

    }
})