function refresh() { //刷新
    ctx.clearRect(0, 0, 500, 500)
    drawLine()
}
function drawLine() { //繪製 x y 軸

    //x軸
    ctx.beginPath();
    ctx.moveTo(0, 250);
    ctx.lineTo(500, 250);
    ctx.stroke()
    ctx.fill()

    //y軸
    ctx.beginPath();
    ctx.moveTo(250, 0);
    ctx.lineTo(250, 500);
    ctx.stroke()
    ctx.fill()

}
function drawDot(x, y) {   //畫點

    ctx.beginPath()
    ctx.arc(x, y, 2, 0, 2 * Math.PI, true)
    ctx.fillStyle = "rgb(255, 39, 39)"
    ctx.fill()
}



tinkerbellBtn.addEventListener('click', function () {
    if (current === 'tinkerbell') {
        return
    }
    getData('tinkerbell').then(res => {
        data = res
        clearInterval(drawInterval)
        refresh()
        selectSystem = new TinkerBell()
        current = 'tinkerbell'
        count = 1
        counter.textContent=`${count}`
        title.textContent = "Tinkerbell Chaos System"
        formula.innerHTML = tinkerbellFormula
    })

})

duffingBtn.addEventListener('click', function () {
    if (current === 'duffing') {
        return
    }
    getData('duffing').then(res => {
        data = res
        clearInterval(drawInterval)
        refresh()
        selectSystem = new Duffing()
        current = 'duffing'
        count = 1
        counter.textContent=`${count}`
        title.textContent = "Duffing Chaos System"
        formula.innerHTML = duffingFormula
    })

})

henonBtn.addEventListener('click', function () {
    if (current === 'henon') {
        return
    }
    getData('henon').then(res => {
        data = res
        clearInterval(drawInterval)
        refresh()
        selectSystem = new Henon()
        current = 'henon'
        count = 1
        counter.textContent=`${count}`
        title.textContent = "Henon Chaos System"
        formula.innerHTML = henonFormula
    })

})

drawBtn.addEventListener('click', function () {
    if (data) {
        const len=data.length
        if(count===1){
            drawInterval=setInterval(() => {
                let coordinate=selectSystem.draw(data[count].x,data[count].y)
                drawDot(coordinate[0],coordinate[1])
                count++
                counter.textContent=`${count}`
                if(count>=len){
                    clearInterval(drawInterval)
                }
            }, 5);
        }else if(count>=len){
            refresh()
            count=1
            counter.textContent=`${count}`
            drawInterval=setInterval(() => {
                let coordinate=selectSystem.draw(data[count].x,data[count].y)
                drawDot(coordinate[0],coordinate[1])
                count++
                counter.textContent=`${count}`
                if(count>=len){
                    clearInterval(drawInterval)
                }
            }, 5);
        }
    }else{
        alert("no data")
    }



})

drawLine()
getData('tinkerbell').then(res => {
    data = res
})
