class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        if (strs.length === 0) return;
        const lengths = []
        let encoded = ''

        for (let str of strs) lengths.push(str.length)
        for (let i = 0; i < strs.length; i++) encoded += `#;${lengths[i]};#${strs[i]}`
        
        return encoded
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        if (str === undefined) return ''

        let [scan, result, tagMatch] = [`${str}`, [], /#;(\d+);#/g]
        let match;
        while (match = tagMatch.exec(scan)){
            let wordStart = match.index + match[0].length
            let wordEnd = wordStart + Number(match[1])
            result.push(scan.slice(wordStart, wordEnd))
        }
        return result
    }
}
