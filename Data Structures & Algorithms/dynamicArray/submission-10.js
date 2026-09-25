class DynamicArray {
    /**
     * @constructor
     * @param {number} capacity
     */
    constructor(capacity) {
        // this.array = new Array(capacity)
               this.capacity = capacity;
        this.length = 0;
        this.arr = new Array(this.capacity).fill(0);
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
                if (this.length === this.capacity) {
            this.resize();
        }
        this.arr[this.length] = n;
        this.length++; 

        // let array = this.array
        // let end = array.length -1;
        // if (isFull(array)) {
        //     array = array.concat(new Array(1))
        //     end += 1;
        // } else if (typeof array[end] !== 'undefined') shiftValuesFromEndUntilEmptySlot(array) 
        // this.array[end] = n;
        // console.log('array after #pushback', array)


        // function isFull(arr) { return arr.flat().length === arr.length }
        // function shiftValuesFromEndUntilEmptySlot(arr) {
        //     console.log('array before shifting left', arr)
        //     arr.splice(array.findLastIndex(el=> el === undefined), 1)
        //     arr = array.concat(new Array(1))
        //     console.log('array after shifting left', arr)
        // }
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
        // this.array = this.array.concat(new Array(this.array.length))
                this.capacity *= 2;
        const newArr = new Array(this.capacity).fill(0);
        for (let i = 0; i < this.length; i++) {
            newArr[i] = this.arr[i];
        }
        this.arr = newArr;
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
