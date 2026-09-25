class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {

        return strs.map(str=> str.split('').reverse().join('')).join(';;')
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        if (!str) return '';
        if (str === '') return '';
        console.log(`str is ${str}`)
        return str.split(';;').map(el=> el.split('').reverse().join(''))
    }
}
