const fs = require('fs')
const ops = {1: '移动', 2: '联通', 3: '电信', 4: '电信虚拟运营商', 5: '联通虚拟运营商', 6: '移动虚拟运营商'}

const init = module.exports = (db) => {
  const buf = fs.readFileSync(db || __dirname + '/phone.dat')
  const start = buf.readUInt32LE(4)
  const total = (buf.length - start) / 9 

  const format = (m) => {
    let idx = buf.readUInt32LE(m * 9 + start + 4)
    const type = buf.readUInt8(m * 9 + start + 8, 1)
    let i = idx;
    for (; buf.readInt8(i) !== 0; i++ ) {}
    const [province, city, postcode, telcode] = buf.toString('utf8', idx, i).split('|')
    return { province, city, postcode, telcode, operator: ops[type]}
  }

  return (no) => {
    if (no.length < 7) throw new Error('mobile phone number length is 7 at least')
    const mobile = no.substr(0, 7) 
    let right = total
    let left = 0
    let i = 0
    while (left < right - 1) {
      const m = left + parseInt((right - left) / 2) 
      const x = buf.readUInt32LE(m * 9 + start)
      if (x > mobile) right = m 
      else if (x < mobile) left = m 
      else return format(m)
    }
    throw new Error('Not Found')
  }
}
