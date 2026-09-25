class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        let encoded = strs.map(str=> str.split('').reverse().join('')).join(';;')
        console.log('encoded:', encoded)
        return encoded
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        console.log(`str is`)
        // if (!str) {
        //     console.log('no string')
        //     return ''
        // };
        if (str === '') {
            console.log('empty string')
            return ['']
        };
        return str.split(';;').map(el=> el.split('').reverse().join(''))
    }
}
