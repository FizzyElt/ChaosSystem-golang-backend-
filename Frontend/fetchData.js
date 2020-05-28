function getData(system = '') {
  return new Promise((resolve, reject) => {
    fetch('/' + system)
      .then(res => res.json())
      .then(res => {
        current = system
        resolve(res)
      })
  })
}
