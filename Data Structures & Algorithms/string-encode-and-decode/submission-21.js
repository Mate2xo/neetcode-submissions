class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        if (strs.length === 0) return;
        let encoded = strs.map(str=> str.split('').reverse().join('')).join(';;')
        return encoded
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        if (str === undefined) return ''
        return str.split(';;').map(el=> el.split('').reverse().join(''))
    }
}
