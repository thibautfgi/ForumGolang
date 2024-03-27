



function likeUpdate() {
    console.log("hello");


    // let inputElement = document.querySelector('.inputUpVote');
    // let currentValue = parseInt(inputElement.value);
    // inputElement.value = currentValue + 1;
    // console.log("here");


    // inputElement.style.fontSize = "30px";



    // PERMET DE RECUP ID DE LA LISTE CLICKER
    var listItems = document.querySelectorAll(".inputUpVote");

    console.log(listItems);


    // listItems.forEach(function (item) {
    //     item.addEventListener("click", function () {
    //         var selectedTopicId = item.querySelector("input[type='radio']").id;
    //         console.log(selectedTopicId);
    //     });
}








// function test(button) {
//     // let computedStyle = window.getComputedStyle(button);
//     // let buttonColor = computedStyle.getPropertyValue('color');

//     if (button.style.color === "rgba(194, 225, 16, 0.755)") {
//         console.log("1");
//     } else {
//         console.log("2");
//     }
// }

// var i = 0;
// var original = document.getElementById('commentBoxBot');

// function newComment() {
//     var clone = original.cloneNode(true); // "deep" clone
//     clone.id = "";
//     // or clone.id = ""; if the divs don't need an ID
//     original.parentNode.appendChild(clone);
// }


// var test = 0;


// function newComment2() {


// var jstestImg =" testImg"
// var jstestTitle = "nomUtilisateur" 
// var jstestInput =" testTitle"
// var jstestUpVote = 5
// var jstestDate = "tabsUser"





// // Créez la div que vous souhaitez copier
// const divToCopy = document.createElement('class');
// divToCopy.innerHTML = '<div class="commentBox"><div class="commentBoxHead"><div class="commentBoxAvatar"><img src="'+
// jstestImg+'"alt="testimgnm" width="50" height="50" ></img></div><div class="commentBoxName">'+
// jstestTitle+'</div></div><div class="commentBoxBody"><h1comment>'+
// jstestInput+'</h1comment></div><div class="commentBoxBottom"><div onclick="count()" class="commentBoxUpVote"><button class="btnUpVote"><i class="material-icons"> star </i></button><input type="number" value="'+
// jstestUpVote+'"class="inputUpVote"></input></div><div class="commentBoxDate">Publié le '+
// jstestDate+'</div></div>';

// const divCopy = divToCopy.cloneNode(true);

// const list = document.getElementById('list');

// list.appendChild(divCopy);
// document.getElementById("inputBox").value = ""; //erase the texto 
// }

