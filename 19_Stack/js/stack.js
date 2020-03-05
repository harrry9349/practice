const leap = new Vue({     
    el: '#stack-stack',  
    data: {
        stack_A: [1,2,3],
        stack_B: [4,5,6],
        stack_C: [7,8,9],

    },
    methods: {
        init() {
            this.stack_A = [1,2,3]
            this.stack_B = [4,5,6]
            this.stack_C = [7,8,9]        
        },

        push_A() {
            // スタックBからスタックAに要素をpushする
            //this.stack_A.length
            if (this.stack_A.length < 5){
                if(this.stack_B.length > 0){
                    let len = this.stack_B.length 
                    this.stack_A.push(this.stack_B[len-1])
                    this.stack_B.pop()

                }else{
                    alert("スタックBが空です。")
                }
            }else{
                alert("スタックAがいっぱいです。")
            }
        },

        pop_A() {
            // スタックAからスタックBに要素をpopする
            if (this.stack_A.length > 0){
                let len = this.stack_A.length 
                this.stack_B.push(this.stack_A[len-1])
                this.stack_A.pop()

            }else{
                alert("スタックAが空です。")
            }
        },

        push_C() {
            // スタックBからスタックCに要素をpushする
            if (this.stack_C.length < 5){
                if(this.stack_B.length > 0){
                    let len = this.stack_B.length 
                    this.stack_C.push(this.stack_B[len-1])
                    this.stack_B.pop()

                }else{
                    alert("スタックBが空です。")
                }
            }else{
                alert("スタックCがいっぱいです。")
            }
        },

        pop_C() {
            // スタックCからスタックBに要素をpopする
            if (this.stack_C.length > 0){
                let len = this.stack_C.length 
                this.stack_B.push(this.stack_C[len-1])
                this.stack_C.pop()

            }else{
                alert("スタックCが空です。")
            }
        }        

    }
})