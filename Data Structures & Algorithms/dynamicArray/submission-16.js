/**
 * Dynamic Array Class
 * Note: JavaScript arrays are dynamic by default,
 * but this class is used to simulate the behavior of a dynamic array.
 */
class DynamicArray {
    /**
     * @constructor
     * @param {number} capacity - Initial capacity of the array.
     */
    constructor(capacity) {
        this.capacity = capacity;
        this.length = 0;
        this.arr = new Array(this.capacity).fill(0);
    }

    /**
     * Get the value at the i-th index.
     * @param {number} i - The index to get the value from.
     * @returns {number} - The value at the i-th index.
     */
    get(i) {
        return this.arr[i];
    }

    /**
     * Set a value at the i-th index.
     * @param {number} i - The index to insert the value at.
     * @param {number} n - The value to insert.
     * @returns {void}
     */
    set(i, n) {
        this.arr[i] = n;
    }

    /**
     * Insert a value in the last position of the array.
     * @param {number} n - The value to insert.
     * @returns {void}
     */
    pushback(n) {
        console.log('array before #pushback', this.arr)
        if (this.length === this.capacity) {
            this.resize();
        }
        this.arr[this.length] = n;
        this.length++;
        console.log('array after #pushback', this.arr)
    }

    /**
     * Remove the last element in the array.
     * @returns {number}
     */
    popback() {
        return 2
        if (this.length > 0) {
            // soft delete the last element
            this.length--;
        }
        return this.arr[this.length];
    }

    /**
     * Resize the array to double its current capacity.
     * @returns {void}
     */
    resize() {
        this.capacity *= 2;
        const newArr = new Array(this.capacity).fill(0);
        for (let i = 0; i < this.length; i++) {
            newArr[i] = this.arr[i];
        }
        this.arr = newArr;
    }

    /**
     * Get the current size of the array.
     * @returns {number} - The current size.
     */
    getSize() {
        return this.length;
    }

    /**
     * Get the current capacity of the array.
     * @returns {number} - The current capacity.
     */
    getCapacity() {
        return this.capacity;
    }
}

// class DynamicArray {
//     /**
//      * @constructor
//      * @param {number} capacity
//      */
//     constructor(capacity) {
//         this.array = new Array(capacity)
//     }

//     /**
//      * @param {number} i
//      * @returns {number}
//      */
//     get(i) {
//         return this.array[i]
//     }

//     /**
//      * @param {number} i
//      * @param {number} n
//      * @returns {void}
//      */
//     set(i, n) {
//         this.array[i] = n;
//     }

//     /**
//      * @param {number} n
//      * @returns {void}
//      */
//     pushback(n) { 
//         let array = this.array
//         let end = array.length -1;
//         if (isFull(array)) {
//             array = array.concat(new Array(1))
//             end += 1;
//         } else if (typeof array[end] !== 'undefined') shiftValuesFromEndUntilEmptySlot(array) 
//         this.array[end] = n;
//         console.log('array after #pushback', array)


//         function isFull(arr) { return arr.flat().length === arr.length }
//         function shiftValuesFromEndUntilEmptySlot(arr) {
//             console.log('array before shifting left', arr)
//             arr.splice(array.findLastIndex(el=> el === undefined), 1)
//             arr = array.concat(new Array(1))
//             console.log('array after shifting left', arr)
//         }
//     }

//     /**
//      * @returns {number}
//      */
//     popback() {
//         const lastValue = this.array.pop()
//         this.array = this.array.concat(new Array(1))

//         return lastValue;
//     }

//     /**
//      * @returns {void}
//      */
//     resize() {
//         this.array = this.array.concat(new Array(this.array.length))
//     }

//     /**
//      * @returns {number}
//      */
//     getSize() {
//         return this.array.flat().length
//     }

//     /**
//      * @returns {number}
//      */
//     getCapacity() {
//         return this.array.length
//     }
// }
