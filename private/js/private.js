let createDefBlock = document.getElementById('createDefBlock');
let x = 1;
if (createDefBlock){
    createDefBlock.onclick = () => {
        let main = document.getElementById("createBlocks");
        let div1 = document.createElement("div");
        let div2 = document.createElement("div");
        let div3 = document.createElement("div");
        let textarea = document.createElement("textarea");
        let span2 = document.createElement("span");
        let btn = document.createElement("div");

        div1.className = "test_block";
        div1.id = ("test_block"+x);
        main.append(div1);

        btn.id = x;
        btn.className = "btn-remove";
        btn.setAttribute("onclick", "remove(this)");
        div1.append(btn);

        textarea.className = "test_textarea";
        textarea.placeholder = "Введите вопрос";
        textarea.id = ("textarea"+x);
        textarea.name = ("textarea"+x);
        div1.append(textarea);

        div2.id = ("defBlock"+x);
        div1.append(div2);

        for (let i = 0; i < 4; i++){
            let input = document.createElement("input");
            input.className = "test_input";
            input.id = ("test_input"+x+"/"+(i+1));
            input.name = ("test_input"+x+"/"+(i+1));
            input.placeholder = ("Вариант ответа "+(i+1));
            div2.append(input);
        }

        span2.innerHTML = "Верный вариант ответа:";
        span2.className = "test_span1";
        div1.append(span2);

        div3.className = "test_radioBlock";
        div3.id = ("test_radioBlock"+x);
        div1.append(div3);

        for (let i = 0; i < 4; i++){
            let radio = document.createElement("input");
            let label = document.createElement("label");
            radio.type = "radio";
            radio.name = ("test_option"+x);
            radio.id = ("test_values"+x+"/"+(i+1));
            radio.value = (i+1);
            div3.append(radio);
            label.className = "radio-style1";
            label.textContent = (i+1);
            label.setAttribute ("for", ("test_values"+x+"/"+(i+1)));
            div3.append(label);
        }

        x++;
    }
}

let createCheckboxBlock = document.getElementById("createCheckboxBlock");
if (createCheckboxBlock) {
    createCheckboxBlock.onclick = () => {
        let main = document.getElementById("createBlocks");
        let div1 = document.createElement("div");
        let div2 = document.createElement("div");
        let textarea = document.createElement("textarea");
        let btn = document.createElement("div");

        div1.className = "test_block";
        div1.id = ("test_block"+x);
        main.append(div1);


        btn.id = x;
        btn.className = "btn-remove";
        btn.setAttribute("onclick", "remove(this)");
        div1.append(btn);

        textarea.className = "test_textarea";
        textarea.placeholder = "Введите вопрос";
        textarea.id = ("textarea"+x);
        textarea.name = ("textarea"+x);
        div1.append(textarea);

        div2.id = ("checkBoxBlock"+x);
        div1.append(div2);

        let checkBox = document.createElement("input");
        let label = document.createElement("label");
        checkBox.type = "checkbox"
        checkBox.id = ("checkBox"+x+"/1");
        checkBox.value = "1"
        label.className = "checkbox-style";
        label.setAttribute ("for", ("checkBox"+x+"/1"));
        div2.append(checkBox);
        div2.append(label);

        let input = document.createElement("input");
        input.className = "test_input";
        input.id = ("test_input"+x+"/1");
        input.name = ("test_input"+x+"/1");
        input.placeholder = ("Вариант ответа 1");
        label.append(input);

        let btnAdd = document.createElement("div");
        btnAdd.id = "add";
        btnAdd.textContent = "Добавить вариант ответа";
        btnAdd.className = "btn-addCheckBox";
        btnAdd.setAttribute("onclick", "addCheckBox(checkBoxBlock"+x+")");
        div1.append(btnAdd);

        x++;
    }
}

function addCheckBox(el){
    let str = (el.id).split("checkBoxBlock");
    let numOfId = str[1];
    let childsLength = el.childNodes.length;
    str = (el.childNodes[childsLength-2].id).split("/");
    let childId = str[1];
    childId = parseInt(childId);
    childId++;

    let checkBox = document.createElement("input");
    let label = document.createElement("label");
    checkBox.type = "checkbox";
    checkBox.id = ("checkBox"+numOfId+"/"+childId);
    checkBox.value = childId;
    label.className = "checkbox-style";
    label.setAttribute ("for", ("checkBox"+numOfId+"/"+childId));
    el.appendChild(checkBox);
    el.appendChild(label);

    let input = document.createElement("input");
    input.className = "test_input";
    input.id = ("test_input"+numOfId+"/"+childId);
    input.name = ("test_input"+numOfId+"/"+childId);
    input.placeholder = ("Вариант ответа "+childId);
    label.appendChild(input);
}

