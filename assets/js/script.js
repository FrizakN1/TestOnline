function Send(method, uri, data, typeData, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, uri);

    xhr.onload = function (event) {
        if (callback && typeof callback === "function") {
            callback(JSON.parse(this.response));
        }
    }
    if (data) {
        if (typeData === 1){
            xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
            xhr.setRequestHeader("X-Requested-With", "XMLHttpRequest");
            xhr.send(JSON.stringify(data));
        } else {
            xhr.send(data);
        }
    } else {
        xhr.send();
    }
}

let email = document.getElementById('email');
if (email){
    let timeout = 0;
    email.oninput = function (e) {
        if (timeout !== 0) {
            clearTimeout(timeout);
            timeout = 0;
        }
        timeout = setTimeout(() => {
            Send("POST", "/registration/checkEmail", {Email: this.value}, 1, response => {
                if (response){
                    if (email.value !== "" && email.value.includes("@") &&  !email.value.includes(" ")){
                        email.dataset["status"] = "true";
                        email.style.border = "unset";
                    } else {
                        email.style.border = "2px solid red";
                    }
                } else {
                    email.dataset["status"] = "false";
                    email.style.border = "2px solid red";
                }
            })
        }, 400)
    }
}

function checkForm(){
    let name = document.getElementById('name');
    let surname = document.getElementById('surname');
    let pass = document.getElementById('password');
    let repass = document.getElementById('repassword');
    let stateMale = document.getElementById('male');
    let stateFemale = document.getElementById('famale');
    let state;
    let check = true;

    if (name.value === "" || name.value.includes(" ")){
        name.style.border = "2px solid red";
        check = false;
    } else {
        name.style.border = "unset";
    }
    if (surname.value === "" || surname.value.includes(" ")){
        surname.style.border = "2px solid red";
        check = false;
    } else {
       surname.style.border = "unset";
    }
    if (stateMale.checked){
        state = "Мужской"
        document.getElementById('stateStyle').style.border = "unset";
    } else if (stateFemale.checked){
        state = "Женский"
        document.getElementById('stateStyle').style.border = "unset";
    } else {
        document.getElementById('stateStyle').style.border = "2px solid red";
        check = false;
    }
    if (email.dataset["status"] === "false"){
        document.getElementById('email').style.border = "2px solid red";
        check = false;
    } else {
        document.getElementById('email').style.border = "unset";
    }
    if (pass.value === "" || pass.value.length > 16 || pass.value.includes(" ")){
        pass.style.border = "2px solid red";
        check = false;
    } else {
        pass.style.border = "unset";
    }
    if (repass.value === "" || repass.value !== pass.value){
        repass.style.border = "2px solid red";
        check = false;
    } else {
        repass.style.border = "unset";
    }
    if (check){
        let userData = {
            Name: name.value,
            Surname: surname.value,
            State: state,
            Email: email.value,
            Password: pass.value,
        }

        Send("POST", "/registration/checkForm", userData, 1, response => {
            if (response){
                let main = document.getElementsByClassName("main")[0];
                for (let i = 1; i < main.childNodes.length; i++){
                    main.childNodes[i].remove()
                }
                let h1 = document.createElement("h1");
                h1.textContent = "На вашу почту отправленно письмо с подтверждением";
                main.append(h1);
            }
        })
    }
}

function login(){
    let emailLog = document.getElementById('emailLogin');
    let password = document.getElementById('password');

    let User = {
        Email: emailLog.value,
        Password: password.value,
    }

    Send("POST", "/login/check", User, 1,response => {
        if (response) {
            window.location.href = "/";
        } else {
            emailLog.style.border = "2px solid red";
            password.style.border = "2px solid red";
        }
    })
}

let exit = document.getElementById('exit');
if (exit){
    exit.onclick = () => {
        Send("POST", "/exit", null, 1,response => {
            if (response){
                window.location.href = "/";
            }

        })
    }
}

function checkTest(){
    let count = 0;
    let TestID = document.body.childNodes[3].id.split('TestCollect')[1];
    let test = document.getElementById('TestCollect'+TestID).childNodes;
    for (let i = 1; i < test.length-2; i = i + 2){
        let temp = 0;
        if (test[i].id.includes("defBlock")){
            for (let j = 3; j < 19; j = j + 5){
                if (test[i].childNodes[j].checked) {
                    break;
                }
                temp++;
            }
            test[i].style.border = "unset";
            if (temp === 4){
                test[i].style.border = "2px solid red";
                count++;
            }
        }
        if (test[i].id.includes("checkBox")){
            for (let j = 3; j < test[i].childNodes.length; j = j + 4){
                if (test[i].childNodes[j].checked) {
                    temp++;
                }
            }
            test[i].style.border = "unset";
            if (temp === 0) {
                test[i].style.border = "2px solid red";
                count++;
            }
        }
    }
    return count === 0;

}

