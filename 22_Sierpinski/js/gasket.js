const btn = document.getElementById("btn")

btn.addEventListener("click", ev =>{

    // キャンバス
    const canvas = document.getElementById("canvas")
    const ctx = canvas.getContext('2d')
    const size = window.innerHeight / 3 * 2
    canvas.width = canvas.height = size
    // 階層
    const num = document.getElementById("num").value
    const level = parseInt(num)
    
    // 三角形を描画する関数
    function Drawtriangle(x1, y1, x2, y2, x3, y3){
      if(canvas.getContext && level > 0) {
        ctx.beginPath()
        ctx.moveTo(x1, y1)
        ctx.lineTo(x2, y2)
        ctx.lineTo(x3, y3)
        ctx.fill()
      }
    }
    
    
    // 三角形を切り取る関数
    function Drawgasket(level, x_top, y_top, x_left, y_left, x_right, y_right){　
      //階層が0になったら完了
      if (level === 0) return 

      // 赤い三角形に内接する白い逆三角の座標を求める
      const x_left_white = (x_top + x_left) / 2 
      const y_left_white = (y_top + y_left) / 2 
      const x_top_white = (x_left + x_right) / 2 
      const y_top_white = (y_left + y_right) / 2 
      const x_right_white = (x_top + x_right) / 2 
      const y_right_white = (y_top + y_right) / 2 
      
      //白い逆三角を描画
      Drawtriangle(x_left_white, y_left_white, x_top_white, y_top_white, x_right_white, y_right_white)
      
      // 白い逆三角で分割された上部の赤三角に対して切り取り
      Drawgasket(level - 1, x_top, y_top, x_left_white, y_left_white, x_right_white, y_right_white)
      // 白い逆三角で分割された左下部の赤三角に対して切り取り
      Drawgasket(level - 1, x_left_white, y_left_white, x_left, y_left, x_top_white, y_top_white)
      // 白い逆三角で分割された右下の赤三角に対して切り取り
      Drawgasket(level - 1, x_right_white, y_right_white, x_top_white, y_top_white, x_right, y_right)
    }
    
    
    // ひとつの大きな赤い三角形を描画
    const trix = size
    const triy = size * Math.sqrt(3) / 2
    ctx.fillStyle = "rgb(255, 10, 30)"
    Drawtriangle(trix/2, 0, 0, triy, trix, triy)
    
    
    // 赤い三角形の切り取り（白い逆三角を描画）
    ctx.fillStyle = "rgb(255, 255 ,255)"
    Drawgasket(level, trix/2, 0, 0, triy, trix, triy)
    
      
    }, false)