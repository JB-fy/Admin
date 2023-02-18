
event.preventDefault(); //阻止冒泡：标签的默认动作
event.stopPropagation();    //阻止冒泡：标签的click事件
console.log('哈哈'.length)