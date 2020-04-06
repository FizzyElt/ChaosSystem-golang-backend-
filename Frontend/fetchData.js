function getData(system = "") {
    return new Promise((resolve, reject) => {
        fetch('/' + system)
            .then(res => res.json())
            .then(res => {
                current = system
                console.log(res)
                resolve(res)
            })
    })
}


