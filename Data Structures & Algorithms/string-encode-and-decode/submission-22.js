class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        console.log('***ENCODING***')
        if (strs.length === 0) return;
        const lengths = []

        // collect string lengths for each word
        for (let str of strs) { lengths.push(str.length) }
        // join each word with #;length;#
        let encoded = ''
        for (let i = 0; i < strs.length; i++) { encoded += `#;${lengths[i]};#${strs[i]}` }
        // let encoded = strs.map(str=> str.split('').reverse().join('')).join(';;')
        console.log({encoded})
        return encoded
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        if (str === undefined) return ''

        let [scan, result, tagMatch] = [`${str}`, [], /#;(\d);#/g]
        let match;
        while (match = tagMatch.exec(scan)){
            let wordStart = match.index + match[0].length
            let wordEnd = wordStart + Number(match[1])
            result.push(scan.slice(wordStart, wordEnd))
        }
        return result
    }
}
