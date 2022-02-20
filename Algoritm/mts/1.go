package main

const randomChars = 'ABaaBCDdeGFAAR';
let result = new Map();
for(let letter of randomChars.toUpperCase().sort().replace(/ /,'').split('')) {
		if (result.has(letter)) {
				result.set(letter, result.get(letter)+1);
		} else {
			result.set(letter, 1);
		}
}
let str = ""
for(let pair of result) {
str += pair.join('')
}
console.log(str)
