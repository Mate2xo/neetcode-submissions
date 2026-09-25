class DynamicArray {
    /**
     * @constructor
     * @param {number} capacity
     */
    constructor(capacity) {
        this.array = new Array(capacity)
    }

    /**
     * @param {number} i
     * @returns {number}
     */
    get(i) {
        return this.array[i]
    }

    /**
     * @param {number} i
     * @param {number} n
     * @returns {void}
     */
    set(i, n) {
        this.array[i] = n;
    }

    /**
     * @param {number} n
     * @returns {void}
     */
    pushback(n) { 
        let array = this.array
        let end = array.length -1;
        if (isFull(array)) {
            array = array.concat(new Array(1))
            end += 1;
        } else if (typeof array[end] !== 'undefined') shiftValuesFromEndUntilEmptySlot(array) 
        this.array[end] = n;


        function isFull(arr) { return arr.flat().length === arr.length }
        function shiftValuesFromEndUntilEmptySlot(arr) {
            arr.splice(array.findLastIndex((new Array(1))[0]), 1)
            arr = array.concat(new Array(1))
        }
    }

    /**
     * @returns {number}
     */
    popback() {
        const lastValue = this.array.pop()
        this.array = this.array.concat(new Array(1))

        return lastValue;
    }

    /**
     * @returns {void}
     */
    resize() {
        this.array = this.array.contact(new Array(this.array.length))
    }

    /**
     * @returns {number}
     */
    getSize() {
        return this.array.flat().length
    }

    /**
     * @returns {number}
     */
    getCapacity() {
        return this.array.length
    }
}
