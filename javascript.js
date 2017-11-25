// ES6

// 변수 정의
const foo = 'ab2v9bc13j5jf4jv21';

// 정규표현식으로 숫자인 것과 홀수인 것을 추출 한 뒤 제곱으로 바꿔서 배열에 담는다
const bar = foo.match(/\d+/g).filter(v => v % 2 != 0).map(v => Math.pow(parseInt(v), 2));

// javascript 에서 제공하는 reduce를 이용해서 배열 내부의 값들을 더한다
console.log(bar.reduce((prevValue, currentValue) => prevValue + currentValue))