function addSelectBlock(el){
    let table = el.getElementsByClassName("test_block_table")[0];
    let div = document.createElement("div");
    let input1 = document.createElement("input");
    let input2 = document.createElement("input");
    div.className = "test_block_bar";
    input1.className = "test_block_s_left";
    input2.className = "test_block_s_right";
    div.append(input1, input2);
    table.append(div)
}

function remove(btn){
    let form = document.getElementById('createBlocks');
    let divDelete = document.getElementById("test_block"+btn.id);
    form.removeChild(divDelete);
}

let createSelectBlock = document.getElementById("createSelectBlock");
if (createSelectBlock){
    createSelectBlock.onclick = () => {
        let main = document.getElementById("createBlocks");
        let div1 = document.createElement("div");
        let div2 = document.createElement("div");
        let div3 = document.createElement("div");
        let div4 = document.createElement("div");
        let div5 = document.createElement("div");
        let div6 = document.createElement("div");
        let div7 = document.createElement("div");
        let input1 = document.createElement("input");
        let input2 = document.createElement("input");
        let btn = document.createElement("div");

        btn.id = x;
        btn.className = "btn-remove-s";
        btn.setAttribute("onclick", "remove(this)");

        div1.id = ("test_block"+x);
        div1.className = "test_block_s";

        div2.className = "test_block_bar";
        div2.id = "selectBlock"+x;
        div3.className = "test_block_s_title";
        div3.textContent = "Вопросы";
        div4.className = "test_block_s_title";
        div4.textContent = "Ответы";

        div5.className = "test_block_table";
        div5.id = "selectBlock"+x;
        div6.className = "test_block_bar";
        input1.className = "test_block_s_left";
        input2.className = "test_block_s_right";

        div7.className = "btn-addCheckBox";
        div7.textContent = "Добавить";
        div7.setAttribute("onclick", "addSelectBlock(test_block"+x+")");

        div6.append(input1, input2);
        div5.append(div6)
        div2.append(div3, div4)
        div1.append(btn, div2, div5,div7);
        main.append(div1)

        x++;
    }
}

function checkCreate(){
    let count = 0;
    let title = document.getElementById('tit');
    if (title.value === ""){
        title.style.border = "3px solid red";
        count++;
    } else {
        title.style.border = "3px solid #d8a929";
    }
    let blocksBody = document.getElementById('createBlocks');
    for (let i = 1; i < blocksBody.childNodes.length; i++){
        let checkedCounter = 0;
        let testBlockType = blocksBody.childNodes[i].childNodes[2].id;
        if (testBlockType.includes("defBlock")){
            if (blocksBody.childNodes[i].childNodes[1].value === ""){
                blocksBody.childNodes[i].childNodes[1].style.border = "3px solid red";
                count++;
            } else {
                blocksBody.childNodes[i].childNodes[1].style.border = "3px solid #d8a929";
            }
            for(let j = 0; j < 4; j++){
                if(blocksBody.childNodes[i].childNodes[2].childNodes[j].value === ""){
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].style.border = "2px solid red";
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].style.borderRadius = "10px";
                    count;
                } else {
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].style.border = "0px solid";
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].style.borderBottom = "2px solid #d8a929";
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].style.borderRadius = "0px";
                }
            }
            for (let j = 0; j < 7; j=j+2){
                if(blocksBody.childNodes[i].childNodes[4].childNodes[j].checked){
                    checkedCounter++;
                }
            }
            if (checkedCounter > 0){
                blocksBody.childNodes[i].childNodes[4].style.border = "0px solid red";
            } else {
                blocksBody.childNodes[i].childNodes[4].style.border = "2px solid red";
                count++;
            }
        } else if (testBlockType.includes("checkBoxBlock")){
            if (blocksBody.childNodes[i].childNodes[1].value === ""){
                blocksBody.childNodes[i].childNodes[1].style.border = "3px solid red";
                count++;
            } else {
                blocksBody.childNodes[i].childNodes[1].style.border = "3px solid #d8a929";
            }
            for(let j = 1; j < blocksBody.childNodes[i].childNodes[2].childNodes.length; j=j+2){
                if(blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].value === ""){
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].style.border = "2px solid red";
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].style.borderRadius = "10px";
                    count++;
                } else {
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].style.border = "0px solid";
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].style.borderBottom = "2px solid #d8a929";
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].style.borderRadius = "0px";
                }
                if(blocksBody.childNodes[i].childNodes[2].childNodes[j-1].checked){
                    checkedCounter++;
                }
            }
            if (checkedCounter > 0){
                for(let k = 1; k < blocksBody.childNodes[i].childNodes[2].childNodes.length; k=k+2){
                    blocksBody.childNodes[i].childNodes[2].childNodes[k].className = "checkbox-style";
                }
            } else {
                count++
                for(let k = 1; k < blocksBody.childNodes[i].childNodes[2].childNodes.length; k=k+2){
                    blocksBody.childNodes[i].childNodes[2].childNodes[k].className = "checkbox-style-2";
                }
            }
        } else if (testBlockType.includes("selectBlock")){
            for (let j = 0; j < blocksBody.childNodes[i].childNodes[2].childNodes.length; j++){
                if (blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].value === ""){
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].style.border = "1px solid red";
                    count++;
                } else {
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].style.border = "1px solid #d8a929";
                }
                if (blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[1].value === ""){
                   blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[1].style.border = "1px solid red";
                   count++;
                } else {
                    blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[1].style.border = "1px solid #d8a929";
                }
            }
        }
    }
    if (count === 0){
        return true
    } else {
        return false
    }
}

