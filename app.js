var _ul = document.createElement("ul");
var _div = document.getElementById("root");
_div.appendChild(_ul);

function add_new_element(input_text) {
    var _li = document.createElement("li");
    var _span = document.createElement("span");
    _span.innerHTML = input_text;
    var _button = document.createElement("button");
    _button.innerHTML = "Удалить";
    _button.addEventListener("click", function () { _li.remove() });
    _li.appendChild(_span);
    _li.appendChild(_button);
    _ul.appendChild(_li);
}

add_new_element("Сделать задание #3 по web-программированию");

var _input = document.createElement("input");
_input.id = "add_task_input";
var _add_button = document.createElement("button");
_add_button.id = "add_task";
_add_button.innerHTML = "Добавить";
_add_button.addEventListener("click", function () { add_new_element(_input.value) });
_div.appendChild(_input);
_div.appendChild(_add_button);