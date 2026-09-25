class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
        const [tFrequencies, sFrequencies] = [{}, {}]
        for (let char of s) {
            sFrequencies[char] ||= 0
            sFrequencies[char]++
        }
        for (let char of t) {
            tFrequencies[char] ||= 0
            tFrequencies[char]++
        }

        for (let key of Object.keys(sFrequencies)) {
            if (sFrenquencies[key] !== tFrequencies[key]) return false
        }
        
        return true
    }
}