let saveTest = document.getElementById('saveTest');
if (saveTest){
    saveTest.onclick = () => {
        if (checkCreate()) {
            let formData = new FormData()
            let title = document.querySelector('#tit').value;
            formData.append("Title", title)
            let blocksBody = document.querySelector('#createBlocks');
            let countSelectBlock = 0;
            for (let i = 1; i < blocksBody.childNodes.length; i++){
                if (blocksBody.childNodes[i].childNodes[2].id.includes("defBlock")){
                    formData.append("defBlock Question", blocksBody.childNodes[i].childNodes[1].value);
                    for (let j = 0; j < 4; j++){
                        formData.append("defBlock Value", blocksBody.childNodes[i].childNodes[2].childNodes[j].value)
                    }
                    for (let j = 0; j < 7; j=j+2){
                        if(blocksBody.childNodes[i].childNodes[4].childNodes[j].checked){
                            formData.append("defBlock Answer", blocksBody.childNodes[i].childNodes[2].childNodes[(blocksBody.childNodes[i].childNodes[4].childNodes[j].value)-1].value);
                        }
                    }
                }
                if (blocksBody.childNodes[i].childNodes[2].id.includes("checkBoxBlock")) {
                    let count = 0;
                    formData.append("checkBoxBlock Question", blocksBody.childNodes[i].childNodes[1].value);
                    formData.append("checkBoxBlock Value Count", ""+(blocksBody.childNodes[i].childNodes[2].childNodes.length/2))
                    for (let j = 1; j < blocksBody.childNodes[i].childNodes[2].childNodes.length; j=j+2){
                        formData.append("checkBoxBlock Value", blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].value);
                    }
                    for (let j = 0; j < blocksBody.childNodes[i].childNodes[2].childNodes.length; j=j+2){
                        if(blocksBody.childNodes[i].childNodes[2].childNodes[j].checked){
                            formData.append("checkBoxBlock Answer", blocksBody.childNodes[i].childNodes[2].childNodes[j+1].childNodes[0].value);
                            count++;
                        }
                    }
                    formData.append("checkBoxBlock Answer Count", ""+count);
                }

                if (blocksBody.childNodes[i].childNodes[2].id.includes("selectBlock")){
                    formData.append("selectBlock ValueAnswer Count", blocksBody.childNodes[i].childNodes[2].childNodes.length)
                    for (let j = 0; j < blocksBody.childNodes[i].childNodes[2].childNodes.length; j++){
                        formData.append("selectBlock Value", blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[0].value);
                        formData.append("selectBlock Answer", blocksBody.childNodes[i].childNodes[2].childNodes[j].childNodes[1].value);
                    }
                    countSelectBlock++;
                }
            }
            if (countSelectBlock > 0){
                formData.append("selectBlock Question Count", countSelectBlock)
            }
            Send("POST", "/create/save", formData,2, response => {
                if (response){
                    window.location.href = "/";
                }
            })
        }
    }
}