let Finish = document.getElementById("Finish");
if (Finish) {
    Finish.onclick = () => {
        if (checkTest()){
            let TestID = document.body.childNodes[3].id.split('TestCollect')[1];
            let test = document.getElementById('TestCollect'+TestID).childNodes;
            let testData = new FormData()
            testData.append("TestID",TestID);
            for (let i = 1; i < test.length-2; i = i + 2){
                if (test[i].id.includes("defBlock")){
                    let answer = "";
                    testData.append("QuestionID defBlock", test[i].id.split("defBlock")[1]);
                    testData.append("QuestionID", test[i].id.split("defBlock")[1]);
                    testData.append("QuestionType", "defBlock");
                    for (let j = 3; j < 19; j = j + 5){
                        if (test[i].childNodes[j].checked) {
                            answer = test[i].childNodes[j].value;
                        }
                    }
                    testData.append("QuestionAnswer defBlock", answer);
                }
                if (test[i].id.includes("checkBox")){
                    let temp = 0;
                    testData.append("QuestionID checkBox", test[i].id.split("checkBox")[1]);
                    testData.append("QuestionID", test[i].id.split("checkBox")[1]);
                    testData.append("QuestionType", "checkBox");
                    for (let j = 3; j < test[i].childNodes.length; j = j + 4){
                        if (test[i].childNodes[j].checked){
                            temp++;
                        }
                    }
                    testData.append("QuestionCount checkBox", temp);
                    for (let j = 3; j < test[i].childNodes.length; j = j + 4){
                        if (test[i].childNodes[j].checked){

                            testData.append("QuestionAnswer checkBox", test[i].childNodes[j].value);
                        }
                    }
                }
                if (test[i].id.includes("selectBlock")){
                    testData.append("QuestionID selectBlock", test[i].id.split("selectBlock")[1]);
                    testData.append("QuestionID", test[i].id.split("selectBlock")[1]);
                    testData.append("QuestionType", "selectBlock");
                    let count = 0;
                    for(let j = 1; j < test[i].childNodes[3].childNodes.length; j+=2){
                        testData.append("QuestionValue selectBlock", test[i].childNodes[3].childNodes[j].childNodes[1].textContent);
                        testData.append("QuestionAnswer selectBlock", test[i].childNodes[3].childNodes[j].childNodes[3].value);
                        count++;
                    }
                    testData.append("QuestionAnswerCount selectBlock", count);
                }
            }
            Send("POST", "/test/result", testData, 2, (response) =>{
                if (response.Result !== "Ошибка"){
                    let main = document.getElementsByClassName("main")[0];
                    for (let i = 1; i < main.childNodes.length; i++){
                        main.childNodes[i].remove();
                    }
                    let resultBlock = document.createElement("div");
                    let resultText1 = document.createElement("p");
                    let resultText2 = document.createElement("p");
                    let btnSubmit = document.createElement("div");
                    let a = document.createElement("a");
                    resultBlock.className = "result_block";
                    resultText1.className = "result_text";
                    resultText1.textContent = "Пользователь: " + response.Name + " " + response.Surname;
                    resultText2.className = "result_text";
                    resultText2.textContent = "Результат: " + response.Result;
                    btnSubmit.className = "btn-sumbit";
                    btnSubmit.textContent = "Завершить";
                    a.href = "/";
                    resultBlock.append(resultText1, resultText2);
                    a.append(btnSubmit);
                    main.append(resultBlock, a);
                }
            })
        }
    }
}

let searchRequest = document.getElementById("searchRequest");
if (searchRequest) {
    searchRequest.onclick = () => {
        Send("POST", "/search_result", {"TestName": document.getElementById("testName").value}, 1, (response) => {
            if (response) {
                let main = document.getElementsByClassName("main")[0];
                for (let i = 1; i < main.childNodes.length; i++){
                    main.childNodes[i].remove();
                }
                for (let i = 0; i < response.length; i++){
                    let bar = document.createElement("div");
                    let result = document.createElement("div");
                    bar.className = "test";
                    bar.textContent = response[i].Name+" "+response[i].Surname;
                    result.className = "test__btn";
                    result.textContent = response[i].Result;
                    bar.append(result);
                    main.append(bar);
                }
            }
        })
    }
}

let editProfile = document.getElementById("editProfile");
if (editProfile){
    editProfile.onclick = () => {
        let userDataBlock = document.getElementsByClassName("user_data")[1];
        let userData = new Array(4);
        for(let i = 1; i < userDataBlock.childNodes.length; i++){
            userData[i-1] = userDataBlock.childNodes[i].textContent;
            userDataBlock.childNodes[i].remove();
        }
        for(let i = 0; i < 4; i++){
            let div = document.createElement("div");
            let input = document.createElement("input");
            input.placeholder = userData[i];
            input.className = "profile_input";
            div.append(input);
            userDataBlock.append(div);
        }
        console.log(userData)
    }
}

