// books界面
function getCookie(name){
	var name = name + "=";
	var ca = document.cookie.split(';');
	for(var i=0; i<ca.length; i++) {
		var c = ca[i].trim();
		if (c.indexOf(name)==0) { 
      return c.substring(name.length,c.length); 
    }
	}
	return "";
}

function getUrl(){
  var url = location.search.substring(1)
  return url
}

function books(){
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("GET","http://10.150.15.145:8080/books",true);
  xmlhttp.send();
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      bookNameDiv(xmlhttp.responseText);
    }
  }
}

function bookNameDiv(data){
  var aa = JSON.parse(data);
  console.log(aa)
  var div1 = document.createElement("div");
  div1.setAttribute("class", "novelslist");
  for (var i=0; i<aa.length; i++){
    var div = document.createElement("div");
    div.setAttribute("class", "divv");
    var a = document.createElement("a");
    a.innerHTML = aa[i].bookName;
    a.href = "chapters.html?"+encodeURI(aa[i].bookName)+"?"+aa[i].bookId.toString();
    div.appendChild(a);
    div1.appendChild(div)
    //document.body.appendChild(div);
  }
  document.body.appendChild(div1);
}



function userCookie(){
  var user = getCookie("name");
  if (user != ""){
    var aa = document.getElementById("user");
    aa.innerHTML = user
    aa.href="login.html"
  }
  return user
}


function searchBook(){
  var num=document.getElementById("in");
  var hash1 = num.value
  var xmlhttp;
  if (window.XMLHttpRequest)
  {
    xmlhttp=new XMLHttpRequest();
  }
  else
  {
    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
  }
  xmlhttp.open("POST","http://10.150.15.145:8080/book",true);
  xmlhttp.send(JSON.stringify({name:hash1}));
  xmlhttp.onreadystatechange=function()
  {
    if (xmlhttp.readyState==4 && xmlhttp.status==200)
    {
        judgeBook(xmlhttp.responseText)
    }
  }
}  

function judgeBook(data){
  var aa = JSON.parse(data)
  var bookName = aa.bookName
  var bookId = aa.bookId
  var ss = encodeURI(bookName)
  if (ss == ""){
    window.location.href="searchBooks.html?"+"The query is not"+"?"+String(bookId)
    return
  }
  window.location.href="searchBooks.html?"+ss+"?"+String(bookId)
}

function search(){
  var hash1 = location.search.substring(1)
  var a = hash1.split("?")
  var name = decodeURI(a[0])
  var num=document.getElementById("name");
  if (a[0] == "The query is not") {
    num.innerHTML = "The query is not"
    return
  }
  num.innerHTML = name
  num.href = "chapters.html?"+a[0]+"?"+a[1]
}

function classification(){
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("GET","http://10.150.15.145:8080/classification",true);
  xmlhttp.send();
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200)
    {
      classdiv(xmlhttp.responseText)
    }
  }
}

function classdiv(data){
  var classdata = JSON.parse(data)
  var h3 = document.createElement("div");
  h3.setAttribute("class", "head");
  for (var i=0; i<classdata.length; i++){
    var h2 = document.createElement("div")
    h2.setAttribute("class", "divv");
    var a = document.createElement("a")
    a.innerHTML = classdata[i];
    a.href = "classBooks.html?"+encodeURI(classdata[i])
    h2.appendChild(a)
    h3.appendChild(h2)
  }
  document.body.appendChild(h3)
}

//chapters 界面

function bookInfo(){
  var url = getUrl()
  var urlList = url.split("?")
  var name = decodeURI(urlList[0])
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("POST","http://10.150.15.145:8080/book",true);
  xmlhttp.send(JSON.stringify({name:name}));
  xmlhttp.onreadystatechange=function()
  {
    if (xmlhttp.readyState==4 && xmlhttp.status==200)
    {
      bookInfoTable(xmlhttp.responseText)
    }
  }
}

function bookInfoTable(data){
  var bookInfo = JSON.parse(data);
  var bookTable = document.getElementsByName("info");
  bookTable[0].innerHTML = bookInfo.bookAuthor;
  bookTable[1].innerHTML = bookInfo.bookName;
  bookTable[2].innerHTML = bookInfo.bookClassification;
  bookTable[3].innerHTML = bookInfo.ChapterNumber;
  bookTable[4].innerHTML = bookInfo.bookSize;
  bookTable[5].innerHTML = bookInfo.likeNumber;
  bookTable[6].innerHTML = bookInfo.readingNumber;
  bookTable[7].innerHTML = bookInfo.description;
}


function chapters(){
  var url = getUrl()
  var book = url.split("?")
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("POST","http://10.150.15.145:8080/chapters",true);
  xmlhttp.send(JSON.stringify({code:book[1]}));
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      chapterNameDiv(xmlhttp.responseText, book)
    }
  }
}

function chapterNameDiv(data, book){
  var bookName = book[0]
  var chapters = JSON.parse(data)
  var div1 = document.createElement("div");
  div1.setAttribute("class", "novelslist");
  for (var i=0; i<chapters.length; i++){
    var div = document.createElement("div");
    div.setAttribute("class", "divv");
    var a = document.createElement("a");
    a.innerHTML = chapters[i].chaName;
    a.href = "chapter.html?"+bookName+"?"+chapters[i].bookId.toString()+"?"+chapters[i].chaId.toString();
    div.appendChild(a);
    div1.appendChild(div);
    //document.body.appendChild(div)
  }
  document.body.appendChild(div1)
}

//chapter 界面

