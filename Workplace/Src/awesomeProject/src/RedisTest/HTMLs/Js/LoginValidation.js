function loadXMLDoc() {
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    if (username == null || password == null || username == "" || password ==""){
        alert("usename or password cannont be empty!");
    }

    var xmlhttp;
    if (window.XMLHttpRequest)
    {
        //  IE7+, Firefox, Chrome, Opera, Safari 浏览器执行代码
        xmlhttp=new XMLHttpRequest();
    }
    else
    {
        // IE6, IE5 浏览器执行代码
        xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlhttp.onreadystatechange=function()
    {
        // alert(xmlhttp.status)
        //
        // if (xmlhttp.readyState==4 && xmlhttp.status==200)
        // {
        //     window.location.href = './index.html'
        // } else if (xmlhttp.readyState==4 && xmlhttp.status==801){
        //     alert("password not matched")
        // } else if (xmlhttp.readyState==4 && xmlhttp.status==800){
        //     alert("No username")
        // } else if (xmlhttp.readyState==4){
        //     alert("refused")
        // }
        if (xmlhttp.readyState == 4) {
                alert(xmlhttp.status)
            switch (xmlhttp.status) {
                case 0:
                    alert("OK")
                    window.location.href = './index.html'
                    break
                case 801:
                    alert("password no matched")
                    break
                case 800:
                    alert("No username")
                    break
                default:
                    alert("refused")
                    break
            }
        }
    }
    xmlhttp.open("post",("http://localhost:9091/login"),true);
    xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xmlhttp.send(JSON.stringify({"username":username,"password":password}));
}