let editElement = document.getElementsByClassName("edit");
if (editElement.length > 0){
    let userDataBlock = document.getElementsByClassName("user_data")[1];
    let userData = new Array(4);
    let count = 0;
    for(let i = 1; i < userDataBlock.childNodes.length; i = i + 2){
        userData[count] = userDataBlock.childNodes[i].childNodes[1].textContent;
        count++;
    }
    function restart(){
        for (let i = 0; i < 4; i++){
            if (editElement[i].textContent === "Изменить"){
                editElement[i].onclick = () => {
                    structProfile(i, userData);
                }
            }
        }
    }
    restart()
}

function structProfile(x, userData) {
    let userDataBlock = document.getElementsByClassName("user_data")[1];
    if (userDataBlock.childNodes[1].childNodes.length>0){
        for(let i = 1; i < userDataBlock.childNodes.length; i++){
            userDataBlock.childNodes[i].remove()
        }
    } else {
        for(let i = 0; i < 4; i++){
            userDataBlock.childNodes[5].remove()
        }
    }
    for(let i = 0; i < 4; i++){
        if (i === x) {
            let div = document.createElement("div");
            let input = document.createElement("input");
            let btn = document.createElement("div");
            div.className = "user_data_element";
            input.className ="profile_input";
            if (x === 2){
                input.placeholder = "Новая почта";
            } else if (x === 3){
                input.type = "password";
                input.placeholder = "Старый пароль";
            } else {
                input.placeholder = userData[i];
            }
            btn.className = "edit";
            btn.textContent = "Применить";
            if (x === 0){
                btn.setAttribute("onclick", "changeName()");
                div.append(input, btn);
            } else if (x === 1) {
                btn.setAttribute("onclick", "changeSurname()");
                div.append(input, btn);
            } else if (x === 2) {
                btn.setAttribute("onclick", "changeEmail()");
                div.append(input, btn);
            } else if (x === 3) {
                let input1 = document.createElement("input");
                let input2 = document.createElement("input");
                input1.className ="profile_input password_input";
                input1.placeholder = "Новый пароль";
                input1.type = "password";
                input2.type = "password";
                input2.className ="profile_input password_input";
                input2.placeholder = "Повторите новый пароль";
                btn.setAttribute("onclick", "changePassword()");
                div.append(input, input1, input2, btn);
            }
            userDataBlock.append(div)
        } else {
            let div = document.createElement("div");
            let div1 = document.createElement("div");
            let btn = document.createElement("div");
            div.className = "user_data_element";
            div1.textContent = userData[i];
            btn.className = "edit";
            btn.textContent = "Изменить";
            div.append(div1, btn);
            userDataBlock.append(div);
        }
    }
    restart();
}

function changeName(){
    let newName = document.getElementsByClassName("profile_input")[0].value;
    Send("POST", "/profile/changeName", {Name: newName}, 1, (response) =>{
        if (response){
            window.location.reload();
        }
    })
}

function changeSurname(){
    let newSurname = document.getElementsByClassName("profile_input")[0].value;
    Send("POST", "/profile/changeSurname", {Surname: newSurname}, 1, (response) =>{
        if (response){
            window.location.reload();
        }
    })
}

function changeEmail(){
    let newEmail = document.getElementsByClassName("profile_input")[0].value;
    Send("POST", "/profile/changeEmail", {Email: newEmail}, 1, (response) =>{
        if (response){
            window.location.reload();
        }
    })
}

function changePassword(){
    let oldPassword = document.getElementsByClassName("profile_input")[0];
    let newPassword = document.getElementsByClassName("profile_input")[1];
    let repeatNewPassword = document.getElementsByClassName("profile_input")[2];
    if (oldPassword.value.length > 2 && oldPassword.value.length < 17){
        oldPassword.style.border = "unset";
        if (newPassword.value !== repeatNewPassword.value) {
            repeatNewPassword.style.border = "2px solid red";
        } else {
            repeatNewPassword.style.border = "unset";
            Send("POST", "/profile/changePassword", {
                OldPassword: oldPassword.value,
                NewPassword: newPassword.value
            }, 1, (response) =>{
                if (response === "Пароль изменен"){
                    window.location.reload();
                } else {
                    oldPassword.style.border = "2px solid red";
                }
            })
        }
    } else {
        oldPassword.style.border = "2px solid red";
    }
}

let search = document.getElementById("search");
if (search){
    let timeout = 0;
    search.oninput = function (e){
        if (timeout !== 0) {
            clearTimeout(timeout);
            timeout = 0;
        }
        timeout = setTimeout(() => {
            Send("POST", "/search", {Search: this.value}, 1, response => {
                let main = document.getElementsByClassName('main')[0];
                main.innerHTML = '';
                for (let i = 0; i < response.length; i++){
                    let div = document.createElement("div");
                    let a = document.createElement("a");
                    div.className = "test";
                    div.id = "test"+response[i].Id;
                    div.textContent = response[i].Name;
                    a.className = "test__btn";
                    a.textContent = "Пройти";
                    a.href = "/test/"+response[i].Id;
                    div.append(a);
                    main.append(div);
                }
            })
        }, 400)
    }
}