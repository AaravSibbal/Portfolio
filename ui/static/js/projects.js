window.addEventListener("DOMContentLoaded", () => {
    colorLinks()
    alternateAlignment()
})

function colorLinks() {
    const linkList = document.getElementsByClassName("p-link")
    for (let i = 0; i < linkList.length; i++) {
        let modVal = i % 3
        if (modVal === 0) {
            linkList[i].style.color = "yellowgreen"
        }
        else if (modVal === 1) {
            linkList[i].style.color = "coral"
        }
        else {
            linkList[i].style.color = "cyan"
        }
    }
}

function alternateAlignment() {
    const projectList = document.getElementsByClassName("project")
    console.log(projectList)
    for(let i=0; i<projectList.length; i++){
        let modVal = i%2
        if(modVal === 0){
            projectList[i].style.textAlign = "left"
        }
        else{
            projectList[i].style.textAlign = "right"
        }
    }
}