function checkCookie(){
  var user = getCookie("name");
  console.log(user)
  if (user != ""){
    var aa = document.getElementById("user");
    aa.innerHTML = user
    aa.href="login.html"
  }
  return user
}
 
function getCookie(name){
  var name = name + "=";
  var ca = document.cookie.split(';');
  for(var i=0; i<ca.length; i++) {
    var c = ca[i].trim();
    if (c.indexOf(name)==0) {
      return c.substring(name.length,c.length)
    }
  }
  return "";
}


function user(){
  name = checkCookie()
  var xmlhttp;
  if (window.XMLHttpRequest){
    xmlhttp=new XMLHttpRequest();
  }else{
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
  }
  xmlhttp.open("GET","http://localhost:8080/getUser",true);
  xmlhttp.send();
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      userData(xmlhttp.responseText, name)
    }
  }
}

function userData(data, name){
  var td = document.getElementsByName("user")
  var data = JSON.parse(data)
  td[0].innerHTML = name
  if (data.Like == null){
    console.log(data.Like)
    td[1].innerHTML = 0;
  }else{
    td[1].innerHTML = data.Like.length;
  }
  if (data.Read == null){
    td[2].innerHTML = 0
  }else{
    td[2].innerHTML = data.Read.length;
    for (var i=0; i<data.Read.length; i++){
      bookDiv(data.Read[i])
    }
  }
}
var bookList = document.body;
function bookDiv(data){
  console.log(data)
  var div = document.createElement("div");
  var a = document.createElement("a")
  a.innerHTML = data[0]
  a.href = "chapters.html?"+data[1]
  div.appendChild(a)
  bookList.appendChild(div);
}
