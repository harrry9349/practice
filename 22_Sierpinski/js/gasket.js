const btn = document.getElementById("btn")

btn.addEventListener("click", ev =>{

    // キャンバス
    const canvas = document.getElementById("canvas")
    const ctx = canvas.getContext('2d')
    const size = window.innerHeight / 3 * 2
    canvas.width = canvas.height = size
    // 階層
    const num = document.getElementById("num").value
    const n = parseInt(num)
    
    // 三角形を描画する関数
    function triangle(x0, y0, x1, y1, x2, y2){
      if(canvas.getContext && n > 0) {
        ctx.beginPath()
        ctx.moveTo(x0, y0)
        ctx.lineTo(x1, y1)
        ctx.lineTo(x2, y2)
        ctx.fill()
      }
    }
    
    
    function gasket(n, x00, y00, x11, y11, x22, y22){
      if (n === 0) return 
      console.assert(n > 0) 
      
      const x01 = (x00 + x11) / 2
      const y01 = (y00 + y11) / 2
      const x12 = (x11 + x22) / 2
      const y12 = (y11 + y22) / 2
      const x02 = (x00 + x22) / 2
      const y02 = (y00 + y22) / 2
      
      triangle(x01, y01, x12, y12, x02, y02)
      
      gasket(n - 1, x00, y00, x01, y01, x02, y02)
      gasket(n - 1, x01, y01, x11, y11, x12, y12)
      gasket(n - 1, x02, y02, x12, y12, x22, y22)
    }
    
    
    const trix = size
    const triy = size * Math.sqrt(3) / 2
    ctx.fillStyle = "rgb(255, 10, 30)"
    triangle(trix/2, 0, 0, triy, trix, triy)
    
    
    ctx.fillStyle = "rgb(255, 255 ,255)"
    gasket(n, trix/2, 0, 0, triy, trix, triy)
    
      
    }, false)