function chapter(){
  var userName = getCookie("name");
  if (userName == ""){
    window.location.href="/login1.html"
    return
  }
  var url = getUrl()
  var urlList = url.split("?")
  bookName = decodeURI(urlList[0])
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("POST","http://10.150.15.145:8080/chapter",true);
  xmlhttp.send(JSON.stringify({bookName:bookName,bookId:urlList[1], chapterId:urlList[2]}));
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      chapterData(xmlhttp.responseText, urlList)
    }
  }
}

function chapterData(data, urlList){
  var chapterData = JSON.parse(data)
  var money = chapterData.userMoney
  var chapter = chapterData.chapter

  if (urlList[2]>3){
    var a=confirm("付费章节，扣除1元");
    if (a){
      if (money<1){
        var b=confirm("余额不足，请充值");
        if (b){
          window.location.href="/topup.html"
        }else{
          window.history.back()
        }
      }else{
        var xmlhttp=new XMLHttpRequest();
        xmlhttp.open("POST","http://10.150.15.145:8080/userdeduct",true);
        xmlhttp.send(JSON.stringify({num:1}));
        xmlhttp.onreadystatechange=function(){
          if (xmlhttp.readyState==4 && xmlhttp.status==200){
            deductMoney(xmlhttp.responseText, chapter)
          }
        }
      }
    }else{
      window.history.back()
    }
  }else{
    var div = document.getElementById("div");
    div.innerHTML = chapter
  }
}

function deductMoney(data, chapter){
  var judge=JSON.parse(data)
  if (judge){
    var div = document.getElementById("div");
    div.innerHTML = chapter 
  }else{
    alert("扣除失败")
  }
}

function nextChapter(){
  var url = getUrl()
  console.log(url)
  var urlList = url.split("?")
  console.log(urlList)
  urlList[2] = String(Number(urlList[2])+1)
  var next = document.getElementById("next");
  next.href = "chapter.html?"+urlList[0]+"?"+urlList[1]+"?"+urlList[2]
}

function onChapter(){
  var url = getUrl()
  var urlList = url.split("?")
  if (urlList[2] == "1"){
    return
  }
  urlList[2] = String(Number(urlList[2])-1)
  var on = document.getElementById("on");
  on.href = "chapter.html?"+urlList[0]+"?"+urlList[1]+"?"+urlList[2]
}

//user 界面


function user(){
  var name = userCookie();
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("GET","http://10.150.15.145:8080/getUser",true);
  xmlhttp.send();
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      userDataTable(xmlhttp.responseText, name)
    }
  }
}

function userDataTable(data, name){
  var userTable = document.getElementsByName("user")
  var userData = JSON.parse(data)
  userTable[0].innerHTML = name
  if (userData.Like == null){
    userTable[1].innerHTML = 0;
  }else{
    userTable[1].innerHTML = userData.Like.length;
  }
  if (userData.Read == null){
    userTable[2].innerHTML = 0
  }else{
    userTable[2].innerHTML = userData.Read.length;
    for (var i=0; i<userData.Read.length; i++){
      readBookDiv(userData.Read[i])
    }
  }
}

function readBookDiv(data){
  var div = document.createElement("div");
  var a = document.createElement("a")
  a.innerHTML = data[0]
  a.href = "chapters.html?"+encodeURI(data[0])+"?"+data[1]
  div.appendChild(a)
  document.body.appendChild(div);
}


function userLogin(){
  var userList = document.getElementsByName("user")
  var name = userList[0].value
  var password = userList[1].value
  console.log(userList)
  console.log(name, password)
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("POST","http://10.150.15.145:8080/login",true);
  xmlhttp.send(JSON.stringify({name:name, password:password}));
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      judgeUser(xmlhttp.responseText)
    }
  }
}

function judgeUser(data){

  var judge=JSON.parse(data)
  if (judge == 1){
    alert("登陆成功")
    window.location.href="/books.html"
  }
  if (judge == 2){
    alert("密码错误请重新输入")
  }
  if (judge == 0){
    var a=confirm("用户不存在，请注册");
    if(a){
      window.location.href="/registration.html";
    }
  }
}

function adminLogin(){
  var adminList = document.getElementsByName("admin");
  var name = adminList[0].value;
  var password = adminList[1].value
  console.log(adminList)
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("POST","http://10.150.15.145:8080/adminlogin",true);
  xmlhttp.send(JSON.stringify({name:name, password:password}));
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      judgeAdmin(xmlhttp.responseText)
    }
  }
}

function judgeAdmin(data){
  var judge=JSON.parse(data)
  console.log(data)
  if (judge == 1){
    alert("登陆成功")
    window.location.href="/books.html"
  }
  if (judge == 2){
    alert("密码错误请重新输入")
  }
  if (judge == 0){
    var a=confirm("用户不存在，请注册");
    if(a){
       window.location.href="/adminregistration.html";
    }
  }
}

function userTopup(){
  var password = document.getElementById("topup");
  var code = Number(password.value)
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("POST","http://10.150.15.145:8080/usertopup",true);
  xmlhttp.send(JSON.stringify({code:code}));
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      judgeUserTopup(xmlhttp.responseText)
    }
  }
}

function judgeUserTopup(data){
  var judge=JSON.parse(data)
  if (judge) {
    alert("充值成功100")
    window.history.back()
  }else{
    alert("充值码错误请重新输入")
  }
}


function classBooks () {
  var url = getUrl()
  var className = decodeURI(url)
  var xmlhttp=new XMLHttpRequest();
  xmlhttp.open("POST","http://10.150.15.145:8080/classbooks",true);
  xmlhttp.send(JSON.stringify({name:className}));
  xmlhttp.onreadystatechange=function(){
    if (xmlhttp.readyState==4 && xmlhttp.status==200){
      bookNameDiv(xmlhttp.responseText)
    }
  }
}

