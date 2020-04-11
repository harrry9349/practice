const leap = new Vue({     
  el: '#quick-sort',  
  data: {
      len: '',
      max: '',
      min: '',
      list_before: [],
      list_after: []
  },
  methods: {
      QuickSort() {

          // 「下限」から「上限」までのランダムな値が「個数」個入った配列を作成する
          let num = new Array()
          for(let i = 0; i < this.len; i++){
              num[i] = parseInt(Math.random() * (this.max - this.min) + this.min)
          }
          
          // ソート前の配列を保存する
          this.list_before = num.slice()

          
          // 分割統治法:大きな問題を小さな問題に分割してそれぞれ解き、結果をまとめて元の問題の解を得る手法
          // 解法が分からないときの手掛かりになるだけでなく、計算時間が問題サイズに従って急激に伸びる場合には直接全体を解くより高速になることが多々ある。
          // 上記引用：https://qiita.com/HMMNRST/items/18a206d56a3425e7717b

          // クイックソートを実行する関数の内容
          function sort(Start,End){
            // ピボットの値　= 初値と終値の平均値
            let pivot = num[Math.floor((Start + End)/2)]
            
            // 初値 = 左端、終値 = 右端　と関数内で暫定的に設定する
            let left = Start
            let right = End
            console.log("left="+left+"right="+right+"pivot="+pivot)
            
            // ピボットより小さい値を左側、大きい値を右側に分割する
            while(true){
              while(num[left]<pivot){
                // 左端がピボットより小さい場合、左端を一つ右に寄せる
                console.log("left++")
                left++
              }

              while(pivot<num[right]){
                // 右端がピボットより大きい場合、右端を一つ左に寄せる
                console.log("right--")
                right--
              }

              if(right<=left){
                // 左端が右端の値を逆転した場合、ループを終了する
                console.log("break")
                break
              }

              //左端が右端の値を逆転していない場合、左端と右端の値を交換する
              console.log("change:num[left]="+num[left]+"num[right]"+num[right])
              let tmp = num[left]
              num[left] = num[right]
              num[right] = tmp
              //交換後に左端を一つ右に、右端を一つ左に寄せる
              left++
              right--
            }

            // 大きな問題を小さな問題に分割して解決する
            if(Start<left-1){
              //左側に分割可能な配列がある場合、再帰する
              console.log("左再帰:Start"+Start+"left-1="+left-1)
              sort(Start,left-1)

            }

            if(right+1<End){
              //右側に分割可能な配列がある場合、再帰する
              console.log("右再帰:right+1"+right+1+"End="+End)
              sort(right+1,End)
            }

          }

          // クイックソートを行う
          sort(0,this.len-1)

          // ソート後の配列を保存する
          this.list_after = num.slice();
      }
  }
})