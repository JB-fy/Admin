
console.log('哈哈'.length)

event.preventDefault(); //阻止冒泡：标签的默认动作
event.stopPropagation();    //阻止冒泡：标签的click事件

window.onscroll = function () {
    let scrollTop = document.documentElement.scrollTop || window.pageYOffset || document.body.scrollTop;
    if (scrollTop + $(window).height() >= $(document).height() - 50) {
        //滚动到距离底部50px时操作
    }
}