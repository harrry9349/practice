const leap = new Vue({     
    el: '#queue-queue',  
    data: {
        queue_A: [1,2,3,4,5],
        queue_B: [6,7,8,9,10],
        queue_C: [11,12,13,14,15] ,

    },
    methods: {
        init() {
            this.queue_A = [1,2,3,4,5]
            this.queue_B = [6,7,8,9,10]
            this.queue_C = [11,12,13,14,15]        
        },

        InQ_A() {
            // キューBからキューAに要素をpushする
            //this.queue_A.length
            if (this.queue_A.length < 10){
                if(this.queue_B.length > 0){
                    this.queue_A.push(this.queue_B[0])
                    this.queue_B.shift()

                }else{
                    alert("キューBが空です。")
                }
            }else{
                alert("キューAがいっぱいです。")
            }
        },

        DeQ_A() {
            // キューAからキューBに要素をpopする
            if (this.queue_A.length > 0){
                this.queue_B.push(this.queue_A[0])
                this.queue_A.shift()

            }else{
                alert("キューAが空です。")
            }
        },

        InQ_C() {
            // キューBからキューCに要素をpushする
            if (this.queue_C.length < 10){
                if(this.queue_B.length > 0){
                    this.queue_C.push(this.queue_B[0])
                    this.queue_B.shift()

                }else{
                    alert("キューBが空です。")
                }
            }else{
                alert("キューCがいっぱいです。")
            }
        },

        DeQ_C() {
            // キューCからキューBに要素をpopする
            if (this.queue_C.length > 0){
                this.queue_B.push(this.queue_C[0])
                this.queue_C.shift()

            }else{
                alert("キューCが空です。")
            }
        }        

    }
})